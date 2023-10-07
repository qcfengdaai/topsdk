package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo struct {
    /*
        规则说明，1500个字符     */
    RuleDesc  *string `json:"rule_desc,omitempty" `

    /*
        Fee_Included：费用包含，跟团游必填； Fee_Excluded：费用不含，所有类目必填； Order_Info：预定须知； Extra_Cost：其他费用，预留；     Shopping：购物说明，预留     */
    RuleType  *string `json:"rule_type,omitempty" `

    /*
        分点组织的规则说明     */
    RuleList  *[]string `json:"rule_list,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo) SetRuleDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo {
    s.RuleDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo) SetRuleType(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo {
    s.RuleType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo) SetRuleList(v []string) *TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo {
    s.RuleList = &v
    return s
}
