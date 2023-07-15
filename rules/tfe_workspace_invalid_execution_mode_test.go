package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestTFEWorkspaceInvalidExecutionMode(t *testing.T) {
	cases := []struct {
		Name string
		Content string
		Expected helper.Issues
	}{
		{
			Name: "valid execution mode",
			Content: `resource "tfe_workspace" "test" {
  name         = "test"
  organization = "test"

  execution_mode        = "local"
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "invalid execution mode 'test'",
			Content: `resource "tfe_workspace" "test" {
  name         = "test"
  organization = "test"

  execution_mode        = "test"
}
`,
			Expected: helper.Issues{
				{
					Rule: NewTfeWorkspaceInvalidExecutionModeRule(),
					Message: "\"test\" is not a valid execution mode",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start: hcl.Pos{Line: 5, Column: 27},
						End:   hcl.Pos{Line: 5, Column: 33},
					},
				},
			},
		},
	}

	rule := NewTfeWorkspaceInvalidExecutionModeRule()

	for _, tt := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tt.Content})

		if err := rule.Check(runner); err != nil {
			t.Errorf("unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tt.Expected, runner.Issues)
	}
}

