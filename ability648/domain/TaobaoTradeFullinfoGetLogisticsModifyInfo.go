package domain


type TaobaoTradeFullinfoGetLogisticsModifyInfo struct {
    /*
        修改关联的订单好     */
    RelatedId  *string `json:"related_id,omitempty" `

    /*
        修改后的发货时效     */
    ConsignTime  *string `json:"consign_time,omitempty" `

    /*
        修改前的发货时效     */
    OriginConsignTime  *string `json:"origin_consign_time,omitempty" `

    /*
        修改时间     */
    ModifyTime  *string `json:"modify_time,omitempty" `

    /*
        修改原因     */
    ModifyReason  *string `json:"modify_reason,omitempty" `

    /*
        成分品的商品id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        成分品的skuId     */
    SkuId  *string `json:"sku_id,omitempty" `

    /*
        成分品的组合id     */
    CombId  *string `json:"comb_id,omitempty" `

}

func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetRelatedId(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.RelatedId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetConsignTime(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetOriginConsignTime(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.OriginConsignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetModifyTime(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.ModifyTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetModifyReason(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.ModifyReason = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetItemId(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.ItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetSkuId(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsModifyInfo) SetCombId(v string) *TaobaoTradeFullinfoGetLogisticsModifyInfo {
    s.CombId = &v
    return s
}
