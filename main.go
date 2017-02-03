package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	if "GET" == r.Method {
		//	title := r.URL.Path[len("/view/"):]
		//	p, _ := loadPage(title)
		//	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
		q := r.URL.Query()["q"][0]
		n := r.URL.Query()["n"][0]
		query := GCMDecrypt(q, n)
		b, _ := r.URL.Parse(query)
		body := b.Query()
		fmt.Println(body["p"][0])
		var temp = make(map[string]string)
		temp["username"] = body["a"][0]
		temp["password"] = body["p"][0]
		temp["server"] = body["s"][0]
		temp["ipAddr"] = "10." + body["a"][0] + ".9" //account is in ###.### format,
		temp["ipNet"] = "10." + body["a"][0] + ".0/24"
		// if it doesn't panic at this point, means the data is correct
		f, err := os.Create("/tmp/config/network")
		if err != nil {
			log.Println("create file: ", err)
			return
		}
		tmpl, err := template.New("test").Parse(network)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(f, temp)
		f.Close()
		f, err = os.Create("/tmp/ipsec.secrets")
		tmpl, err = template.New("test").Parse(secrets)
		err = tmpl.Execute(f, temp)
		f.Close()
		f, err = os.Create("/tmp/ipsec.conf")
		tmpl, err = template.New("test").Parse(conf)
		err = tmpl.Execute(f, temp)
		f.Close()

		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", r.URL, body)
	}
}
func GCMDecrypt(cipherText, n string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("3zTvzr3p67VC61jmV54rIYu1545x4TlY")
	ciphertext, _ := hex.DecodeString(cipherText)

	fmt.Println(n)
	nonce, _ := hex.DecodeString(n)
	//nonce := []byte("60iP0h6vJoEa")
	fmt.Printf("%x\n", nonce)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("localhost:29080", nil)
}
