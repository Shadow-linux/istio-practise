package main

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"io/ioutil"
	"log"
	"os"
)

// openssl genrsa -out mypri.pem 2048
// openssl rsa  -in mypri.pem -pubout -out mypub.pem

//jwk 的字段描述
// alg： 具体的算法
// kty: 秘钥类型
// use: 可选,sig或enc(签名还是加密)
// kid :key的唯一标识
// e:秘钥的模
// n:秘钥指数
// iss: 发行人

func pubKey() []byte {
	f, _ := os.Open("./mypub.pem")
	b, _ := ioutil.ReadAll(f)

	return b
}
func main() {
	key, err := jwk.ParseKey(pubKey(), jwk.WithPEM(true))
	if err != nil {
		log.Fatal(err)
	}

	if pubKey, ok := key.(jwk.RSAPublicKey); ok {

		b, err := json.Marshal(pubKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))

	}

}

//
