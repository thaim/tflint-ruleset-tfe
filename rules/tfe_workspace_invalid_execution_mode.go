package rules

import (
	"fmt"

	"golang.org/x/exp/slices"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var validExecutionModes = []string{"remote", "local", "agent"}

type tfeWorkspaceInvalidExecutionMode struct {
	tflint.DefaultRule

	resourceType string
	attributeName string
}

func NewTfeWorkspaceInvalidExecutionMode() *tfeWorkspaceInvalidExecutionMode {
	return &tfeWorkspaceInvalidExecutionMode{}
}

func (r *tfeWorkspaceInvalidExecutionMode) Name() string {
	return "tfe_workspace_invalid_execution_mode"
}

func (r *tfeWorkspaceInvalidExecutionMode) Enabled() bool {
	return true
}

func (r *tfeWorkspaceInvalidExecutionMode) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *tfeWorkspaceInvalidExecutionMode) Link() string {
	return ""
}

func (r *tfeWorkspaceInvalidExecutionMode) Check(runner tflint.Runner) error {
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
