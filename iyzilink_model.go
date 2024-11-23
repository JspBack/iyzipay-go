package iyzipay

import "gopkg.in/go-playground/validator.v9"

type IyzilinkCreateRequest struct {
	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Ürünün ekranda görünen ismi.
	//
	// zorunlu
	Name string `json:"name" validate:"required"`

	// Ürün fiyatı.
	//
	// zorunlu
	Price string `json:"price" validate:"required"`

	// Taksitli Satış Durumu. Gönderilmezse false değeri alır.
	//
	// zorunlu değil
	InstallmentRequested string `json:"installmentRequested" validate:"omitempty,oneof=true false"`

	// Ürün satın alınırken adres isteme parametresi. Gönderilmese true değerini alır.
	//
	// zorunlu değil
	AddressIgnorable string `json:"addressIgnorable" validate:"omitempty,oneof=true false"`

	// Ürünün para birimi.
	//
	// zorunlu
	CurrencyCode string `json:"currencyCode" validate:"required"`

	// Ürünün base64 formatında resmi.
	//
	// zorunlu
	EncodedImageFile string `json:"encodedImageFile" validate:"required,base64"`

	// Ürünün ekranda görünen açıklaması.
	//
	// zorunlu
	Description string `json:"description" validate:"required"`

	ProductBuyerInfo *ProductBuyerInfo `json:"productBuyerInfo" validate:"omitempty"`

	// Stok sayısı.
	//
	// zorunlu
	StockCount int `json:"stockCount" validate:"required"`

	// Stok onayı, normalde gelen değer FALSE.
	//
	// zorunlu
	StockEnabled bool `json:"stockEnabled" validate:"required"`
}

type ProductBuyerInfo struct {
	// Alıcı ismi
	//
	// zorunlu değil
	BuyerName string `json:"buyerName" validate:"omitempty"`

	// Alıcı soy ismi.
	//
	// zorunlu değil
	BuyerSurname string `json:"buyerSurname" validate:"omitempty"`

	// Alıcı şehri.
	//
	// zorunlu değil
	BuyerCity string `json:"buyerCity" validate:"omitempty"`

	// Alıcı ülkesi.
	//
	// zorunlu değil
	BuyerCountry string `json:"buyerCountry" validate:"omitempty"`

	// Alıcı GSM.
	//
	// zorunlu değil
	BuyerGsmNumber string `json:"buyerGsmNumber" validate:"omitempty"`

	// Alıcı email adresi.
	//
	// zorunlu değil
	BuyerEmail string `json:"buyerEmail" validate:"omitempty,email"`

	// Alıcı adresi.
	//
	// zorunlu değil
	BuyerAddress string `json:"buyerAddress" validate:"omitempty"`
}

type IyzilinkUpdateRequest struct {
	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Ürünün ekranda görünen ismi.
	//
	// zorunlu
	Name string `json:"name" validate:"required"`

	// Ürün fiyatı.
	//
	// zorunlu
	Price string `json:"price" validate:"required"`

	// Taksitli Satış Durumu. Gönderilmezse false değeri alır.
	//
	// zorunlu değil
	InstallmentRequested string `json:"installmentRequested" validate:"omitempty,oneof=true false"`

	// ???
	SoldLimit string `json:"soldLimit" validate:"omitempty"`

	// Ürün satın alınırken adres isteme parametresi. Gönderilmese true değerini alır.
	//
	// zorunlu değil
	AddressIgnorable string `json:"addressIgnorable" validate:"omitempty,oneof=true false"`

	// Ürünün para birimi.
	//
	// zorunlu
	CurrencyCode string `json:"currencyCode" validate:"required"`

	// Ürünün para birimi.
	//
	// zorunlu
	EncodedImageFile string `json:"encodedImageFile" validate:"required,base64"`

	// Ürünün ekranda görünen açıklaması.
	//
	// zorunlu
	Description string `json:"description" validate:"required"`

	// Iyzilink'e ait token.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`
}

type IyzilinkDeleteRequest struct {
	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Iyzilink'e ait token.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`
}

type IyzilinkGetDetailRequest struct {
	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Sayfa numarası.
	//
	// zorunlu
	Page int `json:"page" validate:"required"`

	// Sayfada görüntülenen ürün sayısı.
	//
	// zorunlu
	Count int `json:"count" validate:"required"`

	// Iyzilink'e ait token.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`
}

type IyzilinkGetListRequest struct {
	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Sayfa numarası.
	//
	// zorunlu
	Page int `json:"page" validate:"required"`

	// Sayfada görüntülenen ürün sayısı.
	//
	// zorunlu
	Count int `json:"count" validate:"required"`
}

type IyzilinkStatusUpdateRequest struct {
	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	//
	// zorunlu değil
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// Iyzilink'e ait token.
	//
	// zorunlu
	Token string `json:"token" validate:"required"`

	// Iyzilink'in durumu.
	//
	// zorunlu
	ProductStatus string `json:"productStatus" validate:"required,oneof=ACTIVE PASSIVE"`
}

type IyzilinkResponse struct {
	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	ConversationID string `json:"conversationId"`

	IyzilinkDeleteResponse

	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	Locale string `json:"locale"`

	// Başarılı işlem sonucu oluşan iyzilink verisi.
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	// Başarılı işlem sonucu oluşan iyzilink tokeni.
	Token string `json:"token"`

	// Başarılı işlem sonucu oluşan iyzilink urli.
	Url string `json:"url"`

	// Başarılı işlem sonucu oluşan iyzilink resim urli.
	ImageUrl string `json:"imageUrl"`
}

type IyzilinkDeleteResponse struct {
	// İşlem sonucu başarılı ise success, başarısız ise failure döner.
	Status string `json:"status"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int `json:"systemTime"`
}

type IyzilinkGetDetailResponse struct {
	// Yapılan isteğin sonucunu belirtir. İşlem başarılı ise success değeri döner.
	Status string `json:"status"`

	// Kullanılan dil.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir
	SystemTime int `json:"systemTime"`

	// İşlem sonucu gelen iyzilink verisi.
	Data DetailResponseData `json:"data"`
}

type DetailResponseData struct {
	// Ürün ismi.
	Name string `json:"name"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationID string `json:"conversationId"`

	// Ürün açıklaması.
	Description string `json:"description"`

	// Ürün fiyatı.
	Price float64 `json:"price"`

	// Ürünün para birimi ID'si.
	CurrencyId int `json:"currencyId"`

	// Ürünün para birimi.
	CurrencyCode string `json:"currencyCode"`

	// iyzico tarafından üretilen değer.
	Token string `json:"token"`

	// iyzico tarafından kullanılan tip.
	ProductType string `json:"productType"`

	// Ürünün durumu.
	ProductStatus string `json:"productStatus"`

	// Üye işyeri numarası
	MerchantId int `json:"merchantId"`

	// Ürünün satın alma linki.
	Url string `json:"url"`

	// Ürün resmi.
	ImageUrl string `json:"imageUrl"`

	// Ürün adres durumu.
	AddressIgnorable bool `json:"addressIgnorable"`

	// Satılan ürün sayısı.
	SoldCount int `json:"soldCount"`

	// Taksit durumu.
	InstallmentRequested bool `json:"installmentRequested"`

	// Ürünün stok durumu.
	StockEnabled bool `json:"stockEnabled"`

	// Toplam stok sayısı.
	StockCount int `json:"stockCount"`

	// ???
	PresetPriceValues []int `json:"presetPriceValues"`

	// ???
	FlexibleLink bool `json:"flexibleLink"`

	// ???
	CategoryType string `json:"categoryType"`
}

type IyzilinkGetListResponse struct {
	// İşlem sonucu başarılı ise success, başarısız ise failure döner.
	Status string `json:"status"`

	// Kullanılan dil.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationID string `json:"conversationId"`

	// İşlem sonucu gelen iyzilink verisi.
	Data ListResponseData `json:"data"`
}

type ListResponseData struct {
	// Listeleme durumu.
	ListingReviewed bool `json:"listingReviewed"`

	// Toplam iyzilink sayısı.
	TotalCount int `json:"totalCount"`

	// Sayfa numarası.
	CurrentPage int `json:"currentPage"`

	// Toplam sayfa sayısı.
	PageCount int `json:"pageCount"`

	// Sayfada görüntülenen iyzilink verisi.
	Items []DetailResponseData `json:"items"`
}

type IyzilinkStatusUpdateResponse struct {
	// İşlem sonucu başarılı ise success, başarısız ise failure döner.
	Status string `json:"status"`

	// Kullanılan dil.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationID string `json:"conversationId"`
}

func (r IyzilinkCreateRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r IyzilinkUpdateRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r IyzilinkDeleteRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r IyzilinkGetDetailRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r IyzilinkGetListRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r IyzilinkStatusUpdateRequest) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
