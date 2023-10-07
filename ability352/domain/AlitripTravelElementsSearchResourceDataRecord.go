package domain


type AlitripTravelElementsSearchResourceDataRecord struct {
    /*
        映射id     */
    MappingId  *int64 `json:"mapping_id,omitempty" `

    /*
        项目id     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        次数     */
    Num  *int64 `json:"num,omitempty" `

    /*
        资源类型     */
    Type  *int64 `json:"type,omitempty" `

    /*
        资源id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        扩展属性，包括城市名称，外部编码等     */
    ValueMap  *string `json:"value_map,omitempty" `

}

func (s *AlitripTravelElementsSearchResourceDataRecord) SetMappingId(v int64) *AlitripTravelElementsSearchResourceDataRecord {
    s.MappingId = &v
    return s
}
func (s *AlitripTravelElementsSearchResourceDataRecord) SetProductId(v int64) *AlitripTravelElementsSearchResourceDataRecord {
    s.ProductId = &v
    return s
}
func (s *AlitripTravelElementsSearchResourceDataRecord) SetNum(v int64) *AlitripTravelElementsSearchResourceDataRecord {
    s.Num = &v
    return s
}
func (s *AlitripTravelElementsSearchResourceDataRecord) SetType(v int64) *AlitripTravelElementsSearchResourceDataRecord {
    s.Type = &v
    return s
}
func (s *AlitripTravelElementsSearchResourceDataRecord) SetId(v int64) *AlitripTravelElementsSearchResourceDataRecord {
    s.Id = &v
    return s
}
func (s *AlitripTravelElementsSearchResourceDataRecord) SetValueMap(v string) *AlitripTravelElementsSearchResourceDataRecord {
    s.ValueMap = &v
    return s
}
