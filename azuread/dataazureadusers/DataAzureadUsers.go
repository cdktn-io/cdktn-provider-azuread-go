// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazureadusers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azuread-go/azuread/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azuread-go/azuread/v15/dataazureadusers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/3.8.0/docs/data-sources/users azuread_users}.
type DataAzureadUsers interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmployeeIds() *[]*string
	SetEmployeeIds(val *[]*string)
	EmployeeIdsInput() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreMissing() interface{}
	SetIgnoreMissing(val interface{})
	IgnoreMissingInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MailNicknames() *[]*string
	SetMailNicknames(val *[]*string)
	MailNicknamesInput() *[]*string
	Mails() *[]*string
	SetMails(val *[]*string)
	MailsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	ObjectIds() *[]*string
	SetObjectIds(val *[]*string)
	ObjectIdsInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReturnAll() interface{}
	SetReturnAll(val interface{})
	ReturnAllInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzureadUsersTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserPrincipalNames() *[]*string
	SetUserPrincipalNames(val *[]*string)
	UserPrincipalNamesInput() *[]*string
	Users() DataAzureadUsersUsersList
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DataAzureadUsersTimeouts)
	ResetEmployeeIds()
	ResetId()
	ResetIgnoreMissing()
	ResetMailNicknames()
	ResetMails()
	ResetObjectIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReturnAll()
	ResetTimeouts()
	ResetUserPrincipalNames()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAzureadUsers
type jsiiProxy_DataAzureadUsers struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAzureadUsers) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) EmployeeIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"employeeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) EmployeeIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"employeeIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) IgnoreMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) IgnoreMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) MailNicknames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mailNicknames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) MailNicknamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mailNicknamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Mails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) MailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) ObjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) ObjectIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) ReturnAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) ReturnAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Timeouts() DataAzureadUsersTimeoutsOutputReference {
	var returns DataAzureadUsersTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) UserPrincipalNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userPrincipalNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) UserPrincipalNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userPrincipalNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUsers) Users() DataAzureadUsersUsersList {
	var returns DataAzureadUsersUsersList
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.8.0/docs/data-sources/users azuread_users} Data Source.
func NewDataAzureadUsers(scope constructs.Construct, id *string, config *DataAzureadUsersConfig) DataAzureadUsers {
	_init_.Initialize()

	if err := validateNewDataAzureadUsersParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzureadUsers{}

	_jsii_.Create(
		"@cdktn/provider-azuread.dataAzureadUsers.DataAzureadUsers",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.8.0/docs/data-sources/users azuread_users} Data Source.
func NewDataAzureadUsers_Override(d DataAzureadUsers, scope constructs.Construct, id *string, config *DataAzureadUsersConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azuread.dataAzureadUsers.DataAzureadUsers",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetEmployeeIds(val *[]*string) {
	if err := j.validateSetEmployeeIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"employeeIds",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetIgnoreMissing(val interface{}) {
	if err := j.validateSetIgnoreMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMissing",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetMailNicknames(val *[]*string) {
	if err := j.validateSetMailNicknamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailNicknames",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetMails(val *[]*string) {
	if err := j.validateSetMailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mails",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetObjectIds(val *[]*string) {
	if err := j.validateSetObjectIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectIds",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetReturnAll(val interface{}) {
	if err := j.validateSetReturnAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnAll",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUsers)SetUserPrincipalNames(val *[]*string) {
	if err := j.validateSetUserPrincipalNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPrincipalNames",
		val,
	)
}

// Generates CDKTN code for importing a DataAzureadUsers resource upon running "cdktn plan <stack-name>".
func DataAzureadUsers_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzureadUsers_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azuread.dataAzureadUsers.DataAzureadUsers",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataAzureadUsers_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadUsers_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azuread.dataAzureadUsers.DataAzureadUsers",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadUsers_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadUsers_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azuread.dataAzureadUsers.DataAzureadUsers",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadUsers_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadUsers_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azuread.dataAzureadUsers.DataAzureadUsers",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzureadUsers_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azuread.dataAzureadUsers.DataAzureadUsers",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzureadUsers) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzureadUsers) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzureadUsers) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAzureadUsers) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzureadUsers) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzureadUsers) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzureadUsers) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzureadUsers) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzureadUsers) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzureadUsers) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzureadUsers) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAzureadUsers) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzureadUsers) PutTimeouts(value *DataAzureadUsersTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetEmployeeIds() {
	_jsii_.InvokeVoid(
		d,
		"resetEmployeeIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetIgnoreMissing() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreMissing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetMailNicknames() {
	_jsii_.InvokeVoid(
		d,
		"resetMailNicknames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetMails() {
	_jsii_.InvokeVoid(
		d,
		"resetMails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetObjectIds() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetReturnAll() {
	_jsii_.InvokeVoid(
		d,
		"resetReturnAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) ResetUserPrincipalNames() {
	_jsii_.InvokeVoid(
		d,
		"resetUserPrincipalNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUsers) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUsers) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUsers) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUsers) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUsers) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUsers) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

