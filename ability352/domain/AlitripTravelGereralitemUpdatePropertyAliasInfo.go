package domain


type AlitripTravelGereralitemUpdatePropertyAliasInfo struct {
    /*
        销售属性的pid和vid，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid     */
    Properties  *string `json:"properties,omitempty" `

    /*
        属性具体别名值     */
    Value  *string `json:"value,omitempty" `

}

func (s *AlitripTravelGereralitemUpdatePropertyAliasInfo) SetProperties(v string) *AlitripTravelGereralitemUpdatePropertyAliasInfo {
    s.Properties = &v
    return s
}
func (s *AlitripTravelGereralitemUpdatePropertyAliasInfo) SetValue(v string) *AlitripTravelGereralitemUpdatePropertyAliasInfo {
    s.Value = &v
    return s
}
