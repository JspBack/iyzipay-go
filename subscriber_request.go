package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Abone bilgilerini güncellemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *IyzipayClient) UpdateSubscriber(req UpdateSubscriberRequest) (response SubscriberResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	updateSbscrbrURI := utils.SubscriptionSubscribersURI + "/" + req.CustomerReferenceCode
	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, updateSbscrbrURI)
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

// Abone bilgilerini sorgulamak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *IyzipayClient) GetSubscriberDetail(req GetSubscriberDetailRequest) (response SubscriberResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	getSbscrbrURI := utils.SubscriptionSubscribersURI + "/" + req.CustomerReferenceCode
	httpresp, err := utils.DoRequest(requestData, i.client, "GET", i.baseURI, i.apiKey, i.apiSecret, getSbscrbrURI)
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

// Abone listesini sorgulamak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *IyzipayClient) GetSubscriberList(req GetSubscriberListRequest) (response GetSubscriberListRequestResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "GET", i.baseURI, i.apiKey, i.apiSecret, utils.SubscriptionSubscribersURI)
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
