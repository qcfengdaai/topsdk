package response

import (
)

type TaobaoPromotionCouponAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        优惠券的id
    */
    CouponId  int64 `json:"coupon_id,omitempty" `
}
