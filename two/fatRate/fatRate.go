package main

import (
	"fmt"
	"time"

	fatRate "github.com/hjtcn/benbird-and-annie-go/module_two/practice/refactorFatRate"
)

func main() {
	for {
		fatRate.MainFatRateBody()

		if cont := fatRate.WhetherContinue(); !cont {
			break
		}
	}
	time.Now()

	fmt.Println("已退出体脂计算！\n")
}
