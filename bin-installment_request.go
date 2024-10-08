package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Kart bilgilerini öğrenmek için kullanılır.
//
// Talepte BIN numarası belirtilmemişse, mevcut tüm taksit seçenekleri görüntülenecektir. Ancak, BIN numarası verilirse, yalnızca ilgili ve o karta özel taksit seçeneklerini içerecektir.
func (i *IyzipayClient) BinControlRequest(req *BinRequest) (response BinResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.BinCheckURI)
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

// Kartın taksit bilgilerini sorgulamak için kullanılır.
//
// force3ds değeri 1 olarak dönerse, işlemin 3DS ile işlenmesi gerektiği anlamına gelir. 0 ise tercihlere göre işlem yapılabilir. İşyeri hesabında 3DS zorunlu olarak ayarlanmışsa, bu değer sürekli olarak 1 döndürür.
func (i *IyzipayClient) InstallmentControlRequest(req *InstallmentRequest) (response InstallmentResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.InstallmentCheckURI)
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
