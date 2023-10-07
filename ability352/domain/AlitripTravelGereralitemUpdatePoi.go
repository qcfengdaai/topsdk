package domain


type AlitripTravelGereralitemUpdatePoi struct {
    /*
        POI对应ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        POI对应的名称     */
    Name  *string `json:"name,omitempty" `

}

func (s *AlitripTravelGereralitemUpdatePoi) SetId(v int64) *AlitripTravelGereralitemUpdatePoi {
    s.Id = &v
    return s
}
func (s *AlitripTravelGereralitemUpdatePoi) SetName(v string) *AlitripTravelGereralitemUpdatePoi {
    s.Name = &v
    return s
}
