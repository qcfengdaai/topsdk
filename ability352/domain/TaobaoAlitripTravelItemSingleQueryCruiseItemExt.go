package domain


type TaobaoAlitripTravelItemSingleQueryCruiseItemExt struct {
    /*
        上船地点     */
    ShipUp  *string `json:"ship_up,omitempty" `

    /*
        下船地点     */
    ShipDown  *string `json:"ship_down,omitempty" `

    /*
        邮轮线路     */
    CruiseLine  *string `json:"cruise_line,omitempty" `

    /*
        邮轮公司     */
    CruiseCompany  *string `json:"cruise_company,omitempty" `

    /*
        邮轮船名     */
    ShipName  *string `json:"ship_name,omitempty" `

    /*
        邮轮费用包含     */
    ShipFeeInclude  *string `json:"ship_fee_include,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryCruiseItemExt) SetShipUp(v string) *TaobaoAlitripTravelItemSingleQueryCruiseItemExt {
    s.ShipUp = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryCruiseItemExt) SetShipDown(v string) *TaobaoAlitripTravelItemSingleQueryCruiseItemExt {
    s.ShipDown = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryCruiseItemExt) SetCruiseLine(v string) *TaobaoAlitripTravelItemSingleQueryCruiseItemExt {
    s.CruiseLine = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryCruiseItemExt) SetCruiseCompany(v string) *TaobaoAlitripTravelItemSingleQueryCruiseItemExt {
    s.CruiseCompany = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryCruiseItemExt) SetShipName(v string) *TaobaoAlitripTravelItemSingleQueryCruiseItemExt {
    s.ShipName = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryCruiseItemExt) SetShipFeeInclude(v string) *TaobaoAlitripTravelItemSingleQueryCruiseItemExt {
    s.ShipFeeInclude = &v
    return s
}
