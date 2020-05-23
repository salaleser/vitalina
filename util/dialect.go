package util

type Dialect struct {
	// IETF is a IETF language tag
	IETF string
}

var Dialects = map[string]Dialect{
	"en-US": {"en-US"},
	"en-GB": {"en-GB"},
	"en-AU": {"en-AU"},
	"ru-RU": {"en-RU"},
}
