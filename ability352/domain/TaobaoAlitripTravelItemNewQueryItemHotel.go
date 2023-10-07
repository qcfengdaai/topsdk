package domain


type TaobaoAlitripTravelItemNewQueryItemHotel struct {
    /*
        关联的套餐id     */
    RelatedPackageId  *int64 `json:"related_package_id,omitempty" `

    /*
        晚数     */
    HotelDays  *int64 `json:"hotel_days,omitempty" `

    /*
        酒店描述     */
    HotelDesc  *string `json:"hotel_desc,omitempty" `

    /*
        POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE     */
    PoiResource  *string `json:"poi_resource,omitempty" `

    /*
        酒店的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222     */
    Poi  *string `json:"poi,omitempty" `

    /*
        必填，酒店房型     */
    HouseType  *string `json:"house_type,omitempty" `

    /*
        必填，酒店等级     */
    HotelLevel  *string `json:"hotel_level,omitempty" `

    /*
        必填，酒店名称     */
    CnName  *string `json:"cn_name,omitempty" `

    /*
        必填，所在城市     */
    City  *string `json:"city,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetRelatedPackageId(v int64) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.RelatedPackageId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetHotelDays(v int64) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.HotelDays = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetHotelDesc(v string) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.HotelDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetPoiResource(v string) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.PoiResource = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetPoi(v string) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.Poi = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetHouseType(v string) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.HouseType = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetHotelLevel(v string) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.HotelLevel = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetCnName(v string) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.CnName = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemHotel) SetCity(v string) *TaobaoAlitripTravelItemNewQueryItemHotel {
    s.City = &v
    return s
}
