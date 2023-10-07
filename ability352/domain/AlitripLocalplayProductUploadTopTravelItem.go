package domain

import (
        "topsdk/util"
    )

type AlitripLocalplayProductUploadTopTravelItem struct {
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

func (s *AlitripLocalplayProductUploadTopTravelItem) SetItemId(v int64) *AlitripLocalplayProductUploadTopTravelItem {
    s.ItemId = &v
    return s
}
func (s *AlitripLocalplayProductUploadTopTravelItem) SetOutProductId(v string) *AlitripLocalplayProductUploadTopTravelItem {
    s.OutProductId = &v
    return s
}
func (s *AlitripLocalplayProductUploadTopTravelItem) SetModified(v util.LocalTime) *AlitripLocalplayProductUploadTopTravelItem {
    s.Modified = &v
    return s
}
func (s *AlitripLocalplayProductUploadTopTravelItem) SetExtend(v string) *AlitripLocalplayProductUploadTopTravelItem {
    s.Extend = &v
    return s
}
