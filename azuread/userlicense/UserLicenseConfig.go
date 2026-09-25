// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package userlicense

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UserLicenseConfig struct {
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
	// The unique identifier (GUID) for the SKU (license) to assign to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#sku_id UserLicense#sku_id}
	SkuId *string `field:"required" json:"skuId" yaml:"skuId"`
	// The object ID of the user to which the license should be assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#user_object_id UserLicense#user_object_id}
	UserObjectId *string `field:"required" json:"userObjectId" yaml:"userObjectId"`
	// A set of unique identifiers (GUIDs) for the service plans to disable for this license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#disabled_plans UserLicense#disabled_plans}
	DisabledPlans *[]*string `field:"optional" json:"disabledPlans" yaml:"disabledPlans"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#id UserLicense#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.10.0/docs/resources/user_license#timeouts UserLicense#timeouts}
	Timeouts *UserLicenseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

