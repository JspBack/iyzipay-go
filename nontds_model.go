package iyzipay

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type PaymentRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir. en olarak kullanılabilir.
	//
	// zorunlu değil.
	Locale string `json:"locale" validate:"omitempty,oneof=tr en"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için kullanılabilir. En yaygın kullanış biçimi üye iş yerinin sipariş numarasıdır.
	//
	// zorunlu değil.
	ConversationID string `json:"conversationId" validate:"omitempty"`

	// Ödeme sepet tutarı. Kırılım (sepetin içerisindeki ürünler) tutarlar toplamı sepet tutarına eşit olmalı.
	//
	// zorunlu.
	Price string `json:"price" validate:"required,numeric"`

	// İndirim vade farkı vs. hesaplanmış POS’tan geçecek nihai tutar. Price değerinden küçük, büyük veya eşit olabilir
	//
	// zorunlu.
	PaidPrice string `json:"paidPrice" validate:"required,numeric"`

	// Taksit bilgisi, tek çekim için 1 gönderilmelidir. Geçerli değerler: 1, 2, 3, 6, 9, 12
	//
	// zorunlu.
	Installment int `json:"installment" validate:"required,numeric,gte=1,lte=12"`

	// Ödeme kanalı. Geçerli değerler enum içinde sunulmaktadır: WEB, MOBILE, MOBILE_WEB, MOBILE_IOS, MOBILE_ANDROID, MOBILE_WINDOWS, MOBILE_TABLET, MOBILE_PHONE
	//
	// zorunlu değil.
	PaymentChannel string `json:"paymentChannel" validate:"omitempty,oneof=WEB MOBILE MOBILE_WEB MOBILE_IOS MOBILE_ANDROID MOBILE_WINDOWS MOBILE_TABLET MOBILE_PHONE"`

	// Üye işyeri tarafından ilgili ödemenin sepetini tanımlamak için kullanılan id'dir. Sipariş numarası ya da anlamlı bir değer olabilir.
	//
	// zorunlu değil.
	BasketID string `json:"basketId" validate:"omitempty"`

	// Ödeme grubu, varsayılan PRODUCT. Geçerli değerler enum içinde sunulmaktadır: PRODUCT, LISTING, SUBSCRIPTION, OTHER.
	//
	// zorunlu değil.
	PaymentGroup string `json:"paymentGroup" validate:"omitempty,oneof=PRODUCT LISTING SUBSCRIPTION OTHER"`

	// Müşterinin ödeme işlemi için kullanılacak bilgiler
	//
	// zorunlu.
	Buyer Buyer `json:"buyer" validate:"required,dive"`

	// Ödeme kartı bilgileri
	//
	// zorunlu.
	PaymentCard PaymentCard `json:"paymentCard" validate:"required,dive"`

	// Müşterinin ödeme işlemi sırasında seçtiği teslimat adresi (Sepetteki ürünlerden en az 1 tanesi fiziksel ürün (itemType=PHYSICAL) ise zorunludur.)
	//
	// zorunlu değil.
	ShippingAddress ShippingAddress `json:"shippingAddress,omitempty" validate:"omitempty"`

	// Müşterinin ödeme işlemi sırasında seçtiği fatura adresi
	//
	// zorunlu.
	BillingAddress BillingAddress `json:"billingAddress" validate:"required,dive"`

	// Sepet içerisindeki ürünlerin listesi
	//
	// zorunlu.
	BasketItems []BasketItem `json:"basketItems" validate:"required,dive"`

	// Ödemenin alınacağı para birimi default TL olarak belirlenmiştir. Diğer değerler ise USD, EUR, GBP olmak üzere, farklı para birimleri ile alışverişin hesabınıza tanımlandığından emin olunuz.
	//
	// zorunlu.
	Currency string `json:"currency" validate:"required"`
}

type Buyer struct {
	// Üye işyeri tarafındaki alıcıya ait id.
	//
	// zorunlu.
	ID string `json:"id" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait ad.
	//
	// zorunlu.
	Name string `json:"name" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait soyad.
	//
	// zorunlu.
	Surname string `json:"surname" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait kimlik (TCKN) numarası
	//
	// zorunlu.
	IdentityNumber string `json:"identityNumber" validate:"required,numeric"`

	// Üye işyeri tarafındaki alıcıya ait e-posta bilgisi. E-posta adresi alıcıya ait geçerli ve erişilebilir bir adres olmalıdır
	//
	// zorunlu.
	Email string `json:"email" validate:"required,email"`

	// Üye işyeri tarafındaki alıcıya ait GSM numarası
	//
	// zorunlu değil.
	GSMNumber string `json:"gsmNumber" validate:"omitempty"`

	// Üye işyeri tarafındaki alıcıya ait kayıt tarihi. Tarih formatı 2015-09-17 23:45:06 şeklinde olmalıdır.
	//
	// zorunlu değil.
	RegistrationDate string `json:"registrationDate" validate:"omitempty,iyziDate"`

	// Üye işyeri tarafındaki alıcıya ait son giriş tarihi. Tarih formatı 2015-09-17 23:45:06 şeklinde olmalıdır.
	//
	// zorunlu değil.
	LastLoginDate string `json:"lastLoginDate" validate:"omitempty,iyziDate"`

	// Üye işyeri tarafındaki alıcıya ait kayıt adresi.
	//
	// zorunlu.
	RegistrationAddress string `json:"registrationAddress" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait şehir bilgisi.
	//
	// zorunlu.
	City string `json:"city" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait ülke bilgisi.
	//
	// zorunlu.
	Country string `json:"country" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait posta kodu.
	//
	// zorunlu değil.
	ZipCode string `json:"zipCode" validate:"omitempty,numeric"`

	// Üye işyeri tarafındaki alıcıya ait IP adresi.
	//
	// zorunlu.
	IP string `json:"ip" validate:"required,ip"`
}

type PaymentCard struct {
	// Saklı kartın sahibinin kart saklama servisindeki kullanıcı kimliği. Eğer saklı kart ile ödeme yapılacaksa zorunludur.
	//
	// zorunlu değil.
	CardUserKey string `json:"cardUserKey,omitempty" validate:"omitempty"`

	// Ödemenin alınacağı saklı kartın token’ı. Eğer saklı kart ile ödeme yapılacaksa zorunludur.
	//
	// zorunlu değil.
	CardToken string `json:"cardToken,omitempty" validate:"omitempty"`

	// Ödemenin alınacağı kart sahibinin adı soyadı. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	//
	// zorunlu değil.
	CardHolderName string `json:"cardHolderName,omitempty" omitempty:"required"`

	// Ödemenin alınacağı kart numarası. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	//
	// zorunlu değil.
	CardNumber string `json:"cardNumber,omitempty" validate:"omitempty,creditcard"`

	// Ödemenin alınacağı kartın son kullanma tarihi yılı. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur
	//
	// zorunlu değil.
	ExpireYear string `json:"expireYear,omitempty" validate:"omitempty,numeric,len=4"`

	// Ödemenin alınacağı kartın son kullanma tarihi ayı. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	//
	// zorunlu değil.
	ExpireMonth string `json:"expireMonth,omitempty" validate:"omitempty,numeric,len=2,oneof=01 02 03 04 05 06 07 08 09 10 11 12"`

	// Ödemenin alınacağı kartın güvenlik kodu. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur. Saklı kartla ödeme yapılırken gönderilirse aynen bankaya iletilir
	//
	// zorunlu değil.
	CVC string `json:"cvc,omitempty" validate:"omitempty,numeric,len=3"`

	// Ödeme esnasında kartın kaydedilip kaydedilmeyeceğini belirleyen parametre. Varsayılan değeri 0 olup, geçerli değerler 0 ve 1’dir
	//
	// zorunlu değil.
	RegisterCard int `json:"registerCard,omitempty" validate:"omitempty,oneof=0 1"`
}

type BillingAddress struct {
	// Üye işyeri tarafındaki alıcıya ait adres bilgisi.
	//
	// zorunlu.
	Address string `json:"address" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait posta kodu.
	//
	// zorunlu değil.
	ZipCode string `json:"zipCode,omitempty" validate:"omitempty,numeric"`

	// Üye işyeri tarafındaki alıcıya ait iletişim adı.
	//
	// zorunlu.
	ContactName string `json:"contactName" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait şehir bilgisi.
	//
	// zorunlu.
	City string `json:"city" validate:"required"`

	// Üye işyeri tarafındaki alıcıya ait ülke bilgisi.
	//
	// zorunlu.
	Country string `json:"country" validate:"required"`
}

type ShippingAddress struct {
	// Üye işyeri tarafındaki alıcıya ait adres bilgisi.
	//
	// zorunlu.
	Address string `json:"address" validate:"omitempty"`

	// Üye işyeri tarafındaki alıcıya ait posta kodu.
	//
	// zorunlu değil.
	ZipCode string `json:"zipCode,omitempty" validate:"omitempty,numeric"`

	// Üye işyeri tarafındaki alıcıya ait iletişim adı.
	//
	// zorunlu.
	ContactName string `json:"contactName" validate:"omitempty"`

	// Üye işyeri tarafındaki alıcıya ait şehir bilgisi.
	//
	// zorunlu.
	City string `json:"city" validate:"omitempty"`

	// Üye işyeri tarafındaki alıcıya ait ülke bilgisi.
	//
	// zorunlu.
	Country string `json:"country" validate:"omitempty"`
}

type BasketItem struct {
	// Üye işyeri tarafındaki sepetteki ürüne ait id. Not: Bir ödeme isteğine maksimum 500 basketItem eklenebilir
	//
	// zorunlu.
	ID string `json:"id" validate:"required"`

	// Üye işyeri tarafındaki sepetteki ürüne ait tutar. 0 ve 0’dan küçük olamaz; tutarlar toplamı sepet tutarına (price) eşit olmalıdır.
	//
	// zorunlu.
	Price string `json:"price" validate:"required,numeric"`

	// Üye işyeri tarafındaki sepetteki ürüne ait isim.
	//
	// zorunlu.
	Name string `json:"name" validate:"required"`

	// Üye işyeri tarafındaki sepetteki ürüne ait kategori 1
	//
	// zorunlu.
	Category1 string `json:"category1" validate:"required"`

	// Üye işyeri tarafındaki sepetteki ürüne ait kategori 2
	//
	// zorunlu değil.
	Category2 string `json:"category2" validate:"omitempty"`

	// Üye işyeri tarafındaki sepetteki ürüne ait tip. Geçerli enum değerler: PHYSICAL ve VIRTUAL.
	//
	// zorunlu.
	ItemType string `json:"itemType" validate:"required,oneof=PHYSICAL VIRTUAL"`
}

type Non3DSPaymentResponse struct {
	// Üye işyeri tarafındaki sepetteki ürüne ait tip. Geçerli enum değerler: PHYSICAL ve VIRTUAL.
	Status string `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime int64 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletililir.
	ConversationID string `json:"conversationId"`

	// Ödeme sepet tutarı. Kırılım tutarlar toplamı sepet tutarına eşit olmalı.
	Price float64 `json:"price"`

	// İndirim vade farkı vs. hesaplanmış POS’tan geçen, tahsil edilen, nihai tutar.
	PaidPrice float64 `json:"paidPrice"`

	// Ödemenin taksit bilgisi, tek çekim için 1 döner. Geçerli değerler: 1, 2, 3, 6, 9, 12.
	Installment int `json:"installment"`

	// Ödemeye ait id, üye işyeri tarafından mutlaka saklanmalıdır. Ödemenin iptali ve iyzico ile iletişimde kullanılır.
	PaymentID string `json:"paymentId"`

	// Ödeme işleminin fraud filtrelerine göre durumu. Eğer ödemenin fraud risk skoru düşük ise ödemeye anında onay verilir, bu durumda 1 değeri döner.
	// Eğer fraud risk skoru yüksek ise ödeme işlemi reddedilir ve -1 döner. Eğer ödeme işlemi daha sonradan incelenip karar verilecekse 0 döner.
	// Geçerli değerler: 0, -1 ve 1. Üye işyeri sadece 1 olan işlemlerde ürünü kargoya vermelidir, 0 olan işlemler için bilgilendirme beklemelidir.
	FraudStatus int `json:"fraudStatus"`

	// Üye işyerinin uyguladığı vade/komisyon oranı. Örneğin price=100, paidPrice=110 ise üye işyeri vade/komisyon oranı %10’dur. Bilgi amaçlıdır.
	MerchantCommissionRate float64 `json:"merchantCommissionRate"`

	// Üye işyerinin uyguladığı vade/komisyon tutarı. Örneğin price=100, paidPrice=110 ise üye işyeri vade/komisyon tutarı 10’dur. Bilgi amaçlıdır.
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`

	// Ödemeye ait iyzico işlem komisyon tutarı.
	IyziCommissionRateAmount float64 `json:"iyziCommissionRateAmount"`

	// Ödemeye ait iyzico işlem ücreti
	IyziCommissionFee float64 `json:"iyziCommissionFee"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu kart tipi. Geçerli değerler: CREDIT_CARD, DEBIT_CARD, PREPAID_CARD.
	CardType string `json:"cardType"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu kuruluş. Geçerli değerler: VISA, MASTER_CARD, AMERICAN_EXPRESS, TROY
	CardAssociation string `json:"cardAssociation"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu aile. Geçerli değerler: Bonus, Axess, World, Maximum, Paraf, CardFinans, Advantage
	CardFamily string `json:"cardFamily"`

	// Ödeme yapılan kartın ilk 6 hanesi
	BinNumber string `json:"binNumber"`

	// Ödeme yapılan kartın son 4 hanesi
	LastFourDigits string `json:"lastFourDigits"`

	// Üye işyeri tarafından gönderilen sepet id’si
	BasketID string `json:"basketId"`

	// Ödemenin alındığı para birimi
	Currency string `json:"currency"`

	// Sepetteki ürünlerin listesi
	ItemTransactions []ItemTransaction `json:"itemTransactions"`

	// Metadata (Dökümantasyona eklenmemiş ama response içinde var)
	AuthCode string `json:"authCode"`

	// Metadata (Dökümantasyona eklenmemiş ama response içinde var)
	Phase string `json:"phase"`

	// Metadata (Dökümantasyona eklenmemiş ama response içinde var)
	HostReference string `json:"hostReference"`

	// Metadata (Dökümantasyona eklenmemiş ama response içinde var)
	Signature string `json:"signature"`
}

type ItemTransaction struct {
	// Üye işyeri tarafından iletilen, sepetteki ürüne ait id
	ItemID string `json:"itemId"`

	// Ödeme kırılımına ait id, üye işyeri tarafından mutlaka saklanmalıdır.
	// Ödeme kırılımının iadesi, onayı, onay geri çekmesi ve iyzico ile iletişimde kullanılır. Tercihen itemId ile ilişkili bir şekilde tutulmalıdır.
	PaymentTransactionID string `json:"paymentTransactionId"`

	// Ödeme kırılımının durumu. Ödeme fraud kontrolünde ise 0 değeri döner, bu durumda fraudStatus değeri de 0’dır.
	// Ödeme, fraud kontrolünden sonra reddedilirse -1 döner. Pazaryeri modelinde ürüne onay verilene dek bu değer 1 olarak döner.
	// Pazaryeri modelinde ürüne onay verilmişse bu değer 2 olur. Geçerli değerler: 0, -1, 1 ve 2.
	TransactionStatus int `json:"transactionStatus"`

	// Üye işyeri tarafındaki sepetteki ürüne ait tutar.
	Price float64 `json:"price"`

	// Tahsilat tutarının kırılım bazındaki dağılımı. Üye işyeri tarafından mutlaka saklanmalıdır.
	PaidPrice float64 `json:"paidPrice"`

	// Üye işyerinin uyguladığı vade/komisyon oranının kırılım bazında dağılmış oranı.
	MerchantCommissionRate float64 `json:"merchantCommissionRate"`

	// iyzico işlem komisyon tutarının kırılım bazında dağılmış tutarı.
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`

	// iyzico işlem komisyon tutarının kırılım bazında dağılmış tutarı.
	IyziCommissionRateAmount float64 `json:"iyziCommissionRateAmount"`

	// iyzico işlem ücretinin kırılım bazında dağılmış tutarı.
	IyziCommissionFee float64 `json:"iyziCommissionFee"`

	// Kırılım bazında üye işyeri blokaj oranı. iyzico – üye işyeri anlaşmasına göre, üye işyerine işlem bazında blokaj uygulayabilir.
	// Bu blokaj üye işyeri fraud riskini önlemek içindir, blokaj süresi boyunca para iyzico’da tutulur, bu süre sonrasında üye işyerine gönderilir.
	BlockageRate float64 `json:"blockageRate"`

	// Kırılım bazında üye işyeri blokaj tutarının, üye işyerine yansıyan rakamı. Blokaj tutarı mümkün olduğunca üye işyerine yansıtılır.
	// Eğer blokaj tutarı, üye işyeri tutarından daha büyükse bu durumda alt üye işyerine de yansıtılır.
	BlockageRateAmountMerchant float64 `json:"blockageRateAmountMerchant"`

	// Dökümantasyonda belirtilmemiş
	BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`

	// İşlem bazında blokaj çözülme tarihi. yyyy-MM-dd HH:mm:ss formatındadır, örneğin 2015-10-19 14:36:52
	BlockageResolvedDate string `json:"blockageResolvedDate"`

	// Alt üye işyerine IBAN adresine gönderilmiş tutar.
	SubMerchantPrice float64 `json:"subMerchantPrice"`

	// Dökümantasyonda belirtilmemiş
	SubMerchantPayoutRate float64 `json:"subMerchantPayoutRate"`

	// Dökümantasyonda belirtilmemiş
	SubMerchantPayoutAmount float64 `json:"subMerchantPayoutAmount"`

	// Bu kırılım için, iyzico işlem ücreti, komisyon tutarı ve blokajlar düşüldükten sonra üye işyerine gönderilecek tutar.
	MerchantPayoutAmount float64 `json:"merchantPayoutAmount"`

	// Dökümantasyonda belirtilmemiş
	ConvertedPayout ConvertedPayout `json:"convertedPayout"`
}

type ConvertedPayout struct {
	// Tahsilat tutarının kırılım bazındaki dağılımı. Üye işyeri tarafından mutlaka saklanmalıdır.
	PaidPrice float64 `json:"paidPrice"`

	// iyzico işlem komisyon tutarının kırılım bazında dağılmış tutarı
	IyziCommissionRateAmount float64 `json:"iyziCommissionRateAmount"`

	// iyzico işlem ücretinin kırılım bazında dağılmış tutarı
	IyziCommissionFee float64 `json:"iyziCommissionFee"`

	// Kırılım bazında üye işyeri blokaj tutarının, üye işyerine yansıyan rakamı. Blokaj tutarı mümkün olduğunca üye işyerine yansıtılır.
	// Eğer blokaj tutarı, üye işyeri tutarından daha büyükse bu durumda alt üye işyerine de yansıtılır.
	BlockageRateAmountMerchant float64 `json:"blockageRateAmountMerchant"`

	// Dökümantasyonda belirtilmemiş
	BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`

	// Dökümantasyonda belirtilmemiş
	SubMerchantPayoutAmount float64 `json:"subMerchantPayoutAmount"`

	// Bu kırılım için, iyzico işlem ücreti, komisyon tutarı ve blokajlar düşüldükten sonra üye işyerine gönderilecek tutar.
	MerchantPayoutAmount float64 `json:"merchantPayoutAmount"`

	// Bu kırılım için, iyzico işlem ücreti, komisyon tutarı ve blokajlar düşüldükten sonra üye işyerine gönderilecek tutar.
	IyziConversionRate float64 `json:"iyziConversionRate"`

	// Bu kırılım için, iyzico işlem ücreti, komisyon tutarı ve blokajlar düşüldükten sonra üye işyerine gönderilecek tutar.
	IyziConversionRateAmount float64 `json:"iyziConversionRateAmount"`

	// Ödemenin alındığı para birimi.
	Currency string `json:"currency"`
}

func (r *PaymentRequest) validate() error {
	validate := validator.New()

	// Tarih formatı doğrulaması
	validate.RegisterValidation("iyziDate", func(fl validator.FieldLevel) bool {
		dateFormat := "2006-01-02 15:04:05"
		dateStr := fl.Field().String()
		if dateStr == "" {
			return true
		}
		_, err := time.Parse(dateFormat, dateStr)
		return err == nil
	})

	// Luhn algoritması ile kredi kartı numarası doğrulaması (https://en.wikipedia.org/wiki/Luhn_algorithm)
	validate.RegisterValidation("creditcard", func(fl validator.FieldLevel) bool {
		cardNumber := fl.Field().String()
		cleanNumber := ""
		for _, ch := range cardNumber {
			if ch >= '0' && ch <= '9' {
				cleanNumber += string(ch)
			}
		}

		if len(cleanNumber) < 13 || len(cleanNumber) > 19 {
			return false
		}

		sum := 0
		shouldDouble := false
		for i := len(cleanNumber) - 1; i >= 0; i-- {
			digit := int(cleanNumber[i] - '0')
			if shouldDouble {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
			shouldDouble = !shouldDouble
		}

		return sum%10 == 0
	})

	// Eğer doğrulamada hata oluşursa hata mesajını döndürür.
	err := validate.Struct(r)
	if err != nil {
		var errs []string
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			errs = append(errs, fmt.Sprintf("Validation error at field '%s' with tag '%s'", field, tag))
		}
		return fmt.Errorf("validation errors: %s", strings.Join(errs, ", "))
	}
	return nil
}
