package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo struct {
    /*
        必填，1.可选择纯文本 2.使用xml标签（img和txt）进行图文混排，目前仅支持一段文字和多张图片，如果文件有多段，将会被合并 3.每段行程文字总和小于1500字     */
    Content  *string `json:"content,omitempty" `

    /*
        行程编号，第一天行程为1，第二天行程为2     */
    ItineraryNo  *int64 `json:"itinerary_no,omitempty" `

    /*
        小于等于30字     */
    Title  *string `json:"title,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo) SetContent(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo {
    s.Content = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo) SetItineraryNo(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo {
    s.ItineraryNo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo) SetTitle(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo {
    s.Title = &v
    return s
}
