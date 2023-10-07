package domain


type TaobaoAlitripTravelItemNewQueryItemBaseInfo struct {
    /*
        商品标签     */
    ItemTagContent  *string `json:"item_tag_content,omitempty" `

    /*
        手机端商品描述     */
    WapDesc  *string `json:"wap_desc,omitempty" `

    /*
        最小出行天数     */
    TripMinDays  *int64 `json:"trip_min_days,omitempty" `

    /*
        最大出行天数     */
    TripMaxDays  *int64 `json:"trip_max_days,omitempty" `

    /*
        目的地     */
    ToLocations  *string `json:"to_locations,omitempty" `

    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品亮点     */
    SubTitles  *[]string `json:"sub_titles,omitempty" `

    /*
        宝贝所在地省     */
    Prov  *string `json:"prov,omitempty" `

    /*
        ["xxxx","xxxxx"]     */
    PicUrls  *[]string `json:"pic_urls,omitempty" `

    /*
        商家备注     */
    OuterTitle  *string `json:"outer_title,omitempty" `

    /*
        商户的商品编码     */
    OutId  *string `json:"out_id,omitempty" `

    /*
        商品类型：1-国内跟团游 2-国内自由行 3-出境跟团游 4-出境自由行 5-境外当地玩乐 6-国际邮轮 7-国内当地玩乐 9-国内邮轮 10-同城玩乐     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        出发地     */
    FromLocations  *string `json:"from_locations,omitempty" `

    /*
        商品描述(PC描述)     */
    Desc  *string `json:"desc,omitempty" `

    /*
        宝贝所在地市     */
    City  *string `json:"city,omitempty" `

    /*
        商品类目id     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        行程晚数     */
    AccomNights  *int64 `json:"accom_nights,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetItemTagContent(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.ItemTagContent = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetWapDesc(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.WapDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetTripMinDays(v int64) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.TripMinDays = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetTripMaxDays(v int64) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.TripMaxDays = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetToLocations(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.ToLocations = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetTitle(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.Title = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetSubTitles(v []string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.SubTitles = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetProv(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.Prov = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetPicUrls(v []string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.PicUrls = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetOuterTitle(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.OuterTitle = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetOutId(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.OutId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetItemType(v int64) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.ItemType = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetFromLocations(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.FromLocations = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetDesc(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.Desc = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetCity(v string) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.City = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetCategoryId(v int64) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.CategoryId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryItemBaseInfo) SetAccomNights(v int64) *TaobaoAlitripTravelItemNewQueryItemBaseInfo {
    s.AccomNights = &v
    return s
}
