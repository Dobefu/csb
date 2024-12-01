package utils

func ConstructOutput() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"data":  make(map[string]interface{}),
		"error": nil,
	}
}
