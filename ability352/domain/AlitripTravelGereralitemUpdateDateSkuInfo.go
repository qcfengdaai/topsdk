package domain


type AlitripTravelGereralitemUpdateDateSkuInfo struct {
    /*
        sku销售属性别名；如套餐1  需要调整成其他  需要在这里修改     */
    Alias  *[]AlitripTravelGereralitemUpdatePropertyAliasInfo `json:"alias,omitempty" `

    /*
        SKU的销售价格库存，日历商品使用     */
    DateList  *[]AlitripTravelGereralitemUpdateDateInventoryAndPrice `json:"date_list,omitempty" `

    /*
        sku商品编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        SKU销售属性列表,若未填写,其他sku信息不会生效；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。     */
    Properties  *[]AlitripTravelGereralitemUpdateCatPropInfo `json:"properties,omitempty" `

}

func (s *AlitripTravelGereralitemUpdateDateSkuInfo) SetAlias(v []AlitripTravelGereralitemUpdatePropertyAliasInfo) *AlitripTravelGereralitemUpdateDateSkuInfo {
    s.Alias = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateDateSkuInfo) SetDateList(v []AlitripTravelGereralitemUpdateDateInventoryAndPrice) *AlitripTravelGereralitemUpdateDateSkuInfo {
    s.DateList = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateDateSkuInfo) SetOuterId(v string) *AlitripTravelGereralitemUpdateDateSkuInfo {
    s.OuterId = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateDateSkuInfo) SetProperties(v []AlitripTravelGereralitemUpdateCatPropInfo) *AlitripTravelGereralitemUpdateDateSkuInfo {
    s.Properties = &v
    return s
}
