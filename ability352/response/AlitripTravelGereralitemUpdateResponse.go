package response

import (
    "topsdk/ability352/domain"
)

type AlitripTravelGereralitemUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品发布结果
    */
    TravelItem  domain.AlitripTravelGereralitemUpdateTopTravelItem `json:"travel_item,omitempty" `
}
