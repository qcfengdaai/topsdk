package response

import (
)

type TaobaoUmpToolAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        工具id
    */
    ToolId  int64 `json:"tool_id,omitempty" `
}
