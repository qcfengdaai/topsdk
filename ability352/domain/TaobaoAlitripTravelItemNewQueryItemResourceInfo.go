package domain


type TaobaoAlitripTravelItemNewQueryItemResourceInfo struct {
    /*
        关联的套餐id     */
    RelatedPackageId  *int64 `json:"related_package_id,omitempty" `

    /*
        1-保险2-餐饮3-租车4-签证5-购物点6-赠品7-券99-其他     */
    Type  *int64 `json:"type,omitempty" `

    /*
        小于1500字符     */
    Desc  *string `json:"desc,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryItemResourceInfo) SetRelatedPackageId(v int64) *TaobaoAlitripTravelItemNewQueryItemResourceInfo {
    s.RelatedPackageId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemResourceInfo) SetType(v int64) *TaobaoAlitripTravelItemNewQueryItemResourceInfo {
    s.Type = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemResourceInfo) SetDesc(v string) *TaobaoAlitripTravelItemNewQueryItemResourceInfo {
    s.Desc = &v
    return s
}
