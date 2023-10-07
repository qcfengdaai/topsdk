package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic struct {
    /*
        必填，景点所在城市     */
    City  *string `json:"city,omitempty" `

    /*
        必填，景点名称     */
    CnName  *string `json:"cn_name,omitempty" `

    /*
        景点的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222     */
    Poi  *string `json:"poi,omitempty" `

    /*
        POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE     */
    PoiResource  *string `json:"poi_resource,omitempty" `

    /*
        必填，门票类型     */
    TicketType  *string `json:"ticket_type,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic) SetCity(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic {
    s.City = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic) SetCnName(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic {
    s.CnName = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic) SetPoi(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic {
    s.Poi = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic) SetPoiResource(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic {
    s.PoiResource = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic) SetTicketType(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic {
    s.TicketType = &v
    return s
}
