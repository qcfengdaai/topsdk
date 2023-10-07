package domain


type TaobaoAlitripTravelItemNewQueryBookingRuleInfo struct {
    /*
        规则说明，1500个字符     */
    RuleList  *[]string `json:"rule_list,omitempty" `

    /*
        Fee_Included：费用包含，跟团游必填； Fee_Excluded：费用不含，所有类目必填； Order_Info：预定须知； Extra_Cost：其他费用，预留；     Shopping：购物说明，预留     */
    RuleType  *string `json:"rule_type,omitempty" `

    /*
        规则说明，1500个字符     */
    RuleDesc  *string `json:"rule_desc,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryBookingRuleInfo) SetRuleList(v []string) *TaobaoAlitripTravelItemNewQueryBookingRuleInfo {
    s.RuleList = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryBookingRuleInfo) SetRuleType(v string) *TaobaoAlitripTravelItemNewQueryBookingRuleInfo {
    s.RuleType = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryBookingRuleInfo) SetRuleDesc(v string) *TaobaoAlitripTravelItemNewQueryBookingRuleInfo {
    s.RuleDesc = &v
    return s
}
