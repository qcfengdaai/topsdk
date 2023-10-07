package domain


type TaobaoTradeFullinfoGetIdentifyLogisticsInfo struct {
    /*
        子订单号     */
    DetailOrderId  *int64 `json:"detail_order_id,omitempty" `

    /*
        物流公司     */
    LogisticsCompany  *string `json:"logistics_company,omitempty" `

    /*
        阶段号     */
    StageNo  *int64 `json:"stage_no,omitempty" `

    /*
        运单号     */
    InvoiceNo  *string `json:"invoice_no,omitempty" `

    /*
        退款单号     */
    RefundId  *string `json:"refund_id,omitempty" `

    /*
        是否退货     */
    Refund  *bool `json:"refund,omitempty" `

}

func (s *TaobaoTradeFullinfoGetIdentifyLogisticsInfo) SetDetailOrderId(v int64) *TaobaoTradeFullinfoGetIdentifyLogisticsInfo {
    s.DetailOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyLogisticsInfo) SetLogisticsCompany(v string) *TaobaoTradeFullinfoGetIdentifyLogisticsInfo {
    s.LogisticsCompany = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyLogisticsInfo) SetStageNo(v int64) *TaobaoTradeFullinfoGetIdentifyLogisticsInfo {
    s.StageNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyLogisticsInfo) SetInvoiceNo(v string) *TaobaoTradeFullinfoGetIdentifyLogisticsInfo {
    s.InvoiceNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyLogisticsInfo) SetRefundId(v string) *TaobaoTradeFullinfoGetIdentifyLogisticsInfo {
    s.RefundId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyLogisticsInfo) SetRefund(v bool) *TaobaoTradeFullinfoGetIdentifyLogisticsInfo {
    s.Refund = &v
    return s
}
