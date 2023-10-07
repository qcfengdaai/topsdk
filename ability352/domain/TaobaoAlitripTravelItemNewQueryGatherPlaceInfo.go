package domain


type TaobaoAlitripTravelItemNewQueryGatherPlaceInfo struct {
    /*
        POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE     */
    PoiResource  *string `json:"poi_resource,omitempty" `

    /*
        集合地经纬度，英文逗号分隔。经度在前，纬度在后，英文逗号分隔最多支持6位小数，如120.111222,30.111222     */
    Poi  *string `json:"poi,omitempty" `

    /*
        地点名称     */
    Name  *string `json:"name,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryGatherPlaceInfo) SetPoiResource(v string) *TaobaoAlitripTravelItemNewQueryGatherPlaceInfo {
    s.PoiResource = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryGatherPlaceInfo) SetPoi(v string) *TaobaoAlitripTravelItemNewQueryGatherPlaceInfo {
    s.Poi = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryGatherPlaceInfo) SetName(v string) *TaobaoAlitripTravelItemNewQueryGatherPlaceInfo {
    s.Name = &v
    return s
}
