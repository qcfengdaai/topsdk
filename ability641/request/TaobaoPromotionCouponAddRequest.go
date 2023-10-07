package request

import (
        "topsdk/util"
    )

type TaobaoPromotionCouponAddRequest struct {
    /*
        优惠券的面额，必须是3，5，10，20，50，100     */
    Denominations  *int64 `json:"denominations" required:"true" `
    /*
        优惠券的生效时间     */
    StartTime  *util.LocalTime `json:"start_time,omitempty" required:"false" `
    /*
        优惠券的截止日期     */
    EndTime  *util.LocalTime `json:"end_time" required:"true" `
    /*
        订单满多少元才能用这个优惠券，500就是满500元才能使用     */
    Condition  *int64 `json:"condition,omitempty" required:"false" `
}

func (s *TaobaoPromotionCouponAddRequest) SetDenominations(v int64) *TaobaoPromotionCouponAddRequest {
    s.Denominations = &v
    return s
}
func (s *TaobaoPromotionCouponAddRequest) SetStartTime(v util.LocalTime) *TaobaoPromotionCouponAddRequest {
    s.StartTime = &v
    return s
}
func (s *TaobaoPromotionCouponAddRequest) SetEndTime(v util.LocalTime) *TaobaoPromotionCouponAddRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoPromotionCouponAddRequest) SetCondition(v int64) *TaobaoPromotionCouponAddRequest {
    s.Condition = &v
    return s
}

func (req *TaobaoPromotionCouponAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Denominations != nil) {
        paramMap["denominations"] = *req.Denominations
    }
    if(req.StartTime != nil) {
        paramMap["start_time"] = *req.StartTime
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.Condition != nil) {
        paramMap["condition"] = *req.Condition
    }
    return paramMap
}

func (req *TaobaoPromotionCouponAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}