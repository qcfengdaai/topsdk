package domain


type AlitripTravelGereralskuUpdateCatPropInfo struct {
    /*
        属性PID，调用taobao.itemprops.get取得     */
    Pid  *string `json:"pid,omitempty" `

    /*
        属性VID，调用taobao.itempropvalues.get取得     */
    Vid  *string `json:"vid,omitempty" `

}

func (s *AlitripTravelGereralskuUpdateCatPropInfo) SetPid(v string) *AlitripTravelGereralskuUpdateCatPropInfo {
    s.Pid = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateCatPropInfo) SetVid(v string) *AlitripTravelGereralskuUpdateCatPropInfo {
    s.Vid = &v
    return s
}
