package main

import (
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain(make(map[string]int64))
	defer bc.Save()
	bc.Load()

	handlers := api.NewHandlers(&bc)
	server := api.NewServer(handlers)
	server.SetupRoutes()
	server.Run()

	//privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//signature := GenerateSignature(privateKey)
	//m, _ := json.Marshal("1")
	//h := sha256.Sum256(m)
	//fmt.Println(ecdsa.Verify(&privateKey.PublicKey, h[:], signature.R, signature.S))
}

//type Signature struct {
//	R *big.Int
//	S *big.Int
//}
//
//func GenerateSignature(private *ecdsa.PrivateKey) *Signature {
//	m, _ := json.Marshal("1")
//
//	log.Println("Generate signature", string(m))
//
//	h := sha256.Sum256([]byte(m))
//	r, s, err := ecdsa.Sign(rand.Reader, private, h[:])
//	if err != nil {
//		fmt.Println("Signing signature failed: ", err)
//		return nil
//	}
//	return &Signature{R: r, S: s}
//}
