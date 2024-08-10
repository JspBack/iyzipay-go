package iyzipay

import "gopkg.in/go-playground/validator.v9"

type InquiryRequest struct {
	// Ödemenin ID'si. Tüccarların bu ödeme ID'sini sistemlerinde saklamaları gerekmektedir (bu ID, iptal talepleri için kullanılacaktır).
	//
	// zorunlu.
	PaymentId string `json:"paymentId" validate:"required"`

	// İşlemin gönderildiği IP adresi.
	//
	// zorunlu.
	Ip string `json:"ip" validate:"required,ip"`

	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil.
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// Dil (default: tr).
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Ödeme için kullanılan konuşma ID'si.
	//
	// zorunlu değil.
	PaymentConversationID string `json:"paymentConversationId" validate:"omitempty"`
}

type InquiryResponse Non3DSPaymentResponse

func (r InquiryRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
