package domain


type AlitripGrouptoursProductUploadStructItinerary struct {
    /*
        必填，行程序号，标识是第几天的行程     */
    DayOrder  *int64 `json:"day_order,omitempty" `

    /*
        必填，当天行程包含的多个活动信息     */
    Activities  *[]AlitripGrouptoursProductUploadItineraryActivity `json:"activities,omitempty" `

}

func (s *AlitripGrouptoursProductUploadStructItinerary) SetDayOrder(v int64) *AlitripGrouptoursProductUploadStructItinerary {
    s.DayOrder = &v
    return s
}
func (s *AlitripGrouptoursProductUploadStructItinerary) SetActivities(v []AlitripGrouptoursProductUploadItineraryActivity) *AlitripGrouptoursProductUploadStructItinerary {
    s.Activities = &v
    return s
}
