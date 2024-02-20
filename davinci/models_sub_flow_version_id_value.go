package davinci

import (
	"encoding/json"
	"fmt"
)

var (
	_ DaVinciExportModel = SubFlowVersionIDValue{}
)

type _SubFlowVersionIDValue SubFlowVersionIDValue
type SubFlowVersionIDValue struct {
	ValueFloat64 *float64
	ValueInt     *int32
	ValueString  *string
}

func (o SubFlowVersionIDValue) MarshalJSON() ([]byte, error) {
	if o.ValueFloat64 != nil {
		return json.Marshal(&o.ValueFloat64)
	}

	if o.ValueInt != nil {
		return json.Marshal(&o.ValueInt)
	}

	if o.ValueString != nil {
		return json.Marshal(&o.ValueString)
	}

	return nil, nil // no data in oneOf schemas
}

func (o *SubFlowVersionIDValue) UnmarshalJSON(bytes []byte) (err error) {

	match := false
	// try to unmarshal data into ValueFloat64
	err = newStrictDecoder(bytes).Decode(&o.ValueFloat64)
	if err == nil {
		jsonValueFloat64, _ := json.Marshal(o.ValueFloat64)
		if string(jsonValueFloat64) == "{}" { // empty struct
			o.ValueFloat64 = nil
		} else {
			match = true
		}
	} else {
		o.ValueFloat64 = nil
	}

	if !match {
		// try to unmarshal data into ValueInt
		err = newStrictDecoder(bytes).Decode(&o.ValueInt)
		if err == nil {
			jsonValueInt32, _ := json.Marshal(o.ValueInt)
			if string(jsonValueInt32) == "{}" { // empty struct
				o.ValueInt = nil
			} else {
				match = true
			}
		} else {
			o.ValueInt = nil
		}
	}

	if !match {
		// try to unmarshal data into ValueFloat64
		err = newStrictDecoder(bytes).Decode(&o.ValueFloat64)
		if err == nil {
			jsonValueFloat64, _ := json.Marshal(o.ValueFloat64)
			if string(jsonValueFloat64) == "{}" { // empty struct
				o.ValueFloat64 = nil
			} else {
				match = true
			}
		} else {
			o.ValueFloat64 = nil
		}
	}

	if !match { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubFlowVersionIDValue)")
	}

	return nil
}

// DesignerCuesFields implements DaVinciExportModel.
func (o SubFlowVersionIDValue) DesignerCuesFields() []string {
	return []string{}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o SubFlowVersionIDValue) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o SubFlowVersionIDValue) FlowConfigFields() []string {
	return []string{}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o SubFlowVersionIDValue) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o SubFlowVersionIDValue) VersionMetadataFields() []string {
	return []string{}
}
