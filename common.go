package redis

import "strconv"

type KeyVal struct {
	Key string
	Val interface{}
}

type SortSet struct {
	Score  int64
	Member string
}


func formatBool(rlt interface{}) bool {
	if rlt == nil {
		return false
	}
	i := rlt.(int64)
	if i == 1 {
		return true
	}
	return false
}

func formatInt64(rlt interface{}) int64 {
	if rlt == nil {
		return 0
	}
	return rlt.(int64)
}

func formatString(rlt interface{}) string {
	if rlt == nil {
		return ""
	}
	return string(rlt.([]byte))
}

func formatMap(rlt interface{}, keys ...interface{}) map[string]string {
	i := rlt.([]interface{})
	data := make(map[string]string)
	for k, v := range keys {
		if i[k] != nil {
			data[v.(string)] = string(i[k].([]byte))
		} else {
			data[v.(string)] = ""
		}
	}
	return data
}

func formatSlice(rlt interface{}) []string {
	i := rlt.([]interface{})
	var data []string
	for _, v := range i {
		data = append(data, string(v.([]byte)))
	}
	return data
}

func formatSortSet(rlt interface{}) []SortSet {
	sliceData := rlt.([]interface{})
	var data []SortSet
	for i := 0; i < len(sliceData); i += 2 {
		score, _ := strconv.Atoi(string(sliceData[i+1].([]byte)))
		data = append(data, SortSet{Score: int64(score), Member: string(sliceData[i].([]byte))})
	}
	return data
}
