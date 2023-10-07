package domain


type TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo struct {
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
        必填，套餐操作：1-增加一个套餐，2-删除一个套餐（根据outer_sku_id定位该套餐），3-覆盖修改一个套餐（根据outer_sku_id定位该套餐）     */
    PackageOperation  *int64 `json:"package_operation,omitempty" `

    /*
        商品日历价格库存相关信息     */
    Prices  *[]TaobaoAlitripTravelItemSkuPackageModifyPrices `json:"prices,omitempty" `

    /*
        套餐关联的产品元素列表。以列表中第一个产品元素作为主元素，其他元素作为该套餐的搭配元素。注：新增一个套餐时该参数必填；如果只是要覆盖修改一个已存在套餐的价格库存，则该参数可不填，系统会根据outer_sku_id自动填充该套餐已绑定的产品元素。     */
    Products  *[]TaobaoAlitripTravelItemSkuPackageModifyProduct `json:"products,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo) SetOuterSkuId(v string) *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo {
    s.OuterSkuId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo) SetPackageDesc(v string) *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo {
    s.PackageDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo) SetPackageName(v string) *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo {
    s.PackageName = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo) SetPackageOperation(v int64) *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo {
    s.PackageOperation = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo) SetPrices(v []TaobaoAlitripTravelItemSkuPackageModifyPrices) *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo {
    s.Prices = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo) SetProducts(v []TaobaoAlitripTravelItemSkuPackageModifyProduct) *TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo {
    s.Products = &v
    return s
}
