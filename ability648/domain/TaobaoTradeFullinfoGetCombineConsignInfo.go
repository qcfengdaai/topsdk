package domain


type TaobaoTradeFullinfoGetCombineConsignInfo struct {
    /*
        套餐组合id     */
    CombId  *string `json:"comb_id,omitempty" `

    /*
        套餐内的商品id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        套餐内商品的skuId     */
    SkuId  *string `json:"sku_id,omitempty" `

    /*
        成分品的预计发货时间     */
    ConsignTime  *string `json:"consign_time,omitempty" `

    /*
        成分品的预计发货时间     */
    RenderConsignTime  *string `json:"render_consign_time,omitempty" `

}

func (s *TaobaoTradeFullinfoGetCombineConsignInfo) SetCombId(v string) *TaobaoTradeFullinfoGetCombineConsignInfo {
    s.CombId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineConsignInfo) SetItemId(v string) *TaobaoTradeFullinfoGetCombineConsignInfo {
    s.ItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineConsignInfo) SetSkuId(v string) *TaobaoTradeFullinfoGetCombineConsignInfo {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineConsignInfo) SetConsignTime(v string) *TaobaoTradeFullinfoGetCombineConsignInfo {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineConsignInfo) SetRenderConsignTime(v string) *TaobaoTradeFullinfoGetCombineConsignInfo {
    s.RenderConsignTime = &v
    return s
}
