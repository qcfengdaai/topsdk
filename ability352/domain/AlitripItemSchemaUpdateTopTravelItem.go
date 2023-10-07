package domain

import (
        "topsdk/util"
    )

type AlitripItemSchemaUpdateTopTravelItem struct {
    /*
        商品修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        商品创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

}

func (s *AlitripItemSchemaUpdateTopTravelItem) SetModified(v util.LocalTime) *AlitripItemSchemaUpdateTopTravelItem {
    s.Modified = &v
    return s
}
func (s *AlitripItemSchemaUpdateTopTravelItem) SetCreated(v util.LocalTime) *AlitripItemSchemaUpdateTopTravelItem {
    s.Created = &v
    return s
}
func (s *AlitripItemSchemaUpdateTopTravelItem) SetItemId(v int64) *AlitripItemSchemaUpdateTopTravelItem {
    s.ItemId = &v
    return s
}
