package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
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
func (i *IyzipayClient) InitializeSubscriptionWithCheckoutForm(req InitSubscriptionWithCheckoutFormRequest) (response InitSubscriptionWithCheckoutFormResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	initCFURI := utils.SubscriptionCheckoutFormURI + "/initialize"
	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, initCFURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// CheckoutForm ile Abonelik sorgulama isteği yapmak için kullanılır.
//
// Ödemenin tamamlanması gerekli.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *IyzipayClient) InquirySubscriptionWithCheckoutForm(req InquirySubscriptionWithCheckoutFormRequest) (response SubscriptionPaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	initCFURI := utils.SubscriptionCheckoutFormURI + "/" + req.Token
	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, initCFURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// 3D Secure olmayan Abonelik başlatma isteği yapmak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *IyzipayClient) InitializeSubscriptionWithNonTDS(req InitSubscriptionWithNonTDSRequest) (response SubscriptionPaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.SubscriptionNonTDSURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}
