package domain


type TaobaoTradeFullinfoGetLogisticsAgreement struct {
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

    /*
        ERP应推单时间(主单)     */
    PushTime  *string `json:"push_time,omitempty" `

    /*
        1：不支持子单部分发货     */
    NoDetailPartConsign  *string `json:"no_detail_part_consign,omitempty" `

    /*
        1：代表该订单是预售下沉订单类型为前置表达     */
    SinkType  *string `json:"sink_type,omitempty" `

}

func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetLogisticsServiceMsg(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.LogisticsServiceMsg = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetAsdpBizType(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.AsdpBizType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetAsdpAds(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.AsdpAds = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetSignTime(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.SignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetPromiseSignTime(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.PromiseSignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetPushTime(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.PushTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetNoDetailPartConsign(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.NoDetailPartConsign = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsAgreement) SetSinkType(v string) *TaobaoTradeFullinfoGetLogisticsAgreement {
    s.SinkType = &v
    return s
}
