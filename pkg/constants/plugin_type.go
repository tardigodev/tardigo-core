package constants

type PluginType string

const (
	PluginTypeSourceReader PluginType = "source_reader"
	PluginTypeTargetWriter PluginType = "target_writer"

	PluginTypeSourceParser PluginType = "source_parser"
	PluginTypeTargetParser PluginType = "target_parser"

	PluginTypeProcessor PluginType = "processor"
)
