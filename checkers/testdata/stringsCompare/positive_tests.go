package checker_test

import (
	"strings"
)

func warning() {
	f, b := "aaa", "bbb"

	/*! suggestion: f == b */
	if strings.Compare(f, b) == 0 {
	}

	/*! suggestion: f == b */
	if 0 == strings.Compare(f, b) {
	}

	/*! suggestion: f > b */
	switch foo := strings.Compare(f, b) > 0; foo {
	case true:
		print(0)
	case false:
		print(1)
	}

	/*! suggestion: "s" < "ww" */
	_ = strings.Compare("s", "ww") < 0
	/*! suggestion: "s" == "ww" */
	_ = strings.Compare("s", "ww") == 0
	/*! suggestion: "s" > "ww" */
	_ = strings.Compare("s", "ww") > 0

	/*! suggestion: "s" < "ww" */
	_ = 0 > strings.Compare("s", "ww")
	/*! suggestion: "s" == "ww" */
	_ = 0 == strings.Compare("s", "ww")
	/*! suggestion: "s" > "ww" */
	_ = 0 < strings.Compare("s", "ww")
}
