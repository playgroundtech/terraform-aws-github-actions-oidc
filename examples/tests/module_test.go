// Use test-structure to skip certain steps
// SKIP_destroy=true go test -v -timeout 90m -run TestSimple

package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	// Create a random unique ID for the role name
	randomId := random.UniqueId()
	roleName := fmt.Sprintf("terratest-%v", randomId)
	workingDir := "../simple"

	// Randomize the region
	region := aws.GetRandomRegion(t, []string{"eu-north-1", "us-east-1"}, nil)

	// Terraform destroy
	defer test_structure.RunTestStage(t, "destroy", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.Destroy(t, terraformOptions)
		// clean up saved options
		test_structure.CleanupTestDataFolder(t, workingDir)
	})

	// Terraform init and apply
	test_structure.RunTestStage(t, "init", func() {
		terraformOptions := &terraform.Options{
			TerraformDir: workingDir,
			EnvVars: map[string]string{
				"AWS_DEFAULT_REGION": region,
			},

			Vars: map[string]interface{}{
				"role_name": roleName,
			},
		}
		test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)
		terraform.InitAndApply(t, terraformOptions)
	})

	// Tun tests
	test_structure.RunTestStage(t, "tests", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.ApplyAndIdempotent(t, terraformOptions)

		outputRoleName := terraform.Output(t, terraformOptions, "role_name")
		assert.Equal(t, roleName, outputRoleName, "role_name should be equal")
	})
}
