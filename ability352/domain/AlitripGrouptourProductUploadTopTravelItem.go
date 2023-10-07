package domain

import (
        "topsdk/util"
    )

type AlitripGrouptourProductUploadTopTravelItem struct {
    /*
        itemId     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        outProductId     */
    OutProductId  *string `json:"out_product_id,omitempty" `

    /*
        modified     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        extend     */
    Extend  *string `json:"extend,omitempty" `

}

func (s *AlitripGrouptourProductUploadTopTravelItem) SetItemId(v int64) *AlitripGrouptourProductUploadTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *AlitripGrouptourProductUploadTopTravelItem) SetOutProductId(v string) *AlitripGrouptourProductUploadTopTravelItem {
    s.OutProductId = &v
    return s
}
func (s *AlitripGrouptourProductUploadTopTravelItem) SetModified(v util.LocalTime) *AlitripGrouptourProductUploadTopTravelItem {
    s.Modified = &v
    return s
}
func (s *AlitripGrouptourProductUploadTopTravelItem) SetExtend(v string) *AlitripGrouptourProductUploadTopTravelItem {
    s.Extend = &v
    return s
}
