package response

import (
    "topsdk/ability352/domain"
)

type AlitripTravelElementsSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回对象
    */
    Result  domain.AlitripTravelElementsSearchResourceData `json:"result,omitempty" `
}
