package domain


type AlitripFreetourProductUploadFreeTourPackageInfo struct {
    /*
        必须传，套餐对应的商家编码     */
    OutProductId  *string `json:"out_product_id,omitempty" `

    /*
        包含元素-景区门票，如果该套餐包含景区门票，则需要传这个参数     */
    FreeTourScenicInfoList  *[]AlitripFreetourProductUploadFreeTourScenicInfo `json:"free_tour_scenic_info_list,omitempty" `

    /*
        必填，套餐名称     */
    PackageName  *string `json:"package_name,omitempty" `

    /*
        发布商品时必填。套餐对应的目的地，套餐对应的目的地必须是商品的目的地的子集，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称 ）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性 批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124>）会根据映射关系自动转换成飞猪标准地址     */
    ToLocations  *string `json:"to_locations,omitempty" `

    /*
        返程交通，交通工具必须与商品上的返程交通一致     */
    BackTrafficInfoList  *[]AlitripFreetourProductUploadFreeTourTrafficInfo `json:"back_traffic_info_list,omitempty" `

    /*
        新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    FeeExclude  *[]string `json:"fee_exclude,omitempty" `

    /*
        门票说明     */
    ScenicDesc  *string `json:"scenic_desc,omitempty" `

    /*
        去程交通，交通工具的类型必须与商品上的去程交通一致     */
    GoTrafficInfoList  *[]AlitripFreetourProductUploadFreeTourTrafficInfo `json:"go_traffic_info_list,omitempty" `

    /*
        买家预定须知     */
    OrderInfo  *[]string `json:"order_info,omitempty" `

    /*
        当包含元素选择了餐饮，自驾租车，保险服务，其他费用时，该项必填     */
    ItemResourceInfoList  *[]AlitripFreetourProductUploadItemResourceInfo `json:"item_resource_info_list,omitempty" `

    /*
        套餐对应的出发地，是商品出发地的子集     */
    FromLocations  *string `json:"from_locations,omitempty" `

    /*
        新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    FeeInclude  *[]string `json:"fee_include,omitempty" `

    /*
        当包含元素选择了酒店的时候，该项必填     */
    FreeTourHotelInfoList  *[]AlitripFreetourProductUploadFreeTourHotelInfo `json:"free_tour_hotel_info_list,omitempty" `

    /*
        套餐操作类型，(0:套餐覆盖修改,1:增加套餐,2:删除套餐)===默认为0=== defalutValue:0    */
    PackageOperation  *int64 `json:"package_operation,omitempty" `

}

func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetOutProductId(v string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.OutProductId = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetFreeTourScenicInfoList(v []AlitripFreetourProductUploadFreeTourScenicInfo) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.FreeTourScenicInfoList = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetPackageName(v string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.PackageName = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetToLocations(v string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.ToLocations = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetBackTrafficInfoList(v []AlitripFreetourProductUploadFreeTourTrafficInfo) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.BackTrafficInfoList = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetFeeExclude(v []string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.FeeExclude = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetScenicDesc(v string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.ScenicDesc = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetGoTrafficInfoList(v []AlitripFreetourProductUploadFreeTourTrafficInfo) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.GoTrafficInfoList = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetOrderInfo(v []string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.OrderInfo = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetItemResourceInfoList(v []AlitripFreetourProductUploadItemResourceInfo) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.ItemResourceInfoList = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetFromLocations(v string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.FromLocations = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetFeeInclude(v []string) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.FeeInclude = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetFreeTourHotelInfoList(v []AlitripFreetourProductUploadFreeTourHotelInfo) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.FreeTourHotelInfoList = &v
    return s
}
func (s *AlitripFreetourProductUploadFreeTourPackageInfo) SetPackageOperation(v int64) *AlitripFreetourProductUploadFreeTourPackageInfo {
    s.PackageOperation = &v
    return s
}
