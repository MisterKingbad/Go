package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "713789892198321789231823183187137893113298713297813789137123231723178")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}