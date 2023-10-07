package domain


type TaobaoTradeFullinfoGetCombineLogisticsDetail struct {
    /*
        运单号     */
    InvoiceNo  *string `json:"invoice_no,omitempty" `

    /*
        物流公司     */
    LogisticsCompany  *string `json:"logistics_company,omitempty" `

    /*
        子订单id     */
    SubOrderId  *string `json:"sub_order_id,omitempty" `

    /*
        包裹详情（仅支持成分品）     */
    SendGoodsDetails  *[]TaobaoTradeFullinfoGetSendGoodsDetail `json:"send_goods_details,omitempty" `

}

func (s *TaobaoTradeFullinfoGetCombineLogisticsDetail) SetInvoiceNo(v string) *TaobaoTradeFullinfoGetCombineLogisticsDetail {
    s.InvoiceNo = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineLogisticsDetail) SetLogisticsCompany(v string) *TaobaoTradeFullinfoGetCombineLogisticsDetail {
    s.LogisticsCompany = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineLogisticsDetail) SetSubOrderId(v string) *TaobaoTradeFullinfoGetCombineLogisticsDetail {
    s.SubOrderId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineLogisticsDetail) SetSendGoodsDetails(v []TaobaoTradeFullinfoGetSendGoodsDetail) *TaobaoTradeFullinfoGetCombineLogisticsDetail {
    s.SendGoodsDetails = &v
    return s
}
