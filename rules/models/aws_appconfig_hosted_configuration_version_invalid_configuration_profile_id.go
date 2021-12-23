// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule checks the pattern is valid
type AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule returns new rule with default attributes
func NewAwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule() *AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule {
	return &AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule{
		resourceType:  "aws_appconfig_hosted_configuration_version",
		attributeName: "configuration_profile_id",
		pattern:       regexp.MustCompile(`^[a-z0-9]{4,7}$`),
	}
}

// Name returns the rule name
func (r *AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule) Name() string {
	return "aws_appconfig_hosted_configuration_version_invalid_configuration_profile_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppconfigHostedConfigurationVersionInvalidConfigurationProfileIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]{4,7}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
