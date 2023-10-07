package domain

import (
        "topsdk/util"
    )

type AlitripGrouptoursProductUploadTopTravelItem struct {
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

func (s *AlitripGrouptoursProductUploadTopTravelItem) SetItemId(v int64) *AlitripGrouptoursProductUploadTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *AlitripGrouptoursProductUploadTopTravelItem) SetOutProductId(v string) *AlitripGrouptoursProductUploadTopTravelItem {
    s.OutProductId = &v
    return s
}
func (s *AlitripGrouptoursProductUploadTopTravelItem) SetModified(v util.LocalTime) *AlitripGrouptoursProductUploadTopTravelItem {
    s.Modified = &v
    return s
}
func (s *AlitripGrouptoursProductUploadTopTravelItem) SetExtend(v string) *AlitripGrouptoursProductUploadTopTravelItem {
    s.Extend = &v
    return s
}
