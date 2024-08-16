package davinci_test

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func TestValidFlowJSON(t *testing.T) {

	t.Run("Contains no additional properties", func(t *testing.T) {

		flowFile := "./test/flows/full-basic.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		if ok, errorCode, diff, err := davinci.ValidFlowExport(jsonBytes, davinci.ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		}); !ok {
			log.Printf("Error Code: %v", errorCode)
			log.Printf("Diff: %v", diff)
			log.Printf("Error: %v", err)
			t.Errorf("Expected: %v, Got: %v", false, ok)
		}
	})

	t.Run("Contains additional properties", func(t *testing.T) {

		flowFile := "./test/flows/full-basic-additionalproperties.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		if ok, errorCode, diff, err := davinci.ValidFlowExport(jsonBytes, davinci.ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		}); ok {
			log.Printf("Error Code: %v", errorCode)
			log.Printf("Diff: %v", diff)
			log.Printf("Error: %v", err)
			t.Errorf("Expected: %v, Got: %v", true, ok)
		}
	})
}
