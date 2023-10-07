package domain


type AlitripTravelGereralskuUpdatePropertyAliasInfo struct {
    /*
        销售属性的pid和vid，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid     */
    Properties  *string `json:"properties,omitempty" `

    /*
        属性别名     */
    Value  *string `json:"value,omitempty" `

}

func (s *AlitripTravelGereralskuUpdatePropertyAliasInfo) SetProperties(v string) *AlitripTravelGereralskuUpdatePropertyAliasInfo {
    s.Properties = &v
    return s
}
func (s *AlitripTravelGereralskuUpdatePropertyAliasInfo) SetValue(v string) *AlitripTravelGereralskuUpdatePropertyAliasInfo {
    s.Value = &v
    return s
}
