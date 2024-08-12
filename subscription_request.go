package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Abonelik ürünü oluşturmak için kullanılır.
//
// Şuan hata dönmektedir.
func (i iyzipayClient) CreatSubscriptionProduct(req CreateSubscriptionProductRequest) (response SubscriptionProductResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.SubscriptionProductURI)
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

// Abonelik ürünü güncellemek için kullanılır.
//
// Şuan hata dönmektedir.
func (i iyzipayClient) UpdateSubscriptionProduct(req UpdateSubscriptionProductRequest) (response SubscriptionProductResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	UpdatedPath := utils.SubscriptionProductURI + "/" + req.ProductReferenceCode

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, UpdatedPath)
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

// Abonelik ürünü silmek için kullanılır.
//
// Şuan hata dönmektedir.
func (i iyzipayClient) DeleteSubscriptionProduct(req DeleteSubscriptionProductRequest) (response SubscriptionProductResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	// Sanırım bu şekilde kullanılması lazım ?
	DeletedPath := utils.SubscriptionProductURI + "/" + req.ProductReferenceCode

	httpresp, err := utils.DoRequest(requestData, i.client, "DELETE", i.baseURI, i.apiKey, i.apiSecret, DeletedPath)
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

func (i iyzipayClient) GetSubscriptionProductDetail(req GetSubscriptionProductDetailRequest) (response GetSubscriptionProductDetailResponse, err error) {
	return response, nil
}

func (i iyzipayClient) GetSubscriptionProductList(req GetSubscriptionProductListRequest) (response GetSubscriptionProductListResponse, err error) {
	return response, nil
}
