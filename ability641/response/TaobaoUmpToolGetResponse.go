package response

import (
)

type TaobaoUmpToolGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        工具信息内容，格式为json，可以通过提供给的sdk里面的MarketingBuilder来处理这个内容
    */
    Content  string `json:"content,omitempty" `
}
