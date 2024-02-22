package davinci

var (
	_ DaVinciExportModel = SubFlowValue{}
)

type SubFlowValue LabelValue

// DesignerCuesFields implements DaVinciExportModel.
func (o SubFlowValue) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o SubFlowValue) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o SubFlowValue) FlowConfigFields() []string {
	return []string{
		"Label",
		"Value",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o SubFlowValue) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o SubFlowValue) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o SubFlowValue) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
