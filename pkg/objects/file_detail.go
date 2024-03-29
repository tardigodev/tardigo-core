package objects

import "github.com/tardigodev/tardigo-core/pkg/constants"

type ReaderDetail struct {
	ReaderSource string
	ReaderType   constants.StreamType
	ReaderError  error
}

type WriterDetail struct {
	WriterTarget string
	WriterType   constants.StreamType
	WriterError  error
}
