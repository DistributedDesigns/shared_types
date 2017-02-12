Shared types
====
[![Build Status](https://travis-ci.org/DistributedDesigns/shared_types.svg?branch=master)](https://travis-ci.org/DistributedDesigns/shared_types)

Objects shared between services. Provides common serialization <-> deserialization. Particularly useful when passing objects via RMQ.

#### `QuoteRequest`
```go
type QuoteRequest struct {
	Stock  string
	UserID string
}

qr := QuoteRequest{ "AAPL", "jappleseed"}

qr.ToCSV() // "AAPL,jappleseed"

qr2, error := ParseQuoteRequest("AAPL,jappleseed")
```

#### `Quote`
- `Price` is serialized as a `float`
- `Timestamp` is serialized as milliseconds from the epoch

```go
type Quote struct {
	UserID    string
	Stock     string
	Price     currency.Currency
	Timestamp time.Time
	Cryptokey string
}

tenDollars, _ := currency.NewFromFloat(10.0)
q := Quote{"jappleseed", "AAPL", tenDollars, time.Now(), "abc123="}

q.ToCSV() // "jappleseed,AAPL,10.00,123123456,abc123="

q2, error := ParseQuote("jappleseed,AAPL,10.00,123123456,abc123=")
```
