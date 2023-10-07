package domain


type AlitripDaytoursProductUploadStructItinerary struct {
    /*
        必填，行程序号，标识是第几天的行程     */
    DayOrder  *int64 `json:"day_order,omitempty" `

    /*
        必填，当天行程包含的多个活动信息     */
    Activities  *[]AlitripDaytoursProductUploadItineraryActivity `json:"activities,omitempty" `

}

func (s *AlitripDaytoursProductUploadStructItinerary) SetDayOrder(v int64) *AlitripDaytoursProductUploadStructItinerary {
    s.DayOrder = &v
    return s
}
func (s *AlitripDaytoursProductUploadStructItinerary) SetActivities(v []AlitripDaytoursProductUploadItineraryActivity) *AlitripDaytoursProductUploadStructItinerary {
    s.Activities = &v
    return s
}
