package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo struct {
    /*
        电子凭证有效期 过期类型。1：xxxx-xx-xx 到 xxxx-xx-xx； 2：购买成功日 至 xxxx-xx-xx； 3：购买成功 xx 天内有效     */
    ExpiryDateType  *int64 `json:"expiry_date_type,omitempty" `

    /*
        电子凭证 有效期 开始时间     */
    ExpiryDateStart  *util.LocalTime `json:"expiry_date_start,omitempty" `

    /*
        电子凭证 有效期 结束时间     */
    ExpiryDateEnd  *util.LocalTime `json:"expiry_date_end,omitempty" `

    /*
        电子凭证 有效期 天数     */
    ExpiryDays  *int64 `json:"expiry_days,omitempty" `

    /*
        核销门店库id     */
    PackageId  *int64 `json:"package_id,omitempty" `

    /*
        售中自动退款比例     */
    AutoRefundRate  *int64 `json:"auto_refund_rate,omitempty" `

    /*
        过期自动退款比例     */
    ExpiredRefundRate  *int64 `json:"expired_refund_rate,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo) SetExpiryDateType(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo {
    s.ExpiryDateType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo) SetExpiryDateStart(v util.LocalTime) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo {
    s.ExpiryDateStart = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo) SetExpiryDateEnd(v util.LocalTime) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo {
    s.ExpiryDateEnd = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo) SetExpiryDays(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo {
    s.ExpiryDays = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo) SetPackageId(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo {
    s.PackageId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo) SetAutoRefundRate(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo {
    s.AutoRefundRate = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo) SetExpiredRefundRate(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelItemEleCertInfo {
    s.ExpiredRefundRate = &v
    return s
}
