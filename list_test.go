package redis

import (
	"fmt"
	"testing"
)

func TestListFunc(t *testing.T) {
	key := "list:key"
	fmt.Println(LPush(key, "1")) // 1
	fmt.Println(LPush(key, "2")) // 2
	fmt.Println(LPush(key, "3")) // 3
	fmt.Println(LLen(key))            // 3
	fmt.Println(LRange(key, 0, -1)) // [3 2 1]

	fmt.Println(RPush(key, "1")) // 1
	fmt.Println(RPush(key, "2")) // 2
	fmt.Println(RPush(key, "3")) // 3
	fmt.Println(LLen(key))            // 3
	fmt.Println(LRange(key, 0, -1)) // [3 2 1]

	fmt.Println(LIndex(key, 10)) // [3 2 1]

	fmt.Println(LPop(key)) // asd
	fmt.Println(RPop(key)) // asd

}
