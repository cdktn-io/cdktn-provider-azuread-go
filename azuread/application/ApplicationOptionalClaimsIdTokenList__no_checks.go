// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package application

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationOptionalClaimsIdTokenList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationOptionalClaimsIdTokenList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationOptionalClaimsIdTokenList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsIdTokenList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsIdTokenList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsIdTokenList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsIdTokenList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationOptionalClaimsIdTokenListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

