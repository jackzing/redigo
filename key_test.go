package redis

import (
	"fmt"
	"testing"
)

func TestKeyFunc(t *testing.T) {
	key := "test:key"
	fmt.Print(Set(key, 666))     // true
	fmt.Println("// true")
	fmt.Print(Get(key))          // 666
	fmt.Println("// 666")
	fmt.Print(Decr(key))         // 665
	fmt.Println("// 665")
	fmt.Print(Incr(key))         // 666
	fmt.Println("// 666")
	fmt.Print(IncrBy(key, -333)) // 333
	fmt.Println("// 333")
	fmt.Print(IncrBy(key, 666))  // 999
	fmt.Println("// 999")
	fmt.Print(Del(key))          // true
	fmt.Println("// true")

	fmt.Print(MSet(KeyVal{Key: "one", Val: 123}, KeyVal{Key: "two", Val: 456})) // true
	fmt.Println("// true")
	fmt.Print(MGet("one", "two", key))                                          // map[one:123 test:key: two:456]
	fmt.Println("// map[one:123 test:key: two:456]")

	fmt.Print(SetEx(key, "SetEx", 30))      // OK
	fmt.Println("// OK")
	fmt.Print(SetNx(key, "setNx"))          // false
	fmt.Println("// false")
	fmt.Print(SetNx("set:nx", "setNx", 30)) // true
	fmt.Println("// true")


	lock := make(chan string)
	for i := 0; i < 100; i++ {
		go func() {
			if Lock("lock:key", 30) {
				fmt.Print("lock true") // lock true only once
				fmt.Println("// lock true only once")
			}
			lock <- "lock"
		}()
	}
	for i := 0; i < 100; i++ {
		<-lock
	}
	fmt.Print(UnLock("lock:key")) // true
	fmt.Println("// true")

}
