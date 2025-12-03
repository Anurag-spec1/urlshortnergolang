package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDb = make(map[string]URL)

func generateshorturl(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("encode to string: ", hash)
	fmt.Println("final string: ", hash[:8])
	return hash[:8]
}

func createurl(originalURL string)string{
	shortURL:=generateshorturl(originalURL)
	id:=shortURL
	urlDb[id]=URL{
		ID   :         id,
		OriginalURL  :originalURL,
		ShortURL  :  shortURL,
		CreationDate : time.Now(),
}
return shortURL
	}

	func getURL(id string)(URL,error){
		url,ok:=urlDb[id]
		if !ok{
			return URL{},errors.New("URL NOT FOUND")
		}
		return url,nil
	}

func main() {
	OriginalURL := "https://github.com/Prince-1501/"
	generateshorturl(OriginalURL)
}
