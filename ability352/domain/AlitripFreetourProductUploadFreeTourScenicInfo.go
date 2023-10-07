package domain


type AlitripFreetourProductUploadFreeTourScenicInfo struct {
    /*
        必填，景点名称     */
    CnName  *string `json:"cn_name,omitempty" `

    /*
        必填，门票类型     */
    TicketType  *string `json:"ticket_type,omitempty" `

    /*
        必填，景点所在城市     */
    City  *string `json:"city,omitempty" `

}

func (s *AlitripFreetourProductUploadFreeTourScenicInfo) SetCnName(v string) *AlitripFreetourProductUploadFreeTourScenicInfo {
    s.CnName = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourScenicInfo) SetTicketType(v string) *AlitripFreetourProductUploadFreeTourScenicInfo {
    s.TicketType = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourScenicInfo) SetCity(v string) *AlitripFreetourProductUploadFreeTourScenicInfo {
    s.City = &v
    return s
}
