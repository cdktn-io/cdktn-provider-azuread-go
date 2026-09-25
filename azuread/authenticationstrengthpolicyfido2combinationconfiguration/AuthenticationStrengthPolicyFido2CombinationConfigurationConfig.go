// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package authenticationstrengthpolicyfido2combinationconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AuthenticationStrengthPolicyFido2CombinationConfigurationConfig struct {
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
	// A set of AAGUIDs allowed to be used as part of the `fido2` combination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_fido2_combination_configuration#allowed_aaguids AuthenticationStrengthPolicyFido2CombinationConfiguration#allowed_aaguids}
	AllowedAaguids *[]*string `field:"required" json:"allowedAaguids" yaml:"allowedAaguids"`
	// The object ID of the authentication strength policy to which this combination configuration applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_fido2_combination_configuration#authentication_strength_policy_id AuthenticationStrengthPolicyFido2CombinationConfiguration#authentication_strength_policy_id}
	AuthenticationStrengthPolicyId *string `field:"required" json:"authenticationStrengthPolicyId" yaml:"authenticationStrengthPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_fido2_combination_configuration#id AuthenticationStrengthPolicyFido2CombinationConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_fido2_combination_configuration#timeouts AuthenticationStrengthPolicyFido2CombinationConfiguration#timeouts}
	Timeouts *AuthenticationStrengthPolicyFido2CombinationConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

