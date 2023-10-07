package response

import (
    "topsdk/ability352/domain"
)

type TaobaoAlitripTravelItemSkuPriceModifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        日期级别日历价格库存增量维护
    */
    TravelItem  domain.TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem `json:"travel_item,omitempty" `
}
