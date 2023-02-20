import axios from '@/utils/http'

namespace UserInfo {

    export interface UserInfoReq {}

    export interface UserInfoResp {
        name: string
        age: number
        url: string
        desc: string
    }
}
// 用户登录
export const userInfo = (params: UserInfo.UserInfoReq) => {
    // 返回的数据格式可以和服务端约定
    return axios.post<UserInfo.UserInfoResp>('./data.json', params);
}