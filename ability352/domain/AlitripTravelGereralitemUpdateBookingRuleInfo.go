package domain


type AlitripTravelGereralitemUpdateBookingRuleInfo struct {
    /*
        1500个字     */
    RuleDesc  *string `json:"rule_desc,omitempty" `

    /*
        fee_included：费用包含，跟团游必填； fee_excluded：费用不含，所有类目必填； order_info：预定须知； extra_cost：其他费用，预留；     */
    RuleType  *string `json:"rule_type,omitempty" `

}

func (s *AlitripTravelGereralitemUpdateBookingRuleInfo) SetRuleDesc(v string) *AlitripTravelGereralitemUpdateBookingRuleInfo {
    s.RuleDesc = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateBookingRuleInfo) SetRuleType(v string) *AlitripTravelGereralitemUpdateBookingRuleInfo {
    s.RuleType = &v
    return s
}
