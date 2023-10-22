package wallet

import (
	"encoding/json"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Wallets struct {
	Wallets []*Wallet
}

func NewWallets() *Wallets {
	return &Wallets{
		Wallets: make([]*Wallet, 0),
	}
}
func (w *Wallets) GetWallet(address string) *Wallet {
	for _, wallet := range w.Wallets {
		if wallet.Address == address {
			return wallet
		}
	}
	return nil
}

func (w *Wallets) NewWallet() *Wallet {
	wallet := NewWallet()
	w.Wallets = append(w.Wallets, wallet)
	return wallet
}
func (w *Wallets) Save() {
	m, err := json.Marshal(w)
	if err != nil {
		log.Println("error: ", err)
	}
	err = utils.Save(m, "wallet")
	if err != nil {
		log.Println("error: ", err)
	}

}
func (w *Wallets) Load() {
	err := utils.Load(&w, "wallet")
	if err != nil {
		log.Println("error: ", err)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		w.Save()
		os.Exit(1)
	}()
}
