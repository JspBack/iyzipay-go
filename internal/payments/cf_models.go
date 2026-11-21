package payments

type CFRequest struct {
	InitPWIRequest

	// Ödeme kanalı. Geçerli değerler SHOPIFY, MAGENTO, PRESTASHOP, WOOCOMMERCE, OPENCART
	//
	// zorunlu değil.
	PaymentSource string `json:"paymentSource" validate:"omitempty,oneof=SHOPIFY MAGENTO PRESTASHOP WOOCOMMERCE OPENCART"`

	// Kayıtlı kullanıcı ID’si Veya key'i ? (Dökümantasyon belirsiz)
	//
	// zorunlu değil.
	CardUserKey string `json:"cardUserKey" validate:"omitempty"`
}

type CFResponse struct {
	InitPWIResponse

	// Ödeme sayfası için dönen script tagi
	CheckoutFormContent string `json:"checkoutFormContent"`

	// Ödeme sayfası url bilgisi
	PaymentPageUrl string `json:"paymentPageUrl"`
}

type CFInquiryRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Checkout form için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür. Ödemenin sonucunu öğrenmek için zorunlu bir alandır.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir
	//
	// zorunlu değil.
	ConversationId string `json:"conversationId" validate:"omitempty"`
}

type CFInquiryResponse struct {
	Non3DSPaymentResponse

	//Bankadan dönen değerdir. Sadece ödeme başarısız ise ve işlem 3ds ile yapılmışsa bu değer döner. 0,2,3,4,5,6,7 değerlerini alabilir.
	MdStatus int `json:"mdStatus"`

	// Checkout form için oluşturulan tekil değer. Her istek için özel üretilir ve işyerine dönülür. Ödemenin sonucunu öğrenmek için zorunlu bir alandır.
	Token string `json:"token"`

	// Ödeme isteğinin durumunu gösterir. Success ise karttan ilgili tutar çekilmiştir. SUCCESS, FAILURE, INIT_THREEDS, CALLBACK_THREEDS, BKM_POS_SELECTED, CALLBACK_PECCO
	PaymentStatus string `json:"paymentStatus"`

	// Ödeme akışında üye işyerine başarılı ve hatalı sonucu bildirmek üzere alınan URL adresi.
	CallbackUrl string `json:"callbackUrl"`
}
