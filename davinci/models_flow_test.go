package davinci_test

// import (
// 	"encoding/json"
// 	"io"
// 	"os"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// 	"github.com/google/go-cmp/cmp/cmpopts"
// 	"github.com/samir-gandhi/davinci-client-go/davinci"
// 	"github.com/samir-gandhi/davinci-client-go/davinci/test"
// )

// func TestUnmarshalJSON_Positive(t *testing.T) {

// 	expectedObject := test.Data_FullBasic()

// 	flowFile := "./test/flows/full-basic.json"

// 	jsonFile, err := os.Open(flowFile)
// 	if err != nil {
// 		t.Errorf("Failed to open file: %v", err)
// 	}

// 	jsonBytes, err := io.ReadAll(jsonFile)
// 	if err != nil {
// 		t.Errorf("Failed to read file: %v", err)
// 	}

// 	t.Run("unmarshal full basic", func(t *testing.T) {

// 		var actualObject davinci.Flow

// 		err := json.Unmarshal(jsonBytes, &actualObject)
// 		if err != nil {
// 			t.Errorf("Failed to unmarshal json: %v", err)
// 		}

// 		if !cmp.Equal(actualObject, expectedObject, cmpopts.EquateEmpty()) {
// 			t.Fatalf("Objects unequal: wanted: %#v, got: %#v", expectedObject, actualObject)
// 		}
// 	})
// }
