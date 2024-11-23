package iyzipay

import "gopkg.in/go-playground/validator.v9"

type A2ACreatePaymentRequest struct {
	// Tahsil edilmek istenen tutar.
	//
	// zorunlu
	Amount string `json:"amount" validate:"required"`

	// İşlem yapılacak olan müşterinin eposta bilgisidir.
	//
	// zorunlu
	Email string `json:"email" validate:"required,email"`

	// Üye iş yeri tarafından iletilen ödeme kaynağı bilgisi.
	//
	// zorunlu değil
	PaymentSource string `json:"paymentSource" validate:"omitempty"`

	// İşlem sonrası uygulamanıza geri dönülmesi sağlayan parametredir.
	//
	// zorunlu
	CallBackUrl string `json:"callBackUrl" validate:"required,url"`
}

type A2AListResponse struct {
	// Yapılan isteğin sonucunu bildirir.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir..
	SystemTime int `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir..
	ConversationID string `json:"conversationId"`

	// İlgili üye işyerine ait kullanıcıları listeler.
	UserInfoList []UserInfo `json:"userInfoList"`
}

type UserInfo struct {
	// İlgili üye işyerine ait kullanıcının eposta bilgisidir.
	Email string `json:"email"`

	//İlgili üye işyerine ait kullanıcının ceppos kullanabilme yetkisini gösterir.
	CanPerformAction bool `json:"canPerformAction"`
}

type A2ACreatePaymentResponse struct {
	// Yapılan isteğin sonucunu bildirir.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir..
	SystemTime int `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir..
	ConversationID string `json:"conversationId"`

	// Ceppos app'ini uyandırmak için kullanılan link.
	DeepLinkUrl string `json:"deepLinkUrl"`

	// İşlem başlatılırken kullanılan benzersiz id'dir.
	PaymentSessionToken string `json:"paymentSessionToken"`

	// Ödemenin iyzico tarafındaki işlem id'dir.
	PaymentId string `json:"paymentId"`
}

func (r *A2ACreatePaymentRequest) validate() error {
	validator := validator.New()
	return validator.Struct(r)
}
