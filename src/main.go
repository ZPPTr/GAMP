package main

import (
	"context"
	"github.com/AlekSi/pointer"
	gamp "github.com/olebedev/go-gamp"
	"github.com/olebedev/go-gamp/client/gampops"
	"log"
)

func main() {
	client := gamp.New(context.Background(), "UA-XXXXXXXX-X")
	err := client.Collect(
		gampops.NewCollectParams().
			WithCid(pointer.ToString("42")).
			WithT("event").
			WithEc(pointer.ToString("Category")).
			WithEa(pointer.ToString("Action")),
	)
	if err != nil {
		log.Fatal(err)
	}
}
