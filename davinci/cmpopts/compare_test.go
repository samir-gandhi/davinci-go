package cmpopts_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// 	"github.com/samir-gandhi/davinci-client-go/davinci"
// 	"github.com/samir-gandhi/davinci-client-go/davinci/cmpopts"
// 	"github.com/samir-gandhi/davinci-client-go/davinci/test"
// )

// func TestEqual_Basic(t *testing.T) {

// 	testData := map[string]struct {
// 		parentObject  interface{}
// 		compareObject interface{}
// 		opts          []cmp.Option
// 		result        bool
// 	}{
// 		"equal-no-options-same-object": {
// 			parentObject:  test.Data_FullBasic(),
// 			compareObject: test.Data_FullBasic(),
// 			opts:          nil,
// 			result:        true,
// 		},
// 		"equal-no-options-new-object": {
// 			parentObject:  test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow { r := test.Data_FullBasic(); return r }(),
// 			opts:          nil,
// 			result:        true,
// 		},
// 		"not-equal-modified-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.CompanyID = "my-new-company"
// 				return r
// 			}(),
// 			opts:   nil,
// 			result: false,
// 		},
// 		"not-equal-deleted-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.DeployedDate = nil
// 				return r
// 			}(),
// 			opts:   nil,
// 			result: false,
// 		},
// 		"not-equal-added-attribute": {
// 			parentObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.Settings = map[string]interface{}{
// 					"new-key": "new-value",
// 				}
// 				return r
// 			}(),
// 			compareObject: test.Data_FullBasic(),
// 			opts:          nil,
// 			result:        false,
// 		},
// 	}

// 	for k, v := range testData {
// 		t.Run(fmt.Sprintf("davinci object basic equality - %s", k), func(t *testing.T) {
// 			actualResult := cmp.Equal(v.parentObject, v.compareObject, v.opts...)
// 			if actualResult != v.result {
// 				t.Fatalf("Equality failure (-want, +got):\n%s", cmp.Diff(v.parentObject, v.compareObject, v.opts...))
// 			}
// 		})
// 	}
// }

// func TestEqual_UnmappedPropertiesOption(t *testing.T) {

// 	testData := map[string]struct {
// 		parentObject  interface{}
// 		compareObject interface{}
// 		opts          []cmp.Option
// 		result        bool
// 	}{
// 		"equal-ignore-modified-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.AdditionalProperties = map[string]interface{}{
// 					"modified-key": "modified-key-value",
// 				}
// 				return r
// 			}(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: true,
// 		},
// 		"not-equal-noignore-modified-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.AdditionalProperties = map[string]interface{}{
// 					"modified-key": "modified-key-value",
// 				}
// 				return r
// 			}(),
// 			opts:   nil,
// 			result: false,
// 		},
// 		"equal-ignore-removed-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.AdditionalProperties = nil
// 				return r
// 			}(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: true,
// 		},
// 		"not-equal-noignore-removed-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.AdditionalProperties = nil
// 				return r
// 			}(),
// 			opts:   nil,
// 			result: false,
// 		},
// 		"equal-ignore-added-attribute": {
// 			parentObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.AdditionalProperties = nil
// 				return r
// 			}(),
// 			compareObject: test.Data_FullBasic(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: true,
// 		},
// 		"not-equal-noignore-added-attribute": {
// 			parentObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.AdditionalProperties = nil
// 				return r
// 			}(),
// 			compareObject: test.Data_FullBasic(),
// 			opts:          nil,
// 			result:        false,
// 		},
// 		// unprocessed attribute
// 		//TODO
// 		// nested
// 		"equal-nested-ignore-modified-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.GraphData.AdditionalProperties = map[string]interface{}{
// 					"modified-key": "modified-key-value",
// 				}
// 				return r
// 			}(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: true,
// 		},
// 		"not-equal-nested-noignore-modified-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.GraphData.AdditionalProperties = map[string]interface{}{
// 					"modified-key": "modified-key-value",
// 				}
// 				return r
// 			}(),
// 			opts:   nil,
// 			result: false,
// 		},
// 		"equal-nested-ignore-removed-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.GraphData.AdditionalProperties = nil
// 				return r
// 			}(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: true,
// 		},
// 		"not-equal-nested-noignore-removed-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.GraphData.AdditionalProperties = nil
// 				return r
// 			}(),
// 			opts:   nil,
// 			result: false,
// 		},
// 		"equal-nested-ignore-added-attribute": {
// 			parentObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.GraphData.AdditionalProperties = nil
// 				return r
// 			}(),
// 			compareObject: test.Data_FullBasic(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: true,
// 		},
// 		"not-equal-nested-noignore-added-attribute": {
// 			parentObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.GraphData.AdditionalProperties = nil
// 				return r
// 			}(),
// 			compareObject: test.Data_FullBasic(),
// 			opts:          nil,
// 			result:        false,
// 		},
// 		// test ineffectual option
// 		"not-equal-ineffectual-option-modified-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.FlowStatus = "disabled"
// 				return r
// 			}(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: false,
// 		},
// 		"not-equal-ineffectual-option-deleted-attribute": {
// 			parentObject: test.Data_FullBasic(),
// 			compareObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.UpdatedDate = nil
// 				return r
// 			}(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: false,
// 		},
// 		"not-equal-ineffectual-option-added-attribute": {
// 			parentObject: func() davinci.Flow {
// 				r := test.Data_FullBasic()
// 				r.UpdatedDate = nil
// 				return r
// 			}(),
// 			compareObject: test.Data_FullBasic(),
// 			opts: cmpopts.ExportCmpFilters(cmpopts.ExportCmpOpts{
// 				IgnoreUnmappedProperties:  true,
// 				IgnoreEnvironmentMetadata: false,
// 				IgnoreConfig:              false,
// 				IgnoreDesignerCues:        false,
// 			}),
// 			result: false,
// 		},
// 	}

// 	for k, v := range testData {
// 		t.Run(fmt.Sprintf("davinci object equality - IgnoreExportUnmappedProperties option - %s", k), func(t *testing.T) {
// 			actualResult := cmp.Equal(v.parentObject, v.compareObject, v.opts...)
// 			if actualResult != v.result {
// 				t.Fatalf("Equality failure (-want, +got):\n%s", cmp.Diff(v.parentObject, v.compareObject, v.opts...))
// 			}
// 		})
// 	}
// }

// // func TestEqual_ConfigOption(t *testing.T) {

// // 	testData := map[string]struct {
// // 		parentObject  interface{}
// // 		compareObject interface{}
// // 		opts          []cmp.Option
// // 		result        bool
// // 	}{
// // 		"equal-ignore-modified-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField01 = false
// // 				return r
// // 			}(),
// // 			opts:   IgnoreExportUnmappedProperties(),
// // 			result: true,
// // 		},
// // 		"not-equal-noignore-modified-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField01 = false
// // 				return r
// // 			}(),
// // 			opts:   nil,
// // 			result: false,
// // 		},
// // 		"equal-ignore-removed-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			opts:   IgnoreExportUnmappedProperties(),
// // 			result: true,
// // 		},
// // 		"not-equal-noignore-removed-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			opts:   nil,
// // 			result: false,
// // 		},
// // 		"equal-ignore-added-attribute": {
// // 			parentObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			compareObject: test.Data_FullBasic(),
// // 			opts:          IgnoreExportUnmappedProperties(),
// // 			result:        true,
// // 		},
// // 		"not-equal-noignore-added-attribute": {
// // 			parentObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			compareObject: test.Data_FullBasic(),
// // 			opts:          nil,
// // 			result:        false,
// // 		},
// // 		// unprocessed attribute
// // 		//TODO
// // 		// nested
// // 		"equal-nested-ignore-modified-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField07.ConfigField05 = func() *string { s := "modified"; return &s }()
// // 				return r
// // 			}(),
// // 			opts:   IgnoreExportUnmappedProperties(),
// // 			result: true,
// // 		},
// // 		"not-equal-nested-noignore-modified-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField07.ConfigField05 = func() *string { s := "modified"; return &s }()
// // 				return r
// // 			}(),
// // 			opts:   nil,
// // 			result: false,
// // 		},
// // 		"equal-nested-ignore-removed-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField07.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			opts:   IgnoreExportUnmappedProperties(),
// // 			result: true,
// // 		},
// // 		"not-equal-nested-noignore-removed-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField07.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			opts:   nil,
// // 			result: false,
// // 		},
// // 		"equal-nested-ignore-added-attribute": {
// // 			parentObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField07.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			compareObject: test.Data_FullBasic(),
// // 			opts:          IgnoreExportUnmappedProperties(),
// // 			result:        true,
// // 		},
// // 		"not-equal-nested-noignore-added-attribute": {
// // 			parentObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField07.ConfigField05 = nil
// // 				return r
// // 			}(),
// // 			compareObject: test.Data_FullBasic(),
// // 			opts:          nil,
// // 			result:        false,
// // 		},
// // 		// test ineffectual option
// // 		"not-equal-ineffectual-option-modified-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField02 = true
// // 				return r
// // 			}(),
// // 			opts:   IgnoreExportUnmappedProperties(),
// // 			result: false,
// // 		},
// // 		"not-equal-ineffectual-option-deleted-attribute": {
// // 			parentObject: test.Data_FullBasic(),
// // 			compareObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField06 = nil
// // 				return r
// // 			}(),
// // 			opts:   IgnoreExportUnmappedProperties(),
// // 			result: false,
// // 		},
// // 		"not-equal-ineffectual-option-added-attribute": {
// // 			parentObject: func() davinci.Flow {
// // 				r := test.Data_FullBasic()
// // 				r.ConfigField06 = nil
// // 				return r
// // 			}(),
// // 			compareObject: test.Data_FullBasic(),
// // 			opts:          IgnoreExportUnmappedProperties(),
// // 			result:        false,
// // 		},
// // 	}

// // 	for k, v := range testData {
// // 		t.Run(fmt.Sprintf("davinci object equality - IgnoreExportEnvironmentConfig option - %s", k), func(t *testing.T) {
// // 			actualResult := cmp.Equal(v.parentObject, v.compareObject, v.opts...)
// // 			if actualResult != v.result {
// // 				t.Fatalf("Equality failure: wanted: %#v, got: %#v", v.result, actualResult)
// // 			}
// // 		})
// // 	}
// // }
