package cmpopts

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

type ExportCmpOpts struct {
	IgnoreUnmappedProperties  bool
	IgnoreEnvironmentMetadata bool
	IgnoreConfig              bool
	IgnoreDesignerCues        bool
}

func ExportCmpFilters(opts ExportCmpOpts) []cmp.Option {
	filters := []cmp.Option{}
	if opts.IgnoreUnmappedProperties {
		filters = append(filters, IgnoreExportUnmappedProperties()...)
	}
	if opts.IgnoreEnvironmentMetadata {
		filters = append(filters, IgnoreExportEnvironmentMetadata()...)
	}
	if opts.IgnoreConfig {
		filters = append(filters, IgnoreExportEnvironmentConfig()...)
	}
	if opts.IgnoreDesignerCues {
		filters = append(filters, IgnoreExportDesignerCues()...)
	}
	return filters
}

func IgnoreExportUnmappedProperties() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(davinci.Data{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Edge{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Elements{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Flow{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.FlowVariable{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.GraphData{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Node{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.NodeData{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.OutputSchema{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Pan{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Position{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Properties{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Renderer{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.SaveFlowVariables{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.SubFlowID{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.SubFlowValue{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.SubFlowVersionID{}, "AdditionalProperties"),
		cmpopts.IgnoreFields(davinci.Trigger{}, "AdditionalProperties"),
	}
}

func IgnoreExportEnvironmentMetadata() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(davinci.Flow{},
			"AuthTokenExpireIds",
			"CompanyID",
			"Connections",
			"CreatedDate",
			"CurrentVersion",
			"CustomerID",
			"DeployedDate",
			"FlowID",
			"PublishedVersion",
			"SavedDate",
			"UpdatedDate",
			"VersionID",
		),
	}
}

func IgnoreExportEnvironmentConfig() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreFields(davinci.Flow{},
			"ConnectorIds",
			"Description",
			"EnabledGraphData",
			"FlowColor",
			"FlowStatus",
			"FunctionConnectionID",
			"GraphData",
			"InputSchema",
			"InputSchemaCompiled",
			"IsInputSchemaSaved",
			"IsOutputSchemaSaved",
			"Name",
			"Orx",
			"OutputSchema",
			"OutputSchemaCompiled",
			"Settings",
			"Timeouts",
			"Trigger",
			"Variables",
		),
	}
}

func IgnoreExportDesignerCues() []cmp.Option {
	return []cmp.Option{}
}
