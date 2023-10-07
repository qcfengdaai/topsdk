package domain


type TaobaoAlitripTravelItemSkuPackageModifyProduct struct {
    /*
        关联的元素。以元素的外部商家编码作为元素标示。不传该值则该Product将被自动忽略     */
    ElementId  *string `json:"element_id,omitempty" `

    /*
        绑定销售的元素 份数。不传该值则默认设置1份     */
    Num  *int64 `json:"num,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuPackageModifyProduct) SetElementId(v string) *TaobaoAlitripTravelItemSkuPackageModifyProduct {
    s.ElementId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyProduct) SetNum(v int64) *TaobaoAlitripTravelItemSkuPackageModifyProduct {
    s.Num = &v
    return s
}
