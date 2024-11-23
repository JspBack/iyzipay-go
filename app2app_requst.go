package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// App2App kullanıcı listeleme isteği
// Kullanıcı listeleme servisi üye işyerine tanımlamış olan kullanıcılar listeler. Listelenen kullanıcıların ceppos kullanma yetki durumunu bu servisle öğrenebilirsiniz.
func (i *App2AppClient) App2AppListUsers() (response A2AListResponse, err error) {
	httpresp, err := utils.A2ADoRequest(nil, i.client, utils.RGet, i.baseURI, i.apiKey, i.apiSecret, utils.A2AListURI, i.A2AapiKey, i.A2AsecretKey, i.merchantId, nil)
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

// App2App ödeme oluşturma isteği
// Ödeme başlatma servisinin isteğinde tahsil edilecek tutar ve kullanıcı listeleme servisinde listelenen kullanıcının eposta bilgisi iletilmektedir.
func (i *App2AppClient) App2CreatePayment(req *A2ACreatePaymentRequest) (response A2ACreatePaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	reqRef := requestData
	callRef := req.CallBackUrl

	httpresp, err := utils.A2ADoRequest(&reqRef, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.A2ACreateURI, i.A2AapiKey, i.A2AsecretKey, i.merchantId, &callRef)
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

// App2App iptal/iade isteği

// App2App sorgu isteği
