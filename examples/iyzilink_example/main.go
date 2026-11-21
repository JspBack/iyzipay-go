package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/JspBack/iyzipay-go"
)

func main() {
	ctx := context.Background()
	client, err := iyzipay.New(os.Getenv("IYZIPAY_API_KEY"), os.Getenv("IYZIPAY_API_SECRET"), iyzipay.WithBaseURL(iyzipay.SandboxURL))
	if err != nil {
		fmt.Println(err)
		return
	}

	image, err := os.ReadFile("../../gopher.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	link, err := client.IyzilinkCreate(ctx, &iyzipay.IyzilinkCreateRequest{
		Name:             "Sandbox example",
		Price:            "1.00",
		CurrencyCode:     "TRY",
		EncodedImageFile: base64.StdEncoding.EncodeToString(image),
		Description:      "Create, get, and delete example",
		StockCount:       1,
		StockEnabled:     true,
	})
	if err != nil {
		fmt.Println("Create link:", err)
		return
	}
	token := link.Data.Token
	fmt.Println("Created link:", link.Data.Url)

	detail, err := client.IyzilinkGetDetail(ctx, &iyzipay.IyzilinkGetDetailRequest{
		Token: token,
		Page:  1,
		Count: 50,
	})
	if err != nil {
		fmt.Println("Get link:", err)
		return
	}
	fmt.Printf("Link details: %+v\n", detail.Data)

	_, err = client.IyzilinkDelete(ctx, &iyzipay.IyzilinkDeleteRequest{Token: token})
	if err != nil {
		fmt.Printf("Delete link %s: %v\n", token, err)
		return
	}
	fmt.Println("Deleted link.")
}
