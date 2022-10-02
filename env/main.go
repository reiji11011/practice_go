package env

import (
	"fmt"
	"os"
)

func Test() {
	os.Setenv("test", "Key：testです！")
	test := os.Getenv("test")
	test2 := os.Getenv("LOCAL")
	fmt.Println(test)
	fmt.Println(test2)
}
