// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamFleetInvalidDisplayNameRule checks the pattern is valid
type AwsAppstreamFleetInvalidDisplayNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsAppstreamFleetInvalidDisplayNameRule returns new rule with default attributes
func NewAwsAppstreamFleetInvalidDisplayNameRule() *AwsAppstreamFleetInvalidDisplayNameRule {
	return &AwsAppstreamFleetInvalidDisplayNameRule{
		resourceType:  "aws_appstream_fleet",
		attributeName: "display_name",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Name() string {
	return "aws_appstream_fleet_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"display_name must be 100 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
