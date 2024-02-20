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
	return cmp.Equal(x, y, ExportCmpFilters(cmpOpts)...)
}

func FlowInfoEqual(x, y FlowInfo, cmpOpts ExportCmpOpts) bool {
	return cmp.Equal(x, y, ExportCmpFilters(cmpOpts)...)
}

func FlowEqual(x, y Flow, cmpOpts ExportCmpOpts) bool {
	return cmp.Equal(x, y, ExportCmpFilters(cmpOpts)...)
}

func Equal(x, y interface{}, cmpOpts ExportCmpOpts) bool {
	return cmp.Equal(x, y, ExportCmpFilters(cmpOpts)...)
}

func ExportCmpFilters(opts ExportCmpOpts) []cmp.Option {
	filters := []cmp.Option{}

	for _, v := range flowObjects() {
		if em, ok := v.(DaVinciExportModel); ok {
			filters = append(filters, IgnoreExportFields(em, opts))
		}
	}

	filters = append(filters, standardOptions()...)

	return filters
}

func standardOptions() []cmp.Option {
	r := make([]cmp.Option, 0)

	r = append(r, cmpopts.EquateEmpty())

	return r
}

func flowObjects() []interface{} {
	return []interface{}{
		Data{},
		Edge{},
		EdgeData{},
		Elements{},
		Flow{},
		FlowVariable{},
		FlowVariableFields{},
		GraphData{},
		LabelValue{},
		Node{},
		NodeData{},
		OutputSchema{},
		Pan{},
		Position{},
		Properties{},
		Renderer{},
		SaveFlowVariables{},
		SubFlowID{},
		SubFlowProperties{},
		SubFlowValue{},
		SubFlowVersionID{},
		Trigger{},
	}
}

func IgnoreExportFields(typ DaVinciExportModel, opts ExportCmpOpts) cmp.Option {

	names := make([]string, 0)
	if opts.IgnoreConfig {
		names = append(names, typ.FlowConfigFields()...)
	}
	if opts.IgnoreDesignerCues {
		names = append(names, typ.DesignerCuesFields()...)
	}
	if opts.IgnoreEnvironmentMetadata {
		names = append(names, typ.EnvironmentMetadataFields()...)
	}
	if opts.IgnoreUnmappedProperties {
		names = append(names, "AdditionalProperties")
	}
	if opts.IgnoreVersionMetadata {
		names = append(names, typ.VersionMetadataFields()...)
	}
	if opts.IgnoreFlowMetadata {
		names = append(names, typ.FlowMetadataFields()...)
	}

	return cmpopts.IgnoreFields(typ, names...)
}
