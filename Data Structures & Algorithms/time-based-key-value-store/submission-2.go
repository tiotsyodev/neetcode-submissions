type TimeMap struct {
	store map[string][]struct {
        timestamp int
        value     string
    }
}

func Constructor() TimeMap {
	 return TimeMap{
        store: make(map[string][]struct {
            timestamp int
            value     string
        }),
    }
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.store[key] = append(tm.store[key], struct{
		timestamp int
		value string
	}{timestamp, value})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	list, exists := tm.store[key]
    if !exists {
        return ""
    }

	left, right := 0, len(list)-1
    result := ""
    
    for left <= right {
        mid := left + (right-left)/2
        
        if list[mid].timestamp <= timestamp {
            result = list[mid].value
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}
