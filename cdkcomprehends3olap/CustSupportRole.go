// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/internal"
)

type CustSupportRole interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The ARN of the IAM role.
	RoleArn() *string
	// The unique string identifying the role.
	RoleId() *string
	// The name of the IAM role.
	RoleName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustSupportRole
type jsiiProxy_CustSupportRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustSupportRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustSupportRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustSupportRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustSupportRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}


func NewCustSupportRole(scope constructs.Construct, id *string, props *AdminRoleProps) CustSupportRole {
	_init_.Initialize()

	if err := validateNewCustSupportRoleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustSupportRole{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.CustSupportRole",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCustSupportRole_Override(c CustSupportRole, scope constructs.Construct, id *string, props *AdminRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.CustSupportRole",
		[]interface{}{scope, id, props},
		c,
	)
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
func CustSupportRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustSupportRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.CustSupportRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustSupportRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

