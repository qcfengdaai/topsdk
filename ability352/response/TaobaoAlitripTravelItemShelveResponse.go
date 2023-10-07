package response

import (
)

type TaobaoAlitripTravelItemShelveResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品上下架操作是否成功
    */
    Result  bool `json:"result,omitempty" `
}
