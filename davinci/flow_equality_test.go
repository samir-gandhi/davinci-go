package davinci

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIgnoreExportFields(t *testing.T) {

	t.Run("unmarshal full basic", func(t *testing.T) {

		opts := make([]cmp.Option, 0)

		for _, v := range flowObjects() {
			if em, ok := v.(DaVinciExportModel); ok {
				opts = append(opts, IgnoreExportFields(em, ExportCmpOpts{
					IgnoreConfig:              true,
					IgnoreDesignerCues:        true,
					IgnoreEnvironmentMetadata: true,
					IgnoreUnmappedProperties:  true,
					IgnoreVersionMetadata:     true,
					IgnoreFlowMetadata:        true,
				}))
			}
		}
	})
}
