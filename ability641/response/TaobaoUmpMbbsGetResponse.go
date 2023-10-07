package response

import (
)

type TaobaoUmpMbbsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销积木块内容列表，内容为json格式的，可以通过ump sdk里面的MBB.fromJson来处理
    */
    Mbbs  []string `json:"mbbs,omitempty" `
}
