package assert

import (
	"fmt"
	"testing"
)

func Equal[T int | string](t *testing.T, expected, actual T) {
	if expected == actual {
		return
	}

	t.Errorf("\n%s",
		fmt.Sprintf("Not equal: \n"+
			"expected: %v\n"+
			"actual  : %v", expected, actual))
}
