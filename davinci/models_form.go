package davinci

import (
	"encoding/json"
	"fmt"
)

type _Form Form
type Form struct {
	ValueString  *string
	ValueFormObj *FormObj
}

func (o Form) MarshalJSON() ([]byte, error) {
	return o.marshal()
}

func (o *Form) UnmarshalJSON(bytes []byte) (err error) {
	return o.unmarshal(bytes)
}

func (o Form) MarshalDavinci(_ ExportCmpOpts) ([]byte, error) {
	return o.marshal()
}

func (o *Form) UnmarshalDavinci(bytes []byte, _ ExportCmpOpts) (err error) {
	return o.unmarshal(bytes)
}

func (o Form) marshal() ([]byte, error) {
	if o.ValueFormObj != nil {
		return json.Marshal(&o.ValueFormObj)
	}

	if o.ValueString != nil {
		return json.Marshal(&o.ValueString)
	}

	return nil, nil // no data in oneOf schemas
}

func (o *Form) unmarshal(bytes []byte) (err error) {

	match := false
	// try to unmarshal data into ValueFormObj
	err = newStrictDecoder(bytes).Decode(&o.ValueFormObj)
	if err == nil {
		jsonValueFormObj, _ := json.Marshal(o.ValueFormObj)
		if string(jsonValueFormObj) == "{}" { // empty struct
			o.ValueFormObj = nil
		} else {
			match = true
		}
	} else {
		o.ValueFormObj = nil
	}

	if !match {
		// try to unmarshal data into ValueString
		err = newStrictDecoder(bytes).Decode(&o.ValueString)
		if err == nil {
			jsonValueString, _ := json.Marshal(o.ValueString)
			if string(jsonValueString) == "{}" { // empty struct
				o.ValueString = nil
			} else {
				match = true
			}
		} else {
			o.ValueString = nil
		}
	}

	if !match { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Form)")
	}

	return nil
}
