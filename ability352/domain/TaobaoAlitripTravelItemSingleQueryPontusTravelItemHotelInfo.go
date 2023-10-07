package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo struct {
    /*
        结构化酒店信息 酒店结构化信息和文本描述二选一     */
    HotelList  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel `json:"hotel_list,omitempty" `

    /*
        必填，默认为0 必须大于等于0, 且小于或等于行程晚数     */
    HotelDays  *int64 `json:"hotel_days,omitempty" `

    /*
        酒店描述文本，&lt;1500字符 酒店结构化信息和文本描述二选一     */
    HotelDesc  *string `json:"hotel_desc,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo) SetHotelList(v []TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotel) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo {
    s.HotelList = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo) SetHotelDays(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo {
    s.HotelDays = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo) SetHotelDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemHotelInfo {
    s.HotelDesc = &v
    return s
}
