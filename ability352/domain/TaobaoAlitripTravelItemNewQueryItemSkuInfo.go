package domain


type TaobaoAlitripTravelItemNewQueryItemSkuInfo struct {
    /*
        套餐关联的产品元素信息     */
    Products  *[]TaobaoAlitripTravelItemNewQueryProduct `json:"products,omitempty" `

    /*
        套餐id     */
    PackageId  *int64 `json:"package_id,omitempty" `

    /*
        套餐下面对应商品元素信息 仅针对新版商品     */
    Combos  *string `json:"combos,omitempty" `

    /*
        套餐的日历价格库存。如果是预约商品，只需要填写一个Price，并且，不需要填写Price中的date字段不填，且预约商品只有成人价格和库存。     */
    Prices  *[]TaobaoAlitripTravelItemNewQueryPrices `json:"prices,omitempty" `

    /*
        套餐名称     */
    PackageName  *string `json:"package_name,omitempty" `

    /*
        套餐描述     */
    PackageDesc  *string `json:"package_desc,omitempty" `

    /*
        映射到具体日历价格套餐的外部商家编码     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

    /*
        邮轮房型ID     */
    RoomTypeId  *int64 `json:"room_type_id,omitempty" `

    /*
        邮轮房型类型     */
    RoomType  *int64 `json:"room_type,omitempty" `

    /*
        邮轮房型名称     */
    RoomTypeName  *string `json:"room_type_name,omitempty" `

    /*
        邮轮房型人数     */
    PeopleNumber  *int64 `json:"people_number,omitempty" `

    /*
        邮轮下单是否限制人数和房型人数一致     */
    OrderCountMatch  *bool `json:"order_count_match,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetProducts(v []TaobaoAlitripTravelItemNewQueryProduct) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.Products = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetPackageId(v int64) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.PackageId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetCombos(v string) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.Combos = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetPrices(v []TaobaoAlitripTravelItemNewQueryPrices) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.Prices = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetPackageName(v string) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.PackageName = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetPackageDesc(v string) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.PackageDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetOuterSkuId(v string) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetRoomTypeId(v int64) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.RoomTypeId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetRoomType(v int64) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.RoomType = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetRoomTypeName(v string) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.RoomTypeName = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetPeopleNumber(v int64) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.PeopleNumber = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemSkuInfo) SetOrderCountMatch(v bool) *TaobaoAlitripTravelItemNewQueryItemSkuInfo {
    s.OrderCountMatch = &v
    return s
}
