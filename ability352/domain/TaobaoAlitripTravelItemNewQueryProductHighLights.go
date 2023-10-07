package domain


type TaobaoAlitripTravelItemNewQueryProductHighLights struct {
    /*
        图片列表     */
    PicUrls  *[]string `json:"pic_urls,omitempty" `

    /*
        亮点描述     */
    Desc  *string `json:"desc,omitempty" `

    /*
        亮点标题     */
    Title  *string `json:"title,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryProductHighLights) SetPicUrls(v []string) *TaobaoAlitripTravelItemNewQueryProductHighLights {
    s.PicUrls = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryProductHighLights) SetDesc(v string) *TaobaoAlitripTravelItemNewQueryProductHighLights {
    s.Desc = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryProductHighLights) SetTitle(v string) *TaobaoAlitripTravelItemNewQueryProductHighLights {
    s.Title = &v
    return s
}
