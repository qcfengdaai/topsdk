package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo struct {
    /*
        映射到具体日历价格套餐的外部商家编码     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

    /*
        套餐描述     */
    PackageDesc  *string `json:"package_desc,omitempty" `

    /*
        套餐名称     */
    PackageName  *string `json:"package_name,omitempty" `

    /*
        套餐的日历价格库存。如果是预约商品，只需要填写一个Price，并且，不需要填写Price中的date字段不填，且预约商品只有成人价格和库存。     */
    Prices  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelPrices `json:"prices,omitempty" `

    /*
        套餐关联的产品元素信息     */
    Products  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelProduct `json:"products,omitempty" `

    /*
        套餐下面对应商品元素信息 仅针对新版商品     */
    Combos  *string `json:"combos,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo) SetOuterSkuId(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo) SetPackageDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo {
    s.PackageDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo) SetPackageName(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo {
    s.PackageName = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo) SetPrices(v []TaobaoAlitripTravelItemSingleQueryPontusTravelPrices) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo {
    s.Prices = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo) SetProducts(v []TaobaoAlitripTravelItemSingleQueryPontusTravelProduct) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo {
    s.Products = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo) SetCombos(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo {
    s.Combos = &v
    return s
}
