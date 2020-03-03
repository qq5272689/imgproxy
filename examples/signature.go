package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	key := "363274dd8085e0b1ac8d"
	salt := "bdb101445d746e5fa006"

	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}

	resize := "fit"
	width := 300
	height := 300
	gravity := "sm"
	enlarge := 1
	extension := "jpg"

	url := "local:///google.png"
	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(url))
	apath := fmt.Sprintf("/rs:%s:%d:%d:%d/g:%s/%s.%s", resize, width, height, enlarge, gravity, encodedURL, extension)
	//path := fmt.Sprintf("/%s/%d/%d/%s/%d/%s.%s", resize, width, height, gravity, enlarge, encodedURL, extension)
	//
	//mac := hmac.New(sha256.New, keyBin)
	//mac.Write(saltBin)
	//mac.Write([]byte(path))
	//signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	amac := hmac.New(sha256.New, keyBin)
	amac.Write(saltBin)
	amac.Write([]byte(apath))
	asignature := base64.RawURLEncoding.EncodeToString(amac.Sum(nil))

	fmt.Println(encodedURL)
	//fmt.Println(path)
	//fmt.Printf("/%s%s", signature, path)
	//fmt.Println()
	fmt.Println(apath)
	fmt.Printf("/%s%s", asignature, apath)
}
