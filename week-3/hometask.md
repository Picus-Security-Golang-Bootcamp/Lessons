## Homework | Week 3
`Not: Ödevi yeni bir repoya ekleyeceksiniz. Var olan reponuzda bir güncelleme olmayacak. "homework-2..." şeklinde yeni bir repo üzerinde çalışacaksınız.`


Elimizde bir kitap listesi var. 
Kitap alanları şöyle;
```
- Kitap ID
- Kitap Adı
- Sayfa Sayısı
- Stok Sayısı
- Fiyatı
- Stock Kodu
- ISBN
- Yazar bilgisi (ID ve İsim)
```

1. Tüm kitapları listele (list)
2. Verilen girdi hangi kitap isimlerinde geçiyorsa o kitapları listele (search)
3. ID'ye göre kitabı yazdır
4. IDsi verilen kitabı sil. (Silinen kitabın ID'ye göre geliyor olması gerekiyor.)
5. IDsi verilen kitabı istenilen adet kadar satın al ve kitabın son bilgilerini ekrana yazdır.

Yanlış komut girildiğinde ekrana usage'ı yazdıracak. 


Concurrency ile ilgili medium yazısı yazılacak. 

### list command
```
go run main.go list
```

### search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Ring: The Return of the King
```

### get command
```
go run main.go get <bookID>
go run main.go get 5
```

### delete command
```
go run main.go delete <bookID>
go run main.go delete 5
```

### buy command
```
go run main.go buy <bookID> <quantity>
go run main.go buy 5 2
```

###
# Requirements:
- README
- No third party package(s)
- Everything should be in English (Comments, Function names, File names, etc.)
- Use structs not maps
