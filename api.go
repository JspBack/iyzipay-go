package iyzipay

import (
	"context"

	"github.com/JspBack/iyzipay-go/internal/app2app"
	"github.com/JspBack/iyzipay-go/internal/cards"
	"github.com/JspBack/iyzipay-go/internal/iyzilink"
	"github.com/JspBack/iyzipay-go/internal/marketplace"
	"github.com/JspBack/iyzipay-go/internal/payments"
	"github.com/JspBack/iyzipay-go/internal/subscriptions"
)

// 3D Secure olmayan ödeme işlemi için istek oluşturur
func (i *Client) CreatePayment(ctx context.Context, req *PaymentRequest) (Non3DSPaymentResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreatePayment(ctx, req)
}

// TROY, MASTERCARD, VISA ve AMEX markalı kartlarla yapılan işlemleri destekler.
// BONUS, WORLD, MAXIMUM, AXESS, CARDFINANS, PARAF, ADVANTAGE gibi taksit programlarına katılmış olan kartlara 2, 3, 6, 9 ve 12 taksit seçenekleri sunulmaktadır.
// Response da dönen paymentId saklanmalıdır.
//
// Bincontrol işlemi otomatik olarak yapılyor eğer bin kontrol yapmak istemiyorsanız binRequest parametresini false yapabilirsiniz.(Eğer kayıtlı kart ile ödeme yapılıyorsa bin kontrolü yapılmıyor.)
//
// Html içeriği otomatik olarak decode ediliyor eğer decode etmek istemiyorsanız htmlDecodeRequest parametresini false yapabilirsiniz.
func (i *Client) InitializeTDSPayment(ctx context.Context, req *InitTDSRequest, opts ...tdsOptions) (InitTDSResponse, string, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.InitializeTDSPayment(ctx, req, opts...)
}

// 3DS işlemlerinin tamamlanması için kullanılır.
//
// ÖNEMLİ: Size geri dönen htmlde 3D Secure işlemini tamamlamanız gerekmektedir.
func (i *Client) FinalizeTDSPayment(ctx context.Context, req *TDSPaymentRequest) (TDSPaymentResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.FinalizeTDSPayment(ctx, req)
}

type tdsOptions = payments.TDSOption

// WithBinRequest, otomatik BIN kontrolünü etkinleştirir veya kapatır.
func WithBinRequest(enabled bool) tdsOptions { return payments.WithBinRequest(enabled) }

// WithHtmlDecodeRequest, HTML içeriğinin otomatik çözülmesini belirler.
func WithHtmlDecodeRequest(enabled bool) tdsOptions { return payments.WithHtmlDecodeRequest(enabled) }

// İşlem sorgulama isteği
func (i *Client) PaymentInquiryRequest(ctx context.Context, req *InquiryRequest) (InquiryResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.PaymentInquiryRequest(ctx, req)
}

// Yapılan ödemenin iadesi için kullanılır.
func (i *Client) RefundPayment(ctx context.Context, req *RefundPaymentRequest) (RefundPaymentResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.RefundPayment(ctx, req)
}

// Yapılan ödemenin iadesi için kullanılır.
//
// v1 ile benzerdir sadece paymentTransactionId yerine paymentId kullanır.
func (i *Client) RefundPaymentV2(ctx context.Context, req *RefundPaymentV2Request) (RefundPaymentResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.RefundPaymentV2(ctx, req)
}

// Yapılan ödemenin iptali için kullanılır.
func (i *Client) CancelPayment(ctx context.Context, req *CancelPaymentRequest) (CancelPaymentResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CancelPayment(ctx, req)
}

// Yapılan ödemenin Iyzico hesabına iadesi için kullanılır.
func (i *Client) RefundPaymentToIyzicoAccount(ctx context.Context, req *RefundPaymentIyzicoRequest) (RefundPaymentIyzicoResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.RefundPaymentToIyzicoAccount(ctx, req)
}

// PWI ile ödeme başlatma isteği yapılır.
//
// Response içerisinde gelen token ve payWithIyzicoPageUrl saklanmalıdır.
func (i *Client) InitializePWIPaymentRequest(ctx context.Context, req *InitPWIRequest) (InitPWIResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.InitializePWIPaymentRequest(ctx, req)
}

// PWI ile gerçekleştirilen ödeme sorgulama isteği yapılır.
//
// Response içerisinde gelen paymentId saklanmalıdır.
func (i *Client) CheckPWIPaymentRequest(ctx context.Context, req *PWIInquiryRequest) (PWIInquiryResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CheckPWIPaymentRequest(ctx, req)
}

// Checkout Fom ile ödeme isteği yapmak için kullanılır.
//
// Bu istekte dönecek olan html elementi için 3 farklı CFFormType belirleyebilirsiniz. "responsive", "popup", "iframe".
func (i *Client) CheckoutFormPaymentRequest(ctx context.Context, req *CFRequest, CFFormType string) (*CFResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CheckoutFormPaymentRequest(ctx, req, CFFormType)
}

// Checkout Form ile yapılan ödeme isteği için sorgulama yapmak için kullanılır.
//
// Ödemenin tamamlanması gerekli.
func (i *Client) CheckoutFormPaymentInquiryRequest(ctx context.Context, req *CFInquiryRequest) (*CFInquiryResponse, error) {
	service := payments.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CheckoutFormPaymentInquiryRequest(ctx, req)
}

// Kart bilgilerini öğrenmek için kullanılır.
func (i *Client) BinControlRequest(ctx context.Context, req *BinRequest) (BinResponse, error) {
	service := cards.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.BinControlRequest(ctx, req)
}

// Kartın taksit bilgilerini sorgulamak için kullanılır.
//
// Talepte BIN numarası belirtilmemişse, mevcut tüm taksit seçenekleri görüntülenecektir. Ancak, BIN numarası verilirse, yalnızca ilgili ve o karta özel taksit seçeneklerini içerecektir.
//
// force3ds değeri 1 olarak dönerse, işlemin 3DS ile işlenmesi gerektiği anlamına gelir. 0 ise tercihlere göre işlem yapılabilir. İşyeri hesabında 3DS zorunlu olarak ayarlanmışsa, bu değer sürekli olarak 1 döndürür.
func (i *Client) InstallmentControlRequest(ctx context.Context, req *InstallmentRequest) (InstallmentResponse, error) {
	service := cards.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.InstallmentControlRequest(ctx, req)
}

// Kart saklama işlemi için yeni kullanıcı ile kart saklama isteği
//
// Response içerisinde dönen cardUserKey ve cardToken saklanmalıdır.
func (i *Client) CreateCardStorageNewUser(ctx context.Context, req *CardStorageNewUserRequest) (CardStorageNewUserResponse, error) {
	service := cards.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreateCardStorageNewUser(ctx, req)
}

// Kart saklama işlemi için mevcut kullanıcı ile kart saklama isteği
//
// Response içerisinde dönen cardToken saklanmalıdır.
func (i *Client) CreateCardStorageExUser(ctx context.Context, req *CardStorageExUserRequest) (CardStorageExUserResponse, error) {
	service := cards.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreateCardStorageExUser(ctx, req)
}

// Kart saklama işlemi için kart silme isteği
func (i *Client) DeleteCardStorageCard(ctx context.Context, req *CardStorageDeleteRequest) (CardStorageDeleteResponse, error) {
	service := cards.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.DeleteCardStorageCard(ctx, req)
}

func (i *Client) GetCardStorageCards(ctx context.Context, req *CardStorageRetrieveRequest) (CardStorageRetrieveResponse, error) {
	service := cards.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.GetCardStorageCards(ctx, req)
}

// Pazaryeri çözümünde, ödeme iyzico’dan geçtikten sonra, üye işyeri ödeme içinde yer alan kırılıma / ürüne onay verene dek para korumalı havuz hesapta bekletilir.
// Üye işyeri bu sürede ödemeyi iptal edebilir, ödemenin kırılımını iade edebilir ya da ürün alıcıya ulaştı ve işlem sorunsuz tamamlandıysa para transferi için ürüne onay verebilir veya verdiği ürün onayını geri çekebilir.
//
// Ürün onayı vermek için kullanılır.
func (i *Client) ApproveProduct(ctx context.Context, req *MarketplaceProductRequest) (MarketplaceProductResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.ApproveProduct(ctx, req)
}

// Ürün onayını geri çekmek için kullanılır.
func (i *Client) DisapproveProduct(ctx context.Context, req *MarketplaceProductRequest) (MarketplaceProductResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.DisapproveProduct(ctx, req)
}

// Bireysel Alt Üye Oluşturma İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *Client) CreateIndividualSubMerchant(ctx context.Context, req *IndividualSubMerchantRequest) (SubMerchantResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreateIndividualSubMerchant(ctx, req)
}

// Şahıs Şirketi Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *Client) CreatePrivateSubMerchant(ctx context.Context, req *PrivateSubMerchantRequest) (SubMerchantResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreatePrivateSubMerchant(ctx, req)
}

// Limited/Anonim Şirket Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *Client) CreateLimitedCompanySubMerchant(ctx context.Context, req *LimitedCompanySubMerchantRequest) (SubMerchantResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreateLimitedCompanySubMerchant(ctx, req)
}

// Bireysel Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Client) UpdateIndividualSubMerchant(ctx context.Context, req *IndividualSubMerchantUpdateRequest) (UpdateSubMerchantResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.UpdateIndividualSubMerchant(ctx, req)
}

// Şahıs Şirketi Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Client) UpdatePrivateSubMerchant(ctx context.Context, req *PrivateSubMerchantUpdateRequest) (UpdateSubMerchantResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.UpdatePrivateSubMerchant(ctx, req)
}

// Limited/Anonim Şirket Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Client) UpdateLimitedCompanySubMerchant(ctx context.Context, req *LimitedCompanySubMerchantUpdateRequest) (UpdateSubMerchantResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.UpdateLimitedCompanySubMerchant(ctx, req)
}

// Alt Üye Sorgulama İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Client) SubMerchantInquiry(ctx context.Context, req *SubMerchantInquiryRequest) (SubMerchantInquiryResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.SubMerchantInquiry(ctx, req)
}

func (i *Client) UpdateSubMerchantProduct(ctx context.Context, req *SubMerchantProductUpdateRequest) (SubMerchantProductUpdateResponse, error) {
	service := marketplace.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.UpdateSubMerchantProduct(ctx, req)
}

// Abonelik ürünü oluşturmak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) CreateSubscriptionProduct(ctx context.Context, req CreateSubscriptionProductRequest) (SubscriptionProductResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreateSubscriptionProduct(ctx, req)
}

// Abonelik ürünü güncellemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) UpdateSubscriptionProduct(ctx context.Context, req UpdateSubscriptionProductRequest) (SubscriptionProductResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.UpdateSubscriptionProduct(ctx, req)
}

// Abonelik ürünü silmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) DeleteSubscriptionProduct(ctx context.Context, req DeleteSubscriptionProductRequest) (SubscriptionProductResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.DeleteSubscriptionProduct(ctx, req)
}

// Abonelik ürünü detaylarını getirmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) GetSubscriptionProductDetail(ctx context.Context, req GetSubscriptionProductDetailRequest) (SubscriptionProductResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.GetSubscriptionProductDetail(ctx, req)
}

// Abonelik ürünlerini listelemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) GetSubscriptionProductList(ctx context.Context, req GetSubscriptionProductListRequest) (GetSubscriptionProductListResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.GetSubscriptionProductList(ctx, req)
}

// Abonelik ürünü için fiyatlandırma planı oluşturmak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) CreateSubscriptionPlan(ctx context.Context, req CreateSubscriptionPlanRequest) (SubscriptionPlanResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.CreateSubscriptionPlan(ctx, req)
}

// Abonelik ürünü için fiyatlandırma planı güncellemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) UpdateSubscriptionPlan(ctx context.Context, req UpdateSubscriptionPlanRequest) (SubscriptionPlanResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.UpdateSubscriptionPlan(ctx, req)
}

// Abonelik ürünü için fiyatlandırma planı silmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) DeleteSubscriptionPlan(ctx context.Context, req DeleteSubscriptionPlanRequest) (DeleteSubscriptionPlanResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.DeleteSubscriptionPlan(ctx, req)
}

// Abonelik ürünü için fiyatlandırma planı detaylarını getirmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) GetSubscriptionPlanDetail(ctx context.Context, req GetSubscriptionPlanDetailRequest) (SubscriptionPlanResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.GetSubscriptionPlanDetail(ctx, req)
}

// Abonelik ürünleri için fiyatlandırma planlarını listelemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) GetSubscriptionPlanList(ctx context.Context, req GetSubscriptionPlanListRequest) (GetSubscriptionPlanListResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.GetSubscriptionPlanList(ctx, req)
}

// CheckoutForm ile Abonelik başlatma isteği yapmak için kullanılır.
//
// Abonelik süreci her zaman için ACTIVE veya PENDING durumu ile başlar. Eğer durum PENDING ise veya durum ACTIVE ancak ödeme planında bir deneme süresi belirtilmişse, iyzico abonelik isteğinde sadece kartın validasyonunu gerçekleştirir. Kart validasyonu 1 TL’lik bir çekim ve akabinde iade ile gerçekleşir. Bunun dışında herhangi bir işlem veya ödeme gerçekleşmez.
//
// Istekten dönen script etiketi ile birlikte bu 2 divi kullanabilirsiniz:
//
// <div id="iyzipay-checkout-form" class="responsive"></div>
//
// <div id="iyzipay-checkout-form" class="popup"></div>
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) InitializeSubscriptionWithCheckoutForm(ctx context.Context, req InitSubscriptionWithCheckoutFormRequest) (InitSubscriptionWithCheckoutFormResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.InitializeSubscriptionWithCheckoutForm(ctx, req)
}

// CheckoutForm ile Abonelik sorgulama isteği yapmak için kullanılır.
//
// Ödemenin tamamlanması gerekli.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) InquirySubscriptionWithCheckoutForm(ctx context.Context, req InquirySubscriptionWithCheckoutFormRequest) (SubscriptionPaymentResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.InquirySubscriptionWithCheckoutForm(ctx, req)
}

// 3D Secure olmayan Abonelik başlatma isteği yapmak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) InitializeSubscriptionWithNonTDS(ctx context.Context, req InitSubscriptionWithNonTDSRequest) (SubscriptionPaymentResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.InitializeSubscriptionWithNonTDS(ctx, req)
}

// Abone bilgilerini güncellemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) UpdateSubscriber(ctx context.Context, req *UpdateSubscriberRequest) (SubscriberResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.UpdateSubscriber(ctx, req)
}

// Abone bilgilerini sorgulamak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) GetSubscriberDetail(ctx context.Context, req *GetSubscriberDetailRequest) (SubscriberResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.GetSubscriberDetail(ctx, req)
}

// Abone listesini sorgulamak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Client) GetSubscriberList(ctx context.Context, req *GetSubscriberListRequest) (GetSubscriberListRequestResponse, error) {
	service := subscriptions.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.GetSubscriberList(ctx, req)
}

// Iyzilink oluşturma isteği
func (i *Client) IyzilinkCreate(ctx context.Context, req *IyzilinkCreateRequest) (IyzilinkResponse, error) {
	service := iyzilink.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.IyzilinkCreate(ctx, req)
}

// Iyzilink güncelleme isteği
func (i *Client) IyzilinkUpdate(ctx context.Context, req *IyzilinkUpdateRequest) (IyzilinkResponse, error) {
	service := iyzilink.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.IyzilinkUpdate(ctx, req)
}

// Iyzilink durum güncelleme isteği
func (i *Client) IyzilinkStatusUpdate(ctx context.Context, req *IyzilinkStatusUpdateRequest) (IyzilinkStatusUpdateResponse, error) {
	service := iyzilink.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.IyzilinkStatusUpdate(ctx, req)
}

// Iyzilink silme isteği
func (i *Client) IyzilinkDelete(ctx context.Context, req *IyzilinkDeleteRequest) (IyzilinkDeleteResponse, error) {
	service := iyzilink.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.IyzilinkDelete(ctx, req)
}

// Iyzilink detay sorgulama isteği
// Token bazlı çalışmıyor nedense o yüzden clientta filtreleme ile çözdüm(şimdilik).[liste içerisinden token ile arama yapılıyor]
// Veri sayısı çok olursa bu yöntem performansı düşürebilir.
func (i *Client) IyzilinkGetDetail(ctx context.Context, req *IyzilinkGetDetailRequest) (IyzilinkGetDetailResponse, error) {
	service := iyzilink.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.IyzilinkGetDetail(ctx, req)
}

// Iyzilink listeleme isteği
func (i *Client) IyzilinkGetList(ctx context.Context, req *IyzilinkGetListRequest) (IyzilinkGetListResponse, error) {
	service := iyzilink.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.IyzilinkGetList(ctx, req)
}

// App2App kullanıcı listeleme isteği
// Kullanıcı listeleme servisi üye işyerine tanımlamış olan kullanıcılar listeler. Listelenen kullanıcıların ceppos kullanma yetki durumunu bu servisle öğrenebilirsiniz.
func (i *App2AppClient) App2AppListUsers(ctx context.Context) (A2AListResponse, error) {
	service := app2app.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.App2AppListUsers(ctx)
}

// App2App ödeme oluşturma isteği
// Ödeme başlatma servisinin isteğinde tahsil edilecek tutar ve kullanıcı listeleme servisinde listelenen kullanıcının eposta bilgisi iletilmektedir.
func (i *App2AppClient) App2CreatePayment(ctx context.Context, req *A2ACreatePaymentRequest) (A2ACreatePaymentResponse, error) {
	service := app2app.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.App2CreatePayment(ctx, req)
}

// App2App iptal/iade isteği
// Başarılı gerçekleşen bir ödemeyi iptal veya iade etmek istediğinizde bu servisi kullanabilirsiniz.
//
// İptal / İade işlemi fiziki olarak gerçekleşmek zorundadır. Ceppos işlemleri iyzico paneli üzerinden iade edilememektedir.
func (i *App2AppClient) App2RefundPayment(ctx context.Context, req *A2ARefundPaymentRequest) (A2ARefundPaymentResponse, error) {
	service := app2app.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.App2RefundPayment(ctx, req)
}

// App2App sorgu isteği
// Ödeme işlemi sonrası callback parametre değerindeki url'e yönlendiğinde url'de data ve paymentSessionToken değerleri dönecektir. Bu değerleri kullanarak ödemeyi sorgulayabilirsiniz.
func (i *App2AppClient) App2InquiryPayment(ctx context.Context, req *A2AInquiryPaymentRequest) (A2AInquiryPaymentResponse, error) {
	service := app2app.Service{Transport: i.transport(), Validator: i.config.validator}
	return service.App2InquiryPayment(ctx, req)
}
