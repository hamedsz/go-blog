package base

import (
	"context"
	"github.com/hamedsz/go-blog/internal/database/mongodb"
	"time"
)

func ResetDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := mongodb.GetDB().Drop(ctx)
	if err != nil{
		panic(err)
	}
}
