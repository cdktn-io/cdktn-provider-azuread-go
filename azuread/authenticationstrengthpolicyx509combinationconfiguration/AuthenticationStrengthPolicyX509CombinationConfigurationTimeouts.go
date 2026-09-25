// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package authenticationstrengthpolicyx509combinationconfiguration


type AuthenticationStrengthPolicyX509CombinationConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#create AuthenticationStrengthPolicyX509CombinationConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#delete AuthenticationStrengthPolicyX509CombinationConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#read AuthenticationStrengthPolicyX509CombinationConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/authentication_strength_policy_x509_combination_configuration#update AuthenticationStrengthPolicyX509CombinationConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

