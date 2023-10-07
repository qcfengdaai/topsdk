package domain


type TaobaoAlitripTravelItemSingleQueryTcwlItemExt struct {
    /*
        玩乐主题     */
    TcwlTheme  *int64 `json:"tcwl_theme,omitempty" `

    /*
        主题玩法     */
    TcwlThemePlay  *string `json:"tcwl_theme_play,omitempty" `

    /*
        组织者名称     */
    Organization  *string `json:"organization,omitempty" `

    /*
        组织者介绍     */
    OrgIntroduce  *string `json:"org_introduce,omitempty" `

    /*
        组织者电话     */
    OrgTel  *string `json:"org_tel,omitempty" `

    /*
        组织者旺旺     */
    OrgWangwang  *string `json:"org_wangwang,omitempty" `

    /*
        活动强度。1：底，2：中，3：高     */
    ActivityStrength  *int64 `json:"activity_strength,omitempty" `

    /*
        活动地点     */
    ActivityPlace  *string `json:"activity_place,omitempty" `

    /*
        活动时间     */
    ActivityTime  *string `json:"activity_time,omitempty" `

    /*
        集合地信息     */
    GatherPlaces  *[]TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo `json:"gather_places,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetTcwlTheme(v int64) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.TcwlTheme = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetTcwlThemePlay(v string) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.TcwlThemePlay = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetOrganization(v string) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.Organization = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetOrgIntroduce(v string) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.OrgIntroduce = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetOrgTel(v string) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.OrgTel = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetOrgWangwang(v string) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.OrgWangwang = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetActivityStrength(v int64) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.ActivityStrength = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetActivityPlace(v string) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.ActivityPlace = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetActivityTime(v string) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.ActivityTime = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryTcwlItemExt) SetGatherPlaces(v []TaobaoAlitripTravelItemSingleQueryGatherPlaceInfo) *TaobaoAlitripTravelItemSingleQueryTcwlItemExt {
    s.GatherPlaces = &v
    return s
}
