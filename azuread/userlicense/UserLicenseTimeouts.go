// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package userlicense


type UserLicenseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#create UserLicense#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#delete UserLicense#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#read UserLicense#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

