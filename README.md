# Go Blog

Go Blog, Go dilinde geliştirilmiş basit bir blog uygulamasıdır. Kullanıcılar, gönderileri görebilir ve aratma yapabilir. Admin paneli üzerinden içerikler yönetilebilir.

## Özellikler

- Yanlızca Admin yazı oluşturabilir, güncelleyebilir ve silebilir.
- Kullanıcılar yazıları görüntüleyebilir ve aratma yapabilir.
- Admin paneli, içerik yönetimi sağlar.
- Authentication için JWT (JSON Web Token) kullanılır.
- Veritabanı olarak [GORM](https://gorm.io/) kullanılır.
- Bir veritabanı (Gorm ile uyumlu) gereklidir.
- Postgres veri tabanı kullanılmıştır.
- WYSIWYG editörü olan [Quill.js](https://quilljs.com/) kullanılmıştır.
- Frontend için [html/template](https://pkg.go.dev/html/template) paketi ile birlikte [TailwindCSS](https://tailwindcss.com/) kullanılmıştır.

## Ekran Görüntüleri
 - [examples](https://github.com/fatihesergg/go_blog/tree/main/examples) klasörüne bakınız.

## Başlangıç
<details>
  <summary> Kurulum Adımları</summary>


## Projeyi kendi bilgisayarınızda çalıştırmak için aşağıdaki adımları takip edebilirsiniz:

### Gereksinimler

- Go (derlemek için) yüklü olmalıdır. Yüklemek için [Go'nun resmi web sitesinden](https://golang.org/dl/) en son sürümü indirip kurabilirsiniz.
- Bir veritabanı ,GORM ile uyumlu birçok veritabanını kullanabilirsiniz.

### Proje Kurulumu

1. Bu projeyi kendi bilgisayarınıza klonlayın:
   ```bash
   git clone https://github.com/fatihesergg/go_blog.git
   cd go_blog
   go mod tidy
   ```

2. Veritabanı bağlantınızı main.go içinde bulunan dsn değerine atayın.Ve go_blog adında bir veritabanı oluşturun.

3. Derleme ve çalıştırma:
```bash 
go build -o go_blog ./cmd/go_blog && ./go_blog
```
</details>

