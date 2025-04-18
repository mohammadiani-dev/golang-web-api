package helper

import "golang-web-api/api/validations"

type BaseHttpResponse struct {
	Result           any    `json:"result"`
	Success          bool   `json:"success"`
	ResultCode       int    `json:"resultCode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any    `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int ) BaseHttpResponse {
	return BaseHttpResponse{
		Result: result,
		Success: success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) BaseHttpResponse {
	return BaseHttpResponse{
		Result: result,
		Success: success,
		ResultCode: resultCode,
		Error: err.Error(),
	}
}

func GenerateBaseResponseWithValidationErrors(result any, success bool, resultCode int, err error) BaseHttpResponse {
	return BaseHttpResponse{
		Result: result,
		Success: success,
		ResultCode: resultCode,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}


