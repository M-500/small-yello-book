
//我要用到的一些接口
// import service from "../requests/index";
import service from "@/requests/index";
import {ILoginData, ISendCodeData} from "@/types/login";  //引入接口
 
// 登录接口
export function login(data: ILoginData) {//接口ILoginData接口的变量名
    return service({
        url: "/na/login",
        method: "POST",
        data
    })
}

export function SendCode(data:ISendCodeData) {
    return service({
        url: "/na/email/send",
        method: "POST",
        data
    })
}

