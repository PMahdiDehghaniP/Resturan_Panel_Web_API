package logger

func mapToZapParams(extraKey map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)
	for key, value := range extraKey {
		params = append(params, string(key), value)
	}
	return params
}

func mapToZeroParams(extraKey map[ExtraKey]interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	for key, value := range extraKey {
		params[string(key)] = value
	}
	return params
}
