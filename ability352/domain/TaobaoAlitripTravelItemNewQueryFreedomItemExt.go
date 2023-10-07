package domain


type TaobaoAlitripTravelItemNewQueryFreedomItemExt struct {
    /*
        回程交通信息     */
    BackTrafficInfo  *TaobaoAlitripTravelItemNewQueryItemTrafficInfo `json:"back_traffic_info,omitempty" `

    /*
        去程交通信息     */
    GoTrafficInfo  *TaobaoAlitripTravelItemNewQueryItemTrafficInfo `json:"go_traffic_info,omitempty" `

    /*
        酒店信息     */
    HotelInfos  *TaobaoAlitripTravelItemNewQueryItemHotelInfo `json:"hotel_infos,omitempty" `

    /*
        其他资源信息     */
    OtherInfos  *[]TaobaoAlitripTravelItemNewQueryItemResourceInfo `json:"other_infos,omitempty" `

    /*
        景点信息     */
    ScenicInfos  *TaobaoAlitripTravelItemNewQueryItemScenicInfo `json:"scenic_infos,omitempty" `

    /*
        自由行交通信息详细描述     */
    TrafficDesc  *string `json:"traffic_desc,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryFreedomItemExt) SetBackTrafficInfo(v TaobaoAlitripTravelItemNewQueryItemTrafficInfo) *TaobaoAlitripTravelItemNewQueryFreedomItemExt {
    s.BackTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFreedomItemExt) SetGoTrafficInfo(v TaobaoAlitripTravelItemNewQueryItemTrafficInfo) *TaobaoAlitripTravelItemNewQueryFreedomItemExt {
    s.GoTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFreedomItemExt) SetHotelInfos(v TaobaoAlitripTravelItemNewQueryItemHotelInfo) *TaobaoAlitripTravelItemNewQueryFreedomItemExt {
    s.HotelInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFreedomItemExt) SetOtherInfos(v []TaobaoAlitripTravelItemNewQueryItemResourceInfo) *TaobaoAlitripTravelItemNewQueryFreedomItemExt {
    s.OtherInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFreedomItemExt) SetScenicInfos(v TaobaoAlitripTravelItemNewQueryItemScenicInfo) *TaobaoAlitripTravelItemNewQueryFreedomItemExt {
    s.ScenicInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFreedomItemExt) SetTrafficDesc(v string) *TaobaoAlitripTravelItemNewQueryFreedomItemExt {
    s.TrafficDesc = &v
    return s
}
