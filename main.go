package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
)

const (
	etheriumWalletAdrress = "0xF4Ef9C5413d27466277dB2043C8169b193B7DBF1"
	hourInfomation        = 6
)

func main() {

	var standerdHashRate int
	var varianceAsPresentage int
	var etheriumWalletAdrress string

	flag.IntVar(&standerdHashRate, "hash", 20, "your suspected hash rate")
	flag.IntVar(&varianceAsPresentage, "var", 50, "tolarable variance")
	flag.StringVar(&etheriumWalletAdrress, "wallet", "0xF4Ef9C5413d27466277dB2043C8169b193B7DBF1", "your Etherium Wallet")
	flag.Parse()

	var lowestAcceptableHashRate = float64(standerdHashRate) * (float64(varianceAsPresentage) / 100)

	for {
		poolAPIResponse, err := http.Get("https://api.nanopool.org/v1/eth/avghashratelimited/" +
			etheriumWalletAdrress + "/" + strconv.Itoa(hourInfomation))
		checkPanic(err)
		defer poolAPIResponse.Body.Close()

		responseBodyData, err := ioutil.ReadAll(poolAPIResponse.Body)
		if err == nil {
			averageHashRate, err := jsonparser.GetFloat(responseBodyData, "data")
			checkPanic(err)

			if averageHashRate < lowestAcceptableHashRate {
				fmt.Println("Error With System")
			} else {
				fmt.Println("HashRate Ok")
			}
			fmt.Println("your hash rate is : {0}", averageHashRate)
		} else {
			fmt.Printf("somthing is wrong with nanopool api ")
			panic(err)
		}
		time.Sleep(30 * time.Minute)
	}

}

func checkPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
