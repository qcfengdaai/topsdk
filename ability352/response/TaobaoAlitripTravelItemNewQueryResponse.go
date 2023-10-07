package response

import (
    "topsdk/ability352/domain"
)

type TaobaoAlitripTravelItemNewQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品查询结果
    */
    TravelItem  domain.TaobaoAlitripTravelItemNewQueryFullTravelItem `json:"travel_item,omitempty" `
}
