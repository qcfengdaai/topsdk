package response

import (
    "topsdk/ability352/domain"
)

type AlitripItemSchemaUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果
    */
    Result  domain.AlitripItemSchemaUpdateTopTravelItem `json:"result,omitempty" `
}
