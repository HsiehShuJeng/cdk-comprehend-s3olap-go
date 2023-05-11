package cdkcomprehends3olap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-comprehend-s3olap-go/cdkcomprehends3olap/v2/internal"
)

// Creates the foundation necessary to deploy the S3 Object Lambda Acceess Control Use Case.
type ComprehendS3olab interface {
	constructs.Construct
	// The ARN of the Lambda function combined with Amazon Comprehend for thie administrator role in the redaction case.
	AdminLambdaArn() *string
	// The ARN of the Lambda function combined with Amazon Comprehend for thie billing role in the redaction case.
	BillingLambdaArn() *string
	// The ARN of the Lambda function combined with Amazon Comprehend for thie customer support role in the redaction case.
	CustomerSupportLambdaArn() *string
	// The tree node.
	Node() constructs.Node
	// The ARN of the Lambda function combined with Amazon Comprehend for the general case.
	PiiAccessConrtolLambdaArn() *string
	// The ARN of the S3 Object Lambda for access control.
	S3objectLambdaAccessControlArn() *string
	// The ARN of the S3 Object Lambda for the admin role in the redaction case.
	S3objectLambdaAdminArn() *string
	// The ARN of the S3 Object Lambda for the billing role in the redaction case.
	S3objectLambdaBillingArn() *string
	// The ARN of the S3 Object Lambda for the customer support role in the redaction case.
	S3objectLambdaCustomerSupportArn() *string
	GenerateS3Prefix(length *float64) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ComprehendS3olab
type jsiiProxy_ComprehendS3olab struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ComprehendS3olab) AdminLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) BillingLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) CustomerSupportLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerSupportLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) PiiAccessConrtolLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiAccessConrtolLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) S3objectLambdaAccessControlArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3objectLambdaAccessControlArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) S3objectLambdaAdminArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3objectLambdaAdminArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) S3objectLambdaBillingArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3objectLambdaBillingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendS3olab) S3objectLambdaCustomerSupportArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3objectLambdaCustomerSupportArn",
		&returns,
	)
	return returns
}


func NewComprehendS3olab(scope constructs.Construct, id *string, props *ComprehendS3olabProps) ComprehendS3olab {
	_init_.Initialize()

	if err := validateNewComprehendS3olabParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendS3olab{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.ComprehendS3olab",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewComprehendS3olab_Override(c ComprehendS3olab, scope constructs.Construct, id *string, props *ComprehendS3olabProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.ComprehendS3olab",
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
func ComprehendS3olab_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComprehendS3olab_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.ComprehendS3olab",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendS3olab) GenerateS3Prefix(length *float64) *string {
	if err := c.validateGenerateS3PrefixParameters(length); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"generateS3Prefix",
		[]interface{}{length},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendS3olab) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

