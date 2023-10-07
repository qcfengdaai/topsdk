package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel struct {
    /*
        必填，所在城市     */
    City  *string `json:"city,omitempty" `

    /*
        必填，酒店名称     */
    CnName  *string `json:"cn_name,omitempty" `

    /*
        必填，酒店等级     */
    HotelLevel  *string `json:"hotel_level,omitempty" `

    /*
        必填，酒店房型     */
    HouseType  *string `json:"house_type,omitempty" `

    /*
        酒店的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222     */
    Poi  *string `json:"poi,omitempty" `

    /*
        POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE     */
    PoiResource  *string `json:"poi_resource,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel) SetCity(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel {
    s.City = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel) SetCnName(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel {
    s.CnName = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel) SetHotelLevel(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel {
    s.HotelLevel = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel) SetHouseType(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel {
    s.HouseType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel) SetPoi(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel {
    s.Poi = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel) SetPoiResource(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel {
    s.PoiResource = &v
    return s
}
