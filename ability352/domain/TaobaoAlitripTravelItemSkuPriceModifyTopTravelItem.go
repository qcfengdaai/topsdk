package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem struct {
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

func (s *TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem) SetCreated(v util.LocalTime) *TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem {
    s.Created = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem) SetItemId(v int64) *TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem) SetModified(v util.LocalTime) *TaobaoAlitripTravelItemSkuPriceModifyTopTravelItem {
    s.Modified = &v
    return s
}
