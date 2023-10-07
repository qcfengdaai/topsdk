package domain


type AlitripFreetourProductUploadFreeTourTrafficInfo struct {
    /*
        第几天     */
    Day  *int64 `json:"day,omitempty" `

    /*
        参考班次号，飞机需要填航班号，火车需要填车次号，汽车和船可不填     */
    TrafficNo  *string `json:"traffic_no,omitempty" `

    /*
        交通公司名，飞机选填     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        飞机机型，飞机选填     */
    PlaneType  *string `json:"plane_type,omitempty" `

    /*
        出发城市     */
    Departure  *string `json:"departure,omitempty" `

    /*
        到达城市     */
    Destination  *string `json:"destination,omitempty" `

    /*
        出发时间，当地时间HH:mm     */
    DepartureTime  *string `json:"departure_time,omitempty" `

    /*
        到达时间，当地时间HH:mm     */
    ArrivalTime  *string `json:"arrival_time,omitempty" `

    /*
        交通说明，针对交通类型是汽车，轮船和其他     */
    TrafficDesc  *string `json:"traffic_desc,omitempty" `

    /*
        是否直飞，飞机选填，1-直飞；0-不是直飞     */
    NonStop  *int64 `json:"non_stop,omitempty" `

    /*
        第几组交通信息，请必传，不传可能导致商品详情页交通信息不展示。每一组交通信息包含一组去程交通和返程交通，当在页> 面上点击【添加交通信息】按钮后，就会出现第二组交通信息，第一组交>通信息group=1，第二组交通信息group取值为2，以此类推     */
    Group  *int64 `json:"group,omitempty" `

    /*
        是否经停     */
    StopOver  *bool `json:"stop_over,omitempty" `

    /*
        经停城市     */
    StopCity  *string `json:"stop_city,omitempty" `

}

func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetDay(v int64) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.Day = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetTrafficNo(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.TrafficNo = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetVendor(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.Vendor = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetPlaneType(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.PlaneType = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetDeparture(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.Departure = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetDestination(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.Destination = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetDepartureTime(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.DepartureTime = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetArrivalTime(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.ArrivalTime = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetTrafficDesc(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.TrafficDesc = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetNonStop(v int64) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.NonStop = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetGroup(v int64) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.Group = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetStopOver(v bool) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.StopOver = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourTrafficInfo) SetStopCity(v string) *AlitripFreetourProductUploadFreeTourTrafficInfo {
    s.StopCity = &v
    return s
}
