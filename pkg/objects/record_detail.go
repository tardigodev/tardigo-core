package objects

import "github.com/tardigodev/tardigo-core/pkg/constants"

type RecordDetail struct {
	ReaderDetail
	WriterDetail
	RecordType   constants.RecordType
	RecordErrors []error
}
