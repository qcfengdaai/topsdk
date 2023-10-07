package domain


type AlitripGrouptourProductUploadGroupTourPackageInfo struct {
    /*
        行程信息，新发布商品必填，数据，可以一次提交多天的行程信息     */
    TripElementList  *[]AlitripGrouptourProductUploadGroupTourTripElement `json:"trip_element_list,omitempty" `

    /*
        返程交通，交通工具的类型必须与商品上的去程交通一致     */
    BackTrafficInfoList  *[]AlitripGrouptourProductUploadGroupTourTrafficInfo `json:"back_traffic_info_list,omitempty" `

    /*
        去程交通，交通工具的类型必须与商品上的去程交通一致     */
    GoTrafficInfoList  *[]AlitripGrouptourProductUploadGroupTourTrafficInfo `json:"go_traffic_info_list,omitempty" `

    /*
        新发布商品时必填。自费项目。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    SelfExplanation  *[]string `json:"self_explanation,omitempty" `

    /*
        新发布商品时必填。买家预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    OrderInfo  *[]string `json:"order_info,omitempty" `

    /*
        新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    FeeExclude  *[]string `json:"fee_exclude,omitempty" `

    /*
        新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔     */
    FeeInclude  *[]string `json:"fee_include,omitempty" `

    /*
        套餐对应的出发地，是商品出发地的子集     */
    FromLocations  *string `json:"from_locations,omitempty" `

    /*
        必填，跟团游套餐对应的商家编码     */
    OutProductId  *string `json:"out_product_id,omitempty" `

    /*
        新发布商品必填，跟团游套餐名     */
    PackageName  *string `json:"package_name,omitempty" `

    /*
        套餐操作类型(0:套餐覆盖修改,1:增加套餐,2:删除套餐)===默认为0=== defalutValue:0    */
    PackageOperation  *int64 `json:"package_operation,omitempty" `

    /*
        套餐级别行程天数，必填。     */
    TripDays  *int64 `json:"trip_days,omitempty" `

    /*
        套餐级别行程晚数，必填。     */
    AccomNights  *int64 `json:"accom_nights,omitempty" `

}

func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetTripElementList(v []AlitripGrouptourProductUploadGroupTourTripElement) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.TripElementList = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetBackTrafficInfoList(v []AlitripGrouptourProductUploadGroupTourTrafficInfo) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.BackTrafficInfoList = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetGoTrafficInfoList(v []AlitripGrouptourProductUploadGroupTourTrafficInfo) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.GoTrafficInfoList = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetSelfExplanation(v []string) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.SelfExplanation = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetOrderInfo(v []string) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.OrderInfo = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetFeeExclude(v []string) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.FeeExclude = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetFeeInclude(v []string) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.FeeInclude = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetFromLocations(v string) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.FromLocations = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetOutProductId(v string) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.OutProductId = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetPackageName(v string) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.PackageName = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetPackageOperation(v int64) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.PackageOperation = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetTripDays(v int64) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.TripDays = &v
    return s
}
func (s *AlitripGrouptourProductUploadGroupTourPackageInfo) SetAccomNights(v int64) *AlitripGrouptourProductUploadGroupTourPackageInfo {
    s.AccomNights = &v
    return s
}
