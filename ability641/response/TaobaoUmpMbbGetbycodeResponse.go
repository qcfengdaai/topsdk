package response

import (
)

type TaobaoUmpMbbGetbycodeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销积木块的内容，通过ump sdk来进行处理
    */
    Mbb  string `json:"mbb,omitempty" `
}
