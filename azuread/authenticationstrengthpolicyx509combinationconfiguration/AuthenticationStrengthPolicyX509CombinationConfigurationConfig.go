// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package authenticationstrengthpolicyx509combinationconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AuthenticationStrengthPolicyX509CombinationConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The x509 certificate authentication method combinations this configuration applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#applies_to_combinations AuthenticationStrengthPolicyX509CombinationConfiguration#applies_to_combinations}
	AppliesToCombinations *[]*string `field:"required" json:"appliesToCombinations" yaml:"appliesToCombinations"`
	// The object ID of the authentication strength policy to which this combination configuration applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#authentication_strength_policy_id AuthenticationStrengthPolicyX509CombinationConfiguration#authentication_strength_policy_id}
	AuthenticationStrengthPolicyId *string `field:"required" json:"authenticationStrengthPolicyId" yaml:"authenticationStrengthPolicyId"`
	// A set of allowed certificate issuer subject key identifier values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#allowed_issuer_skis AuthenticationStrengthPolicyX509CombinationConfiguration#allowed_issuer_skis}
	AllowedIssuerSkis *[]*string `field:"optional" json:"allowedIssuerSkis" yaml:"allowedIssuerSkis"`
	// A set of allowed certificate policy OIDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#allowed_policy_oids AuthenticationStrengthPolicyX509CombinationConfiguration#allowed_policy_oids}
	AllowedPolicyOids *[]*string `field:"optional" json:"allowedPolicyOids" yaml:"allowedPolicyOids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#id AuthenticationStrengthPolicyX509CombinationConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#timeouts AuthenticationStrengthPolicyX509CombinationConfiguration#timeouts}
	Timeouts *AuthenticationStrengthPolicyX509CombinationConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

