package cdkcomprehends3olap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/internal"
)

type AccessConrtolLambda interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The name of the underlying resoure in the serverless application.
	StackName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AccessConrtolLambda
type jsiiProxy_AccessConrtolLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AccessConrtolLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessConrtolLambda) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}


func NewAccessConrtolLambda(scope constructs.Construct, id *string, props *AccessConrtolLambdaProps) AccessConrtolLambda {
	_init_.Initialize()

	if err := validateNewAccessConrtolLambdaParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessConrtolLambda{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.AccessConrtolLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAccessConrtolLambda_Override(a AccessConrtolLambda, scope constructs.Construct, id *string, props *AccessConrtolLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.AccessConrtolLambda",
		[]interface{}{scope, id, props},
		a,
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
func AccessConrtolLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessConrtolLambda_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.AccessConrtolLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessConrtolLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

