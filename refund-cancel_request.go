package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Yapılan ödemenin iadesi için kullanılır.
func (i *IyzipayClient) RefundPaymentV1(req *RefundPaymentV1Request) (response RefundPaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.RefundV1URI)
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

// Yapılan ödemenin iadesi için kullanılır.
//
// v1 ile benzerdir sadece paymentTransactionId yerine paymentId kullanır.
func (i *IyzipayClient) RefundPaymentV2(req *RefundPaymentV2Request) (response RefundPaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.RefundV2URI)
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

// Yapılan ödemenin iptali için kullanılır.
func (i *IyzipayClient) CancelPayment(req *CancelPaymentRequest) (response CancelPaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.CancelURI)
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

// Yapılan ödemenin Iyzico hesabına iadesi için kullanılır.
func (i *IyzipayClient) RefundPaymentToIyzicoAccount(req *RefundPaymentIyzicoRequest) (response RefundPaymentIyzicoResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.RefundToIyzicoAccountURI)
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
