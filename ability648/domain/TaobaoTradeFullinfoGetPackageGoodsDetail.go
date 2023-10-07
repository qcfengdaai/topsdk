package domain


type TaobaoTradeFullinfoGetPackageGoodsDetail struct {
    /*
        商品数字编号     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        sku_id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        数量     */
    Amount  *int64 `json:"amount,omitempty" `

}

func (s *TaobaoTradeFullinfoGetPackageGoodsDetail) SetItemId(v int64) *TaobaoTradeFullinfoGetPackageGoodsDetail {
    s.ItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPackageGoodsDetail) SetSkuId(v int64) *TaobaoTradeFullinfoGetPackageGoodsDetail {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetPackageGoodsDetail) SetAmount(v int64) *TaobaoTradeFullinfoGetPackageGoodsDetail {
    s.Amount = &v
    return s
}
