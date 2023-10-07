package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt struct {
    /*
        回程交通信息     */
    BackTrafficInfo  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty" `

    /*
        去程交通信息     */
    GoTrafficInfo  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty" `

    /*
        酒店信息     */
    HotelInfos  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo `json:"hotel_infos,omitempty" `

    /*
        其他资源信息     */
    OtherInfos  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelItemResourceInfo `json:"other_infos,omitempty" `

    /*
        景点信息     */
    ScenicInfos  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo `json:"scenic_infos,omitempty" `

    /*
        自由行交通信息详细描述     */
    TrafficDesc  *string `json:"traffic_desc,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt) SetBackTrafficInfo(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt {
    s.BackTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt) SetGoTrafficInfo(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt {
    s.GoTrafficInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt) SetHotelInfos(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt {
    s.HotelInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt) SetOtherInfos(v []TaobaoAlitripTravelItemSingleQueryPontusTravelItemResourceInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt {
    s.OtherInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt) SetScenicInfos(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt {
    s.ScenicInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt) SetTrafficDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt {
    s.TrafficDesc = &v
    return s
}
