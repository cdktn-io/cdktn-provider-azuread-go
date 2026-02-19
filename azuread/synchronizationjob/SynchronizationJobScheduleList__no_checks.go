// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package synchronizationjob

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SynchronizationJobScheduleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SynchronizationJobScheduleList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SynchronizationJobScheduleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SynchronizationJobScheduleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SynchronizationJobScheduleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SynchronizationJobScheduleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSynchronizationJobScheduleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

