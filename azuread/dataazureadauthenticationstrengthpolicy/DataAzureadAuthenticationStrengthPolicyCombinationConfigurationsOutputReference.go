// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazureadauthenticationstrengthpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azuread-go/azuread/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-azuread-go/azuread/v17/dataazureadauthenticationstrengthpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference interface {
	cdktn.ComplexObject
	AllowedAaguids() *[]*string
	AllowedIssuerSkis() *[]*string
	AllowedPolicyOids() *[]*string
	AppliesToCombinations() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAzureadAuthenticationStrengthPolicyCombinationConfigurations
	SetInternalValue(val *DataAzureadAuthenticationStrengthPolicyCombinationConfigurations)
	ObjectId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference
type jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) AllowedAaguids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAaguids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) AllowedIssuerSkis() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIssuerSkis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) AllowedPolicyOids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPolicyOids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) AppliesToCombinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appliesToCombinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) InternalValue() *DataAzureadAuthenticationStrengthPolicyCombinationConfigurations {
	var returns *DataAzureadAuthenticationStrengthPolicyCombinationConfigurations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azuread.dataAzureadAuthenticationStrengthPolicy.DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference_Override(d DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azuread.dataAzureadAuthenticationStrengthPolicy.DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference)SetInternalValue(val *DataAzureadAuthenticationStrengthPolicyCombinationConfigurations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadAuthenticationStrengthPolicyCombinationConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

