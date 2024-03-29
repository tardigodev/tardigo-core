package constants

type PluginType string

const (
	PluginTypeSourceStorage PluginType = "source_storage"
	PluginTypeTargetStorage PluginType = "target_storage"

	PluginTypeSourceParser PluginType = "source_parser"
	PluginTypeTargetParser PluginType = "target_parser"

	PluginTypeProcessor PluginType = "processor"
)
