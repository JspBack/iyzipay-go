package payments

type InitTDSRequest struct {
	PaymentRequest

	// 3D Secure işlemi başlatıldığında, banka tarafından kullanıcıya gösterilecek sayfa URL'si.(https olmalıdır.)
	//
	// zorunlu.
	CallbackUrl string `json:"callbackUrl" validate:"required,httpsurl"`
}

type InitTDSResponse struct {
	// Servis yanıt sonucu (başarılı/başarısız)
	Status string `json:"status"`

	// Dil bilgisi
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int64 `json:"systemTime"`

	// Konuşma Id'si
	ConversationID string `json:"conversationId"`

	// 3D Secure işlem için html içeriği (base64 şifreli)
	ThreeDSHtmlContent string `json:"threeDSHtmlContent"`

	// İstek sonrası dönen İşlem Id'si
	PaymentID string `json:"paymentId"`

	// İstek sonrası dönen imza
	Signature string `json:"signature"`
}

type TDSPaymentRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin sipariş numarasıdır
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`

	// 3D dönüş bağlantı adresinizden alacağınız paymentid değeridir.
	//
	// zorunlu.
	PaymentId string `json:"paymentId" validate:"required"`

	// 3DInit isteği oluştururken kullandığınız conversationId değeridir.
	//
	// zorunlu.
	PaymentConversationId string `json:"paymentConversationId" validate:"required"`
}

type TDSPaymentResponse struct {
	Non3DSPaymentResponse

	//Bankadan dönen değerdir. Sadece ödeme başarısız ise ve işlem 3ds ile yapılmışsa bu değer döner. 0,2,3,4,5,6,7 değerlerini alabilir.
	MdStatus int `json:"mdStatus"`
}
