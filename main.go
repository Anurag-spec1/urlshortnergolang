package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

	func handler(w http.ResponseWriter,r*http.Request){
		fmt.Fprintf(w,"Anurag's Territory")
	}
	
	func shorturlhanndler(w http.ResponseWriter,r *http.Request){
		var data struct{
			URL string `json:"url"`
		}
		err:=json.NewDecoder(r.Body).Decode(&data)
		if err!=nil{
			http.Error(w,"Invalid request body",http.StatusBadRequest)
			return
		}
		shortURL_:=createurl(data.URL)
		// fmt.Fprintf(w,shortURL)
		response:=struct{
			ShortURL string `json:"short_url"`
		}{ShortURL: shortURL_}

		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(response)
	}
	

	func redirecturlhandler(w http.ResponseWriter,r*http.Request){
		id:=r.URL.Path[len("/redirect/"):]
		url,err:=getURL(id)
		if err !=nil{
			http.Error(w,"Invalid request",http.StatusNotFound)
		}
		http.Redirect(w,r,url.OriginalURL,http.StatusFound)
	}

func main() {
	// OriginalURL := "https://github.com/Prince-1501/"
	// generateshorturl(OriginalURL)

	http.HandleFunc("/",handler)
	http.HandleFunc("/shorten",shorturlhanndler)
	http.HandleFunc("/redirect/",redirecturlhandler)

	fmt.Println("Server starting on 3000")
	err:=http.ListenAndServe(":3000",nil)
	if err!=nil{
		fmt.Println("Error on starting server:",err)
	}
}
