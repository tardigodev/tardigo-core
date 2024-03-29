package constants

type RecordType uint8

const (
	RecordTypeOk RecordType = iota
	RecordTypeFailed
	RecordTypeSkip
	RecordTypeSchema
	RecordTypeEnd
)
