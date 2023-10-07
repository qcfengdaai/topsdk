package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo struct {
    /*
        结构化景点信息 景点结构化信息和文本描述二选一     */
    ScenicList  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic `json:"scenic_list,omitempty" `

    /*
        景点描述文本，<1500字符 景点结构化信息和文本描述二选一     */
    ScenicDesc  *string `json:"scenic_desc,omitempty" `

    /*
        门票套餐名称     */
    TicketPackageName  *string `json:"ticket_package_name,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo) SetScenicList(v []TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenic) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo {
    s.ScenicList = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo) SetScenicDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo {
    s.ScenicDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo) SetTicketPackageName(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemScenicInfo {
    s.TicketPackageName = &v
    return s
}
