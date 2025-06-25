package regex

type regexTokenType uint8

const (
	literal regexTokenType = iota
	or                     = iota
)

type regexToken struct {
	tokenType regexTokenType
}
