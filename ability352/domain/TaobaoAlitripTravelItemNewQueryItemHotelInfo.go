package domain


type TaobaoAlitripTravelItemNewQueryItemHotelInfo struct {
    /*
        结构化酒店信息 酒店结构化信息和文本描述二选一     */
    HotelList  *[]TaobaoAlitripTravelItemNewQueryItemHotel `json:"hotel_list,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryItemHotelInfo) SetHotelList(v []TaobaoAlitripTravelItemNewQueryItemHotel) *TaobaoAlitripTravelItemNewQueryItemHotelInfo {
    s.HotelList = &v
    return s
}
