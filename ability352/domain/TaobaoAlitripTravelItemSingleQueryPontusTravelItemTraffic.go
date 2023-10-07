package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic struct {
    /*
        到达时间，当地时间HH:mm     */
    ArrivalTime  *string `json:"arrival_time,omitempty" `

    /*
        出发城市     */
    Departure  *string `json:"departure,omitempty" `

    /*
        出发时间，当地时间HH:mm     */
    DepartureTime  *string `json:"departure_time,omitempty" `

    /*
        到达城市     */
    Destination  *string `json:"destination,omitempty" `

    /*
        飞机机型，飞机选填     */
    PlaneType  *string `json:"plane_type,omitempty" `

    /*
        参考班次号，飞机需要填航班号，火车需要填车次号，汽车和船可不填     */
    TrafficNo  *string `json:"traffic_no,omitempty" `

    /*
        交通公司名，飞机选填     */
    Vendor  *string `json:"vendor,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) SetArrivalTime(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic {
    s.ArrivalTime = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) SetDeparture(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic {
    s.Departure = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) SetDepartureTime(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic {
    s.DepartureTime = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) SetDestination(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic {
    s.Destination = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) SetPlaneType(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic {
    s.PlaneType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) SetTrafficNo(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic {
    s.TrafficNo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) SetVendor(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic {
    s.Vendor = &v
    return s
}
