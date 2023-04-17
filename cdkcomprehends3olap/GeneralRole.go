// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/internal"
)

// The role that you are going to assume (switch role).
//
// Explores how the S3 Object Lambda works.
type GeneralRole interface {
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

// The jsii proxy struct for GeneralRole
type jsiiProxy_GeneralRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_GeneralRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeneralRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeneralRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeneralRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}


func NewGeneralRole(scope constructs.Construct, id *string, props *GeneralRoleProps) GeneralRole {
	_init_.Initialize()

	if err := validateNewGeneralRoleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GeneralRole{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.GeneralRole",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGeneralRole_Override(g GeneralRole, scope constructs.Construct, id *string, props *GeneralRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.GeneralRole",
		[]interface{}{scope, id, props},
		g,
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
func GeneralRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGeneralRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.GeneralRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeneralRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
