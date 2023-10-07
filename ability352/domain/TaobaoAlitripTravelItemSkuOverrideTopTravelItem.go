package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuOverrideTopTravelItem struct {
    /*
        商品创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuOverrideTopTravelItem) SetCreated(v util.LocalTime) *TaobaoAlitripTravelItemSkuOverrideTopTravelItem {
    s.Created = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverrideTopTravelItem) SetItemId(v int64) *TaobaoAlitripTravelItemSkuOverrideTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverrideTopTravelItem) SetModified(v util.LocalTime) *TaobaoAlitripTravelItemSkuOverrideTopTravelItem {
    s.Modified = &v
    return s
}
