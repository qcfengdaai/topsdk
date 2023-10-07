package domain


type TaobaoUmpRangeGetRange struct {
    /*
        营销范围参与者id。即MarketingRangeDO中的participateId。     */
    ParticipateId  *int64 `json:"participate_id,omitempty" `

    /*
        营销范围参与者类型。MarketingRangeDO中的participateType。(1:商品;2:店铺;3:seller;4:sku;5:category;6:shopCategory)     */
    ParticipateType  *int64 `json:"participate_type,omitempty" `

}

func (s *TaobaoUmpRangeGetRange) SetParticipateId(v int64) *TaobaoUmpRangeGetRange {
    s.ParticipateId = &v
    return s
}
func (s *TaobaoUmpRangeGetRange) SetParticipateType(v int64) *TaobaoUmpRangeGetRange {
    s.ParticipateType = &v
    return s
}
