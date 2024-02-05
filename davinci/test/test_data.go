package test

import "github.com/samir-gandhi/davinci-client-go/davinci"

func Data_FullBasic() davinci.Flow {

	currentVersion := 8
	deployedDate := 1706709739837
	publishedVersion := 8
	savedDate := 1706708769645
	updatedDate := 1706709739837
	description := ""
	flowColor := "#E3F0FF"
	isOutputSchemaSaved := false
	flowVariable2value := "10"

	return davinci.Flow{
		AdditionalProperties: map[string]interface{}{
			"custom-attribute-1": "custom-attribute-1-value",
			"custom-attribute-2": "custom-attribute-2-value",
		},
		AuthTokenExpireIds: []interface{}{},
		CompanyID:          "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb",
		Connections:        []interface{}{},
		CreatedDate:        1706708769850,
		CurrentVersion:     &currentVersion,
		CustomerID:         "db5f4450b2bd8a56ce076dec0c358a9a",
		DeployedDate:       &deployedDate,
		FlowID:             "c7062a8857740ee2185694bb855f8f21",
		PublishedVersion:   &publishedVersion,
		SavedDate:          savedDate,
		UpdatedDate:        &updatedDate,
		VersionID:          8,
		ConnectorIds: []string{
			"httpConnector",
			"functionsConnector",
			"errorConnector",
			"flowConnector",
			"variablesConnector",
		},
		Description:          &description,
		EnabledGraphData:     nil,
		FlowColor:            &flowColor,
		FlowStatus:           "enabled",
		FunctionConnectionID: nil,
		GraphData: davinci.GraphData{
			AdditionalProperties: map[string]interface{}{
				"custom-attribute-1": "custom-attribute-1-value",
				"custom-attribute-2": "custom-attribute-2-value",
			},
			Elements: davinci.Elements{
				AdditionalProperties: map[string]interface{}{
					"custom-attribute-1": "custom-attribute-1-value",
					"custom-attribute-2": "custom-attribute-2-value",
				},
				Nodes: []davinci.Node{
					{
						AdditionalProperties: map[string]interface{}{
							"custom-attribute-1": "custom-attribute-1-value",
							"custom-attribute-2": "custom-attribute-2-value",
						},
						Data: davinci.NodeData{
							AdditionalProperties: map[string]interface{}{
								"custom-attribute-1": "custom-attribute-1-value",
								"custom-attribute-2": "custom-attribute-2-value",
							},
							ID:             "1u2m5vzr49",
							NodeType:       "CONNECTION",
							ConnectionID:   "867ed4363b2bc21c860085ad2baa817d",
							ConnectorID:    "httpConnector",
							Name:           "Http",
							Label:          "Http",
							Status:         "configured",
							CapabilityName: "customHtmlMessage",
							Type:           "trigger",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									AdditionalProperties: map[string]interface{}{
										"custom-attribute-1": "custom-attribute-1-value",
										"custom-attribute-2": "custom-attribute-2-value",
										"message": map[string]interface{}{
											"custom-attribute-1": "custom-attribute-1-value",
											"custom-attribute-2": "custom-attribute-2-value",
											"value":              "[\n  {\n    \"children\": [\n      {\n        \"text\": \"Hello, world?\"\n      }\n    ]\n  }\n]",
										},
									},
								}
							}(),
						},
						Position: davinci.Position{
							AdditionalProperties: map[string]interface{}{
								"custom-attribute-1": "custom-attribute-1-value",
								"custom-attribute-2": "custom-attribute-2-value",
							},
							X: 277,
							Y: 336,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:       "8fvg7tfr8j",
							NodeType: "EVAL",
							Label:    "Evaluator",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 426.5,
							Y: 337.25,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:             "nx0o1b2cmw",
							NodeType:       "CONNECTION",
							ConnectionID:   "de650ca45593b82c49064ead10b9fe17",
							ConnectorID:    "functionsConnector",
							Name:           "Functions",
							Label:          "Functions",
							Status:         "configured",
							CapabilityName: "AEqualsB",
							Type:           "trigger",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									AdditionalProperties: map[string]interface{}{
										"leftValueA": map[string]interface{}{
											"value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"1\"\n      }\n    ]\n  }\n]",
										},
										"rightValueB": map[string]interface{}{
											"value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"1\"\n      }\n    ]\n  }\n]",
										},
									},
								}
							}(),
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 576,
							Y: 338.5,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:       "cdcw8k7dnx",
							NodeType: "EVAL",
							Label:    "Evaluator",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									AdditionalProperties: map[string]interface{}{
										"vsp1ewtr9m": map[string]interface{}{
											"value": "allTriggersFalse",
										},
										"xb74p6rkd8": map[string]interface{}{
											"value": "anyTriggersFalse",
										},
									},
								}
							}(),
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 717,
							Y: 326,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:             "ikt13crnhy",
							NodeType:       "CONNECTION",
							ConnectionID:   "867ed4363b2bc21c860085ad2baa817d",
							ConnectorID:    "httpConnector",
							Name:           "Http",
							Label:          "Http",
							Status:         "configured",
							CapabilityName: "createSuccessResponse",
							Type:           "action",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									// AdditionalProperties: map[string]interface{}{},
								}
							}(),
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 1197,
							Y: 266,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:             "vsp1ewtr9m",
							NodeType:       "CONNECTION",
							ConnectionID:   "53ab83a4a4ab919d9f2cb02d9e111ac8",
							ConnectorID:    "errorConnector",
							Name:           "Error Message",
							Label:          "Error Message",
							Status:         "configured",
							CapabilityName: "customErrorMessage",
							Type:           "action",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									AdditionalProperties: map[string]interface{}{
										"errorMessage": map[string]interface{}{
											"value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"Error\"\n      }\n    ]\n  }\n]",
										},
									},
								}
							}(),
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 1197,
							Y: 416,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:             "xb74p6rkd8",
							NodeType:       "CONNECTION",
							ConnectionID:   "33329a264e268ab31fb19637debf1ea3",
							ConnectorID:    "flowConnector",
							Name:           "Flow Conductor",
							Label:          "Flow Conductor",
							Status:         "configured",
							CapabilityName: "startUiSubFlow",
							Type:           "trigger",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									// AdditionalProperties: map[string]interface{}{},
									SubFlowID: func() *davinci.SubFlowID {
										return &davinci.SubFlowID{
											Value: davinci.SubFlowValue{
												Label: "subflow 2",
												Value: "07503fed5c02849dbbd5ee932da654b2",
											},
										}
									}(),
									SubFlowVersionID: func() *davinci.SubFlowVersionID {
										return &davinci.SubFlowVersionID{
											Value: davinci.SubFlowVersionIDValue{
												ValueInt: func() *int {
													v := -1
													return &v
												}(),
											},
										}
									}(),
								}
							}(),
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 867,
							Y: 446,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:             "kq5ybvwvro",
							NodeType:       "CONNECTION",
							ConnectionID:   "33329a264e268ab31fb19637debf1ea3",
							ConnectorID:    "flowConnector",
							Name:           "Flow Conductor",
							Label:          "Flow Conductor",
							Status:         "configured",
							CapabilityName: "startUiSubFlow",
							Type:           "trigger",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									// AdditionalProperties: map[string]interface{}{},
									SubFlowID: func() *davinci.SubFlowID {
										return &davinci.SubFlowID{
											Value: davinci.SubFlowValue{
												Label: "subflow 1",
												Value: "00f66e8926ced6ef5b83619fde4a314a",
											},
										}
									}(),
									SubFlowVersionID: func() *davinci.SubFlowVersionID {
										return &davinci.SubFlowVersionID{
											Value: davinci.SubFlowVersionIDValue{
												ValueInt: func() *int {
													v := -1
													return &v
												}(),
											},
										}
									}(),
								}
							}(),
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 867,
							Y: 236,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:       "j74pmg6577",
							NodeType: "EVAL",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 1017,
							Y: 236,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:         "pensvkew7y",
							NodeType:   "EVAL",
							Properties: nil,
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 1032,
							Y: 431,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:             "3zvjdgdljx",
							NodeType:       "CONNECTION",
							ConnectionID:   "06922a684039827499bdbdd97f49827b",
							ConnectorID:    "variablesConnector",
							Name:           "Variables",
							Label:          "Variables",
							Status:         "configured",
							CapabilityName: "saveFlowValue",
							Type:           "trigger",
							Properties: func() *davinci.Properties {
								return &davinci.Properties{
									// AdditionalProperties: map[string]interface{}{},
									SaveFlowVariables: func() *davinci.SaveFlowVariables {
										return &davinci.SaveFlowVariables{
											Value: func() []davinci.FlowVariable {
												return []davinci.FlowVariable{
													{
														// AdditionalProperties: map[string]interface{}{},
														Name:  "fdgdfgfdg",
														Value: "[\n  {\n    \"children\": [\n      {\n        \"text\": \"test124\"\n      }\n    ]\n  }\n]",
														Key:   0.8936786494474329,
														Label: "fdgdfgfdg (string - flow)",
														Type:  "string",
													},
													{
														// AdditionalProperties: map[string]interface{}{},
														Name:  "test123",
														Value: "[\n  {\n    \"children\": [\n      {\n        \"text\": \"test456\"\n      }\n    ]\n  }\n]",
														Key:   0.379286774724122,
														Label: "test123 (number - flow)",
														Type:  "number",
													},
												}
											}(),
										}
									}(),
								}
							}(),
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 270,
							Y: 180,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.NodeData{
							// AdditionalProperties: map[string]interface{}{},
							ID:       "bbemfztdyk",
							NodeType: "EVAL",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 273.5,
							Y: 258,
						},
						Group:      "nodes",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   false,
						Classes:    "",
					},
				},
				Edges: []davinci.Edge{
					{
						AdditionalProperties: map[string]interface{}{
							"custom-attribute-1": "custom-attribute-1-value",
							"custom-attribute-2": "custom-attribute-2-value",
						},
						Data: davinci.Data{
							AdditionalProperties: map[string]interface{}{
								"custom-attribute-1": "custom-attribute-1-value",
								"custom-attribute-2": "custom-attribute-2-value",
							},
							ID:     "hseww5vtf0",
							Source: "1u2m5vzr49",
							Target: "8fvg7tfr8j",
						},
						Position: davinci.Position{
							AdditionalProperties: map[string]interface{}{
								"custom-attribute-1": "custom-attribute-1-value",
								"custom-attribute-2": "custom-attribute-2-value",
							},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "ljavni2nky",
							Source: "8fvg7tfr8j",
							Target: "nx0o1b2cmw",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "0o2fqy3mf3",
							Source: "nx0o1b2cmw",
							Target: "cdcw8k7dnx",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "493yd0jbi6",
							Source: "cdcw8k7dnx",
							Target: "kq5ybvwvro",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "pn2kixnzms",
							Source: "j74pmg6577",
							Target: "ikt13crnhy",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "0sb4quzlgx",
							Source: "kq5ybvwvro",
							Target: "j74pmg6577",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "v5p4i55lt9",
							Source: "cdcw8k7dnx",
							Target: "xb74p6rkd8",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "k0trrhjqt6",
							Source: "xb74p6rkd8",
							Target: "pensvkew7y",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "2g0chago4l",
							Source: "pensvkew7y",
							Target: "vsp1ewtr9m",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "gs1fx4x303",
							Source: "3zvjdgdljx",
							Target: "bbemfztdyk",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
					{
						// AdditionalProperties: map[string]interface{}{},
						Data: davinci.Data{
							// AdditionalProperties: map[string]interface{}{},
							ID:     "cum544luro",
							Source: "bbemfztdyk",
							Target: "1u2m5vzr49",
						},
						Position: davinci.Position{
							// AdditionalProperties: map[string]interface{}{},
							X: 0,
							Y: 0,
						},
						Group:      "edges",
						Removed:    false,
						Selected:   false,
						Selectable: true,
						Locked:     false,
						Grabbable:  true,
						Pannable:   true,
						Classes:    "",
					},
				},
			},
			Data: davinci.Data{
				AdditionalProperties: map[string]interface{}{
					"custom-attribute-1": "custom-attribute-1-value",
					"custom-attribute-2": "custom-attribute-2-value",
				},
			},
			ZoomingEnabled:     true,
			UserZoomingEnabled: true,
			Zoom:               1,
			MinZoom:            1e-50,
			MaxZoom:            1e+50,
			PanningEnabled:     true,
			UserPanningEnabled: true,
			Pan: davinci.Pan{
				AdditionalProperties: map[string]interface{}{
					"custom-attribute-1": "custom-attribute-1-value",
					"custom-attribute-2": "custom-attribute-2-value",
				},
				X: 0,
				Y: 0,
			},
			BoxSelectionEnabled: true,
			Renderer: davinci.Renderer{
				AdditionalProperties: map[string]interface{}{
					"custom-attribute-1": "custom-attribute-1-value",
					"custom-attribute-2": "custom-attribute-2-value",
				},
				Name: "null",
			},
		},
		InputSchema:          nil,
		InputSchemaCompiled:  nil,
		IsInputSchemaSaved:   nil,
		IsOutputSchemaSaved:  &isOutputSchemaSaved,
		Name:                 "full-basic",
		Orx:                  nil,
		OutputSchema:         nil,
		OutputSchemaCompiled: nil,
		Settings:             nil,
		Timeouts:             "null",
		Trigger:              nil,
		Variables: []davinci.FlowVariable{
			{
				AdditionalProperties: map[string]interface{}{
					"custom-attribute-1": "custom-attribute-1-value",
					"custom-attribute-2": "custom-attribute-2-value",
				},
				Context:     "flow",
				CreatedDate: 1706708735989,
				CustomerID:  "db5f4450b2bd8a56ce076dec0c358a9a",
				Fields: davinci.FlowVariableFields{
					AdditionalProperties: map[string]interface{}{
						"custom-attribute-1": "custom-attribute-1-value",
						"custom-attribute-2": "custom-attribute-2-value",
					},
					Type:        "string",
					DisplayName: "",
					Value:       nil,
					Mutable:     true,
					Min:         0,
					Max:         2000,
				},
				FlowID:      "c7062a8857740ee2185694bb855f8f21",
				Type:        "property",
				UpdatedDate: nil,
				Visibility:  "private",
				Name:        "fdgdfgfdg##SK##flow##SK##c7062a8857740ee2185694bb855f8f21",
				CompanyID:   "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb",
			},
			{
				// AdditionalProperties: map[string]interface{}{},
				Context:     "flow",
				CreatedDate: 1706708761083,
				CustomerID:  "db5f4450b2bd8a56ce076dec0c358a9a",
				Fields: davinci.FlowVariableFields{
					// AdditionalProperties: map[string]interface{}{},
					Type:        "number",
					DisplayName: "test123",
					Value:       &flowVariable2value,
					Mutable:     true,
					Min:         4,
					Max:         20,
				},
				FlowID:      "c7062a8857740ee2185694bb855f8f21",
				Type:        "property",
				UpdatedDate: nil,
				Visibility:  "private",
				Name:        "test123##SK##flow##SK##c7062a8857740ee2185694bb855f8f21",
				CompanyID:   "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb",
			},
		},
	}
}
