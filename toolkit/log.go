package toolkit

import (
	"fmt"
)

func Output(key, val string) {
	fmt.Printf("::set-output name=%s::%s\n", key, val)
}
