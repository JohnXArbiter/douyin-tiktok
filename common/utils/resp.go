package utils

func GenOkResp() map[string]interface{} {
	return map[string]interface{}{
		"status_code": 0,
		"status_msg":  "ok",
	}
}

func GenErrorResp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status_code": -1,
		"status_msg":  msg,
	}
}
