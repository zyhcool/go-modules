package hello

import "rsc.io/quote/v3"

func HelloV2() string {
	return quote.HelloV3()
}

func ProverbV2() string {
	return quote.Concurrency()
}
