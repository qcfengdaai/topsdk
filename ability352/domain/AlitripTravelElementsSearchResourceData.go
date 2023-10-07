package domain


type AlitripTravelElementsSearchResourceData struct {
    /*
        资源列表     */
    Datas  *[]AlitripTravelElementsSearchResourceDataRecord `json:"datas,omitempty" `

    /*
        资源名称     */
    Name  *string `json:"name,omitempty" `

    /*
        资源类型     */
    Type  *int64 `json:"type,omitempty" `

}

func (s *AlitripTravelElementsSearchResourceData) SetDatas(v []AlitripTravelElementsSearchResourceDataRecord) *AlitripTravelElementsSearchResourceData {
    s.Datas = &v
    return s
}
func (s *AlitripTravelElementsSearchResourceData) SetName(v string) *AlitripTravelElementsSearchResourceData {
    s.Name = &v
    return s
}
func (s *AlitripTravelElementsSearchResourceData) SetType(v int64) *AlitripTravelElementsSearchResourceData {
    s.Type = &v
    return s
}
