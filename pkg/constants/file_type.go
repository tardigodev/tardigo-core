package constants

type StreamType uint8

const (
	StreamTypeOk StreamType = iota
	StreamTypeFailed
	StreamTypeEnd
)
