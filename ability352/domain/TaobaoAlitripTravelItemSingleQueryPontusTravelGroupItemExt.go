package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt struct {
    /*
        必填，回程交通信息     */
    BackTrafficInfo  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty" `

    /*
        是否支持电子合同，默认不支持     */
    Electronic  *bool `json:"electronic,omitempty" `

    /*
        集合地信息     */
    GatherPlaces  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelGatherPlaceInfo `json:"gather_places,omitempty" `

    /*
        必填，去程交通信息     */
    GoTrafficInfo  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty" `

    /*
        必填，线路类型，0 为目的地参团 	1为出发地参团     */
    RouteType  *int64 `json:"route_type,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt) SetBackTrafficInfo(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt {
    s.BackTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt) SetElectronic(v bool) *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt {
    s.Electronic = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt) SetGatherPlaces(v []TaobaoAlitripTravelItemSingleQueryPontusTravelGatherPlaceInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt {
    s.GatherPlaces = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt) SetGoTrafficInfo(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt {
    s.GoTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt) SetRouteType(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt {
    s.RouteType = &v
    return s
}
