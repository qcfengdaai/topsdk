package response

import (
    "topsdk/ability352/domain"
)

type AlitripTravelGereralskuUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果
    */
    FirstResult  domain.AlitripTravelGereralskuUpdateTopTravelItem `json:"first_result,omitempty" `
}
