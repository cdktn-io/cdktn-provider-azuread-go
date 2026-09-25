// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grouplicense


type GroupLicenseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/group_license#create GroupLicense#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/group_license#delete GroupLicense#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/group_license#read GroupLicense#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

