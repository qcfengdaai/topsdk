package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemNewQueryFullTravelItem struct {
    /*
        商品基本信息     */
    BaseInfo  *TaobaoAlitripTravelItemNewQueryItemBaseInfo `json:"base_info,omitempty" `

    /*
        预定规则     */
    BookingRules  *[]TaobaoAlitripTravelItemNewQueryBookingRuleInfo `json:"booking_rules,omitempty" `

    /*
        创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        自由行相关信息     */
    FreedomItemExt  *TaobaoAlitripTravelItemNewQueryFreedomItemExt `json:"freedom_item_ext,omitempty" `

    /*
        跟团游相关信息     */
    GroupItemExt  *TaobaoAlitripTravelItemNewQueryGroupItemExt `json:"group_item_ext,omitempty" `

    /*
        商品类型     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        商品状态。0,1正常;-1:用户删除;-2:用户下架;-3 小二下架;-4 小二删除;-5 从未上架;-9 被处罚     */
    ItemStatus  *int64 `json:"item_status,omitempty" `

    /*
        商品id1     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        行程信息     */
    Itineraries  *[]TaobaoAlitripTravelItemNewQueryItemItineraryInfo `json:"itineraries,omitempty" `

    /*
        修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        退改规则信息     */
    RefundInfo  *TaobaoAlitripTravelItemNewQueryItemRefundInfo `json:"refund_info,omitempty" `

    /*
        销售属性信息     */
    SaleInfo  *TaobaoAlitripTravelItemNewQueryItemSaleInfo `json:"sale_info,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        卖家id     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        sku信息     */
    SkuInfos  *[]TaobaoAlitripTravelItemNewQueryItemSkuInfo `json:"sku_infos,omitempty" `

    /*
        同城活动商品相关信息     */
    TcwlItemExt  *TaobaoAlitripTravelItemNewQueryTcwlItemExt `json:"tcwl_item_ext,omitempty" `

    /*
        航旅度假TOP API3.0 邮轮扩展信息结构     */
    CruiseItemExt  *TaobaoAlitripTravelItemNewQueryCruiseItemExt `json:"cruise_item_ext,omitempty" `

    /*
        商品特征数据     */
    Features  *string `json:"features,omitempty" `

    /*
        新版行程信息     */
    RefTrip  *string `json:"ref_trip,omitempty" `

    /*
        产品亮点     */
    HighLights  *[]TaobaoAlitripTravelItemNewQueryProductHighLights `json:"high_lights,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetBaseInfo(v TaobaoAlitripTravelItemNewQueryItemBaseInfo) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.BaseInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetBookingRules(v []TaobaoAlitripTravelItemNewQueryBookingRuleInfo) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.BookingRules = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetCreated(v util.LocalTime) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.Created = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetFreedomItemExt(v TaobaoAlitripTravelItemNewQueryFreedomItemExt) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.FreedomItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetGroupItemExt(v TaobaoAlitripTravelItemNewQueryGroupItemExt) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.GroupItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetItemType(v int64) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.ItemType = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetItemStatus(v int64) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.ItemStatus = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetItemId(v int64) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetItineraries(v []TaobaoAlitripTravelItemNewQueryItemItineraryInfo) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.Itineraries = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetModified(v util.LocalTime) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.Modified = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetRefundInfo(v TaobaoAlitripTravelItemNewQueryItemRefundInfo) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.RefundInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetSaleInfo(v TaobaoAlitripTravelItemNewQueryItemSaleInfo) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.SaleInfo = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetSellerNick(v string) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.SellerNick = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetSellerId(v int64) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.SellerId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetSkuInfos(v []TaobaoAlitripTravelItemNewQueryItemSkuInfo) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.SkuInfos = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetTcwlItemExt(v TaobaoAlitripTravelItemNewQueryTcwlItemExt) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.TcwlItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetCruiseItemExt(v TaobaoAlitripTravelItemNewQueryCruiseItemExt) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.CruiseItemExt = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetFeatures(v string) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.Features = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetRefTrip(v string) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.RefTrip = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryFullTravelItem) SetHighLights(v []TaobaoAlitripTravelItemNewQueryProductHighLights) *TaobaoAlitripTravelItemNewQueryFullTravelItem {
    s.HighLights = &v
    return s
}
