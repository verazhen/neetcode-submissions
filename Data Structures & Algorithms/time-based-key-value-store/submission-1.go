type TimeMap struct {
	kvStore map[string][]mapValue
}

type mapValue struct{
	value string
	ts int
}

/*
{
	"alice": [
		{
			"value": happy,
			"ts": 1
		},
		{
			"value": sad,
			"ts": 3
		},
	]
}
*/

func Constructor() TimeMap {
	return TimeMap{map[string][]mapValue{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.kvStore[key] = append(this.kvStore[key],mapValue{value:value,ts:timestamp})

}

func (this *TimeMap) Get(key string, timestamp int) string {
	list := this.kvStore[key]
	if len(list) == 0{
		return ""
	}

	i := 0
	j := len(list) - 1
	for i < j{
		m := (i+j+1)/2
		if list[m].ts == timestamp{
			return list[m].value
		}else if list[m].ts > timestamp{
			j = m - 1
		}else{
			i = m
		}
	}

	if list[i].ts <= timestamp{
		return list[i].value
	}

	return ""
}
