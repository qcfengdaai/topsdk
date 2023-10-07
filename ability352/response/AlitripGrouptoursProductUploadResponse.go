package response

import (
    "topsdk/ability352/domain"
)

type AlitripGrouptoursProductUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        firstResult
    */
    FirstResult  domain.AlitripGrouptoursProductUploadTopTravelItem `json:"first_result,omitempty" `
}
