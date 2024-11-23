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

type A2ARefundPaymentRequest struct {
	// İptal/İade edilmek istenen tutar.
	//
	// zorunlu
	RefundAmount string `json:"refundAmount" validate:"required"`

	// Ödemenin iyzico tarafındaki işlem id'sidir.
	//
	// zorunlu
	PaymentId int `json:"paymentId" validate:"required"`

	// İlgili üye işyerine ait kullanıcının eposta bilgisidir.
	//
	// zorunlu
	Email string `json:"email" validate:"required,email"`

	// İşlem sonrası uygulamanıza geri dönülmesi sağlayan parametredir.
	//
	// zorunlu
	CallBackUrl string `json:"callBackUrl" validate:"required,url"`
}

type A2AInquiryPaymentRequest struct {
	// Url'de dönen şifrelenmiş değer.
	//
	// zorunlu
	Data string `json:"data" validate:"required"`

	// Ödeme başlatılırken kullanılan benzersiz id'dir.
	//
	// zorunlu
	PaymentSessionToken string `json:"paymentSessionToken" validate:"required"`
}

type A2AInquiryPaymentResponse struct {
	// Yapılan isteğin sonucunu bildirir.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationID string `json:"conversationId"`

	// Tamamlan ceppos işlemi değerlerini içerir.
	InStoreCompleteOperation InStoreCompleteOperation `json:"inStoreCompleteOperation"`
}

type InStoreCompleteOperation struct {
	// Tamamlan ceppos işlemi
	Transaction Transaction `json:"transaction"`

	// Ödeme Oturumu
	PaymentSession *string `json:"paymentSession"`

	// İşlem Tipi
	TransactionType *string `json:"transactionType"`

	// Ödeme Başarısız Sonucu
	PaymentFailedResult *string `json:"paymentFailedResult"`
}

type Transaction struct {
	// Tamamlan ceppos işlemi tarihi
	TransactionDate string `json:"transactionDate"`

	// Tamamlan ceppos işlemi tarihinin ISO değeri
	TransactionDateISO string `json:"transactionDateISO"`

	// Tamamlan ceppos işleminin fiş bilgileri
	Receipt Receipt `json:"receipt"`

	// Ekran Mesajı
	ScreenMessage string `json:"screenMessage"`

	// Yerleşik Kayıt Numarası
	Rrn string `json:"rrn"`

	// Tahsil Edilen Tutar
	Amount float64 `json:"amount"`

	// Kur Kodu
	CurrencyCode string `json:"currencyCode"`

	// İşlem Kodu
	TransactionCode string `json:"transactionCode"`

	// Kayıt No
	RecordId int `json:"recordId"`

	// Maskeli Kart No
	MaskedPan string `json:"maskedPan"`

	// Otorizasyon Kodu
	AuthorizationCode string `json:"authorizationCode"`

	// Yanıt Kodu Değeri
	ResponseCodeValue string `json:"responseCodeValue"`

	// Durum Kodu
	StatusCode string `json:"statusCode"`

	// Sonuç Metni
	ResultText string `json:"resultText"`

	// İptal Edilebilme
	IsVoidable bool `json:"isVoidable"`

	// İade Edilebilme
	IsRefundable bool `json:"isRefundable"`

	// Banka Yanıt Kodu Değeri
	BankResponseCodeValue string `json:"bankResponseCodeValue"`

	// Issuer Bin No
	IssuerBin string `json:"issuerBin"`

	// Acquirer No
	AcquirerId string `json:"acquirerId"`

	// Temizleme tarihi
	ClearingDate *string `json:"clearingDate"`

	// İndirim Durumu
	HasDiscount bool `json:"hasDiscount"`

	// Taksit Durumu
	HasInstallment bool `json:"hasInstallment"`
}

type Receipt struct {
	// Tamamlan ceppos işleminin fiş bilgilerinin detayları
	Detail []DetailItem `json:"detail"`

	// Tamamlan ceppos işleminin fiş bilgileri
	PaymentBrandUrl string `json:"paymentBrandUrl"`

	// Üye İşyeri Logo Linki
	MerchantLogoURL *string `json:"merchantLogoURL"`

	// Kart Kuruluşu
	SchemaName string `json:"schemaName"`

	// Onay Durumu
	Approved bool `json:"approved"`

	// İade Onay Durumu
	RefundApproved bool `json:"refundApproved"`

	// İptal Onay Durumu
	VoidApproved bool `json:"voidApproved"`
}

type DetailItem struct {
	// Tamamlanan ceppos işleminin fiş bilgi başlığı
	Key string `json:"key"`

	// Tamamlanan ceppos işleminin fiş bilgi değeri
	Value string `json:"value"`
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

type A2ARefundPaymentResponse struct {
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
}

type A2ACreatePaymentResponse struct {
	A2ARefundPaymentResponse

	// Ödemenin iyzico tarafındaki işlem id'dir.
	PaymentId string `json:"paymentId"`
}

func (r *A2ACreatePaymentRequest) validate() error {
	validator := validator.New()
	return validator.Struct(r)
}

func (r *A2ARefundPaymentRequest) validate() error {
	validator := validator.New()
	return validator.Struct(r)
}

func (r *A2AInquiryPaymentRequest) validate() error {
	validator := validator.New()
	return validator.Struct(r)
}
