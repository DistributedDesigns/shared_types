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
	ID         uint64
}

qr := QuoteRequest{ "AAPL", "jappleseed", true, 1}

qr.ToCSV() // "AAPL,jappleseed,true,1"

qr2, err := ParseQuoteRequest("AAPL,jappleseed,true,1")
```

#### `Quote`
- `Price` is serialized as a `float`
- `Timestamp` is serialized as **milliseconds** from the epoch

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

q.ToCSV() // "10.00,AAPL,jappleseed,123456789,abc123="

q2, err := ParseQuote("10.00,AAPL,jappleseed,123456789,abc123=")
```

#### `AuditEvent`
- Format for passing items through the `audit_event` RMQ channel
- Should map to the columns in the `Logs` table
- `EventType` should be an enum with entries
	- `command`: User commands. One per transaction!
	- `account_action`: Changes to a user's account state.
	- `quote`: Shouldn't have to be used since quotes aren't logged using the `audit_event` RMQ channel
- `Content` can be the formatted xml for commands or quotes, or semi-structured text for account actions

```go
type AuditEvent struct {
	UserID    string
	ID        uint64
	EventType string
	Content   string
}

ae := AuditEvent{"jappleseed", uint64(1), "command", "DO THIS NOW"}

ae.ToCSV() // "jappleseed,1,command,DO THIS NOW"

ae2, err := ParseAuditEvent("jappleseed,1,command,DO THIS NOW")
```

#### `DumplogRequest`
- `Filename` specifies the name for the create file on the audit log server

```go
type DumplogRequest struct {
	UserID   string
	Filename string
}

dr := DumplogRequest{"jappleseed", "ja-dumplog"}

dr.ToCSV() // "jappleseed,ja-dumplog"

dr2, err := ParseDumplogRequest("jappleseed,ja-dumplog")
```
