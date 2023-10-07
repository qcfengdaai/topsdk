package response

import (
)

type TaobaoUmpMbbGetbyidResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销积木块定义信息，可以通过ump sdk里面的MBB.fromJson来处理
    */
    Mbb  string `json:"mbb,omitempty" `
}
