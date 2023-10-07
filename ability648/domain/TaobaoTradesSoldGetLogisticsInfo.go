package domain


type TaobaoTradesSoldGetLogisticsInfo struct {
    /*
        订单推荐配送类型      * 0：子单无配建议；ERP按照自己的逻辑进行择配。      * 1：子单有推荐配list，erp可按需参考。      * 2：子单有推荐配list，erp必须在推荐配list中选择配品牌。      * 3：子单有禁用配list，erp需要过滤配品牌。     */
    BizDeliveryType  *int64 `json:"biz_delivery_type,omitempty" `

    /*
        未使用仓建议报错     */
    UnusedWarehouseErrorMsg  *string `json:"unused_warehouse_error_msg,omitempty" `

    /*
        未使用配建议报错     */
    UnusedDeliveryErrorMsg  *string `json:"unused_delivery_error_msg,omitempty" `

    /*
        使用禁止配报错     */
    UsedBlackDeliveryErrorMsg  *string `json:"used_black_delivery_error_msg,omitempty" `

    /*
        承诺/最晚出库时间     */
    PromiseOutboundTime  *string `json:"promise_outbound_time,omitempty" `

    /*
        主交易号     */
    TradeId  *int64 `json:"trade_id,omitempty" `

    /*
        子交易号     */
    SubTradeId  *int64 `json:"sub_trade_id,omitempty" `

    /*
        承诺/最晚揽收时间     */
    PromiseCollectTime  *string `json:"promise_collect_time,omitempty" `

}

func (s *TaobaoTradesSoldGetLogisticsInfo) SetBizDeliveryType(v int64) *TaobaoTradesSoldGetLogisticsInfo {
    s.BizDeliveryType = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsInfo) SetUnusedWarehouseErrorMsg(v string) *TaobaoTradesSoldGetLogisticsInfo {
    s.UnusedWarehouseErrorMsg = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsInfo) SetUnusedDeliveryErrorMsg(v string) *TaobaoTradesSoldGetLogisticsInfo {
    s.UnusedDeliveryErrorMsg = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsInfo) SetUsedBlackDeliveryErrorMsg(v string) *TaobaoTradesSoldGetLogisticsInfo {
    s.UsedBlackDeliveryErrorMsg = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsInfo) SetPromiseOutboundTime(v string) *TaobaoTradesSoldGetLogisticsInfo {
    s.PromiseOutboundTime = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsInfo) SetTradeId(v int64) *TaobaoTradesSoldGetLogisticsInfo {
    s.TradeId = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsInfo) SetSubTradeId(v int64) *TaobaoTradesSoldGetLogisticsInfo {
    s.SubTradeId = &v
    return s
}
func (s *TaobaoTradesSoldGetLogisticsInfo) SetPromiseCollectTime(v string) *TaobaoTradesSoldGetLogisticsInfo {
    s.PromiseCollectTime = &v
    return s
}
