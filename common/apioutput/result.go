package apioutput


// Result 接口统一返回格式
type Result struct {
	Code    int    `json:"code"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

func ResultFail(mes interface{}) *Result{
	return &Result{
		Code:    -1,
		Data:    nil,
		Message: mes,
	}
}

func ResultSuccess(data interface{}) *Result{
	return &Result{
		Code:    0,
		Data:    data,
		Message: "Success",
	}
}

func TokenInvalid(mes interface{}) *Result{
	return &Result{
		Code:    -2,
		Data:    nil,
		Message: mes,
	}
}



