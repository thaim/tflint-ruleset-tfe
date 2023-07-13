package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-tfe/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "tfe",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTfeWorkspaceInvalidExecutionMode(),
			},
		},
	})
}
