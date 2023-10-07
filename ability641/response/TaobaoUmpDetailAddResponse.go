package response

import (
)

type TaobaoUmpDetailAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        活动详情的id
    */
    DetailId  int64 `json:"detail_id,omitempty" `
}
