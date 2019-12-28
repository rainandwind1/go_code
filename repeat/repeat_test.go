package repeat

import (
	"fmt"
	"testing"
)

func Testrepeat(t *testing.T) {
	result := repeat(0, 5)
	d := 10
	if result == d {
		fmt.Printf("%d", result)
	}

}
