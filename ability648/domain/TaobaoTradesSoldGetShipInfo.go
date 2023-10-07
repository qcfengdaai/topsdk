package domain


type TaobaoTradesSoldGetShipInfo struct {
    /*
        发货方式（小写）     */
    TransportType  *string `json:"transport_type,omitempty" `

}

func (s *TaobaoTradesSoldGetShipInfo) SetTransportType(v string) *TaobaoTradesSoldGetShipInfo {
    s.TransportType = &v
    return s
}
