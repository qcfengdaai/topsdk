package domain


type AlitripDaytoursProductUploadProductHighLights struct {
    /*
        产品亮点标题     */
    Title  *string `json:"title,omitempty" `

    /*
        产品亮点描述     */
    Desc  *string `json:"desc,omitempty" `

    /*
        产品亮点图片     */
    PicUrls  *[]string `json:"pic_urls,omitempty" `

}

func (s *AlitripDaytoursProductUploadProductHighLights) SetTitle(v string) *AlitripDaytoursProductUploadProductHighLights {
    s.Title = &v
    return s
}
func (s *AlitripDaytoursProductUploadProductHighLights) SetDesc(v string) *AlitripDaytoursProductUploadProductHighLights {
    s.Desc = &v
    return s
}
func (s *AlitripDaytoursProductUploadProductHighLights) SetPicUrls(v []string) *AlitripDaytoursProductUploadProductHighLights {
    s.PicUrls = &v
    return s
}
