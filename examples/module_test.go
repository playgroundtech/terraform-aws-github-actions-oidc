// Use test-structure to skip certain steps
// SKIP_destroy=true go test -v -timeout 90m -run TestSimple

package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"testing"
)

func TestSimple(t *testing.T) {
	// Create a random unique ID for the VPC
	randomId := random.UniqueId()
	workingDir := "../examples/simple"

	defer test_structure.RunTestStage(t, "destroy", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.Destroy(t, terraformOptions)
		// clean up saved options
		test_structure.CleanupTestDataFolder(t, workingDir)
	})

	test_structure.RunTestStage(t, "init", func() {
		terraformOptions := &terraform.Options{
			TerraformDir: workingDir,

			Vars: map[string]interface{}{
				"role_name": fmt.Sprintf("terratest-%v", randomId),
			},
		}
		test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)
		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "tests", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.ApplyAndIdempotent(t, terraformOptions)
	})
}
