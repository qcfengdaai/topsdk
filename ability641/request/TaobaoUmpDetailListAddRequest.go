package request


type TaobaoUmpDetailListAddRequest struct {
    /*
        营销活动id。     */
    ActId  *int64 `json:"act_id" required:"true" `
    /*
        营销详情的列表。此列表由detail的json字符串组成。最多插入为10个。     */
    Details  *string `json:"details" required:"true" `
}

func (s *TaobaoUmpDetailListAddRequest) SetActId(v int64) *TaobaoUmpDetailListAddRequest {
    s.ActId = &v
    return s
}
func (s *TaobaoUmpDetailListAddRequest) SetDetails(v string) *TaobaoUmpDetailListAddRequest {
    s.Details = &v
    return s
}

func (req *TaobaoUmpDetailListAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    if(req.Details != nil) {
        paramMap["details"] = *req.Details
    }
    return paramMap
}

func (req *TaobaoUmpDetailListAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}