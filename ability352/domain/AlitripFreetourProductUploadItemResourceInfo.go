package domain


type AlitripFreetourProductUploadItemResourceInfo struct {
    /*
        对应的说明     */
    Desc  *string `json:"desc,omitempty" `

    /*
        1-保险 2-餐饮 3-租车 4-签证 5-购物点，预留  6-赠品，预留 7-券，预留  8-其他费用     */
    Type  *int64 `json:"type,omitempty" `

}

func (s *AlitripFreetourProductUploadItemResourceInfo) SetDesc(v string) *AlitripFreetourProductUploadItemResourceInfo {
    s.Desc = &v
    return s
}
func (s *AlitripFreetourProductUploadItemResourceInfo) SetType(v int64) *AlitripFreetourProductUploadItemResourceInfo {
    s.Type = &v
    return s
}
