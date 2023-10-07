package domain


type TaobaoTradesSoldGetLogisticsAgreement struct {
    /*
        服务承诺/异常文案     */
    LogisticsServiceMsg  *string `json:"logistics_service_msg,omitempty" `

    /*
        物流服务业务身份     */
    AsdpBizType  *string `json:"asdp_biz_type,omitempty" `

    /*
        物流信息，多个值时用英文逗号隔开     */
    AsdpAds  *string `json:"asdp_ads,omitempty" `

    /*
        计划送达时间     */
    SignTime  *string `json:"sign_time,omitempty" `

    /*
        承诺/最晚送达时间     */
    PromiseSignTime  *string `json:"promise_sign_time,omitempty" `

}

func (s *TaobaoTradesSoldGetLogisticsAgreement) SetLogisticsServiceMsg(v string) *TaobaoTradesSoldGetLogisticsAgreement {
    s.LogisticsServiceMsg = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsAgreement) SetAsdpBizType(v string) *TaobaoTradesSoldGetLogisticsAgreement {
    s.AsdpBizType = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsAgreement) SetAsdpAds(v string) *TaobaoTradesSoldGetLogisticsAgreement {
    s.AsdpAds = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsAgreement) SetSignTime(v string) *TaobaoTradesSoldGetLogisticsAgreement {
    s.SignTime = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsAgreement) SetPromiseSignTime(v string) *TaobaoTradesSoldGetLogisticsAgreement {
    s.PromiseSignTime = &v
    return s
}
