// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlueMlTransformInvalidWorkerTypeRule checks the pattern is valid
type AwsGlueMlTransformInvalidWorkerTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlueMlTransformInvalidWorkerTypeRule returns new rule with default attributes
func NewAwsGlueMlTransformInvalidWorkerTypeRule() *AwsGlueMlTransformInvalidWorkerTypeRule {
	return &AwsGlueMlTransformInvalidWorkerTypeRule{
		resourceType:  "aws_glue_ml_transform",
		attributeName: "worker_type",
		enum: []string{
			"Standard",
			"G.1X",
			"G.2X",
			"G.025X",
		},
	}
}

// Name returns the rule name
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Name() string {
	return "aws_glue_ml_transform_invalid_worker_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as worker_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
