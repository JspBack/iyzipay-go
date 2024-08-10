package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

func (i Iyzipay) BinControlRequest(req *BinRequest) (response BinResponse, err error) {
	if err = req.validate(); err != nil {
		return BinResponse{}, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return BinResponse{}, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, i.baseURI, i.apiKey, i.apiSecret, utils.BinControlURI)
	if err != nil {
		return BinResponse{}, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil {
		return BinResponse{}, errors.New("failed to unmarshal response")
	}

	if response.Status == "failure" {
		var respError errorModel
		err = json.Unmarshal(httpresp, &respError)
		if err != nil {
			return BinResponse{}, errors.New("failed to unmarshal response")
		}
		return BinResponse{}, errors.New(
			"{" +
				"Status: " + respError.Status + "," +
				"ErrorCode: " + respError.ErrorCode + "," +
				"ErrorMessage: " + respError.ErrorMessage + "," +
				"Locale: " + respError.Locale + "," +
				"SystemTime: " + string(rune(respError.SystemTime)) + "," +
				"ConversationID: " + respError.ConversationID +
				"}",
		)
	}

	return response, nil
}
