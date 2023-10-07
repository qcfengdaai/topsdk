package response

import (
)

type TaobaoUmpDetailUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        调用是否成功
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
