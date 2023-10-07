package domain


type TaobaoAlitripTravelItemNewQueryProduct struct {
    /*
        关联的套餐id     */
    PackageId  *int64 `json:"package_id,omitempty" `

    /*
        是否主元素     */
    MainProduct  *bool `json:"main_product,omitempty" `

    /*
        描述     */
    Descr  *string `json:"descr,omitempty" `

    /*
        数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        资源元素的外部商家编码     */
    ElementId  *string `json:"element_id,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryProduct) SetPackageId(v int64) *TaobaoAlitripTravelItemNewQueryProduct {
    s.PackageId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryProduct) SetMainProduct(v bool) *TaobaoAlitripTravelItemNewQueryProduct {
    s.MainProduct = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryProduct) SetDescr(v string) *TaobaoAlitripTravelItemNewQueryProduct {
    s.Descr = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryProduct) SetNum(v int64) *TaobaoAlitripTravelItemNewQueryProduct {
    s.Num = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryProduct) SetElementId(v string) *TaobaoAlitripTravelItemNewQueryProduct {
    s.ElementId = &v
    return s
}
