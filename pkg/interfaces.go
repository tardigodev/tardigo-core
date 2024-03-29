package pkg

import (
	"io"

	"github.com/tardigodev/tardigo-core/pkg/objects"

	dtypes "github.com/tardigodev/tardigo-core/pkg/dtypes"
)

type AddRecord func(any, objects.RecordDetail) error

type Plugin interface {
	GetPluginDetail() objects.PluginDetail
}

type SourceParserPlugin interface {
	Plugin
	GetRecord(io.Reader, objects.ReaderDetail, AddRecord) error
	GetSchema(io.Reader, objects.ReaderDetail) (dtypes.Schema, error)
}

type SourceReaderPlugin interface {
	Plugin
	GetReader(func(io.Reader, objects.ReaderDetail) error) error
}

type ProcessorPlugin interface {
	Plugin
	ProcessRecord(any, objects.RecordDetail, AddRecord)
}

type TargetParserPlugin interface {
	Plugin
	PutRecord(io.Writer, objects.WriterDetail, any, objects.RecordDetail, dtypes.Schema, AddRecord) error
	ConvertSchema(dtypes.Schema) (dtypes.Schema, error)
}

type TargetWriterPlugin interface {
	Plugin
	GetWriter(func(io.Writer, objects.WriterDetail) error) error
}
