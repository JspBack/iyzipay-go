package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Checkout Fom ile ödeme isteği yapmak için kullanılır.
//
// Bu istekte dönecek olan html elementi için 3 farklı CFFormType belirleyebilirsiniz. "responsive", "popup", "iframe".
func (i *IyzipayClient) CheckoutFormPaymentRequest(req *CFRequest, CFFormType string) (response *CFResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.CheckoutFormURI)
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

	if utils.AcceptedCFTypes[CFFormType] {
		if CFFormType == "iframe" {
			response.PaymentPageUrl = response.PaymentPageUrl + "&iframe=true"
			return response, nil
		}
	}

	return response, errors.New("CFFormType is not valid")
}

// Checkout Form ile yapılan ödeme isteği için sorgulama yapmak için kullanılır.
//
// Ödemenin tamamlanması gerekli.
func (i *IyzipayClient) CheckoutFormPaymentInquiryRequest(req *CFInquiryRequest) (response *CFInquiryResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.CheckoutFormInquiryURI)
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

	return response, err
}
