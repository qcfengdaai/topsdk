package response

import (
    "topsdk/ability352/domain"
)

type AlitripFreetourProductUploadResponse struct {

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
    FirstResult  domain.AlitripFreetourProductUploadTopTravelItem `json:"first_result,omitempty" `
}
