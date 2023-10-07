package response

import (
    "topsdk/ability352/domain"
)

type AlitripTravelPoiSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        POI详情
    */
    Results  []domain.AlitripTravelPoiSearchPOI `json:"results,omitempty" `
}
