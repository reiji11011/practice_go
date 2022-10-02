package volatiletech

import (
	"fmt"
	"github.com/volatiletech/null"
)

func Test() string {
	num := null.NewUint32(1, false)
	fmt.Println(num)
	fmt.Println(num.Uint32)

	num1 := null.Uint32FromPtr(nil)
	fmt.Println(num1.Uint32)

	var num2 null.Uint32
	fmt.Println(num2)
	fmt.Println(num2.Uint32)

	n := null.Uint32{}
	fmt.Println(n)

	return ""
}
