package redis

// SAdd 添加一个或多个成员
func SAdd(setKey string, members ...string) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v)
	}
	rlt, err := doCommand("SAdd", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// SRem 移除一个或多个成员
func SRem(setKey string, members ...string) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v)
	}
	rlt, err := doCommand("SRem", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// SCard 获取集合成员数
func SCard(setKey string) (int64, error) {
	rlt, err := doCommand("SCard", setKey)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// SDiff 返回其他集合不存在的成员列表
func SDiff(setKeys ...interface{}) ([]string, error) {
	rlt, err := doCommand("SDiff", setKeys...)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// SInter 返回多个集合交集
func SInter(setKeys ...interface{}) ([]string, error) {
	rlt, err := doCommand("SInter", setKeys...)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// SUnion 返回多个集合并集
func SUnion(setKeys ...interface{}) ([]string, error) {
	rlt, err := doCommand("SUnion", setKeys...)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// SIsMember 判断集合是否存在成员
func SIsMember(setKey string, member string) (bool, error) {
	rlt, err := doCommand("SIsMember", setKey, member)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

// SMove 把成员member从source集合移动到destination集合
func SMove(source string, destination string, member string) (bool, error) {
	rlt, err := doCommand("SMove", source, destination, member)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

// SPop 随机获取并移除一个成员
func SPop(setKey string) (string, error) {
	rlt, err := doCommand("SPop", setKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// SMembers 获取集合成员列表
func SMembers(setKey string) ([]string, error) {
	rlt, err := doCommand("SMembers", setKey)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// SRandMember 随机获取 count 个成员
func SRandMember(setKey string, count int64) ([]string, error) {
	rlt, err := doCommand("SRandMember", setKey, count)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

// 有序集合

// ZAdd 添加一个或多个有序集合
func ZAdd(setKey string, members ...SortSet) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v.Score)
		params = append(params, v.Member)
	}
	rlt, err := doCommand("ZAdd", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// ZCard 集合成员数
func ZCard(setKey string) (int64, error) {
	rlt, err := doCommand("ZCard", setKey)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// ZCount 计算score 之间的数量
func ZCount(setKey string, min int64, max int64) (int64, error) {
	rlt, err := doCommand("ZCount", setKey, min, max)
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

// ZIncrBy 对成员分数修改
func ZIncrBy(setKey string, member string, sort int64) (string, error) {
	rlt, err := doCommand("ZIncrBy", setKey, sort, member)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// ZRangeMap 取范围 0 - -1 返回 map
func ZRangeMap(setKey string, start int64, stop int64) (map[string]string, error) {
	rlt, err := doCommand("ZRange", setKey, start, stop, "WithScores")
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

// ZRange 取范围 0 - -1 返回 排序 递增
func ZRange(setKey string, start int64, stop int64) ([]SortSet, error) {
	rlt, err := doCommand("ZRange", setKey, start, stop, "WithScores")
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// ZRevRange 取范围 0 - -1 返回 排序 递减
func ZRevRange(setKey string, start int64, stop int64) ([]SortSet, error) {
	rlt, err := doCommand("ZRevRange", setKey, start, stop, "WithScores")
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// ZRangeByScoreMap 取值 score 区间 -inf => (开始) +inf => (结束) (666 => (大于666) 666 => (大于等于666) 返回 map
func ZRangeByScoreMap(setKey string, start string, stop string, offsetAndCount ...string) (map[string]string, error) {
	var args []interface{}
	args = append(args, setKey, start, stop, "WithScores")
	if len(offsetAndCount) > 0 {
		args = append(args, "limit", offsetAndCount[0], offsetAndCount[1])
	}
	rlt, err := doCommand("ZRangeByScore", args...)
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

// ZRangeByScore 取值 score 区间 -inf => (开始) +inf => (结束) (666 => (大于666) 666 => (大于等于666) 递增
func ZRangeByScore(setKey string, start string, stop string, offsetAndCount ...string) ([]SortSet, error) {
	var args []interface{}
	args = append(args, setKey, start, stop, "WithScores")
	if len(offsetAndCount) > 0 {
		args = append(args, "limit", offsetAndCount[0], offsetAndCount[1])
	}
	rlt, err := doCommand("ZRangeByScore", args...)
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// ZRevRangeByScore 取值 score 区间 -inf => (开始) +inf => (结束) (666 => (大于666) 666 => (大于等于666) 递减
func ZRevRangeByScore(setKey string, start string, stop string,offsetAndCount ...string) ([]SortSet, error) {
	var args []interface{}
	args = append(args, setKey, start, stop, "WithScores")
	if len(offsetAndCount) > 0 {
		args = append(args, "limit", offsetAndCount[0], offsetAndCount[1])
	}
	rlt, err := doCommand("ZRevRangeByScore", args...)
	if err != nil {
		return nil, err
	}
	return formatSortSet(rlt), err
}

// ZRank 获取member 排名 递增
func ZRank(setKey string, member string) (int64, error) {
	rlt, err := doCommand("ZRank", setKey, member)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// ZRevRank 获取member 排名 递减
func ZRevRank(setKey string, member string) (int64, error) {
	rlt, err := doCommand("ZRevRank", setKey, member)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// ZRem 移除有序集合一个或多个成员
func ZRem(setKey string, members ...interface{}) (int64, error) {
	var params []interface{}
	params = append(params, setKey)
	for _, v := range members {
		params = append(params, v)
	}
	rlt, err := doCommand("ZRem", params...)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// ZRemByRank 移除有序集合 指定排名区间的成员
func ZRemByRank(setKey string, min int64, max int64) (int64, error) {
	rlt, err := doCommand("ZRemRangeByRank", setKey, min, max)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// ZRemByScore 移除有序集合 指定分数区间的成员
func ZRemByScore(setKey string, min int64, max int64) (int64, error) {
	rlt, err := doCommand("ZRemRangeByScore", setKey, min, max)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// ZScore 返回有序集合成员的分数
func ZScore(setKey string, member string) (string, error) {
	rlt, err := doCommand("ZScore", setKey, member)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}
