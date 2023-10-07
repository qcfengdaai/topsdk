package domain

import (
        "topsdk/util"
    )

type AlitripTravelGereralitemUpdateItemEleCertInfo struct {
    /*
        可选，售中自动退款比例，0~100。默认为0，即不支持售中自动退款；当值为1~100时表示售中自动退款的比例     */
    AutoRefundRate  *int64 `json:"auto_refund_rate,omitempty" `

    /*
        可选，过期自动退款比例，0~100。默认为0，即不支持过期自动退款比例；当值为1~100时表示过期自动退款的比例     */
    ExpiredRefundRate  *int64 `json:"expired_refund_rate,omitempty" `

    /*
        殊必填（expiryDateType为1或2时必填），电子凭证 有效期 结束时间     */
    ExpiryDateEnd  *util.LocalTime `json:"expiry_date_end,omitempty" `

    /*
        特殊必填（expiryDateType为1时必填），电子凭证 有效期 开始时间     */
    ExpiryDateStart  *util.LocalTime `json:"expiry_date_start,omitempty" `

    /*
        必填，电子凭证有效期 过期类型。1：xxxx-xx-xx 到 xxxx-xx-xx； 2：购买成功日 至 xxxx-xx-xx； 3：购买成功 xx 天内有效     */
    ExpiryDateType  *int64 `json:"expiry_date_type,omitempty" `

    /*
        特殊必填（expiryDateType为3时必填），电子凭证 有效期 天数     */
    ExpiryDays  *int64 `json:"expiry_days,omitempty" `

    /*
        必填，核销门店库id     */
    PackageId  *int64 `json:"package_id,omitempty" `

}

func (s *AlitripTravelGereralitemUpdateItemEleCertInfo) SetAutoRefundRate(v int64) *AlitripTravelGereralitemUpdateItemEleCertInfo {
    s.AutoRefundRate = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateItemEleCertInfo) SetExpiredRefundRate(v int64) *AlitripTravelGereralitemUpdateItemEleCertInfo {
    s.ExpiredRefundRate = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateItemEleCertInfo) SetExpiryDateEnd(v util.LocalTime) *AlitripTravelGereralitemUpdateItemEleCertInfo {
    s.ExpiryDateEnd = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateItemEleCertInfo) SetExpiryDateStart(v util.LocalTime) *AlitripTravelGereralitemUpdateItemEleCertInfo {
    s.ExpiryDateStart = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateItemEleCertInfo) SetExpiryDateType(v int64) *AlitripTravelGereralitemUpdateItemEleCertInfo {
    s.ExpiryDateType = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateItemEleCertInfo) SetExpiryDays(v int64) *AlitripTravelGereralitemUpdateItemEleCertInfo {
    s.ExpiryDays = &v
    return s
}
func (s *AlitripTravelGereralitemUpdateItemEleCertInfo) SetPackageId(v int64) *AlitripTravelGereralitemUpdateItemEleCertInfo {
    s.PackageId = &v
    return s
}
