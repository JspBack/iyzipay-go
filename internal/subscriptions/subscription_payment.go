package subscriptions

import (
	"context"
	"net/http"
	"net/url"
)

// CheckoutForm ile Abonelik başlatma isteği yapmak için kullanılır.
//
// Abonelik süreci her zaman için ACTIVE veya PENDING durumu ile başlar. Eğer durum PENDING ise veya durum ACTIVE ancak ödeme planında bir deneme süresi belirtilmişse, iyzico abonelik isteğinde sadece kartın validasyonunu gerçekleştirir. Kart validasyonu 1 TL’lik bir çekim ve akabinde iade ile gerçekleşir. Bunun dışında herhangi bir işlem veya ödeme gerçekleşmez.
//
// Istekten dönen script etiketi ile birlikte bu 2 divi kullanabilirsiniz:
//
// <div id="iyzipay-checkout-form" class="responsive"></div>
//
// <div id="iyzipay-checkout-form" class="popup"></div>
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) InitializeSubscriptionWithCheckoutForm(ctx context.Context, req InitSubscriptionWithCheckoutFormRequest) (response InitSubscriptionWithCheckoutFormResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	initCFURI := subscriptionCheckoutFormURI + "/initialize"
	err = i.Transport.Do(ctx, http.MethodPost, initCFURI, req, &response)
	return response, err
}

// CheckoutForm ile Abonelik sorgulama isteği yapmak için kullanılır.
//
// Ödemenin tamamlanması gerekli.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) InquirySubscriptionWithCheckoutForm(ctx context.Context, req InquirySubscriptionWithCheckoutFormRequest) (response SubscriptionPaymentResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	initCFURI := subscriptionCheckoutFormURI + "/" + url.PathEscape(req.Token)
	err = i.Transport.Do(ctx, http.MethodPost, initCFURI, req, &response)
	return response, err
}

// 3D Secure olmayan Abonelik başlatma isteği yapmak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) InitializeSubscriptionWithNonTDS(ctx context.Context, req InitSubscriptionWithNonTDSRequest) (response SubscriptionPaymentResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, subscriptionNonTDSURI, req, &response)
	return response, err
}
