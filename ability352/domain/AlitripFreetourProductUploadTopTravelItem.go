package domain

import (
        "topsdk/util"
    )

type AlitripFreetourProductUploadTopTravelItem struct {
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
        扩展信息     */
    Extend  *string `json:"extend,omitempty" `

}

func (s *AlitripFreetourProductUploadTopTravelItem) SetItemId(v int64) *AlitripFreetourProductUploadTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *AlitripFreetourProductUploadTopTravelItem) SetOutProductId(v string) *AlitripFreetourProductUploadTopTravelItem {
    s.OutProductId = &v
    return s
}
func (s *AlitripFreetourProductUploadTopTravelItem) SetModified(v util.LocalTime) *AlitripFreetourProductUploadTopTravelItem {
    s.Modified = &v
    return s
}
func (s *AlitripFreetourProductUploadTopTravelItem) SetExtend(v string) *AlitripFreetourProductUploadTopTravelItem {
    s.Extend = &v
    return s
}
