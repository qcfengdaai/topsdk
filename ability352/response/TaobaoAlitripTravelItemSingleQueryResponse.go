package response

import (
    "topsdk/ability352/domain"
)

type TaobaoAlitripTravelItemSingleQueryResponse struct {

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
    TravelItem  domain.TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem `json:"travel_item,omitempty" `
}
