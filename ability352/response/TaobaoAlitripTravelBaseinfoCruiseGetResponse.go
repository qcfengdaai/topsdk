package response

import (
)

type TaobaoAlitripTravelBaseinfoCruiseGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        邮轮类目扩展信息的json格式字符串
    */
    CruiseExtInfos  string `json:"cruise_ext_infos,omitempty" `
}
