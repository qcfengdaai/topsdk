package response

import (
)

type TaobaoAlitripTravelBaseinfoScenicsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回详细景点信息，返回数据为json数组结构的字符串
    */
    ScenicInfos  string `json:"scenic_infos,omitempty" `
}
