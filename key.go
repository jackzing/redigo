package redis

// Set string action
func Set(key string, val interface{}) (bool, error) {
	rlt, err := doCommand("Set", key, val)
	if err != nil {
		return false, err
	}
	if rlt != nil {
		return true, err
	}
	return false, err
}

func MSet(keyVal ...KeyVal) (bool, error) {
	var params []interface{}
	for _, v := range keyVal {
		params = append(params, v.Key)
		params = append(params, v.Val)
	}
	rlt, err := doCommand("MSet", params...)
	if err != nil {
		return false, err
	}
	if rlt != nil {
		return true, err
	}
	return false, err
}

func Get(key string) (string, error) {
	rlt, err := doCommand("Get", key)
	if err != nil {
		return "", err
	}
	return formatString(rlt), err
}

func MGet(keys ...interface{}) (map[string]string, error) {
	rlt, err := doCommand("MGet", keys...)
	if err != nil {
		return nil, err
	}
	return formatMap(rlt, keys...), err
}

func Del(key string) (bool, error) {
	rlt, err := doCommand("Del", key)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

func Expire(key string, ttl int64) (bool, error) {
	rlt, err := doCommand("EXPIRE", key, ttl)
	if err != nil {
		return false, err
	}
	return formatBool(rlt), err
}

func SetEx(key string, val string, ttl int64) (string, error) {
	rlt, err := doCommand("SetEx", key, ttl, val)
	if err != nil {
		return "", err
	}
	return rlt.(string), err
}

func SetNx(key string, val string, ttl ...int64) (bool, error) {
	rlt, err := doCommand("SetNx", key, val)
	if err != nil {
		return false, err
	}
	if formatBool(rlt) {
		if len(ttl) > 0 {
			Expire(key, ttl[0])
		}
		return true, err
	}
	return false, err
}

func Decr(key string) (int64, error) {
	rlt, err := doCommand("Decr", key)
	if err != nil {
		return 0, err
	}
	return rlt.(int64), err
}

func Incr(key string) (int64, error) {
	rlt, err := doCommand("Incr", key)
	if err != nil {
		return 0, err
	}
	return rlt.(int64), err
}

func IncrBy(key string, num int64) (int64, error) {
	rlt, err := doCommand("IncrBy", key, num)
	if err != nil {
		return 0, err
	}
	return rlt.(int64), err
}

func Lock(key string, ttl ...int64) bool {
	rlt, _ := doCommand("SetNx", key, 1)
	if formatBool(rlt) {
		if len(ttl) > 0 {
			Expire(key, ttl[0])
		}
		return true
	}
	return false
}

func UnLock(key string) (bool, error) {
	return Del(key)
}
