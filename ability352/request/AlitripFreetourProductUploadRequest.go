package request

import (
        "topsdk/ability352/domain"
        "topsdk/util"
    )

type AlitripFreetourProductUploadRequest struct {
    /*
        新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船，100-其他     */
    GoTrafficType  *int64 `json:"go_traffic_type,omitempty" required:"false" `
    /*
        套餐信息,数组类型，支持上传多个套餐信息     */
    FreeTourPackageInfo  *[]domain.AlitripFreetourProductUploadFreeTourPackageInfo `json:"free_tour_package_info,omitempty" required:"false" `
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
        新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他     */
    BackTrafficType  *int64 `json:"back_traffic_type,omitempty" required:"false" `
    /*
        PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。     */
    DescXml  *string `json:"desc_xml,omitempty" required:"false" `
    /*
        可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认     */
    ConfirmTime  *int64 `json:"confirm_time,omitempty" required:"false" `
    /*
        新发布商品时必填。商品标题，30个中文字符以内     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        （已废弃，建议使用refund_regulations_json）特殊可选，当refund_type=1或7时，需要上传自定义退改内容。自定义退改规则，最多可含5组规则，每组规则间以英文逗号分隔。 1）当refund_type为1时格式为：a_b_num,b-1_c_num。含义：提前a天至提前b天发起退款，买家需支付num比例违约费。 2）当refund_type为7时格式为：a_b_num1_num2_0,b-1_c_num1_num2_0。含义：提前a天至提前b天发起退款，买家需支付num1比例违约费，卖家需支付num2比例违约费，最后一个数字代表是否节假日规则（0-不是，1-是）。特别注意，当refund_type为7时，自定义退改规则必须设置 n天以上违约规则 以及 行程当日违约规则，即第一组规则需要以-1_a_num1_num2_0或-1_a_num1_num2_1开头，且最后一组规则需要以0_0_num1_num2_0或0_0_num1_num2_1结尾。     */
    RefundRegulations  *[]string `json:"refund_regulations,omitempty" required:"false" `
    /*
        可选，资源确认类型。1-即时确认，2-二次确认。不传默认1     */
    ConfirmType  *int64 `json:"confirm_type,omitempty" required:"false" `
    /*
        可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）     */
    ItemCustomTag  *string `json:"item_custom_tag,omitempty" required:"false" `
    /*
        商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。     */
    OutProductId  *string `json:"out_product_id,omitempty" required:"false" `
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
        PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。     */
    DescHtml  *string `json:"desc_html,omitempty" required:"false" `
    /*
        新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”     */
    FromLocations  *string `json:"from_locations,omitempty" required:"false" `
    /*
        可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔     */
    SubTitles  *[]string `json:"sub_titles,omitempty" required:"false" `
    /*
        可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。     */
    TravellerTemplateId  *int64 `json:"traveller_template_id,omitempty" required:"false" `
    /*
        新发布商品时必填。是否出境游，0-不是，1-是。     */
    IsOverseasTour  *int64 `json:"is_overseas_tour,omitempty" required:"false" `
    /*
        特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90     */
    RefundRegulationsJson  *string `json:"refund_regulations_json,omitempty" required:"false" `
    /*
        0：使用上传的套餐信息（free_tour_package_info）覆盖商品上原有的套餐信息（此时free_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（free_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0=== defalutValue��0    */
    PackageOperation  *int64 `json:"package_operation,omitempty" required:"false" `
    /*
        关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65     */
    SellerCids  *[]string `json:"seller_cids,omitempty" required:"false" `
    /*
        商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)     */
    SecondKill  *string `json:"second_kill,omitempty" required:"false" `
    /*
        是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false defalutValue��false    */
    HasDiscount  *bool `json:"has_discount,omitempty" required:"false" `
}

func (s *AlitripFreetourProductUploadRequest) SetGoTrafficType(v int64) *AlitripFreetourProductUploadRequest {
    s.GoTrafficType = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetFreeTourPackageInfo(v []domain.AlitripFreetourProductUploadFreeTourPackageInfo) *AlitripFreetourProductUploadRequest {
    s.FreeTourPackageInfo = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetTripDay(v int64) *AlitripFreetourProductUploadRequest {
    s.TripDay = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetSubStock(v int64) *AlitripFreetourProductUploadRequest {
    s.SubStock = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetWapDesc(v string) *AlitripFreetourProductUploadRequest {
    s.WapDesc = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetBackTrafficType(v int64) *AlitripFreetourProductUploadRequest {
    s.BackTrafficType = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetDescXml(v string) *AlitripFreetourProductUploadRequest {
    s.DescXml = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetConfirmTime(v int64) *AlitripFreetourProductUploadRequest {
    s.ConfirmTime = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetTitle(v string) *AlitripFreetourProductUploadRequest {
    s.Title = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetRefundRegulations(v []string) *AlitripFreetourProductUploadRequest {
    s.RefundRegulations = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetConfirmType(v int64) *AlitripFreetourProductUploadRequest {
    s.ConfirmType = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetItemCustomTag(v string) *AlitripFreetourProductUploadRequest {
    s.ItemCustomTag = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetOutProductId(v string) *AlitripFreetourProductUploadRequest {
    s.OutProductId = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetToLocations(v string) *AlitripFreetourProductUploadRequest {
    s.ToLocations = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetTripNight(v int64) *AlitripFreetourProductUploadRequest {
    s.TripNight = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetPicUrls(v []string) *AlitripFreetourProductUploadRequest {
    s.PicUrls = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetItemId(v int64) *AlitripFreetourProductUploadRequest {
    s.ItemId = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetReserveLimit(v string) *AlitripFreetourProductUploadRequest {
    s.ReserveLimit = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetRefundType(v int64) *AlitripFreetourProductUploadRequest {
    s.RefundType = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetDescHtml(v string) *AlitripFreetourProductUploadRequest {
    s.DescHtml = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetFromLocations(v string) *AlitripFreetourProductUploadRequest {
    s.FromLocations = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetSubTitles(v []string) *AlitripFreetourProductUploadRequest {
    s.SubTitles = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetTravellerTemplateId(v int64) *AlitripFreetourProductUploadRequest {
    s.TravellerTemplateId = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetIsOverseasTour(v int64) *AlitripFreetourProductUploadRequest {
    s.IsOverseasTour = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetRefundRegulationsJson(v string) *AlitripFreetourProductUploadRequest {
    s.RefundRegulationsJson = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetPackageOperation(v int64) *AlitripFreetourProductUploadRequest {
    s.PackageOperation = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetSellerCids(v []string) *AlitripFreetourProductUploadRequest {
    s.SellerCids = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetSecondKill(v string) *AlitripFreetourProductUploadRequest {
    s.SecondKill = &v
    return s
}
func (s *AlitripFreetourProductUploadRequest) SetHasDiscount(v bool) *AlitripFreetourProductUploadRequest {
    s.HasDiscount = &v
    return s
}

func (req *AlitripFreetourProductUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GoTrafficType != nil) {
        paramMap["go_traffic_type"] = *req.GoTrafficType
    }
    if(req.FreeTourPackageInfo != nil) {
        paramMap["free_tour_package_info"] = util.ConvertStructList(*req.FreeTourPackageInfo)
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
    if(req.DescXml != nil) {
        paramMap["desc_xml"] = *req.DescXml
    }
    if(req.ConfirmTime != nil) {
        paramMap["confirm_time"] = *req.ConfirmTime
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.RefundRegulations != nil) {
        paramMap["refund_regulations"] = util.ConvertBasicList(*req.RefundRegulations)
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
    if(req.DescHtml != nil) {
        paramMap["desc_html"] = *req.DescHtml
    }
    if(req.FromLocations != nil) {
        paramMap["from_locations"] = *req.FromLocations
    }
    if(req.SubTitles != nil) {
        paramMap["sub_titles"] = util.ConvertBasicList(*req.SubTitles)
    }
    if(req.TravellerTemplateId != nil) {
        paramMap["traveller_template_id"] = *req.TravellerTemplateId
    }
    if(req.IsOverseasTour != nil) {
        paramMap["is_overseas_tour"] = *req.IsOverseasTour
    }
    if(req.RefundRegulationsJson != nil) {
        paramMap["refund_regulations_json"] = *req.RefundRegulationsJson
    }
    if(req.PackageOperation != nil) {
        paramMap["package_operation"] = *req.PackageOperation
    }
    if(req.SellerCids != nil) {
        paramMap["seller_cids"] = util.ConvertBasicList(*req.SellerCids)
    }
    if(req.SecondKill != nil) {
        paramMap["second_kill"] = *req.SecondKill
    }
    if(req.HasDiscount != nil) {
        paramMap["has_discount"] = *req.HasDiscount
    }
    return paramMap
}

func (req *AlitripFreetourProductUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}