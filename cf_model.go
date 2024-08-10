package iyzipay

type CFRequest struct {
	InitPWIRequest `validate:"dive"`

	// Ödeme kanalı. Geçerli değerler SHOPIFY, MAGENTO, PRESTASHOP, WOOCOMMERCE, OPENCART
	//
	// zorunlu değil.
	PaymentSource string `json:"paymentSource" validate:"omitempty,oneof=SHOPIFY MAGENTO PRESTASHOP WOOCOMMERCE OPENCART"`
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

type CFInquiryResponse struct{}
