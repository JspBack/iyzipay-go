// Package iyzipay, Iyzico API ile iletişim kurmak için bir istemci sağlar.
//
// İstemcileri New ile oluşturun; ortamı ve HTTP istemcisini WithBaseURL ve
// WithHTTPClient ile yapılandırın. API metotları context parametresiyle istek iptalini
// ve son tarihleri destekler.
// İstemciler, yapılandırıldıktan sonra goroutine'ler arasında paylaşılabilir.
// İstekler sürerken verilen HTTP istemcisini veya dışa aktarılan App2App
// kimlik bilgilerini değiştirmeyin.
// Kimlik bilgileri uygulama tarafından sağlanır; bu paket .env dosyalarını yüklemez.
package iyzipay
