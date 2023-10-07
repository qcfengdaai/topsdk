package domain


type TaobaoAlitripTravelItemNewQueryTcwlItemExt struct {
    /*
        集合地信息     */
    GatherPlaces  *[]TaobaoAlitripTravelItemNewQueryGatherPlaceInfo `json:"gather_places,omitempty" `

    /*
        活动时间     */
    ActivityTime  *string `json:"activity_time,omitempty" `

    /*
        活动地点     */
    ActivityPlace  *string `json:"activity_place,omitempty" `

    /*
        活动强度。1：底，2：中，3：高     */
    ActivityStrength  *int64 `json:"activity_strength,omitempty" `

    /*
        组织者旺旺     */
    OrgWangwang  *string `json:"org_wangwang,omitempty" `

    /*
        组织者电话     */
    OrgTel  *string `json:"org_tel,omitempty" `

    /*
        组织者介绍     */
    OrgIntroduce  *string `json:"org_introduce,omitempty" `

    /*
        组织者名称     */
    Organization  *string `json:"organization,omitempty" `

    /*
        主题玩法     */
    TcwlThemePlay  *string `json:"tcwl_theme_play,omitempty" `

    /*
        玩乐主题     */
    TcwlTheme  *int64 `json:"tcwl_theme,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetGatherPlaces(v []TaobaoAlitripTravelItemNewQueryGatherPlaceInfo) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.GatherPlaces = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetActivityTime(v string) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.ActivityTime = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetActivityPlace(v string) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.ActivityPlace = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetActivityStrength(v int64) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.ActivityStrength = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetOrgWangwang(v string) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.OrgWangwang = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetOrgTel(v string) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.OrgTel = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetOrgIntroduce(v string) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.OrgIntroduce = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetOrganization(v string) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.Organization = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetTcwlThemePlay(v string) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.TcwlThemePlay = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryTcwlItemExt) SetTcwlTheme(v int64) *TaobaoAlitripTravelItemNewQueryTcwlItemExt {
    s.TcwlTheme = &v
    return s
}
