package utilities

func CheckList(key , list interface{}) bool {
	switch list.(type) {
	case []string:
		for _, v := range list.([]string) {
			if key.(string) == v {
				return true
			}
		}
	case []int:
		for _, v := range list.([]int) {
			if key.(int) == v {
				return true
			}
		}
	}

	return false
}