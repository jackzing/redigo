package redis

func HSet(hashKey string, key string, val interface{}) (bool, error) {
	rlt, err := doCommand("HSet", hashKey, key, val)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

func HMSet(hashKey string, keyVal ...KeyVal) (string, error) {
	var params []interface{}
	params = append(params, hashKey)
	for _, v := range keyVal {
		params = append(params, v.Key)
		params = append(params, v.Val)
	}
	rlt, err := doCommand("HMSet", params...)
	if err != nil {
		return "", err
	}
	return rlt.(string), err
}

func HGet(hashKey string, key string) (string, error) {
	rlt, err := doCommand("HGet", hashKey, key)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func HGetAll(hashKey string) (map[string]string, error) {
	rlt, err := doCommand("HGetAll", hashKey)
	if err != nil {
		return nil, err
	}
	i := rlt.([]interface{})
	data := make(map[string]string)
	for k := 0; k < len(i); k += 2 {
		if i[k] != nil {
			data[string(i[k].([]byte))] = string(i[k+1].([]byte))
		} else {
			data[string(i[k].([]byte))] = ""
		}
	}
	return data, err
}

func HMGet(hashKey string, keys ...interface{}) (map[string]string, error) {
	var params []interface{}
	params = append(params, hashKey)
	for _, v := range keys {
		params = append(params, v)
	}
	rlt, err := doCommand("HMGet", params...)
	if err != nil {
		return nil, err
	}
	return formatMap(rlt, keys...), err
}

func HDel(hashKey string, key string) (bool, error) {
	rlt, err := doCommand("HDel", hashKey, key)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

func HExists(hashKey string, key string) (bool, error) {
	rlt, err := doCommand("HExists", hashKey, key)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

func HIncrBy(hashKey string, key string, num int64) (int64, error) {
	rlt, err := doCommand("HIncrBy", hashKey, key, num)
	if err != nil {
		return 0, err
	}
	return rlt.(int64), err
}

func HKeys(hashKey string) ([]string, error) {
	rlt, err := doCommand("HKeys", hashKey)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}
func HVals(hashKey string) ([]string, error) {
	rlt, err := doCommand("HVals", hashKey)
	if err != nil {
		return nil, err
	}
	return formatSlice(rlt), err
}

func HLen(hashKey string) (int64, error) {
	rlt, err := doCommand("HLen", hashKey)
	if err != nil {
		return 0, err
	}
	return rlt.(int64), err
}
