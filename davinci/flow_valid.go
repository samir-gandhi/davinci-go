package davinci

import (
	"encoding/json"
)

func validateFlowJSONType(data []byte, exportType interface{}) (ok bool) {

	if ok := json.Valid(data); !ok {
		return ok
	}

	if err := json.Unmarshal(data, &exportType); err != nil {
		return false
	}

	jsonBytes, err := json.Marshal(exportType)
	if err != nil {
		return false
	}

	if string(jsonBytes) == "{}" {
		return false
	}

	return true
}

func ValidFlowsInfoJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := validateFlowJSONType(data, FlowsInfo{}); !ok {
		return false
	}

	var flowTypeObject FlowsInfo

	if err := json.Unmarshal([]byte(data), &flowTypeObject); err != nil {
		return false
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := FlowsInfo{}

		if ok := FlowsInfoEqual(flowTypeObject, empty, ExportCmpOpts{
			IgnoreUnmappedProperties:  false,
			IgnoreEnvironmentMetadata: true,
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
		}); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidFlowInfoJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := validateFlowJSONType(data, FlowInfo{}); !ok {
		return false
	}

	var flowTypeObject FlowInfo

	if err := json.Unmarshal([]byte(data), &flowTypeObject); err != nil {
		return false
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := FlowInfo{}

		if ok := FlowInfoEqual(flowTypeObject, empty, ExportCmpOpts{
			IgnoreUnmappedProperties:  false,
			IgnoreEnvironmentMetadata: true,
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
		}); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidFlowJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := validateFlowJSONType(data, Flow{}); !ok {
		return false
	}

	var flowTypeObject Flow

	if err := json.Unmarshal([]byte(data), &flowTypeObject); err != nil {
		return false
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := Flow{}

		if ok := FlowEqual(flowTypeObject, empty, ExportCmpOpts{
			IgnoreUnmappedProperties:  false,
			IgnoreEnvironmentMetadata: true,
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
		}); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	return ValidFlowJSON(data, cmpOpts) || ValidFlowInfoJSON(data, cmpOpts) || ValidFlowsInfoJSON(data, cmpOpts)
}
