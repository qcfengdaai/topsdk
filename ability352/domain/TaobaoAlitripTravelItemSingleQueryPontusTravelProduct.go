package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelProduct struct {
    /*
        资源元素的外部商家编码     */
    ElementId  *string `json:"element_id,omitempty" `

    /*
        元素份数     */
    Num  *int64 `json:"num,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelProduct) SetElementId(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelProduct {
    s.ElementId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelProduct) SetNum(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelProduct {
    s.Num = &v
    return s
}
