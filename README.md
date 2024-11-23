# Go Client for Iyzico API

<img align="right" width="300" src="gopher.png" alt="gopher">

![Proje Durumu](https://img.shields.io/badge/version-1.0.0-green.svg)
[![Go Rapor Kartı](https://goreportcard.com/badge/github.com/JspBack/iyzipay-go)](https://goreportcard.com/report/github.com/JspBack/iyzipay-go)
[![GoDoc](https://godoc.org/github.com/JspBack/iyzipay-go?status.svg)](https://pkg.go.dev/github.com/JspBack/iyzipay-go)
![Lisans](https://img.shields.io/badge/license-MIT-blue.svg)

`iyzipay` paketi, Iyzico API ile etkileşim kurmak için bir Go istemcisi sağlar. Bu paket, Non3DS istekleri, BIN kontrolü ve ödeme sorgulamaları gibi çeşitli işlevleri destekler.

## Kurulum

Paketi yüklemek için `go get` komutunu kullanın:

```bash
go get github.com/JspBack/iyzipay-go
```

## Özellikler

- Non3DS İstekleri: Non3DS ödeme isteklerini işleyin.
- BIN ve taksit Kontrolü: Kredi kartı BIN'lerini ve taksitlerini doğrulayın.
- Ödeme Sorgulama: Ödeme durumlarını sorgulayın.
- 3DS İstekleri: 3D Secure işlemleri desteği.
- PWI Desteği: Ödeme Penceresi Entegrasyonu (ödemek için yeni bir ekran açar).
- Ödeme Formu Entegrasyonu: Sorunsuz ödeme işleme için ödeme formlarını entegre edin.
- MarketPlace API'leri: MarketPlace API'leri ile etkileşim. (yada MerchantPlace)
- Kart Saklama: Kullanıcıların kredi kartı bilgilerini güvenli bir şekilde saklayın ve tekrar eden ödemeler için kullanın.
- Abonelik API'leri: Abonelik hizmetlerini yönetin.
- Iyzilink API'leri: Iyzilink hizmetleri ile entegrasyon.
- Ceppos App2App entegrasyonu

## Planlanan Özellikler

- Ekstra Hizmetler: Diğer ekstra hizmetler eklenecektir.
- Alışveriş Kredisi entegrasyonu

## Basit Kullanımı

```go
// Olabildiğince basit tuttum 😎
apikey := os.Getenv("IYZIPAY_API_KEY")
apiSecret := os.Getenv("IYZIPAY_API_SECRET")

client, err := iyzipay.New(apikey, apiSecret)
if err != nil {
    fmt.Println(err)
}

binReq := &iyzipay.BinRequest{
    Locale:         "tr",
    BinNumber:      "454671",
    ConversationId: "123456789",
}

binRes, err := client.BinControlRequest(binReq)
if err != nil {
    fmt.Println(err)
}

if binRes.Status == "success" {
    fmt.Println("Bin Number: ", binRes.CardAssociation) // VISA
    return
}
```

## Örnekler

Örnekler /examples dosyası içerisinde.

## Notlar

Şu an hala geliştirme aşamasında muhtemelen hatalara denk gelebilirsiniz.(Ayrıca ilk go paketim 😄). \
App2App entegrasyonu mobile özel geliştirilmiş sanırım ama go'da da mobile build alınabildiği için ekledim.

## Katkıda Bulunma

Bir pull request oluşturarak projeye destek olabilirsiniz. 🙂

## Bilinen Problemler

[GENEL]

- Unauthorized (401) hataları panic oluşturuyor (hata formatı farklı olduğu için).
- Pazaryeri, Abonelik, İptal ve iade, Iyzilink, App2App örnekleri yok (test edilmediler veya hatalı çalışıyorlar).

[IYZILINK]

- Iyzilink entegrasyonunda IyzilinkUpdate,IyzilinkDelete çalışmıyor 'sistem hatası' dönüyor.
- Iyzilink entegrasyonunda IyzilinkGetDetail, IyzilinkGetList'ten gelen veri filtrelenerek alınıyor. Bu doğru bir kullanım değil ama geçici bir çözüm, çünkü IyzilinkGetDetail'den dönen veri IyzilinkGetList'ten dönen veri tipinde yani doğru çalışmıyor.

## Lisans

MIT Lisansı altında dağıtılmaktadır, daha fazla ayrıntı için lütfen kod içindeki lisans dosyasına bakın.

[Read this in english](en.README.md)
