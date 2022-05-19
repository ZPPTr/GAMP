package main

import (
	"context"
	"fmt"
	"github.com/AlekSi/pointer"
	gampClient "github.com/olebedev/go-gamp"
	"github.com/olebedev/go-gamp/client/gampops"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func sendData(gampClient *gampops.Client) error {
	fmt.Println("...initialized")

	ch := make(chan error)

	url := "https://api.apilayer.com/exchangerates_data/latest?symbols=usd,chf,jpy,gbp&base=eur"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", os.Getenv("APILAYER_API_KEY"))

	go func() {
		for {
			res, err := client.Do(req)
			if res.Body != nil {
				defer res.Body.Close()
			}
			body, err := ioutil.ReadAll(res.Body)

			err = gampClient.Collect(
				gampops.NewCollectParams().
					WithCid(pointer.ToString("266035594")).
					WithT("event").
					WithEc(pointer.ToString(string(body))).
					WithEa(pointer.ToString("data_sent")).
					WithUa(pointer.ToString("local pc")).
					WithDp(pointer.ToString("/test")),
			)
			if err != nil {
				ch <- err
				break
			}

			fmt.Printf("%v\n ...sleeping an hour", string(body))
			time.Sleep(10 * time.Minute)
		}
	}()

	return <-ch
}

func main() {
	err := sendData(gampClient.New(context.Background(), "UA-104390508-1"))

	fmt.Println(err)
}
