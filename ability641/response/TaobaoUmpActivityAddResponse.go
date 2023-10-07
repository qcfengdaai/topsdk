package response

import (
)

type TaobaoUmpActivityAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        活动id
    */
    ActId  int64 `json:"act_id,omitempty" `
}
