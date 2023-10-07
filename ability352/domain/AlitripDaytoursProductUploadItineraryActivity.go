package domain


type AlitripDaytoursProductUploadItineraryActivity struct {
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
        活动预计时长，小时数     */
    Hour  *int64 `json:"hour,omitempty" `

    /*
        活动预计时长，分钟数     */
    Minute  *int64 `json:"minute,omitempty" `

}

func (s *AlitripDaytoursProductUploadItineraryActivity) SetTitle(v string) *AlitripDaytoursProductUploadItineraryActivity {
    s.Title = &v
    return s
}
func (s *AlitripDaytoursProductUploadItineraryActivity) SetTxt(v string) *AlitripDaytoursProductUploadItineraryActivity {
    s.Txt = &v
    return s
}
func (s *AlitripDaytoursProductUploadItineraryActivity) SetImages(v []string) *AlitripDaytoursProductUploadItineraryActivity {
    s.Images = &v
    return s
}
func (s *AlitripDaytoursProductUploadItineraryActivity) SetHour(v int64) *AlitripDaytoursProductUploadItineraryActivity {
    s.Hour = &v
    return s
}
func (s *AlitripDaytoursProductUploadItineraryActivity) SetMinute(v int64) *AlitripDaytoursProductUploadItineraryActivity {
    s.Minute = &v
    return s
}
