package redis

func RPopPush(listKey string, toListKey string) (string, error) {
	rlt, err := doCommand("RPopPush", listKey, toListKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// RPush queue right
func RPush(listKey string, val string) (int64, error) {
	rlt, err := doCommand("RPush", listKey, val)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// RPop queue consumption
func RPop(listKey string) (string, error) {
	rlt, err := doCommand("RPop", listKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

// LPush queue right
func LPush(listKey string, val string) (int64, error) {
	rlt, err := doCommand("LPush", listKey, val)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

// LPop queue consumption
func LPop(listKey string) (string, error) {
	rlt, err := doCommand("LPop", listKey)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func LSet(listKey string, index int64, val string) (string, error) {
	rlt, err := doCommand("LSet", listKey, index, val)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func LIndex(listKey string, i int64) (string, error) {
	rlt, err := doCommand("LIndex", listKey, i)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func LInsert(listKey string, target string, val string) (int64, error) {
	rlt, err := doCommand("LInsert", listKey, "BEFORE", target, val)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

func LLen(listKey string) (int64, error) {
	rlt, err := doCommand("LLen", listKey)
	if err != nil {
		return 0, err
	}
	return formatInt64(rlt), err
}

func LRange(listKey string, start int64, end int64) ([]string, error) {
	rlt, err := doCommand("LRange", listKey, start, end)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}
