package checkers

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"os"

	"github.com/quasilyte/go-ruleguard/ruleguard"

	"github.com/go-critic/go-critic/checkers/rulesdata"
	"github.com/go-critic/go-critic/framework/linter"
)

//go:generate go run ./rules/precompile.go -rules ./rules/rules.go -o ./rulesdata/rulesdata.go

func InitEmbeddedRules() error {
	filename := "rules/rules.go"

	fset := token.NewFileSet()
	var groups []ruleguard.GoRuleGroup

	var buildContext *build.Context

	ruleguardDebug := os.Getenv("GOCRITIC_RULEGUARD_DEBUG") != ""

	// First we create an Engine to parse all rules.
	// We need it to get the structured info about our rules
	// that will be used to generate checkers.
	// We introduce an extra scope in hope that rootEngine
	// will be garbage-collected after we don't need it.
	// LoadedGroups() returns a slice copy and that's all what we need.
	{
		rootEngine := ruleguard.NewEngine()
		rootEngine.InferBuildContext()
		buildContext = rootEngine.BuildContext

		loadContext := &ruleguard.LoadContext{
			Fset:         fset,
			DebugImports: ruleguardDebug,
			DebugPrint: func(s string) {
				fmt.Println("debug:", s)
			},
		}
		if err := rootEngine.LoadFromIR(loadContext, filename, rulesdata.PrecompiledRules); err != nil {
			return fmt.Errorf("load embedded ruleguard rules: %w", err)
		}
		groups = rootEngine.LoadedGroups()
	}

	parseContext := &ruleguard.LoadContext{
		Fset:         fset,
		DebugImports: ruleguardDebug,
		DebugPrint: func(s string) {
			fmt.Println("debug:", s)
		},
	}
	engine := ruleguard.NewEngine()
	engine.BuildContext = buildContext
	err := engine.LoadFromIR(parseContext, filename, rulesdata.PrecompiledRules)
	if err != nil {
		return err
	}

	// For every rule's group we create a new checker and a separate engine.
	// That dedicated ruleguard engine will contain rules only from one group.
	for i := range groups {
		g := groups[i]
		info := &linter.CheckerInfo{
			Name:    g.Name,
			Summary: g.DocSummary,
			Before:  g.DocBefore,
			After:   g.DocAfter,
			Note:    g.DocNote,
			Tags:    g.DocTags,

			EmbeddedRuleguard: true,
		}

		engine := *engine
		engine.FilterLoadedGroups(g.Name)

		collection.AddChecker(info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
			return &embeddedRuleguardChecker{
				ctx:    ctx,
				engine: &engine,
			}, nil
		})
	}

	return nil
}

type embeddedRuleguardChecker struct {
	ctx    *linter.CheckerContext
	engine *ruleguard.Engine
}

func (c *embeddedRuleguardChecker) WalkFile(f *ast.File) {
	runRuleguardEngine(c.ctx, f, c.engine, &ruleguard.RunContext{
		Pkg:         c.ctx.Pkg,
		Types:       c.ctx.TypesInfo,
		Sizes:       c.ctx.SizesInfo,
		Fset:        c.ctx.FileSet,
		TruncateLen: 100,
	})
}
