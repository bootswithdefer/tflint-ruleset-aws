package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/project"
)

// AwsElastiCacheReplicationGroupPreviousTypeRule checks whether the resource uses previous generation node type
type AwsElastiCacheReplicationGroupPreviousTypeRule struct {
	resourceType      string
	attributeName     string
	previousNodeTypes map[string]bool
}

// NewAwsElastiCacheReplicationGroupPreviousTypeRule returns new rule with default attributes
func NewAwsElastiCacheReplicationGroupPreviousTypeRule() *AwsElastiCacheReplicationGroupPreviousTypeRule {
	return &AwsElastiCacheReplicationGroupPreviousTypeRule{
		resourceType:  "aws_elasticache_replication_group",
		attributeName: "node_type",
		previousNodeTypes: map[string]bool{
			"cache.m1.small":   true,
			"cache.m1.medium":  true,
			"cache.m1.large":   true,
			"cache.m1.xlarge":  true,
			"cache.m2.xlarge":  true,
			"cache.m2.2xlarge": true,
			"cache.m2.4xlarge": true,
			"cache.m3.medium":  true,
			"cache.m3.large":   true,
			"cache.m3.xlarge":  true,
			"cache.m3.2xlarge": true,
			"cache.r3.large":   true,
			"cache.r3.xlarge":  true,
			"cache.r3.2xlarge": true,
			"cache.r3.4xlarge": true,
			"cache.r3.8xlarge": true,
			"cache.c1.xlarge":  true,
			"cache.t1.micro":   true,
		},
	}
}

// Name returns the rule name
func (r *AwsElastiCacheReplicationGroupPreviousTypeRule) Name() string {
	return "aws_elasticache_replication_group_previous_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheReplicationGroupPreviousTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheReplicationGroupPreviousTypeRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AwsElastiCacheReplicationGroupPreviousTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the resource's `node_type` is included in the list of previous generation node type
func (r *AwsElastiCacheReplicationGroupPreviousTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var nodeType string
		err := runner.EvaluateExpr(attribute.Expr, &nodeType, nil)

		return runner.EnsureNoError(err, func() error {
			if r.previousNodeTypes[nodeType] {
				runner.EmitIssue(
					r,
					fmt.Sprintf("\"%s\" is previous generation node type.", nodeType),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
