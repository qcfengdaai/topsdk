package domain


type AlitripGrouptoursProductUploadGatherPlaceInfo struct {
    /*
        集合地点名称     */
    Name  *string `json:"name,omitempty" `

    /*
        集合地经纬度，英文逗号分隔。经度在前，纬度在后，英文逗号分隔最多支持6位小数，如120.111222,30.111222     */
    Poi  *string `json:"poi,omitempty" `

    /*
        AMAP/GOOGLE/OTHERS。高德（AMAP），GOOGLE，其他（OTHERS）     */
    PoiResource  *string `json:"poi_resource,omitempty" `

}

func (s *AlitripGrouptoursProductUploadGatherPlaceInfo) SetName(v string) *AlitripGrouptoursProductUploadGatherPlaceInfo {
    s.Name = &v
    return s
}
func (s *AlitripGrouptoursProductUploadGatherPlaceInfo) SetPoi(v string) *AlitripGrouptoursProductUploadGatherPlaceInfo {
    s.Poi = &v
    return s
}
func (s *AlitripGrouptoursProductUploadGatherPlaceInfo) SetPoiResource(v string) *AlitripGrouptoursProductUploadGatherPlaceInfo {
    s.PoiResource = &v
    return s
}
