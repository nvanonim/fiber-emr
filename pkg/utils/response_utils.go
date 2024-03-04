package utils

// Response Body Struct
type ResponseBody struct {
	ResponseCode    string      `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	Data            interface{} `json:"data"`
}

// Response Code Enum
const (
	RC_Success             = "00"
	RC_DataNotFound        = "01"
	RC_Unauthorized        = "02"
	RC_InvalidRequest      = "03"
	RC_DataAlreadyExist    = "04"
	RC_InternalServerError = "99"
)

// Response Message Enum
const (
	RM_Success = "Success"
)

// GenerateResponse generates a response body. Data is optional
func GenerateResponse(responseCode string, responseMessage string, data ...interface{}) ResponseBody {
	if len(data) > 0 {
		return ResponseBody{
			ResponseCode:    responseCode,
			ResponseMessage: responseMessage,
			Data:            data[0],
		}
	}
	return ResponseBody{
		ResponseCode:    responseCode,
		ResponseMessage: responseMessage,
		Data:            []string{},
	}
}

// GenerateErrorResponse generates an error response body (empty array for data)
func GenerateErrorResponse(responseCode string, responseMessage string) ResponseBody {
	return ResponseBody{
		ResponseCode:    responseCode,
		ResponseMessage: responseMessage,
		Data:            []string{},
	}
}
