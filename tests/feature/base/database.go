package base

import (
	"context"
	"golang-blog/database/mongodb"
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
