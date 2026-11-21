package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JspBack/iyzipay-go"
)

// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func main() {
	ctx := context.Background()
	client, err := iyzipay.New(os.Getenv("IYZIPAY_API_KEY"), os.Getenv("IYZIPAY_API_SECRET"), iyzipay.WithBaseURL(iyzipay.SandboxURL))
	if err != nil {
		fmt.Println(err)
		return
	}

	product, err := client.CreateSubscriptionProduct(ctx, iyzipay.CreateSubscriptionProductRequest{
		Name:        fmt.Sprintf("Sandbox product %d", time.Now().UnixNano()),
		Description: "Create, get, and delete example",
	})
	if err != nil {
		fmt.Println("Create product:", err)
		return
	}
	reference := product.Data.ReferenceCode
	fmt.Println("Created product:", reference)

	detail, err := client.GetSubscriptionProductDetail(ctx, iyzipay.GetSubscriptionProductDetailRequest{
		ProductReferenceCode: reference,
	})
	if err != nil {
		fmt.Println("Get product:", err)
		return
	}
	fmt.Println("Product name:", detail.Data.Name)

	_, err = client.DeleteSubscriptionProduct(ctx, iyzipay.DeleteSubscriptionProductRequest{
		ProductReferenceCode: reference,
	})
	if err != nil {
		fmt.Printf("Delete product %s: %v\n", reference, err)
		return
	}
	fmt.Println("Deleted product:", reference)
}
