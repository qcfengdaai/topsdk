package domain


type TaobaoAlitripTravelItemNewQueryGroupItemExt struct {
    /*
        必填，回程交通信息     */
    BackTrafficInfo  *TaobaoAlitripTravelItemNewQueryItemTrafficInfo `json:"back_traffic_info,omitempty" `

    /*
        是否支持电子合同，默认不支持     */
    Electronic  *bool `json:"electronic,omitempty" `

    /*
        集合地信息     */
    GatherPlaces  *[]TaobaoAlitripTravelItemNewQueryGatherPlaceInfo `json:"gather_places,omitempty" `

    /*
        必填，去程交通信息     */
    GoTrafficInfo  *TaobaoAlitripTravelItemNewQueryItemTrafficInfo `json:"go_traffic_info,omitempty" `

    /*
        必填，线路类型，0 为目的地参团 	1为出发地参团     */
    RouteType  *int64 `json:"route_type,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryGroupItemExt) SetBackTrafficInfo(v TaobaoAlitripTravelItemNewQueryItemTrafficInfo) *TaobaoAlitripTravelItemNewQueryGroupItemExt {
    s.BackTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryGroupItemExt) SetElectronic(v bool) *TaobaoAlitripTravelItemNewQueryGroupItemExt {
    s.Electronic = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryGroupItemExt) SetGatherPlaces(v []TaobaoAlitripTravelItemNewQueryGatherPlaceInfo) *TaobaoAlitripTravelItemNewQueryGroupItemExt {
    s.GatherPlaces = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryGroupItemExt) SetGoTrafficInfo(v TaobaoAlitripTravelItemNewQueryItemTrafficInfo) *TaobaoAlitripTravelItemNewQueryGroupItemExt {
    s.GoTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryGroupItemExt) SetRouteType(v int64) *TaobaoAlitripTravelItemNewQueryGroupItemExt {
    s.RouteType = &v
    return s
}
