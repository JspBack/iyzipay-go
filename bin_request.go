package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Kart taksit bilgileri gibi bilgilerin alınabilmesi için kullanılır.
func (i iyzipayClient) BinControlRequest(req *BinRequest) (response BinResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.BinControlURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}
