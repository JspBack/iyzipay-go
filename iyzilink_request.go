package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Iyzilink oluşturma isteği
func (i *IyzipayClient) IyzilinkCreate(req *IyzilinkCreateRequest) (response IyzilinkResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.IyzilinkURI)
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

// Iyzilink güncelleme isteği
func (i *IyzipayClient) IyzilinkUpdate(req *IyzilinkUpdateRequest) (response IyzilinkResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPut, i.baseURI, i.apiKey, i.apiSecret, utils.IyzilinkURI)
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

// Iyzilink durum güncelleme isteği
func (i *IyzipayClient) IyzilinkStatusUpdate(req *IyzilinkStatusUpdateRequest) (response IyzilinkStatusUpdateResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	updateStatusURI := utils.IyzilinkURI + "/" + req.Token + "/status/" + req.ProductStatus

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPtc, i.baseURI, i.apiKey, i.apiSecret, updateStatusURI)
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

// Iyzilink silme isteği
func (i *IyzipayClient) IyzilinkDelete(req *IyzilinkDeleteRequest) (response IyzilinkDeleteResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RDel, i.baseURI, i.apiKey, i.apiSecret, utils.IyzilinkURI)
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

// Iyzilink detay sorgulama isteği
// Token bazlı çalışmıyor nedense o yüzden clientta filtreleme ile çözdüm(şimdilik).[liste içerisinden token ile arama yapılıyor]
// Veri sayısı çok olursa bu yöntem performansı düşürebilir.
func (i *IyzipayClient) IyzilinkGetDetail(req *IyzilinkGetDetailRequest) (response IyzilinkGetDetailResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	resp, err := i.IyzilinkGetList(&IyzilinkGetListRequest{
		ConversationID: req.ConversationID,
		Locale:         req.Locale,
		Page:           req.Page,
		Count:          req.Count,
	})
	if err != nil {
		return response, err
	}

	for _, item := range resp.Data.Items {
		if item.Token == req.Token {
			response = IyzilinkGetDetailResponse{
				Status: "success",
				Locale: req.Locale,
				Data:   item,
			}
		} else {
			return response, errors.New("token not found")
		}
	}
	return response, nil

}

// Iyzilink listeleme isteği
func (i *IyzipayClient) IyzilinkGetList(req *IyzilinkGetListRequest) (response IyzilinkGetListResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RGet, i.baseURI, i.apiKey, i.apiSecret, utils.IyzilinkURI)
	if err != nil {
		return response, err
	}

	// var prettyJSON bytes.Buffer
	// err = json.Indent(&prettyJSON, httpresp, "", "  ")
	// if err != nil {
	// 	return response, err
	// }

	// log.Fatal(prettyJSON.String())

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
