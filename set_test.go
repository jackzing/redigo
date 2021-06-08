package redis

import (
	"fmt"
	"testing"
)

func TestSetFunc(t *testing.T) {
	key := "sets:key"
	key1 := "sets:key1"
	fmt.Println(SAdd(key, "m1", "m2", "m3", "m4", "m5", "m6"))  // 6（成功添加个数） 或 0
	fmt.Println(SAdd(key1, "k1", "k2", "k3", "k4", "m5", "m6")) // 6（成功添加个数） 或 0
	fmt.Println(SCard(key))                                     // 6
	fmt.Println(SCard(key1))                                    // 6
	fmt.Println(SMembers(key))                                  // [m2 m1 m6 m5 m4 m3]
	fmt.Println(SMembers(key1))                                 // [k1 k2 k3 k4 m6 m5]

	fmt.Println(SDiff(key, key1))      // [m2 m4 m1 m3]
	fmt.Println(SInter(key, key1))     // [m6 m5]
	fmt.Println(SUnion(key, key1))     // [m1 m2 m3 m4 m6 m5 k1 k2 k3 k4]
	fmt.Println(SIsMember(key, "m2"))  // true
	fmt.Println(SIsMember(key1, "m2")) // false

	fmt.Println(SMove(key1, key, "k2")) // true
	fmt.Println(SPop(key))              // 随机出现key中的成员
	fmt.Println(SRandMember(key1, 3))   // 随机出现key1中的3成员
	fmt.Println(SRem(key1, "k1", "k2")) // 1（成功移除个数） k2 不存在了

	// 有序集合
	sk := "sortSet:key"
	fmt.Println(ZAdd(sk, SortSet{Score: 5, Member: "five"}, SortSet{Score: 1, Member: "one"}, SortSet{Score: 3, Member: "three"}, SortSet{Score: 2, Member: "two"}, SortSet{Score: 9, Member: "four"}))
	fmt.Println(ZIncrBy(sk, "four", -5)) // 4
	fmt.Println(ZCard(sk))               // 5
	fmt.Println(ZCount(sk, 2, 5))        // 4
	//fmt.Println(ZLexCount(sk,"[1","[5"))
	fmt.Println(ZRangeMap(sk, 0, -1))            // map[1:one 2:two 3:three 4:four 5:five] <nil>
	fmt.Println(ZRange(sk, 0, 2))                // [{1 one} {2 two} {3 three}] <nil>
	fmt.Println(ZRevRange(sk, 0, 2))             // [{5 five} {4 four} {3 three}] <nil>
	fmt.Println(ZRangeByScoreMap(sk, "3", "5"))  // map[3:three 4:four 5:five] <nil>
	fmt.Println(ZRangeByScore(sk, "0", "3"))     // [{1 one} {2 two} {3 three}] <nil>
	fmt.Println(ZRevRangeByScore(sk, "1", "(3")) // [{2 two} {1 one}] <nil>
	fmt.Println(ZRank(sk, "two"))                // 1
	fmt.Println(ZRevRank(sk, "two"))             // 3
	fmt.Println(ZRem(sk, "five"))                // 1
	fmt.Println(ZRemByRank(sk, 0, 1))            // 2
	fmt.Println(ZRemByScore(sk, 2, 3))           // 1
	fmt.Println(ZScore(sk, "two"))               // nil

}
