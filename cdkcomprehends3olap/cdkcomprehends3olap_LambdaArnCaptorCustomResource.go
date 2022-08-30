// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/internal"
)

type LambdaArnCaptorCustomResource interface {
	constructs.Construct
	// The ARN of the general Lambda function created from the serverless application.
	// See: https://github.com/aws/aws-cdk/issues/8760
	//
	LambdaArn() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for LambdaArnCaptorCustomResource
type jsiiProxy_LambdaArnCaptorCustomResource struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaArnCaptorCustomResource) LambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaArnCaptorCustomResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewLambdaArnCaptorCustomResource(scope constructs.Construct, id *string, props *LambdaArnCaptorResourceProps) LambdaArnCaptorCustomResource {
	_init_.Initialize()

	if err := validateNewLambdaArnCaptorCustomResourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaArnCaptorCustomResource{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.LambdaArnCaptorCustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaArnCaptorCustomResource_Override(l LambdaArnCaptorCustomResource, scope constructs.Construct, id *string, props *LambdaArnCaptorResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.LambdaArnCaptorCustomResource",
		[]interface{}{scope, id, props},
		l,
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
func LambdaArnCaptorCustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaArnCaptorCustomResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.LambdaArnCaptorCustomResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaArnCaptorCustomResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

