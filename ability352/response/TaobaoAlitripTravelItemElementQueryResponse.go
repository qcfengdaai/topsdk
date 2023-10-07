package response

import (
    "topsdk/ability352/domain"
)

type TaobaoAlitripTravelItemElementQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        资源元素列表
    */
    Results  []domain.TaobaoAlitripTravelItemElementQueryTopElementParam `json:"results,omitempty" `
}
