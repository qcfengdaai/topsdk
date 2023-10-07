package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem struct {
    /*
        商品基本信息     */
    BaseInfo  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo `json:"base_info,omitempty" `

    /*
        预定规则     */
    BookingRules  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo `json:"booking_rules,omitempty" `

    /*
        创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        自由行相关信息     */
    FreedomItemExt  *TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt `json:"freedom_item_ext,omitempty" `

    /*
        跟团游相关信息     */
    GroupItemExt  *TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt `json:"group_item_ext,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品状态。0,1正常;-1:用户删除;-2:用户下架;-3 小二下架;-4 小二删除;-5 从未上架;-9 被处罚     */
    ItemStatus  *int64 `json:"item_status,omitempty" `

    /*
        商品类型     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        行程信息     */
    Itineraries  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo `json:"itineraries,omitempty" `

    /*
        修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        退改规则信息     */
    RefundInfo  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemRefundInfo `json:"refund_info,omitempty" `

    /*
        销售属性信息     */
    SaleInfo  *TaobaoAlitripTravelItemSingleQueryPontusTravelItemSaleInfo `json:"sale_info,omitempty" `

    /*
        卖家id     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        sku信息     */
    SkuInfos  *[]TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo `json:"sku_infos,omitempty" `

    /*
        同城活动商品相关信息     */
    TcwlItemExt  *TaobaoAlitripTravelItemSingleQueryTcwlItemExt `json:"tcwl_item_ext,omitempty" `

    /*
        航旅度假TOP API3.0 邮轮扩展信息结构     */
    CruiseItemExt  *TaobaoAlitripTravelItemSingleQueryCruiseItemExt `json:"cruise_item_ext,omitempty" `

    /*
        商品特征数据     */
    Features  *string `json:"features,omitempty" `

    /*
        新版行程信息     */
    RefTrip  *string `json:"ref_trip,omitempty" `

    /*
        产品亮点     */
    HighLights  *[]TaobaoAlitripTravelItemSingleQueryHighLights `json:"high_lights,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetBaseInfo(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemBaseInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.BaseInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetBookingRules(v []TaobaoAlitripTravelItemSingleQueryPontusTravelBookingRuleInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.BookingRules = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetCreated(v util.LocalTime) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.Created = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetFreedomItemExt(v TaobaoAlitripTravelItemSingleQueryPontusTravelFreedomItemExt) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.FreedomItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetGroupItemExt(v TaobaoAlitripTravelItemSingleQueryPontusTravelGroupItemExt) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.GroupItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetItemId(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetItemStatus(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.ItemStatus = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetItemType(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.ItemType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetItineraries(v []TaobaoAlitripTravelItemSingleQueryPontusTravelItemItineraryInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.Itineraries = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetModified(v util.LocalTime) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.Modified = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetRefundInfo(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemRefundInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.RefundInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetSaleInfo(v TaobaoAlitripTravelItemSingleQueryPontusTravelItemSaleInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.SaleInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetSellerId(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.SellerId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetSellerNick(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.SellerNick = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetSkuInfos(v []TaobaoAlitripTravelItemSingleQueryPontusTravelItemSkuInfo) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.SkuInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetTcwlItemExt(v TaobaoAlitripTravelItemSingleQueryTcwlItemExt) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.TcwlItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetCruiseItemExt(v TaobaoAlitripTravelItemSingleQueryCruiseItemExt) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.CruiseItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetFeatures(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.Features = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetRefTrip(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.RefTrip = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem) SetHighLights(v []TaobaoAlitripTravelItemSingleQueryHighLights) *TaobaoAlitripTravelItemSingleQueryPontusTravelFullTravelItem {
    s.HighLights = &v
    return s
}
