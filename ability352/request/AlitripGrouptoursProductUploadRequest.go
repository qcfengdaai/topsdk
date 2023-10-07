package request

import (
        "topsdk/ability352/domain"
        "topsdk/util"
    )

type AlitripGrouptoursProductUploadRequest struct {
    /*
        新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船     */
    GoTrafficType  *int64 `json:"go_traffic_type,omitempty" required:"false" `
    /*
        新发布商品时必填。旅游天数     */
    TripDay  *int64 `json:"trip_day,omitempty" required:"false" `
    /*
        可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0     */
    SubStock  *int64 `json:"sub_stock,omitempty" required:"false" `
    /*
        可选，手机端详情描述，xml格式，格式详见示例。     */
    WapDesc  *string `json:"wap_desc,omitempty" required:"false" `
    /*
        新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船     */
    BackTrafficType  *int64 `json:"back_traffic_type,omitempty" required:"false" `
    /*
        可选，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔     */
    Itineraries  *[]string `json:"itineraries,omitempty" required:"false" `
    /*
        PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。     */
    DescXml  *string `json:"desc_xml,omitempty" required:"false" `
    /*
        新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    FeeExclude  *[]string `json:"fee_exclude,omitempty" required:"false" `
    /*
        新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    OrderInfo  *[]string `json:"order_info,omitempty" required:"false" `
    /*
        可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认     */
    ConfirmTime  *int64 `json:"confirm_time,omitempty" required:"false" `
    /*
        新发布商品时必填。商品标题，30个中文字符以内     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        新发布商品时必填。参团线路类型。0 -目的地参团，1-为出发地参团     */
    RouteType  *int64 `json:"route_type,omitempty" required:"false" `
    /*
        特殊可选，当refund_type=1或7时，需要上传自定义退改内容。自定义退改规则，最多可含5组规则，每组规则间以英文逗号分隔。 1）当refund_type为1时格式为：a_b_num,b-1_c_num。含义：提前a天至提前b天发起退款，买家需支付num比例违约费。 2）当refund_type为7时格式为：a_b_num1_num2_0,b-1_c_num1_num2_0。含义：提前a天至提前b天发起退款，买家需支付num1比例违约费，卖家需支付num2比例违约费，最后一个数字代表是否节假日规则（0-不是，1-是）。特别注意，当refund_type为7时，自定义退改规则必须设置 n天以上违约规则 以及 行程当日违约规则，即第一组规则需要以-1_a_num1_num2_0或-1_a_num1_num2_1开头，且最后一组规则需要以0_0_num1_num2_0或0_0_num1_num2_1结尾。     */
    RefundRegulations  *[]string `json:"refund_regulations,omitempty" required:"false" `
    /*
        可选，跟团时的集合地点，列表中每一个元素对应一个集合地点     */
    GatherPlaces  *[]domain.AlitripGrouptoursProductUploadGatherPlaceInfo `json:"gather_places,omitempty" required:"false" `
    /*
        新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    FeeInclude  *[]string `json:"fee_include,omitempty" required:"false" `
    /*
        可选，资源确认类型。1-即时确认，2-二次确认。不传默认1     */
    ConfirmType  *int64 `json:"confirm_type,omitempty" required:"false" `
    /*
        可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）     */
    ItemCustomTag  *string `json:"item_custom_tag,omitempty" required:"false" `
    /*
        商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。     */
    OutProductId  *string `json:"out_product_id" required:"true" `
    /*
        行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。     */
    StructItineraries  *[]domain.AlitripGrouptoursProductUploadStructItinerary `json:"struct_itineraries,omitempty" required:"false" `
    /*
        新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址     */
    ToLocations  *string `json:"to_locations,omitempty" required:"false" `
    /*
        可选，旅游晚数，不传默认旅游天数-1     */
    TripNight  *int64 `json:"trip_night,omitempty" required:"false" `
    /*
        新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔     */
    PicUrls  *[]string `json:"pic_urls,omitempty" required:"false" `
    /*
        可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定     */
    ReserveLimit  *string `json:"reserve_limit,omitempty" required:"false" `
    /*
        可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0     */
    RefundType  *int64 `json:"refund_type,omitempty" required:"false" `
    /*
        新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”     */
    FromLocations  *string `json:"from_locations,omitempty" required:"false" `
    /*
        PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。     */
    DescHtml  *string `json:"desc_html,omitempty" required:"false" `
    /*
        可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔     */
    SubTitles  *[]string `json:"sub_titles,omitempty" required:"false" `
    /*
        可选，电子合同信息设置。     */
    ElectronContract  *domain.AlitripGrouptoursProductUploadElectronContract `json:"electron_contract,omitempty" required:"false" `
    /*
        可选，出行人模板id，预留，暂不支持     */
    TravellerTemplateId  *int64 `json:"traveller_template_id,omitempty" required:"false" `
    /*
        可选，是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”     */
    PurePlay  *int64 `json:"pure_play,omitempty" required:"false" `
    /*
        新发布商品时必填。是否出境游，0-不是，1-是。     */
    IsOverseasTour  *int64 `json:"is_overseas_tour,omitempty" required:"false" `
}

func (s *AlitripGrouptoursProductUploadRequest) SetGoTrafficType(v int64) *AlitripGrouptoursProductUploadRequest {
    s.GoTrafficType = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetTripDay(v int64) *AlitripGrouptoursProductUploadRequest {
    s.TripDay = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetSubStock(v int64) *AlitripGrouptoursProductUploadRequest {
    s.SubStock = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetWapDesc(v string) *AlitripGrouptoursProductUploadRequest {
    s.WapDesc = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetBackTrafficType(v int64) *AlitripGrouptoursProductUploadRequest {
    s.BackTrafficType = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetItineraries(v []string) *AlitripGrouptoursProductUploadRequest {
    s.Itineraries = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetDescXml(v string) *AlitripGrouptoursProductUploadRequest {
    s.DescXml = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetFeeExclude(v []string) *AlitripGrouptoursProductUploadRequest {
    s.FeeExclude = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetOrderInfo(v []string) *AlitripGrouptoursProductUploadRequest {
    s.OrderInfo = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetConfirmTime(v int64) *AlitripGrouptoursProductUploadRequest {
    s.ConfirmTime = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetTitle(v string) *AlitripGrouptoursProductUploadRequest {
    s.Title = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetRouteType(v int64) *AlitripGrouptoursProductUploadRequest {
    s.RouteType = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetRefundRegulations(v []string) *AlitripGrouptoursProductUploadRequest {
    s.RefundRegulations = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetGatherPlaces(v []domain.AlitripGrouptoursProductUploadGatherPlaceInfo) *AlitripGrouptoursProductUploadRequest {
    s.GatherPlaces = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetFeeInclude(v []string) *AlitripGrouptoursProductUploadRequest {
    s.FeeInclude = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetConfirmType(v int64) *AlitripGrouptoursProductUploadRequest {
    s.ConfirmType = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetItemCustomTag(v string) *AlitripGrouptoursProductUploadRequest {
    s.ItemCustomTag = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetOutProductId(v string) *AlitripGrouptoursProductUploadRequest {
    s.OutProductId = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetStructItineraries(v []domain.AlitripGrouptoursProductUploadStructItinerary) *AlitripGrouptoursProductUploadRequest {
    s.StructItineraries = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetToLocations(v string) *AlitripGrouptoursProductUploadRequest {
    s.ToLocations = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetTripNight(v int64) *AlitripGrouptoursProductUploadRequest {
    s.TripNight = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetPicUrls(v []string) *AlitripGrouptoursProductUploadRequest {
    s.PicUrls = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetItemId(v int64) *AlitripGrouptoursProductUploadRequest {
    s.ItemId = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetReserveLimit(v string) *AlitripGrouptoursProductUploadRequest {
    s.ReserveLimit = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetRefundType(v int64) *AlitripGrouptoursProductUploadRequest {
    s.RefundType = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetFromLocations(v string) *AlitripGrouptoursProductUploadRequest {
    s.FromLocations = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetDescHtml(v string) *AlitripGrouptoursProductUploadRequest {
    s.DescHtml = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetSubTitles(v []string) *AlitripGrouptoursProductUploadRequest {
    s.SubTitles = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetElectronContract(v domain.AlitripGrouptoursProductUploadElectronContract) *AlitripGrouptoursProductUploadRequest {
    s.ElectronContract = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetTravellerTemplateId(v int64) *AlitripGrouptoursProductUploadRequest {
    s.TravellerTemplateId = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetPurePlay(v int64) *AlitripGrouptoursProductUploadRequest {
    s.PurePlay = &v
    return s
}
func (s *AlitripGrouptoursProductUploadRequest) SetIsOverseasTour(v int64) *AlitripGrouptoursProductUploadRequest {
    s.IsOverseasTour = &v
    return s
}

func (req *AlitripGrouptoursProductUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoTrafficType != nil) {
        paramMap["go_traffic_type"] = *req.GoTrafficType
    }
    if(req.TripDay != nil) {
        paramMap["trip_day"] = *req.TripDay
    }
    if(req.SubStock != nil) {
        paramMap["sub_stock"] = *req.SubStock
    }
    if(req.WapDesc != nil) {
        paramMap["wap_desc"] = *req.WapDesc
    }
    if(req.BackTrafficType != nil) {
        paramMap["back_traffic_type"] = *req.BackTrafficType
    }
    if(req.Itineraries != nil) {
        paramMap["itineraries"] = util.ConvertBasicList(*req.Itineraries)
    }
    if(req.DescXml != nil) {
        paramMap["desc_xml"] = *req.DescXml
    }
    if(req.FeeExclude != nil) {
        paramMap["fee_exclude"] = util.ConvertBasicList(*req.FeeExclude)
    }
    if(req.OrderInfo != nil) {
        paramMap["order_info"] = util.ConvertBasicList(*req.OrderInfo)
    }
    if(req.ConfirmTime != nil) {
        paramMap["confirm_time"] = *req.ConfirmTime
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.RouteType != nil) {
        paramMap["route_type"] = *req.RouteType
    }
    if(req.RefundRegulations != nil) {
        paramMap["refund_regulations"] = util.ConvertBasicList(*req.RefundRegulations)
    }
    if(req.GatherPlaces != nil) {
        paramMap["gather_places"] = util.ConvertStructList(*req.GatherPlaces)
    }
    if(req.FeeInclude != nil) {
        paramMap["fee_include"] = util.ConvertBasicList(*req.FeeInclude)
    }
    if(req.ConfirmType != nil) {
        paramMap["confirm_type"] = *req.ConfirmType
    }
    if(req.ItemCustomTag != nil) {
        paramMap["item_custom_tag"] = *req.ItemCustomTag
    }
    if(req.OutProductId != nil) {
        paramMap["out_product_id"] = *req.OutProductId
    }
    if(req.StructItineraries != nil) {
        paramMap["struct_itineraries"] = util.ConvertStructList(*req.StructItineraries)
    }
    if(req.ToLocations != nil) {
        paramMap["to_locations"] = *req.ToLocations
    }
    if(req.TripNight != nil) {
        paramMap["trip_night"] = *req.TripNight
    }
    if(req.PicUrls != nil) {
        paramMap["pic_urls"] = util.ConvertBasicList(*req.PicUrls)
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.ReserveLimit != nil) {
        paramMap["reserve_limit"] = *req.ReserveLimit
    }
    if(req.RefundType != nil) {
        paramMap["refund_type"] = *req.RefundType
    }
    if(req.FromLocations != nil) {
        paramMap["from_locations"] = *req.FromLocations
    }
    if(req.DescHtml != nil) {
        paramMap["desc_html"] = *req.DescHtml
    }
    if(req.SubTitles != nil) {
        paramMap["sub_titles"] = util.ConvertBasicList(*req.SubTitles)
    }
    if(req.ElectronContract != nil) {
        paramMap["electron_contract"] = util.ConvertStruct(*req.ElectronContract)
    }
    if(req.TravellerTemplateId != nil) {
        paramMap["traveller_template_id"] = *req.TravellerTemplateId
    }
    if(req.PurePlay != nil) {
        paramMap["pure_play"] = *req.PurePlay
    }
    if(req.IsOverseasTour != nil) {
        paramMap["is_overseas_tour"] = *req.IsOverseasTour
    }
    return paramMap
}

func (req *AlitripGrouptoursProductUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}