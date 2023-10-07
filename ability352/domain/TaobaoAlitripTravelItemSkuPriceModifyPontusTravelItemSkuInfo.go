package domain


type TaobaoAlitripTravelItemSkuPriceModifyPontusTravelItemSkuInfo struct {
    /*
        套餐的外部商家编码。注意：如果没有为该商品设置过套餐商家编码，则无法调用本接口更新该套餐的日历价格。请手动去网页端后台编辑下该商品：为该商品的每个套餐添加设置一个商家编码；或者调用sku.override接口给每个套餐设置一个商家编码。     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

    /*
        商品日历价格库存相关信息     */
    Prices  *[]TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices `json:"prices,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelItemSkuInfo) SetOuterSkuId(v string) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelItemSkuInfo {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelItemSkuInfo) SetPrices(v []TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelItemSkuInfo {
    s.Prices = &v
    return s
}
