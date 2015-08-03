# timestamps
Small Golang Library for Simplifying the Creation of Timestamps and Datetimes

Intended use of this library is two fold:
1) Reduced code since calls to get a timestamp or datetime become short 1-liners.
2) Reduced formatting mistakes since there is no memorization of formats.


### Unix()
Returns an int64 representation of the seconds since the unix epoch.

### ISO8601()
Returns a string representation of the current datetime in ISO8601 minimal format.
Ex.: 2006-02-01T15:04:05.000Z
