package validations

func IsEmpty(param string) bool {
	if param == "" {
		return true;
	}
	return false
}
