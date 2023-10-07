package response

import (
)

type TaobaoUmpDetailGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        活动详情信息，可以通过ump sdk中的MarketingTool来进行处理
    */
    Content  string `json:"content,omitempty" `
}
