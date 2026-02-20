// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyEligibleAssignmentRules struct {
	// Must the assignment have an expiry date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.8.0/docs/resources/group_role_management_policy#expiration_required GroupRoleManagementPolicy#expiration_required}
	ExpirationRequired interface{} `field:"optional" json:"expirationRequired" yaml:"expirationRequired"`
	// The duration after which assignments expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.8.0/docs/resources/group_role_management_policy#expire_after GroupRoleManagementPolicy#expire_after}
	ExpireAfter *string `field:"optional" json:"expireAfter" yaml:"expireAfter"`
}

