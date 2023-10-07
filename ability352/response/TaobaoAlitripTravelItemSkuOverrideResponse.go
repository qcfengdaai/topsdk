package response

import (
    "topsdk/ability352/domain"
)

type TaobaoAlitripTravelItemSkuOverrideResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品sku修改结果
    */
    TravelItem  domain.TaobaoAlitripTravelItemSkuOverrideTopTravelItem `json:"travel_item,omitempty" `
}
