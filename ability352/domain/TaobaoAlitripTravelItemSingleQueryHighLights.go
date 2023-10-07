package domain


type TaobaoAlitripTravelItemSingleQueryHighLights struct {
    /*
        亮点标题     */
    Title  *string `json:"title,omitempty" `

    /*
        亮点描述     */
    Desc  *string `json:"desc,omitempty" `

    /*
        图片列表     */
    PicUrls  *[]string `json:"pic_urls,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryHighLights) SetTitle(v string) *TaobaoAlitripTravelItemSingleQueryHighLights {
    s.Title = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryHighLights) SetDesc(v string) *TaobaoAlitripTravelItemSingleQueryHighLights {
    s.Desc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryHighLights) SetPicUrls(v []string) *TaobaoAlitripTravelItemSingleQueryHighLights {
    s.PicUrls = &v
    return s
}
