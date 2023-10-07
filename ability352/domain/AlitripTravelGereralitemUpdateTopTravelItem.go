package domain

import (
        "topsdk/util"
    )

type AlitripTravelGereralitemUpdateTopTravelItem struct {
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

func (s *AlitripTravelGereralitemUpdateTopTravelItem) SetCreated(v util.LocalTime) *AlitripTravelGereralitemUpdateTopTravelItem {
    s.Created = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateTopTravelItem) SetItemId(v int64) *AlitripTravelGereralitemUpdateTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateTopTravelItem) SetModified(v util.LocalTime) *AlitripTravelGereralitemUpdateTopTravelItem {
    s.Modified = &v
    return s
}
