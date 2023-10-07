package response

import (
)

type TaobaoUmpToolsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        工具列表，单个内容为json格式，需要通过ump的sdk提供的MarketingBuilder来进行处理
    */
    Tools  []string `json:"tools,omitempty" `
}
