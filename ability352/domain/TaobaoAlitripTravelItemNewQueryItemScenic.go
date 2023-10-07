package domain


type TaobaoAlitripTravelItemNewQueryItemScenic struct {
    /*
        关联的套餐id     */
    RelatedPackageId  *string `json:"related_package_id,omitempty" `

    /*
        必填，门票类型     */
    TicketType  *string `json:"ticket_type,omitempty" `

    /*
        POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE     */
    PoiResource  *string `json:"poi_resource,omitempty" `

    /*
        景点的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222     */
    Poi  *string `json:"poi,omitempty" `

    /*
        必填，景点名称     */
    CnName  *string `json:"cn_name,omitempty" `

    /*
        必填，景点所在城市     */
    City  *string `json:"city,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryItemScenic) SetRelatedPackageId(v string) *TaobaoAlitripTravelItemNewQueryItemScenic {
    s.RelatedPackageId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemScenic) SetTicketType(v string) *TaobaoAlitripTravelItemNewQueryItemScenic {
    s.TicketType = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemScenic) SetPoiResource(v string) *TaobaoAlitripTravelItemNewQueryItemScenic {
    s.PoiResource = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemScenic) SetPoi(v string) *TaobaoAlitripTravelItemNewQueryItemScenic {
    s.Poi = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemScenic) SetCnName(v string) *TaobaoAlitripTravelItemNewQueryItemScenic {
    s.CnName = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemScenic) SetCity(v string) *TaobaoAlitripTravelItemNewQueryItemScenic {
    s.City = &v
    return s
}
