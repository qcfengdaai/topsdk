package response

import (
    "topsdk/ability352/domain"
)

type AlitripDaytoursProductUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品维护结果
    */
    FirstResult  domain.AlitripDaytoursProductUploadTopTravelItem `json:"first_result,omitempty" `
}
