package domain


type AlitripTravelPoiSearchPOI struct {
    /*
        POI对应id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        POI对应名称     */
    Name  *string `json:"name,omitempty" `

}

func (s *AlitripTravelPoiSearchPOI) SetId(v int64) *AlitripTravelPoiSearchPOI {
    s.Id = &v
    return s
}
func (s *AlitripTravelPoiSearchPOI) SetName(v string) *AlitripTravelPoiSearchPOI {
    s.Name = &v
    return s
}
