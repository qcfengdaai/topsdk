package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemResourceInfo struct {
    /*
        小于1500字符     */
    Desc  *string `json:"desc,omitempty" `

    /*
        1-保险2-餐饮3-租车4-签证5-购物点6-赠品7-券99-其他     */
    Type  *int64 `json:"type,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemResourceInfo) SetDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemResourceInfo {
    s.Desc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemResourceInfo) SetType(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemResourceInfo {
    s.Type = &v
    return s
}
