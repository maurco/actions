package toolkit

import (
	"fmt"
)

func Output(key, val interface{}) {
	fmt.Printf("::set-output name=%v::%v\n", key, val)
}
