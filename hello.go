package gomodhello

import "rsc.io/quote/v3"

// Hello ...
func Hello() string {
    return quote.HelloV3()
}

// MyHello ...
func MyHello() string {
    return "Hello from Rostov!"
}

// Proverb ...
func Proverb() string {
    return quote.Concurrency()
}
