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
- BIN Kontrolü: Kredi kartı BIN'lerini doğrulayın.
- Ödeme Sorgulama: Ödeme durumlarını sorgulayın.
- 3DS İstekleri: 3D Secure işlemleri desteği.
- PWI Desteği: Ödeme Penceresi Entegrasyonu (ödemek için yeni bir ekran açar).
- Ödeme Formu Entegrasyonu: Sorunsuz ödeme işleme için ödeme formlarını entegre edin.

## Planlanan Özellikler

- MerchantPlace API'leri: MerchantPlace API'leri ile etkileşim.
- Abonelik API'leri: Abonelik hizmetlerini yönetin.
- Iyzilink API'leri: Iyzilink hizmetleri ile entegrasyon.
- EFT API'leri: Elektronik Fon Transferi desteği.
- Ekstra Hizmetler: Diğer ekstra hizmetler eklenecektir.

## Örnekler

Örnekler /examples dosyası içerisinde.

## Notlar

Şu an hala geliştirme aşamasında muhtemelen hatalara denk gelebilirsiniz.(Ayrıca ilk go paketim 😄)

## Katkıda Bulunma

Bir pull request oluşturarak projeye destek olabilrisiniz. 🙂

## Bilinen Problemler

Saklı kartlarla ödeme için ayrı bir yöntem kullanılmalı.

## Lisans

MIT Lisansı altında dağıtılmaktadır, daha fazla ayrıntı için lütfen kod içindeki lisans dosyasına bakın.

[Read this in english](en.README.md)
