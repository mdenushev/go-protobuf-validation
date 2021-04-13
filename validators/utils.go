package validators

func IsInArray(val interface{}, allowedArr []interface{}) bool {
	for _, allowed := range allowedArr {
		if allowed == val {
			return true
		}
	}
	return false
}
