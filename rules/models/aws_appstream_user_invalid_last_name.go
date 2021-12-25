// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamUserInvalidLastNameRule checks the pattern is valid
type AwsAppstreamUserInvalidLastNameRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsAppstreamUserInvalidLastNameRule returns new rule with default attributes
func NewAwsAppstreamUserInvalidLastNameRule() *AwsAppstreamUserInvalidLastNameRule {
	return &AwsAppstreamUserInvalidLastNameRule{
		resourceType:  "aws_appstream_user",
		attributeName: "last_name",
		max:           2048,
		pattern:       regexp.MustCompile(`^[A-Za-z0-9_\-\s]+$`),
	}
}

// Name returns the rule name
func (r *AwsAppstreamUserInvalidLastNameRule) Name() string {
	return "aws_appstream_user_invalid_last_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamUserInvalidLastNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamUserInvalidLastNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamUserInvalidLastNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamUserInvalidLastNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"last_name must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9_\-\s]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
