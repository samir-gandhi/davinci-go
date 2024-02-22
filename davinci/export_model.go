package davinci

type DaVinciExportModel interface {
	EnvironmentMetadataFields() []string
	FlowMetadataFields() []string
	VersionMetadataFields() []string
	FlowConfigFields() []string
	DesignerCuesFields() []string
	SetAdditionalProperties(map[string]interface{})
}
