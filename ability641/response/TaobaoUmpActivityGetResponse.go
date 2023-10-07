package response

import (
)

type TaobaoUmpActivityGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销活动的内容，可以通过ump sdk中的marketingTool来完成对该内容的处理
    */
    Content  string `json:"content,omitempty" `
}
