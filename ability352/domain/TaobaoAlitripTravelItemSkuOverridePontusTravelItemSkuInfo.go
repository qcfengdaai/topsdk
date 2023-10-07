package domain


type TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo struct {
    /*
        必填，外部商家映射到具体套餐的SKU ID，一般为外部商家具体商品ID。     */
    OuterSkuId  *string `json:"outer_sku_id,omitempty" `

    /*
        套餐描述     */
    PackageDesc  *string `json:"package_desc,omitempty" `

    /*
        套餐名称，特别说明：出境邮轮请按照发布后台规范传递套餐名称，会自动解析容纳人数与房型     */
    PackageName  *string `json:"package_name,omitempty" `

    /*
        商品日历价格库存相关信息     */
    Prices  *[]TaobaoAlitripTravelItemSkuOverridePontusTravelPrices `json:"prices,omitempty" `

    /*
        套餐关联的产品元素列表。以列表中第一个产品元素作为主元素，其他元素作为该套餐的搭配元素。注：新增一个套餐时该参数必填；如果只是要覆盖修改一个已存在套餐的价格库存，则该参数可不填，系统会根据outer_sku_id自动填充该套餐已绑定的产品元素。     */
    Products  *[]TaobaoAlitripTravelItemSkuOverrideProduct `json:"products,omitempty" `

    /*
        邮轮房型ID，新版本有值     */
    RoomTypeId  *int64 `json:"room_type_id,omitempty" `

    /*
        邮轮房型类型     */
    RoomType  *int64 `json:"room_type,omitempty" `

    /*
        邮轮房型名称     */
    RoomTypeName  *string `json:"room_type_name,omitempty" `

    /*
        人数     */
    PeopleNumber  *int64 `json:"people_number,omitempty" `

    /*
        下单人数是否与房型人数一致     */
    OrderCountMatch  *bool `json:"order_count_match,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetOuterSkuId(v string) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetPackageDesc(v string) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.PackageDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetPackageName(v string) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.PackageName = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetPrices(v []TaobaoAlitripTravelItemSkuOverridePontusTravelPrices) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.Prices = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetProducts(v []TaobaoAlitripTravelItemSkuOverrideProduct) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.Products = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetRoomTypeId(v int64) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.RoomTypeId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetRoomType(v int64) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.RoomType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetRoomTypeName(v string) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.RoomTypeName = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetPeopleNumber(v int64) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.PeopleNumber = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) SetOrderCountMatch(v bool) *TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo {
    s.OrderCountMatch = &v
    return s
}
