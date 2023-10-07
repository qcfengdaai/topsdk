package domain

import (
        "topsdk/util"
    )

type AlitripTravelGereralskuUpdateTopTravelItem struct {
    /*
        sku创建日期     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        itemId     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        sku修改日期     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        skuId     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        outerId     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *AlitripTravelGereralskuUpdateTopTravelItem) SetCreated(v util.LocalTime) *AlitripTravelGereralskuUpdateTopTravelItem {
    s.Created = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateTopTravelItem) SetItemId(v int64) *AlitripTravelGereralskuUpdateTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateTopTravelItem) SetModified(v util.LocalTime) *AlitripTravelGereralskuUpdateTopTravelItem {
    s.Modified = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateTopTravelItem) SetSkuId(v int64) *AlitripTravelGereralskuUpdateTopTravelItem {
    s.SkuId = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateTopTravelItem) SetOuterId(v string) *AlitripTravelGereralskuUpdateTopTravelItem {
    s.OuterId = &v
    return s
}
