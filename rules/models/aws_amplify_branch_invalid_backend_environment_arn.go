// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyBranchInvalidBackendEnvironmentArnRule checks the pattern is valid
type AwsAmplifyBranchInvalidBackendEnvironmentArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAmplifyBranchInvalidBackendEnvironmentArnRule returns new rule with default attributes
func NewAwsAmplifyBranchInvalidBackendEnvironmentArnRule() *AwsAmplifyBranchInvalidBackendEnvironmentArnRule {
	return &AwsAmplifyBranchInvalidBackendEnvironmentArnRule{
		resourceType:  "aws_amplify_branch",
		attributeName: "backend_environment_arn",
		max:           1000,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAmplifyBranchInvalidBackendEnvironmentArnRule) Name() string {
	return "aws_amplify_branch_invalid_backend_environment_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyBranchInvalidBackendEnvironmentArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyBranchInvalidBackendEnvironmentArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyBranchInvalidBackendEnvironmentArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyBranchInvalidBackendEnvironmentArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"backend_environment_arn must be 1000 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"backend_environment_arn must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}