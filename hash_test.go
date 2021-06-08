package redis

import (
	"fmt"
	"testing"
)
func init() {
	r := Config{
		Host:        "127.0.0.7:6379",
		Password:    "123456",
		MaxActive:   100,
		MaxIdle:     20,
		IdleTimeout: 300,
	}
	ConRedis(r)
}git status
func TestHashFunc(t *testing.T) {
	hasKey := "hash:key"
	fmt.Print(HSet(hasKey, "one", 123))                                                 // true
	fmt.Println("// true")
	fmt.Print(HMSet(hasKey, KeyVal{Key: "two", Val: 456}, KeyVal{Key: "wu", Val: 789})) // OK
	fmt.Println("// OK")
	fmt.Print(HGet(hasKey, "two"))                                                      // 456
	fmt.Println("// 456")
	fmt.Print(HMGet(hasKey, "one", "two"))                                              // map[one:123 two:456]
	fmt.Println("// map[one:123 two:456]")
	fmt.Print(HGetAll(hasKey))                                                          // map[one:123 two:456 wu:789]
	fmt.Println("// map[one:123 two:456 wu:789]")

	fmt.Print(HDel(hasKey, "wu"))         // true
	fmt.Println("// true")
	fmt.Print(HExists(hasKey, "wu"))      // false
	fmt.Println("// false")
	fmt.Print(HIncrBy(hasKey, "one", 1))  // 124
	fmt.Println("// 124")
	fmt.Print(HIncrBy(hasKey, "one", -1)) // 123
	fmt.Println("// 123")
	fmt.Print(HKeys(hasKey))              // [one two]
	fmt.Println("// [one two]")
	fmt.Print(HVals(hasKey))              // [123 456]
	fmt.Println("// [123 456]")
	fmt.Print(HLen(hasKey))               // 2
	fmt.Println("// 2")

}
