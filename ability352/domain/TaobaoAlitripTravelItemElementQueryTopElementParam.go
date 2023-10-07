package domain


type TaobaoAlitripTravelItemElementQueryTopElementParam struct {
    /*
        资源元素类型。1-景点，2-酒店，999-其他     */
    ElementType  *int64 `json:"element_type,omitempty" `

    /*
        元素所在城市     */
    City  *string `json:"city,omitempty" `

    /*
        元素名称     */
    Name  *string `json:"name,omitempty" `

    /*
        元素的子类型     */
    Type  *string `json:"type,omitempty" `

    /*
        元素的外部商家编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        元素的说明描述     */
    Desc  *string `json:"desc,omitempty" `

}

func (s *TaobaoAlitripTravelItemElementQueryTopElementParam) SetElementType(v int64) *TaobaoAlitripTravelItemElementQueryTopElementParam {
    s.ElementType = &v
    return s
}
func (s *TaobaoAlitripTravelItemElementQueryTopElementParam) SetCity(v string) *TaobaoAlitripTravelItemElementQueryTopElementParam {
    s.City = &v
    return s
}
func (s *TaobaoAlitripTravelItemElementQueryTopElementParam) SetName(v string) *TaobaoAlitripTravelItemElementQueryTopElementParam {
    s.Name = &v
    return s
}
func (s *TaobaoAlitripTravelItemElementQueryTopElementParam) SetType(v string) *TaobaoAlitripTravelItemElementQueryTopElementParam {
    s.Type = &v
    return s
}
func (s *TaobaoAlitripTravelItemElementQueryTopElementParam) SetOuterId(v string) *TaobaoAlitripTravelItemElementQueryTopElementParam {
    s.OuterId = &v
    return s
}
func (s *TaobaoAlitripTravelItemElementQueryTopElementParam) SetDesc(v string) *TaobaoAlitripTravelItemElementQueryTopElementParam {
    s.Desc = &v
    return s
}
