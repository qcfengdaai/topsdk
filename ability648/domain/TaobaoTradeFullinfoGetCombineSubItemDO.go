package domain


type TaobaoTradeFullinfoGetCombineSubItemDO struct {
    /*
        商品数字编号     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品的最小库存单位Sku的id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        商品名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        商家外部编码(可与商家外部系统对接)。     */
    OuterIid  *string `json:"outer_iid,omitempty" `

    /*
        SKU标题     */
    SkuTitle  *string `json:"sku_title,omitempty" `

    /*
        成分品原价     */
    OriginFee  *int64 `json:"origin_fee,omitempty" `

    /*
        成分品套餐原价     */
    CombineSubItemFee  *int64 `json:"combine_sub_item_fee,omitempty" `

    /*
        套餐购是否商品主子成分品       */
    Ismain  *bool `json:"ismain,omitempty" `

    /*
        套餐购成分品预计承诺时效(如果为时间格式为yyyy-MM-dd 则为绝对时间，为数字则为相对时间，比如7天内发货)     */
    EstconTime  *string `json:"estcon_time,omitempty" `

    /*
        outer_sku_id     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

    /*
        教育优惠原价     */
    EduOriginalFee  *int64 `json:"edu_original_fee,omitempty" `

}

func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetItemId(v int64) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.ItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetSkuId(v int64) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetQuantity(v int64) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.Quantity = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetItemName(v string) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.ItemName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetOuterIid(v string) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.OuterIid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetSkuTitle(v string) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.SkuTitle = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetOriginFee(v int64) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.OriginFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetCombineSubItemFee(v int64) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.CombineSubItemFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetIsmain(v bool) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.Ismain = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetEstconTime(v string) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.EstconTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetOuterSkuId(v string) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetCombineSubItemDO) SetEduOriginalFee(v int64) *TaobaoTradeFullinfoGetCombineSubItemDO {
    s.EduOriginalFee = &v
    return s
}
