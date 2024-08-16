package davinci

import (
	"encoding/json"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type EnumValidFlowObjError string

const (
	ENUM_VALID_FLOW_OBJ_ERROR_NONE                           EnumValidFlowObjError = "NONE"
	ENUM_VALID_FLOW_OBJ_ERROR_INVALID_JSON                   EnumValidFlowObjError = "INVALID_JSON"
	ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_UNMARSHAL            EnumValidFlowObjError = "CANNOT_DV_UNMARSHAL"
	ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_MARSHAL              EnumValidFlowObjError = "CANNOT_DV_MARSHAL"
	ENUM_VALID_FLOW_OBJ_ERROR_EMPTY_FLOW                     EnumValidFlowObjError = "EMPTY_FLOW"
	ENUM_VALID_FLOW_OBJ_ERROR_OBJECT_EQUATES_EMPTY           EnumValidFlowObjError = "OBJECT_EQUATES_EMPTY"
	ENUM_VALID_FLOW_OBJ_ERROR_NO_FLOW_DEF                    EnumValidFlowObjError = "NO_FLOW_DEF"
	ENUM_VALID_FLOW_OBJ_ERROR_MULTIPLE_FLOW_DEF              EnumValidFlowObjError = "MULTIPLE_FLOW_DEF"
	ENUM_VALID_FLOW_OBJ_ERROR_INVALID_REQUIRED_FLOW_DEF      EnumValidFlowObjError = "INVALID_REQUIRED_FLOW_DEF"
	ENUM_VALID_FLOW_OBJ_ERROR_UNKNOWN_ADDITIONAL_JSON_VALUES EnumValidFlowObjError = "UNKNOWN_ADDITIONAL_JSON_VALUES"
)

func ValidFlowsInfoExport(data []byte, cmpOpts ExportCmpOpts) (ok bool, errorCode EnumValidFlowObjError, diff *string, err error) {
	if ok := json.Valid(data); !ok {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_INVALID_JSON, nil, nil
	}

	var flowTypeObject FlowsInfo

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_UNMARSHAL, nil, err
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_MARSHAL, nil, err
	}

	if string(jsonBytes) == "{}" {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_EMPTY_FLOW, nil, nil
	}

	if cmp.Equal(flowTypeObject, FlowsInfo{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(flowTypeObject, FlowsInfo{}, cmpopts.EquateEmpty())
		return false, ENUM_VALID_FLOW_OBJ_ERROR_OBJECT_EQUATES_EMPTY, &diff, nil
	}

	if flowTypeObject.Flow == nil {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_NO_FLOW_DEF, nil, nil
	}

	if len(flowTypeObject.Flow) > 1 {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_MULTIPLE_FLOW_DEF, nil, nil
	}

	for _, flow := range flowTypeObject.Flow {
		if ok, diff := validateRequiredFlowAttributes(flow, cmpOpts); !ok {
			return false, ENUM_VALID_FLOW_OBJ_ERROR_INVALID_REQUIRED_FLOW_DEF, diff, nil
		}
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := FlowsInfo{
			Flow: []Flow{
				{
					FlowConfiguration: FlowConfiguration{
						FlowUpdateConfiguration: FlowUpdateConfiguration{
							GraphData: &GraphData{
								Elements: &Elements{
									Nodes: []Node{
										{
											Data: &NodeData{
												AdditionalProperties: map[string]interface{}{
													"test1": "test1", // to overcome odd behaviours with cmpopts
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		cmpOpts := ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreFlowMetadata:        true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
		}

		opts := cmpopts.IgnoreFields(Elements{}, "Nodes")

		if ok := Equal(empty, flowTypeObject, cmpOpts, opts); !ok {
			diff := Diff(empty, flowTypeObject, cmpOpts, opts)
			return false, ENUM_VALID_FLOW_OBJ_ERROR_UNKNOWN_ADDITIONAL_JSON_VALUES, &diff, nil
		}
	}

	// TODO validate required struct attributes
	return true, ENUM_VALID_FLOW_OBJ_ERROR_NONE, nil, nil
}

func ValidFlowInfoExport(data []byte, cmpOpts ExportCmpOpts) (ok bool, errorCode EnumValidFlowObjError, diff *string, err error) {
	if ok := json.Valid(data); !ok {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_INVALID_JSON, nil, nil
	}

	var flowTypeObject FlowInfo

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_UNMARSHAL, nil, err
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_MARSHAL, nil, err
	}

	if string(jsonBytes) == "{}" {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_EMPTY_FLOW, nil, nil
	}

	if cmp.Equal(flowTypeObject, FlowInfo{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(flowTypeObject, FlowInfo{}, cmpopts.EquateEmpty())
		return false, ENUM_VALID_FLOW_OBJ_ERROR_OBJECT_EQUATES_EMPTY, &diff, nil
	}

	if ok, diff := validateRequiredFlowAttributes(flowTypeObject.Flow, cmpOpts); !ok {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_INVALID_REQUIRED_FLOW_DEF, diff, nil
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := FlowInfo{
			Flow: Flow{
				FlowConfiguration: FlowConfiguration{
					FlowUpdateConfiguration: FlowUpdateConfiguration{
						GraphData: &GraphData{
							Elements: &Elements{
								Nodes: []Node{
									{
										Data: &NodeData{
											AdditionalProperties: map[string]interface{}{
												"test1": "test1", // to overcome odd behaviours with cmpopts
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		cmpOpts := ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		}

		opts := cmpopts.IgnoreFields(Elements{}, "Nodes")

		if ok := Equal(empty, flowTypeObject, cmpOpts, opts); !ok {
			diff := Diff(empty, flowTypeObject, cmpOpts, opts)
			return false, ENUM_VALID_FLOW_OBJ_ERROR_UNKNOWN_ADDITIONAL_JSON_VALUES, &diff, nil
		}
	}

	// TODO validate required struct attributes
	return true, ENUM_VALID_FLOW_OBJ_ERROR_NONE, nil, nil
}

func ValidFlowExport(data []byte, cmpOpts ExportCmpOpts) (ok bool, errorCode EnumValidFlowObjError, diff *string, err error) {
	if ok := json.Valid(data); !ok {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_INVALID_JSON, nil, nil
	}

	var flowTypeObject Flow

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_UNMARSHAL, nil, err
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_CANNOT_DV_MARSHAL, nil, err
	}

	if string(jsonBytes) == "{}" {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_EMPTY_FLOW, nil, nil
	}

	if cmp.Equal(flowTypeObject, Flow{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(flowTypeObject, Flow{}, cmpopts.EquateEmpty())
		return false, ENUM_VALID_FLOW_OBJ_ERROR_OBJECT_EQUATES_EMPTY, &diff, nil
	}

	if ok, diff := validateRequiredFlowAttributes(flowTypeObject, cmpOpts); !ok {
		return false, ENUM_VALID_FLOW_OBJ_ERROR_INVALID_REQUIRED_FLOW_DEF, diff, nil
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := Flow{
			FlowConfiguration: FlowConfiguration{
				FlowUpdateConfiguration: FlowUpdateConfiguration{
					GraphData: &GraphData{
						Elements: &Elements{
							Nodes: []Node{
								{
									Data: &NodeData{
										AdditionalProperties: map[string]interface{}{
											"test1": "test1", // to overcome odd behaviours with cmpopts
										},
									},
								},
							},
						},
					},
				},
			},
		}

		cmpOpts := ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
			IgnoreFlowVariables:       true,
		}

		opts := cmpopts.IgnoreFields(Elements{}, "Nodes")

		if ok := Equal(empty, flowTypeObject, cmpOpts, opts); !ok {
			diff := Diff(empty, flowTypeObject, cmpOpts, opts)
			return false, ENUM_VALID_FLOW_OBJ_ERROR_UNKNOWN_ADDITIONAL_JSON_VALUES, &diff, nil
		}
	}

	// TODO validate required struct attributes
	return true, ENUM_VALID_FLOW_OBJ_ERROR_NONE, nil, nil
}

func ValidExport(data []byte, cmpOpts ExportCmpOpts) (ok bool, errorCode EnumValidFlowObjError, diff *string, err error) {

	if ok, code, diff, err := ValidFlowExport(data, cmpOpts); !ok {
		return ok, code, diff, err
	}

	if ok, code, diff, err := ValidFlowInfoExport(data, cmpOpts); !ok {
		return ok, code, diff, err
	}

	if ok, code, diff, err := ValidFlowsInfoExport(data, cmpOpts); !ok {
		return ok, code, diff, err
	}

	return true, ENUM_VALID_FLOW_OBJ_ERROR_NONE, nil, nil
}

func validateRequiredFlowAttributes(v Flow, opts ExportCmpOpts) (ok bool, diff *string) {

	if !opts.IgnoreConfig && cmp.Equal(v.FlowConfiguration, FlowConfiguration{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(v.FlowConfiguration, FlowConfiguration{}, cmpopts.EquateEmpty())
		return false, &diff
	}

	if !opts.IgnoreFlowMetadata && cmp.Equal(v.FlowMetadata, FlowMetadata{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(v.FlowMetadata, FlowMetadata{}, cmpopts.EquateEmpty())
		return false, &diff
	}

	// TODO - anything else to validate?

	return true, nil
}
