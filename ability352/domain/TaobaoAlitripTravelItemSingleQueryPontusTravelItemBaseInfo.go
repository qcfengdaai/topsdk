package domain


type TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo struct {
    /*
        行程晚数     */
    AccomNights  *int64 `json:"accom_nights,omitempty" `

    /*
        商品类目id     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        宝贝所在地市     */
    City  *string `json:"city,omitempty" `

    /*
        商品描述(PC描述)     */
    Desc  *string `json:"desc,omitempty" `

    /*
        出发地     */
    FromLocations  *string `json:"from_locations,omitempty" `

    /*
        商品类型：1-国内跟团游 2-国内自由行 3-出境跟团游 4-出境自由行 5-境外当地玩乐 6-国际邮轮 7-国内当地玩乐 9-国内邮轮 10-同城玩乐     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        商户的商品编码     */
    OutId  *string `json:"out_id,omitempty" `

    /*
        商家备注     */
    OuterTitle  *string `json:"outer_title,omitempty" `

    /*
        商品主图     */
    PicUrls  *[]string `json:"pic_urls,omitempty" `

    /*
        宝贝所在地省     */
    Prov  *string `json:"prov,omitempty" `

    /*
        商品亮点     */
    SubTitles  *[]string `json:"sub_titles,omitempty" `

    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        目的地     */
    ToLocations  *string `json:"to_locations,omitempty" `

    /*
        最大出行天数     */
    TripMaxDays  *int64 `json:"trip_max_days,omitempty" `

    /*
        最小出行天数     */
    TripMinDays  *int64 `json:"trip_min_days,omitempty" `

    /*
        手机端商品描述     */
    WapDesc  *string `json:"wap_desc,omitempty" `

    /*
        商品标签     */
    ItemTagContent  *string `json:"item_tag_content,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetAccomNights(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.AccomNights = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetCategoryId(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.CategoryId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetCity(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.City = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.Desc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetFromLocations(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.FromLocations = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetItemType(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.ItemType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetOutId(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.OutId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetOuterTitle(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.OuterTitle = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetPicUrls(v []string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.PicUrls = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetProv(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.Prov = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetSubTitles(v []string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.SubTitles = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetTitle(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.Title = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetToLocations(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.ToLocations = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetTripMaxDays(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.TripMaxDays = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetTripMinDays(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.TripMinDays = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetWapDesc(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.WapDesc = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) SetItemTagContent(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo {
    s.ItemTagContent = &v
    return s
}
