package types

type Value int

const (
	Array Value = iota
	SimpleArray
	JsonArray
	Json
	BigInt
	Boolean
	Datetime
	DatetimeInmutable
	DatetimeTZ
	Date
	DateInmutable
	Time
	TimeInmutable
	Decimal
	Integer
	Object
	SmallInt
	String
	Text
	Binary
	Blob
	Float
	Guid
	DateInterval
)
