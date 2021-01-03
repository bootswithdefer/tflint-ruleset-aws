// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule checks the pattern is valid
type AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule returns new rule with default attributes
func NewAwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule() *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule {
	return &AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule{
		resourceType:  "aws_redshift_snapshot_schedule_association",
		attributeName: "cluster_identifier",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Name() string {
	return "aws_redshift_snapshot_schedule_association_invalid_cluster_identifier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"cluster_identifier must be 2147483647 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
