package domain


type AlitripGrouptoursProductUploadItineraryActivity struct {
    /*
        活动标题     */
    Title  *string `json:"title,omitempty" `

    /*
        活动内容文本描述     */
    Txt  *string `json:"txt,omitempty" `

    /*
        活动图片列表，多个图片以英文逗号分隔     */
    Images  *[]string `json:"images,omitempty" `

    /*
        活动预计时长     */
    Hour  *int64 `json:"hour,omitempty" `

}

func (s *AlitripGrouptoursProductUploadItineraryActivity) SetTitle(v string) *AlitripGrouptoursProductUploadItineraryActivity {
    s.Title = &v
    return s
}
func (s *AlitripGrouptoursProductUploadItineraryActivity) SetTxt(v string) *AlitripGrouptoursProductUploadItineraryActivity {
    s.Txt = &v
    return s
}
func (s *AlitripGrouptoursProductUploadItineraryActivity) SetImages(v []string) *AlitripGrouptoursProductUploadItineraryActivity {
    s.Images = &v
    return s
}
func (s *AlitripGrouptoursProductUploadItineraryActivity) SetHour(v int64) *AlitripGrouptoursProductUploadItineraryActivity {
    s.Hour = &v
    return s
}
