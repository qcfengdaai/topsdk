package domain

import (
        "topsdk/util"
    )

type AlitripDaytoursProductUploadTopTravelItem struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商家自定义商品编码     */
    OutProductId  *string `json:"out_product_id,omitempty" `

    /*
        商品修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        扩展信息     */
    Extend  *string `json:"extend,omitempty" `

}

func (s *AlitripDaytoursProductUploadTopTravelItem) SetItemId(v int64) *AlitripDaytoursProductUploadTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *AlitripDaytoursProductUploadTopTravelItem) SetOutProductId(v string) *AlitripDaytoursProductUploadTopTravelItem {
    s.OutProductId = &v
    return s
}
func (s *AlitripDaytoursProductUploadTopTravelItem) SetModified(v util.LocalTime) *AlitripDaytoursProductUploadTopTravelItem {
    s.Modified = &v
    return s
}
func (s *AlitripDaytoursProductUploadTopTravelItem) SetExtend(v string) *AlitripDaytoursProductUploadTopTravelItem {
    s.Extend = &v
    return s
}
