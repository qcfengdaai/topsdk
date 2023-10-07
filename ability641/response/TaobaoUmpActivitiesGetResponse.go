package response

import (
)

type TaobaoUmpActivitiesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销活动内容，可以通过ump sdk来进行处理
    */
    Contents  []string `json:"contents,omitempty" `
    /*
        记录总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
