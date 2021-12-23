// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppconfigEnvironmentInvalidDescriptionRule checks the pattern is valid
type AwsAppconfigEnvironmentInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsAppconfigEnvironmentInvalidDescriptionRule returns new rule with default attributes
func NewAwsAppconfigEnvironmentInvalidDescriptionRule() *AwsAppconfigEnvironmentInvalidDescriptionRule {
	return &AwsAppconfigEnvironmentInvalidDescriptionRule{
		resourceType:  "aws_appconfig_environment",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsAppconfigEnvironmentInvalidDescriptionRule) Name() string {
	return "aws_appconfig_environment_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppconfigEnvironmentInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppconfigEnvironmentInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppconfigEnvironmentInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppconfigEnvironmentInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 1024 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
