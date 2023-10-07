package response

import (
    "topsdk/ability641/domain"
)

type TaobaoUmpRangeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销范围类列表！
    */
    Ranges  []domain.TaobaoUmpRangeGetRange `json:"ranges,omitempty" `
}
