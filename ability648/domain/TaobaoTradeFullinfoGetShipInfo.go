package domain


type TaobaoTradeFullinfoGetShipInfo struct {
    /*
        发货方式（小写）     */
    TransportType  *string `json:"transport_type,omitempty" `

}

func (s *TaobaoTradeFullinfoGetShipInfo) SetTransportType(v string) *TaobaoTradeFullinfoGetShipInfo {
    s.TransportType = &v
    return s
}
