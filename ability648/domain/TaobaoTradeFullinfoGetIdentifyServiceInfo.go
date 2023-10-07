package domain


type TaobaoTradeFullinfoGetIdentifyServiceInfo struct {
    /*
        子订单号     */
    DetailOrderId  *int64 `json:"detail_order_id,omitempty" `

    /*
        标的物     */
    UnitId  *string `json:"unit_id,omitempty" `

    /*
        服务单号     */
    ServiceId  *string `json:"service_id,omitempty" `

}

func (s *TaobaoTradeFullinfoGetIdentifyServiceInfo) SetDetailOrderId(v int64) *TaobaoTradeFullinfoGetIdentifyServiceInfo {
    s.DetailOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyServiceInfo) SetUnitId(v string) *TaobaoTradeFullinfoGetIdentifyServiceInfo {
    s.UnitId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyServiceInfo) SetServiceId(v string) *TaobaoTradeFullinfoGetIdentifyServiceInfo {
    s.ServiceId = &v
    return s
}
