package rules

import (
	"fmt"

	"golang.org/x/exp/slices"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var validExecutionModes = []string{"remote", "local", "agent"}

type tfeWorkspaceInvalidExecutionModeRule struct {
	tflint.DefaultRule

	resourceType string
	attributeName string
}

func NewTfeWorkspaceInvalidExecutionModeRule() *tfeWorkspaceInvalidExecutionModeRule {
	return &tfeWorkspaceInvalidExecutionModeRule{
		resourceType: "tfe_workspace",
		attributeName: "execution_mode",
	}
}

func (r *tfeWorkspaceInvalidExecutionModeRule) Name() string {
	return "tfe_workspace_invalid_execution_mode"
}

func (r *tfeWorkspaceInvalidExecutionModeRule) Enabled() bool {
	return true
}

func (r *tfeWorkspaceInvalidExecutionModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *tfeWorkspaceInvalidExecutionModeRule) Link() string {
	return ""
}

func (r *tfeWorkspaceInvalidExecutionModeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(mode string) error {
			if !slices.Contains(validExecutionModes, mode) {
				runner.EmitIssue(
					r,
					fmt.Sprintf("\"%s\" is not a valid execution mode", mode),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
