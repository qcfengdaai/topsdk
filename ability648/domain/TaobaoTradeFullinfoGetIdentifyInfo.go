package domain


type TaobaoTradeFullinfoGetIdentifyInfo struct {
    /*
        三方鉴定物流相关信息     */
    IdentifyLogisticsInfos  *[]TaobaoTradeFullinfoGetIdentifyLogisticsInfo `json:"identify_logistics_infos,omitempty" `

    /*
        三方鉴定服务相关信息     */
    IdentifyServiceInfos  *[]TaobaoTradeFullinfoGetIdentifyServiceInfo `json:"identify_service_infos,omitempty" `

}

func (s *TaobaoTradeFullinfoGetIdentifyInfo) SetIdentifyLogisticsInfos(v []TaobaoTradeFullinfoGetIdentifyLogisticsInfo) *TaobaoTradeFullinfoGetIdentifyInfo {
    s.IdentifyLogisticsInfos = &v
    return s
}
func (s *TaobaoTradeFullinfoGetIdentifyInfo) SetIdentifyServiceInfos(v []TaobaoTradeFullinfoGetIdentifyServiceInfo) *TaobaoTradeFullinfoGetIdentifyInfo {
    s.IdentifyServiceInfos = &v
    return s
}
