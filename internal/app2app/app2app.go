package app2app

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// App2App kullanıcı listeleme isteği
// Kullanıcı listeleme servisi üye işyerine tanımlamış olan kullanıcılar listeler. Listelenen kullanıcıların ceppos kullanma yetki durumunu bu servisle öğrenebilirsiniz.
func (i *Service) App2AppListUsers(ctx context.Context) (response A2AListResponse, err error) {
	err = i.Transport.Do(ctx, http.MethodGet, a2AListURI, nil, &response)
	return response, err
}

// App2App ödeme oluşturma isteği
// Ödeme başlatma servisinin isteğinde tahsil edilecek tutar ve kullanıcı listeleme servisinde listelenen kullanıcının eposta bilgisi iletilmektedir.
func (i *Service) App2CreatePayment(ctx context.Context, req *A2ACreatePaymentRequest) (response A2ACreatePaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, a2ACreateURI, req, &response, http.Header{"X-Callback-Url": {req.CallBackUrl}})
	return response, err
}

// App2App iptal/iade isteği
// Başarılı gerçekleşen bir ödemeyi iptal veya iade etmek istediğinizde bu servisi kullanabilirsiniz.
//
// İptal / İade işlemi fiziki olarak gerçekleşmek zorundadır. Ceppos işlemleri iyzico paneli üzerinden iade edilememektedir.
func (i *Service) App2RefundPayment(ctx context.Context, req *A2ARefundPaymentRequest) (response A2ARefundPaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, a2ARefundURI, req, &response, http.Header{"X-Callback-Url": {req.CallBackUrl}})
	return response, err
}

// App2App sorgu isteği
// Ödeme işlemi sonrası callback parametre değerindeki url'e yönlendiğinde url'de data ve paymentSessionToken değerleri dönecektir. Bu değerleri kullanarak ödemeyi sorgulayabilirsiniz.
func (i *Service) App2InquiryPayment(ctx context.Context, req *A2AInquiryPaymentRequest) (response A2AInquiryPaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, a2AInquiryURI, req, &response)
	return response, err
}
