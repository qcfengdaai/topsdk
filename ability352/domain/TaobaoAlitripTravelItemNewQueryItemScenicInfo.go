package domain


type TaobaoAlitripTravelItemNewQueryItemScenicInfo struct {
    /*
        结构化景点信息 景点结构化信息和文本描述二选一     */
    ScenicList  *[]TaobaoAlitripTravelItemNewQueryItemScenic `json:"scenic_list,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryItemScenicInfo) SetScenicList(v []TaobaoAlitripTravelItemNewQueryItemScenic) *TaobaoAlitripTravelItemNewQueryItemScenicInfo {
    s.ScenicList = &v
    return s
}
