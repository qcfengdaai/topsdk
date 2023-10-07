package domain


type TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo struct {
    /*
        地点名称     */
    Name  *string `json:"name,omitempty" `

    /*
        经纬度     */
    Poi  *string `json:"poi,omitempty" `

    /*
        经纬度来源     */
    PoiResource  *string `json:"poi_resource,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo) SetName(v string) *TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo {
    s.Name = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo) SetPoi(v string) *TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo {
    s.Poi = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo) SetPoiResource(v string) *TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo {
    s.PoiResource = &v
    return s
}
