package response

import (
    "topsdk/ability352/domain"
)

type AlitripGrouptourProductUploadResponse struct {

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
    FirstResult  domain.AlitripGrouptourProductUploadTopTravelItem `json:"first_result,omitempty" `
}
