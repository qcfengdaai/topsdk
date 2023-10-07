package domain


type TaobaoTradeFullinfoGetSendGoodsDetail struct {
    /*
        0普通1组合商品2组套商品3后置赠品4状态推进     */
    Type  *int64 `json:"type,omitempty" `

    /*
        0全部发货1部分发货     */
    ConsignStatus  *int64 `json:"consign_status,omitempty" `

    /*
        数量     */
    Amount  *int64 `json:"amount,omitempty" `

    /*
        包裹详情     */
    GoodsDetails  *[]TaobaoTradeFullinfoGetPackageGoodsDetail `json:"goods_details,omitempty" `

}

func (s *TaobaoTradeFullinfoGetSendGoodsDetail) SetType(v int64) *TaobaoTradeFullinfoGetSendGoodsDetail {
    s.Type = &v
    return s
}
func (s *TaobaoTradeFullinfoGetSendGoodsDetail) SetConsignStatus(v int64) *TaobaoTradeFullinfoGetSendGoodsDetail {
    s.ConsignStatus = &v
    return s
}
func (s *TaobaoTradeFullinfoGetSendGoodsDetail) SetAmount(v int64) *TaobaoTradeFullinfoGetSendGoodsDetail {
    s.Amount = &v
    return s
}
func (s *TaobaoTradeFullinfoGetSendGoodsDetail) SetGoodsDetails(v []TaobaoTradeFullinfoGetPackageGoodsDetail) *TaobaoTradeFullinfoGetSendGoodsDetail {
    s.GoodsDetails = &v
    return s
}
