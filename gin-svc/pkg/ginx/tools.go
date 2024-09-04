package ginx

import (
	"fmt"
	"strings"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		fmt.Println(err)
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

//func HandleValidatorError(ctx *gin.Context, err error) {
//	errs, ok := err.(validator.ValidationErrors)
//	if !ok {
//		JsonErrorResp(ctx, err)
//		return
//	}
//	JsonErrorInterfaceResp(ctx, removeTopStruct(errs.Translate(validators.InstanceTrans)))
//}
