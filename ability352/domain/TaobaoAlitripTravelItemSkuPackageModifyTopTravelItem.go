package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuPackageModifyTopTravelItem struct {
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

func (s *TaobaoAlitripTravelItemSkuPackageModifyTopTravelItem) SetCreated(v util.LocalTime) *TaobaoAlitripTravelItemSkuPackageModifyTopTravelItem {
    s.Created = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyTopTravelItem) SetItemId(v int64) *TaobaoAlitripTravelItemSkuPackageModifyTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyTopTravelItem) SetModified(v util.LocalTime) *TaobaoAlitripTravelItemSkuPackageModifyTopTravelItem {
    s.Modified = &v
    return s
}
