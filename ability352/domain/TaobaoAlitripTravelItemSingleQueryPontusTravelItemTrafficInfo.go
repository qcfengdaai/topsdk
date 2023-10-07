package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo struct {
    /*
        必填，交通类型。1/2/3/4 分别对应 飞机/火车/汽车/船     */
    TrafficType  *int64 `json:"traffic_type,omitempty" `

    /*
        详细交通信息结构。【注意】当traffic_type选择飞机或火车时，该字段为必填，汽车或轮船时该字段不用填。     */
    Traffics  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic `json:"traffics,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo) SetTrafficType(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo {
    s.TrafficType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo) SetTraffics(v []TaobaoAlitripTravelItemSingleQueryPontusTravelItemTraffic) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemTrafficInfo {
    s.Traffics = &v
    return s
}
