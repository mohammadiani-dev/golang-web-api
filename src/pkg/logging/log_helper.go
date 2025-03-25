package logging

func mapTozapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0 , len(extra))
	for k,v := range extra{
		params = append(params, string(k))
		params = append(params, v)
	}

	return params
}

func mapToZeroParams(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := make(map[string]interface{}, len(extra))
	for k,v := range extra{
		params[string(k)] = v
	}

	return params
}

