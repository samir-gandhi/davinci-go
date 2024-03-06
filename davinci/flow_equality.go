package davinci

import (
	"fmt"

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

func Equal(x, y interface{}, cmpOpts ExportCmpOpts, opts ...cmp.Option) bool {

	xBytes, err := Marshal(x, cmpOpts)
	if err != nil {
		return false
	}

	yBytes, err := Marshal(y, cmpOpts)
	if err != nil {
		return false
	}

	switch x.(type) {
	case FlowsInfo:
		var xFlowsInfo, yFlowsInfo FlowsInfo

		if err := Unmarshal(xBytes, &xFlowsInfo, cmpOpts); err != nil {
			return false
		}

		if err := Unmarshal(yBytes, &yFlowsInfo, cmpOpts); err != nil {
			return false
		}

		return cmp.Equal(xFlowsInfo, yFlowsInfo, standardOptions(opts...)...)
	case FlowInfo:
		var xFlowInfo, yFlowInfo FlowInfo

		if err := Unmarshal(xBytes, &xFlowInfo, cmpOpts); err != nil {
			return false
		}

		if err := Unmarshal(yBytes, &yFlowInfo, cmpOpts); err != nil {
			return false
		}

		return cmp.Equal(xFlowInfo, yFlowInfo, standardOptions(opts...)...)
	case Flow:
		var xFlow, yFlow Flow

		if err := Unmarshal(xBytes, &xFlow, cmpOpts); err != nil {
			return false
		}

		if err := Unmarshal(yBytes, &yFlow, cmpOpts); err != nil {
			return false
		}

		return cmp.Equal(xFlow, yFlow, standardOptions(opts...)...)
	}

	return cmp.Equal(x, y, standardOptions(opts...)...)
}

func Diff(x, y interface{}, cmpOpts ExportCmpOpts, opts ...cmp.Option) string {

	xBytes, err := Marshal(x, cmpOpts)
	if err != nil {
		panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
	}

	yBytes, err := Marshal(y, cmpOpts)
	if err != nil {
		panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
	}

	switch x.(type) {
	case FlowsInfo:
		var xFlowsInfo, yFlowsInfo FlowsInfo

		if err := Unmarshal(xBytes, &xFlowsInfo, cmpOpts); err != nil {
			panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
		}

		if err := Unmarshal(yBytes, &yFlowsInfo, cmpOpts); err != nil {
			panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
		}

		return cmp.Diff(xFlowsInfo, yFlowsInfo, standardOptions(opts...)...)
	case FlowInfo:
		var xFlowInfo, yFlowInfo FlowInfo

		if err := Unmarshal(xBytes, &xFlowInfo, cmpOpts); err != nil {
			panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
		}

		if err := Unmarshal(yBytes, &yFlowInfo, cmpOpts); err != nil {
			panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
		}

		return cmp.Diff(xFlowInfo, yFlowInfo, standardOptions(opts...)...)
	case Flow:
		var xFlow, yFlow Flow

		if err := Unmarshal(xBytes, &xFlow, cmpOpts); err != nil {
			panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
		}

		if err := Unmarshal(yBytes, &yFlow, cmpOpts); err != nil {
			panic(fmt.Sprintf("inconsistent difference and equality results %s", err))
		}

		return cmp.Diff(xFlow, yFlow, standardOptions(opts...)...)
	}

	return cmp.Diff(x, y, standardOptions(opts...)...)
}

func standardOptions(opts ...cmp.Option) []cmp.Option {
	r := make([]cmp.Option, 0)

	r = append(r, cmpopts.EquateEmpty())

	for _, opt := range opts {
		r = append(r, opt)
	}

	return r
}
