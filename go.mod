module github.com/go-critic/go-critic

go 1.16

require (
	github.com/cristalhq/acmd v0.8.1
	github.com/go-toolsmith/astcast v1.0.0
	github.com/go-toolsmith/astcopy v1.0.2
	github.com/go-toolsmith/astequal v1.0.3
	github.com/go-toolsmith/astfmt v1.0.0
	github.com/go-toolsmith/astp v1.0.0
	github.com/go-toolsmith/pkgload v1.0.2-0.20220101231613-e814995d17c5
	github.com/go-toolsmith/strparse v1.0.0
	github.com/go-toolsmith/typep v1.0.2
	github.com/google/go-cmp v0.5.9
	github.com/quasilyte/go-ruleguard v0.3.18
	github.com/quasilyte/go-ruleguard/dsl v0.3.21
	github.com/quasilyte/gogrep v0.0.0-20221002170714-e78263da2dd3
	github.com/quasilyte/regex/syntax v0.0.0-20200407221936-30656e2c4a95
	golang.org/x/exp/typeparams v0.0.0-20221002003631-540bb7301a08
	golang.org/x/tools v0.1.12
)

replace github.com/quasilyte/go-ruleguard v0.3.18 => github.com/peakle/go-ruleguard v0.3.22
