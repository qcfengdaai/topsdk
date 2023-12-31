package domain


type AlitripGrouptourProductUploadGroupTourTrafficInfo struct {
    /*
        交通类型 1：飞机， 2：火车，3：汽车，4：轮船     */
    TransportWay  *int64 `json:"transport_way,omitempty" `

    /*
        第几组交通信息，每一组交通信息包含一组去程交通和返程交通，当在页> 面上点击【添加交通信息】按钮后，就会出现第二组交通信息，第一组交>通信息group=1，第二组交通信息group取值为2，以此类推     */
    Group  *int64 `json:"group,omitempty" `

    /*
        是否直飞，飞机选填，1-直飞；0-不是直飞     */
    NonStop  *int64 `json:"non_stop,omitempty" `

    /*
        交通说明，针对交通类型是汽车，轮船和其他     */
    TrafficDesc  *string `json:"traffic_desc,omitempty" `

    /*
        到达时间，当地时间HH:mm     */
    ArrivalTime  *string `json:"arrival_time,omitempty" `

    /*
        出发时间，当地时间HH:mm     */
    DepartureTime  *string `json:"departure_time,omitempty" `

    /*
        到达城市     */
    Destination  *string `json:"destination,omitempty" `

    /*
        出发城市     */
    Departure  *string `json:"departure,omitempty" `

    /*
        飞机机型，飞机选填     */
    PlaneType  *string `json:"plane_type,omitempty" `

    /*
        交通公司名，飞机选填     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        参考班次号，飞机需要填航班号，火车需要填车次号，汽车和船可不填     */
    TrafficNo  *string `json:"traffic_no,omitempty" `

    /*
        第几天     */
    Day  *int64 `json:"day,omitempty" `

    /*
        是否经停     */
    StopOver  *bool `json:"stop_over,omitempty" `

    /*
        经停城市     */
    StopCity  *string `json:"stop_city,omitempty" `

    /*
        是否是"非红眼航班"。【红眼航班】定义：凌晨一点至六点起飞，且飞行时间少于少于正常睡眠需求（8小时）的航班。     */
    IsNonRedEyeFlight  *bool `json:"is_non_red_eye_flight,omitempty" `

}

func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetTransportWay(v int64) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.TransportWay = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetGroup(v int64) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.Group = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetNonStop(v int64) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.NonStop = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetTrafficDesc(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.TrafficDesc = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetArrivalTime(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.ArrivalTime = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetDepartureTime(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.DepartureTime = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetDestination(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.Destination = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetDeparture(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.Departure = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetPlaneType(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.PlaneType = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetVendor(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.Vendor = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetTrafficNo(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.TrafficNo = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetDay(v int64) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.Day = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetStopOver(v bool) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.StopOver = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetStopCity(v string) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.StopCity = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourTrafficInfo) SetIsNonRedEyeFlight(v bool) *AlitripGrouptourProductUploadGroupTourTrafficInfo {
    s.IsNonRedEyeFlight = &v
    return s
}
