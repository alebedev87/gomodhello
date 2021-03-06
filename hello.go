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

// TagCheck ...
func TagCheck() string {
    return "Tag check"
}

// TagFormat ...
func TagFormat() string {
    return "Can be w/o v"
}

// Proverb ...
func Proverb() string {
    return quote.Concurrency()
}
