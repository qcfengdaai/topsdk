package request

import (
        "topsdk/ability352/domain"
        "topsdk/util"
    )

type AlitripTravelGereralitemUpdateRequest struct {
    /*
        必填，商品基本信息     */
    BaseInfo  *domain.AlitripTravelGereralitemUpdateBaseInfo `json:"base_info,omitempty" required:"false" `
    /*
        必填，预定规则结构。示例： [{ "rule_type": "fee_excluded", "rule_desc": "费用包含描述"},{ "rule_type": "fee_included", "rule_desc": "费用不含描述"},{ "rule_type": "order_info", "rule_desc": "预定须知描述"}]     */
    BookingRules  *[]domain.AlitripTravelGereralitemUpdateBookingRuleInfo `json:"booking_rules" required:"true" `
    /*
        更新sku信息，仅限日历商品使用     */
    DateSkuInfoList  *[]domain.AlitripTravelGereralitemUpdateDateSkuInfo `json:"date_sku_info_list,omitempty" required:"false" `
    /*
        新版电子凭证信息。如果传递了此参数，则sales_info中旧版电子凭证信息将被忽略，否则电子凭证信息将以旧版电子凭证参数为准。如果新、旧版参数都没传，则该商品不支持电子凭证     */
    ItemEleCertInfo  *domain.AlitripTravelGereralitemUpdateItemEleCertInfo `json:"item_ele_cert_info,omitempty" required:"false" `
    /*
        选填，退款规则结构     */
    ItemRefundInfo  *domain.AlitripTravelGereralitemUpdateItemRefundInfo `json:"item_refund_info,omitempty" required:"false" `
    /*
        更新sku信息，仅限非日历（普通）商品使用     */
    CommonSkuList  *[]domain.AlitripTravelGereralitemUpdateNoDateSkuInfo `json:"common_sku_list,omitempty" required:"false" `
    /*
        poi信息，个别类目必填，如演艺类目下场馆信息     */
    Poi  *domain.AlitripTravelGereralitemUpdatePoi `json:"poi,omitempty" required:"false" `
}

func (s *AlitripTravelGereralitemUpdateRequest) SetBaseInfo(v domain.AlitripTravelGereralitemUpdateBaseInfo) *AlitripTravelGereralitemUpdateRequest {
    s.BaseInfo = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateRequest) SetBookingRules(v []domain.AlitripTravelGereralitemUpdateBookingRuleInfo) *AlitripTravelGereralitemUpdateRequest {
    s.BookingRules = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateRequest) SetDateSkuInfoList(v []domain.AlitripTravelGereralitemUpdateDateSkuInfo) *AlitripTravelGereralitemUpdateRequest {
    s.DateSkuInfoList = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateRequest) SetItemEleCertInfo(v domain.AlitripTravelGereralitemUpdateItemEleCertInfo) *AlitripTravelGereralitemUpdateRequest {
    s.ItemEleCertInfo = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateRequest) SetItemRefundInfo(v domain.AlitripTravelGereralitemUpdateItemRefundInfo) *AlitripTravelGereralitemUpdateRequest {
    s.ItemRefundInfo = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateRequest) SetCommonSkuList(v []domain.AlitripTravelGereralitemUpdateNoDateSkuInfo) *AlitripTravelGereralitemUpdateRequest {
    s.CommonSkuList = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateRequest) SetPoi(v domain.AlitripTravelGereralitemUpdatePoi) *AlitripTravelGereralitemUpdateRequest {
    s.Poi = &v
    return s
}

func (req *AlitripTravelGereralitemUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BaseInfo != nil) {
        paramMap["base_info"] = util.ConvertStruct(*req.BaseInfo)
    }
    if(req.BookingRules != nil) {
        paramMap["booking_rules"] = util.ConvertStructList(*req.BookingRules)
    }
    if(req.DateSkuInfoList != nil) {
        paramMap["date_sku_info_list"] = util.ConvertStructList(*req.DateSkuInfoList)
    }
    if(req.ItemEleCertInfo != nil) {
        paramMap["item_ele_cert_info"] = util.ConvertStruct(*req.ItemEleCertInfo)
    }
    if(req.ItemRefundInfo != nil) {
        paramMap["item_refund_info"] = util.ConvertStruct(*req.ItemRefundInfo)
    }
    if(req.CommonSkuList != nil) {
        paramMap["common_sku_list"] = util.ConvertStructList(*req.CommonSkuList)
    }
    if(req.Poi != nil) {
        paramMap["poi"] = util.ConvertStruct(*req.Poi)
    }
    return paramMap
}

func (req *AlitripTravelGereralitemUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}