// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsOrganizationsAccountInvalidRoleNameRule checks the pattern is valid
type AwsOrganizationsAccountInvalidRoleNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsAccountInvalidRoleNameRule returns new rule with default attributes
func NewAwsOrganizationsAccountInvalidRoleNameRule() *AwsOrganizationsAccountInvalidRoleNameRule {
	return &AwsOrganizationsAccountInvalidRoleNameRule{
		resourceType:  "aws_organizations_account",
		attributeName: "role_name",
		pattern:       regexp.MustCompile(`^[\w+=,.@-]{1,64}$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Name() string {
	return "aws_organizations_account_invalid_role_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`role_name does not match valid pattern ^[\w+=,.@-]{1,64}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}