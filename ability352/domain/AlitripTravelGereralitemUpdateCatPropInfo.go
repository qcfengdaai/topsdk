package domain


type AlitripTravelGereralitemUpdateCatPropInfo struct {
    /*
        属性PID，调用taobao.itemprops.get取得     */
    Pid  *string `json:"pid,omitempty" `

    /*
        属性VID，调用taobao.itempropvalues.get取得     */
    Vid  *string `json:"vid,omitempty" `

}

func (s *AlitripTravelGereralitemUpdateCatPropInfo) SetPid(v string) *AlitripTravelGereralitemUpdateCatPropInfo {
    s.Pid = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateCatPropInfo) SetVid(v string) *AlitripTravelGereralitemUpdateCatPropInfo {
    s.Vid = &v
    return s
}
