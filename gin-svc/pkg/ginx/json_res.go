package ginx

func Success() JsonResult {
	return JsonResult{
		Code: 0,
		Msg:  "OK",
		Data: nil,
	}
}

func SuccessItemList(data any) JsonResult {
	return JsonResult{
		Code: 0,
		Msg:  "OK",
		Data: data,
	}
}

func SuccessPageList(data any, total int) JsonResult {
	return JsonResult{
		Code: 0,
		Msg:  "OK",
		Data: map[string]interface{}{
			"list":  data,
			"total": total,
		},
	}
}

func Error(code int, msg string) JsonResult {
	return JsonResult{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
