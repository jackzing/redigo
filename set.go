package redis

// 添加一个或多个成员
func SAdd(setKey string, members ...string) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v)
	}
	rlt, err := DB.Get().Do("SAdd", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 移除一个或多个成员
func SRem(setKey string, members ...string) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v)
	}
	rlt, err := DB.Get().Do("SRem", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 获取集合成员数
func SCard(setKey string) (int64, error) {
	rlt, err := DB.Get().Do("SCard", setKey)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 返回其他集合不存在的成员列表
func SDiff(setKeys ...interface{}) ([]string, error) {
	rlt, err := DB.Get().Do("SDiff", setKeys...)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// 返回多个集合交集
func SInter(setKeys ...interface{}) ([]string, error) {
	rlt, err := DB.Get().Do("SInter", setKeys...)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// 返回多个集合并集
func SUnion(setKeys ...interface{}) ([]string, error) {
	rlt, err := DB.Get().Do("SUnion", setKeys...)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// 判断集合是否存在成员
func SIsMember(setKey string, member string) (bool, error) {
	rlt, err := DB.Get().Do("SIsMember", setKey, member)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

// 把成员member从source集合移动到destination集合
func SMove(source string, destination string, member string) (bool, error) {
	rlt, err := DB.Get().Do("SMove", source, destination, member)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

// 随机获取并移除一个成员
func SPop(setKey string) (string, error) {
	rlt, err := DB.Get().Do("SPop", setKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// 获取集合成员列表
func SMembers(setKey string) ([]string, error) {
	rlt, err := DB.Get().Do("SMembers", setKey)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// 随机获取 count 个成员
func SRandMember(setKey string, count int64) ([]string, error) {
	rlt, err := DB.Get().Do("SRandMember", setKey, count)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// 有序集合

// 添加一个或多个有序集合
func ZAdd(setKey string, members ...SortSet) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v.Score)
		params = append(params, v.Member)
	}
	rlt, err := DB.Get().Do("ZAdd", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 集合成员数
func ZCard(setKey string) (int64, error) {
	rlt, err := DB.Get().Do("ZCard", setKey)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 计算score 之间的数量
func ZCount(setKey string, min int64, max int64) (int64, error) {
	rlt, err := DB.Get().Do("ZCount", setKey, min, max)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 计算member 之间的数量
//func ZLexCount(setKey string, start string, stop string) (int64, error) {
//	rlt, err := DB.Get().Do("ZLexCount", setKey, start, stop)
//	if err != nil {
//		return 0, err
//	}
//	return formatInt64(rlt), err
//}

// 对成员分数修改
func ZIncrBy(setKey string, member string, sort int64) (string, error) {
	rlt, err := DB.Get().Do("ZIncrBy", setKey, sort, member)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// 取范围 0 - -1 返回 map
func ZRangeMap(setKey string, start int64, stop int64) (map[string]string, error) {
	rlt, err := DB.Get().Do("ZRange", setKey, start, stop, "WithScores")
	if err != nil {
		return nil, err
	}
	data := formatSlice(rlt)
	mapRlt := make(map[string]string)
	for i := 0; i < len(data); i += 2 {
		mapRlt[data[i+1]] = data[i]
	}
	return mapRlt, err
}

// 取范围 0 - -1 返回 排序 递增
func ZRange(setKey string, start int64, stop int64) ([]SortSet, error) {
	rlt, err := DB.Get().Do("ZRange", setKey, start, stop, "WithScores")
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// 取范围 0 - -1 返回 排序 递减
func ZRevRange(setKey string, start int64, stop int64) ([]SortSet, error) {
	rlt, err := DB.Get().Do("ZRevRange", setKey, start, stop, "WithScores")
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// 取值 score 区间 -inf => (开始) +inf => (结束) (666 => (大于666) 666 => (大于等于666) 返回 map
func ZRangeByScoreMap(setKey string, start string, stop string) (map[string]string, error) {
	rlt, err := DB.Get().Do("ZRangeByScore", setKey, start, stop, "WithScores")
	if err != nil {
		return nil, err
	}
	data := formatSlice(rlt)
	mapRlt := make(map[string]string)
	for i := 0; i < len(data); i += 2 {
		mapRlt[data[i+1]] = data[i]
	}
	return mapRlt, err
}

// 取值 score 区间 -inf => (开始) +inf => (结束) (666 => (大于666) 666 => (大于等于666) 递增
func ZRangeByScore(setKey string, start string, stop string) ([]SortSet, error) {
	rlt, err := DB.Get().Do("ZRangeByScore", setKey, start, stop, "WithScores")
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// 取值 score 区间 -inf => (开始) +inf => (结束) (666 => (大于666) 666 => (大于等于666) 递减
func ZRevRangeByScore(setKey string, start string, stop string) ([]SortSet, error) {
	rlt, err := DB.Get().Do("ZRevRangeByScore", setKey, stop, start, "WithScores")
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// 获取member 排名 递增
func ZRank(setKey string, member string) (int64, error) {
	rlt, err := DB.Get().Do("ZRank", setKey, member)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 获取member 排名 递减
func ZRevRank(setKey string, member string) (int64, error) {
	rlt, err := DB.Get().Do("ZRevRank", setKey, member)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 移除有序集合一个或多个成员
func ZRem(setKey string, members ...interface{}) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v)
	}
	rlt, err := DB.Get().Do("ZRem", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 移除有序集合 指定排名区间的成员
func ZRemByRank(setKey string, min int64, max int64) (int64, error) {
	rlt, err := DB.Get().Do("ZRemRangeByRank", setKey, min, max)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 移除有序集合 指定分数区间的成员
func ZRemByScore(setKey string, min int64, max int64) (int64, error) {
	rlt, err := DB.Get().Do("ZRemRangeByScore", setKey, min, max)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// 返回有序集合成员的分数
func ZScore(setKey string, member string) (string, error) {
	rlt, err := DB.Get().Do("ZScore", setKey, member)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}
