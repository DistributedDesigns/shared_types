Shared types
====
[![Build Status](https://travis-ci.org/DistributedDesigns/shared_types.svg?branch=master)](https://travis-ci.org/DistributedDesigns/shared_types)

Objects shared between services. Provides common serialization <-> deserialization. Particularly useful when passing objects via RMQ.

#### `QuoteRequest`
```go
type QuoteRequest struct {
	Stock      string
	UserID     string
	AllowCache bool
}

qr := QuoteRequest{ "AAPL", "jappleseed", true}

qr.ToCSV() // "AAPL,jappleseed,true"

qr2, error := ParseQuoteRequest("AAPL,jappleseed,true")
```

#### `Quote`
- `Price` is serialized as a `float`
- `Timestamp` is serialized as milliseconds from the epoch

```go
type Quote struct {
	Price     currency.Currency
	Stock     string
	UserID    string
	Timestamp time.Time
	Cryptokey string
}

tenDollars, _ := currency.NewFromFloat(10.0)
q := Quote{tenDollars, "AAPL", "jappleseed", time.Now(), "abc123="}

q.ToCSV() // "10.00,AAPL,jappleseed,123123456,abc123="

q2, error := ParseQuote("10.00,AAPL,jappleseed,123123456,abc123=")
```
