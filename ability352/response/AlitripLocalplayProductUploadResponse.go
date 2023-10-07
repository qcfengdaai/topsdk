package response

import (
    "topsdk/ability352/domain"
)

type AlitripLocalplayProductUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品发布/更新结果
    */
    FirstResult  domain.AlitripLocalplayProductUploadTopTravelItem `json:"first_result,omitempty" `
}
