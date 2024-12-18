package utils

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Random string oluşturmak için kullanılan karakterler
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits

	RPst = "POST"   // POST methodu
	RGet = "GET"    // GET methodu
	RDel = "DELETE" // DELETE methodu
	RPut = "PUT"    // PUT methodu
	RPtc = "PATCH"  // PATCH methodu

	SandboxAPIURL = "https://sandbox-api.iyzipay.com" // Sandbox API URL
	APIURL        = "https://api.iyzipay.com"         // API URL

	NonTDSURI              = "/payment/auth"                                      // 3DS olmayan ödeme isteği için URI yolu
	InquiryURI             = "/payment/detail"                                    // Ödeme sorgulama isteği için URI yolu
	BinCheckURI            = "/payment/bin/check"                                 // Bin kontrol isteği için URI yolu
	InstallmentCheckURI    = "/payment/iyzipos/installment"                       // Taksit kontrol isteği için URI yolu
	TDSInitilizeURI        = "/payment/3dsecure/initialize"                       // 3DS ödeme başlatma isteği için URI yolu
	TDSFinalizeURI         = "/payment/3dsecure/auth"                             // 3DS ödeme isteği için URI yolu
	InitPWIURI             = "/payment/pay-with-iyzico/initialize"                // PWI başlatma isteği için URI yolu
	CheckPWIURI            = "/payment/iyzipos/checkoutform/auth/ecom/detail"     // PWI ödeme isteği için URI yolu
	CheckoutFormURI        = "/payment/iyzipos/checkoutform/initialize/auth/ecom" // Checkout form ödeme isteği için URI yolu
	CheckoutFormInquiryURI = "/payment/iyzipos/checkoutform/auth/ecom/detail"     // Checkout form ödeme sorgulama isteği için URI yolu

	SubMerchantURI        = "/onboarding/submerchant"          // Alt üye işyeri oluşturma isteği için URI yolu
	SubMerchantInquiryURI = "/onboarding/submerchant/detail"   // Alt üye işyeri sorgulama isteği için URI yolu
	SubMerchantProductURI = "/payment/item"                    // Alt üye işyeri ürün oluşturma isteği için URI yolu
	ApproveProductURI     = "/payment/iyzipos/item/approve"    // Ürün onaylama isteği için URI yolu
	DisapproveProductURI  = "/payment/iyzipos/item/disapprove" // Ürün onay kaldırma isteği için URI yolu

	SubscriptionProductURI       = "/v2/subscription/products"      // Abonelik ürün oluşturma, düzenleme ve silme isteği için URI yolu
	SubscriptionPlanURI          = "/v2/subscription/pricing-plans" // Abonelik ürün fiyatlandırma planı oluşturma, düzenleme ve silme isteği için URI yolu
	SubscriptionCheckoutFormURI  = "/v2/subscription/checkoutform/" // Abonelik ödeme başlatma isteği için URI yolu
	SubscriptionNonTDSURI        = "/v2/subscription/initialize"    // Abonelik ödeme isteği için URI yolu
	SubscriptionSubscriptionsURI = "/v2/subscription/subscriptions" // Abonelik sorgulama isteği için URI yolu
	SubscriptionSubscribersURI   = "/v2/subscription/customers"     // Abonelik müşteri oluşturma ve sorgulama isteği için URI yolu

	CardStorageURI  = "/cardstorage/card"  // Kart saklama isteği için URI yolu
	CardsStorageURI = "/cardstorage/cards" // Kart saklama sorgulama isteği için URI yolu

	RefundV1URI              = "/payment/refund"               // İade isteği için URI yolu
	RefundV2URI              = "/v2/payment/refund"            // İade isteği için URI yolu
	CancelURI                = "/payment/cancel"               // İptal isteği için URI yolu
	RefundToIyzicoAccountURI = "/payment/iyzipos/item/approve" // Iyzico hesabına iade isteği için URI yolu

	IyzilinkURI = "/v2/iyzilink/products" // Iyzilink oluşturma isteği için URI yolu

	A2AListURI    = "/v2/in-store/user-info/list" // App2App kullanıcı listeleme işlemi için URI yolu
	A2ACreateURI  = "/v2/in-store/payment"        // App2App ödeme oluşturma işlemi için URI yolu
	A2ARefundURI  = "/v2/in-store/payment/refund" // App2App iade işlemi için URI yolu
	A2AInquiryURI = "/v2/in-store/crypt/decrypt"  // App2App ödeme sorgulama işlemi için URI yolu

)

var (
	// 3DS işlemi için desteklenen kart tipleri
	TDSAcceptedCardAssociations = map[string]bool{
		"TROY":        true,
		"MASTER_CARD": true,
		"MASTERCARD":  true,
		"VISA":        true,
		"AMEX":        true,
	}

	// Checkout form için desteklenen form tipleri
	AcceptedCFTypes = map[string]bool{
		"iframe":     true,
		"responsive": true,
		"popup":      true,
	}
)
