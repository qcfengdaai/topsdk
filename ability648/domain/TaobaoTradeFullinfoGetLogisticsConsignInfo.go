package domain


type TaobaoTradeFullinfoGetLogisticsConsignInfo struct {
    /*
        时效关联的订单好     */
    RelatedId  *string `json:"related_id,omitempty" `

    /*
        商家的预计发货时间     */
    ConsignTime  *string `json:"consign_time,omitempty" `

    /*
        商家的预计发货时间     */
    RenderConsignTime  *string `json:"render_consign_time,omitempty" `

    /*
        成分品发货时效     */
    CombineItem  *[]TaobaoTradeFullinfoGetCombineConsignInfo `json:"combine_item,omitempty" `

}

func (s *TaobaoTradeFullinfoGetLogisticsConsignInfo) SetRelatedId(v string) *TaobaoTradeFullinfoGetLogisticsConsignInfo {
    s.RelatedId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsConsignInfo) SetConsignTime(v string) *TaobaoTradeFullinfoGetLogisticsConsignInfo {
    s.ConsignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsConsignInfo) SetRenderConsignTime(v string) *TaobaoTradeFullinfoGetLogisticsConsignInfo {
    s.RenderConsignTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsConsignInfo) SetCombineItem(v []TaobaoTradeFullinfoGetCombineConsignInfo) *TaobaoTradeFullinfoGetLogisticsConsignInfo {
    s.CombineItem = &v
    return s
}
