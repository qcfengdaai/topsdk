package response

import (
)

type TaobaoUmpActivitiesListGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销活动列表！
    */
    Activities  []string `json:"activities,omitempty" `
}
