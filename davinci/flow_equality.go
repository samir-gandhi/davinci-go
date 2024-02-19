package davinci

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type ExportCmpOpts struct {
	IgnoreConfig              bool
	IgnoreDesignerCues        bool
	IgnoreEnvironmentMetadata bool
	IgnoreUnmappedProperties  bool
	IgnoreVersionMetadata     bool
	IgnoreFlowMetadata        bool
}

func FlowsInfoEqual(x, y FlowsInfo, cmpOpts ExportCmpOpts) bool {
	return cmp.Equal(x, y, exportCmpFilters(cmpOpts)...)
}

func FlowInfoEqual(x, y FlowInfo, cmpOpts ExportCmpOpts) bool {
	return cmp.Equal(x, y, exportCmpFilters(cmpOpts)...)
}

func FlowEqual(x, y Flow, cmpOpts ExportCmpOpts) bool {
	return cmp.Equal(x, y, exportCmpFilters(cmpOpts)...)
}

func Equal(x, y interface{}, cmpOpts ExportCmpOpts) bool {
	return cmp.Equal(x, y, exportCmpFilters(cmpOpts)...)
}

func exportCmpFilters(opts ExportCmpOpts) []cmp.Option {
	filters := []cmp.Option{}
	if opts.IgnoreConfig {
		filters = append(filters, IgnoreExportFlowConfig()...)
	}
	if opts.IgnoreDesignerCues {
		filters = append(filters, IgnoreExportDesignerCues()...)
	}
	if opts.IgnoreEnvironmentMetadata {
		filters = append(filters, IgnoreExportEnvironmentMetadata()...)
	}
	if opts.IgnoreUnmappedProperties {
		filters = append(filters, IgnoreExportUnmappedProperties()...)
	}
	if opts.IgnoreVersionMetadata {
		filters = append(filters, IgnoreExportVersionMetadata()...)
	}
	if opts.IgnoreFlowMetadata {
		filters = append(filters, IgnoreExportFlowMetadata()...)
	}

	filters = append(filters, standardOptions()...)

	return filters
}

func IgnoreExportUnmappedProperties() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(Data{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Edge{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Elements{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Flow{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(FlowVariable{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(GraphData{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Node{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(NodeData{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(OutputSchema{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Pan{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Position{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Properties{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Renderer{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(SaveFlowVariables{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(SubFlowID{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(SubFlowValue{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(SubFlowVersionID{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(Trigger{}, "AdditionalProperties"),
	}
}

func IgnoreExportEnvironmentMetadata() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(FlowEnvironmentMetadata{},
			"CompanyID",
			"CreatedDate",
			"CustomerID",
			"FlowID",
		),
	}
}

func IgnoreExportFlowMetadata() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(FlowMetadata{},
			"authTokenExpireIds",
			"connections",
			"connectorIds",
			"description",
			"enabledGraphData",
			"functionConnectionId",
			"name",
			"orx",
			"timeouts",
			"variables",
		),
	}
}

func IgnoreExportVersionMetadata() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(FlowVersionMetadata{},
			"currentVersion",
			"deployedDate",
			"flowStatus",
			"publishedVersion",
			"savedDate",
			"updatedDate",
			"versionId",
		),
	}
}

func IgnoreExportFlowConfig() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(FlowConfiguration{},
			"flowColor",
			"inputSchemaCompiled",
			"isInputSchemaSaved",
			"isOutputSchemaSaved",
			"outputSchemaCompiled",
		),
		cmpopts.IgnoreFields(FlowUpdateConfiguration{},
			"graphData",
			"inputSchema",
			"outputSchema",
			"settings",
			"trigger",
		),
	}
}

func IgnoreExportDesignerCues() []cmp.Option {
	return []cmp.Option{}
}

func standardOptions() []cmp.Option {
	return []cmp.Option{}
}
