# Golang API with MongoDB

Bu proje MongoDB veritabanı kullanarak kitaplar ve kullanıcılar için bir API sağlar.

## Kurulum

### Gereksinimler

- Go 1.24.7 veya üzeri
- MongoDB (yerel veya uzak)

### MongoDB Kurulumu

```bash
# Ubuntu/Debian için
sudo apt-get install mongodb

# macOS için (Homebrew ile)
brew install mongodb-community

# Docker ile
docker run -d -p 27017:27017 --name mongodb mongo:latest
```

### Projeyi Çalıştırma

1. Bağımlılıkları yükleyin:

```bash
go mod tidy
```

2. MongoDB'yi başlatın (eğer yerel kurulum varsa):

```bash
sudo systemctl start mongodb
# veya
mongod
```

3. Uygulamayı çalıştırın:

```bash
go run main.go
```

## Yapılandırma

Varsayılan olarak uygulama `mongodb://localhost:27017` adresindeki MongoDB'ye bağlanır.

Çevre değişkenleri ile yapılandırma:

```bash
export MONGODB_URI="mongodb://localhost:27017"
export DB_NAME="library_db"
export BOOKS_COLLECTION="books"
export USERS_COLLECTION="users"
```

## Veritabanı Yapısı

### Database: `library_db`

#### Collection: `books`

```json
{
	"_id": "ObjectId",
	"title": "string",
	"author": "string",
	"is_taken": "boolean"
}
```

#### Collection: `users`

```json
{
	"_id": "ObjectId",
	"username": "string",
	"password": "string",
	"books_taken": ["ObjectId"]
}
```

## Kullanım

### Örnek Kullanım

```bash
go run examples/example_usage.go
```

Bu örnek:

- MongoDB'ye bağlanır
- Örnek bir kitap ekler
- Örnek bir kullanıcı ekler

### Programatik Kullanım

```go
import "main/database"

// Veritabanına bağlan
err := database.ConnectToMongoDB()
if err != nil {
    log.Fatal(err)
}

// Collection'ları al
booksCollection := database.GetBooksCollection()
usersCollection := database.GetUsersCollection()

// İşlemlerinizi gerçekleştirin...
```

## Proje Yapısı

```
golang_api/
├── main.go              # Ana uygulama
├── models/              # Veri modelleri
│   ├── books.go        # Kitap modeli
│   └── users.go        # Kullanıcı modeli
├── database/           # Veritabanı bağlantısı
│   └── connection.go   # MongoDB bağlantı yönetimi
├── config/             # Yapılandırma
│   └── config.go       # Uygulama yapılandırması
└── examples/           # Örnek kullanımlar
    └── example_usage.go
```
