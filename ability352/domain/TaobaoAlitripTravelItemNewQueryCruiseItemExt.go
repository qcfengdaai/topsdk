package domain


type TaobaoAlitripTravelItemNewQueryCruiseItemExt struct {
    /*
        邮轮费用包含     */
    ShipFeeInclude  *string `json:"ship_fee_include,omitempty" `

    /*
        邮轮船名     */
    ShipName  *string `json:"ship_name,omitempty" `

    /*
        邮轮公司     */
    CruiseCompany  *string `json:"cruise_company,omitempty" `

    /*
        邮轮线路     */
    CruiseLine  *string `json:"cruise_line,omitempty" `

    /*
        下船地点     */
    ShipDown  *string `json:"ship_down,omitempty" `

    /*
        上船地点     */
    ShipUp  *string `json:"ship_up,omitempty" `

    /*
        邮轮数据版本     */
    CruiseItemVersion  *string `json:"cruise_item_version,omitempty" `

    /*
        邮轮具体航线     */
    CruiseSubLine  *string `json:"cruise_sub_line,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetShipFeeInclude(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.ShipFeeInclude = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetShipName(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.ShipName = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetCruiseCompany(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.CruiseCompany = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetCruiseLine(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.CruiseLine = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetShipDown(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.ShipDown = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetShipUp(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.ShipUp = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetCruiseItemVersion(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.CruiseItemVersion = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryCruiseItemExt) SetCruiseSubLine(v string) *TaobaoAlitripTravelItemNewQueryCruiseItemExt {
    s.CruiseSubLine = &v
    return s
}
