/*
Package timestamps is a simple library that creates timestamps and datetimes in a specific format
Its use is two fold:
1) Simplify code since a one liner can now be used instead of the code blocks below (very minimal difference, I know).
2) Reduce formatting mistakes since calls to these func will always return a datetime or timestamp with the same format.
*/

package timestamps

import "time"

//*********************************************************************************************************************************
//FUNCS

//GENERATE UNIX TIMESTAMP
func Unix() int {
	t := time.Now()
	s := t.Unix()
	return int(s)
}

//GENERATE ISO8601 DATETIME
//format: YY-MM-DDTHH:MM:SS.mmmZ
func ISO8601() string {
	t := time.Now().UTC()
	s := t.Format("2006-01-02T15:04:05.000Z")
	return s
}
