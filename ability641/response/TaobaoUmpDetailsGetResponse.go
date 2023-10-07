package response

import (
)

type TaobaoUmpDetailsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        活动详情的信息
    */
    Contents  []string `json:"contents,omitempty" `
    /*
        记录总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
