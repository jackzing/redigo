package redis

func RPopPush(listKey string, toListKey string) (string, error) {
	rlt, err := DB.Get().Do("RPopPush", listKey, toListKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// queue right
func RPush(listKey string, val string) (int64, error) {
	rlt, err := DB.Get().Do("RPush", listKey, val)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// queue consumption
func RPop(listKey string) (string, error) {
	rlt, err := DB.Get().Do("RPop", listKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// queue right
func LPush(listKey string, val string) (int64, error) {
	rlt, err := DB.Get().Do("LPush", listKey, val)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// queue consumption
func LPop(listKey string) (string, error) {
	rlt, err := DB.Get().Do("LPop", listKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func LSet(listKey string, index int64, val string) (string, error) {
	rlt, err := DB.Get().Do("LSet", listKey, index, val)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func LIndex(listKey string, i int64) (string, error) {
	rlt, err := DB.Get().Do("LIndex", listKey, i)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func LInsert(listKey string, target string, val string) (int64, error) {
	rlt, err := DB.Get().Do("LInsert", listKey, "BEFORE", target, val)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

func LLen(listKey string) (int64, error) {
	rlt, err := DB.Get().Do("LLen", listKey)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

func LRange(listKey string, start int64, end int64) ([]string, error) {
	rlt, err := DB.Get().Do("LRange", listKey, start, end)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}
