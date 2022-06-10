// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
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

type AccessConrtolLambdaProps struct {
	// The minimum prediction confidence score above which PII classification and detection would be considered as final answer.
	//
	// Valid range (0.5 to 1.0).
	ConfidenceThreshold *string `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// Number of threads to use for calling Comprehend's ContainsPiiEntities API.
	//
	// This controls the number of simultaneous calls that will be made from this Lambda.
	ContainsPiiEntitiesThreadCount *string `field:"optional" json:"containsPiiEntitiesThreadCount" yaml:"containsPiiEntitiesThreadCount"`
	// Default language of the text to be processed.
	//
	// This code will be used for interacting with Comprehend.
	DefaultLanguageCode *string `field:"optional" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// Default maximum document size (in bytes) that this function can process otherwise will throw exception for too large document size.
	DocumentMaxSize *string `field:"optional" json:"documentMaxSize" yaml:"documentMaxSize"`
	// Maximum document size (in bytes) to be used for making calls to Comprehend's ContainsPiiEntities API.
	DocumentMaxSizeContainsPiiEntities *string `field:"optional" json:"documentMaxSizeContainsPiiEntities" yaml:"documentMaxSizeContainsPiiEntities"`
	// Whether to support partial objects or not.
	//
	// Accessing partial object through http headers such byte-range can corrupt the object and/or affect PII detection accuracy.
	IsPartialObjectSupported *string `field:"optional" json:"isPartialObjectSupported" yaml:"isPartialObjectSupported"`
	// Log level for Lambda function logging, e.g., ERROR, INFO, DEBUG, etc.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Maximum characters to overlap among segments of a document in case chunking is needed because of maximum document size limit.
	MaxCharsOverlap *string `field:"optional" json:"maxCharsOverlap" yaml:"maxCharsOverlap"`
	// List of comma separated PII entity types to be considered for access control.
	//
	// Refer Comprehend's documentation page for list of supported PII entity types.
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// True if publish metrics to Cloudwatch, false otherwise.
	//
	// See README.md for details on CloudWatch metrics.
	PublishCloudWatchMetrics *string `field:"optional" json:"publishCloudWatchMetrics" yaml:"publishCloudWatchMetrics"`
	// The version of the serverless application.
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
	// Number of tokens/words to overlap among segments of a document in case chunking is needed because of maximum document size limit.
	SubsegmentOverlappingTokens *string `field:"optional" json:"subsegmentOverlappingTokens" yaml:"subsegmentOverlappingTokens"`
	// Handling logic for Unsupported files.
	//
	// Valid values are PASS and FAIL.
	UnsupportedFileHandling *string `field:"optional" json:"unsupportedFileHandling" yaml:"unsupportedFileHandling"`
}

type AdminRole interface {
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

// The jsii proxy struct for AdminRole
type jsiiProxy_AdminRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AdminRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}


func NewAdminRole(scope constructs.Construct, id *string, props *AdminRoleProps) AdminRole {
	_init_.Initialize()

	j := jsiiProxy_AdminRole{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.AdminRole",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAdminRole_Override(a AdminRole, scope constructs.Construct, id *string, props *AdminRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.AdminRole",
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
func AdminRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.AdminRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AdminRoleProps struct {
	// The name of the IAM role.
	IamRoleName *string `field:"optional" json:"iamRoleName" yaml:"iamRoleName"`
	// The name of the object Lambda access point, which will be the same as the S3 acceess point for the S3 bucket in the demostration.
	ObjectLambdaAccessPointName *string `field:"optional" json:"objectLambdaAccessPointName" yaml:"objectLambdaAccessPointName"`
	// The name of the IAM policy for the IAM role.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

type BillingRole interface {
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

// The jsii proxy struct for BillingRole
type jsiiProxy_BillingRole struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BillingRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}


func NewBillingRole(scope constructs.Construct, id *string, props *AdminRoleProps) BillingRole {
	_init_.Initialize()

	j := jsiiProxy_BillingRole{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.BillingRole",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBillingRole_Override(b BillingRole, scope constructs.Construct, id *string, props *AdminRoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.BillingRole",
		[]interface{}{scope, id, props},
		b,
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
func BillingRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.BillingRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BillingRoleProps struct {
	// The name of the IAM role.
	IamRoleName *string `field:"optional" json:"iamRoleName" yaml:"iamRoleName"`
	// The name of the object Lambda access point, which will be the same as the S3 acceess point for the S3 bucket in the demostration.
	ObjectLambdaAccessPointName *string `field:"optional" json:"objectLambdaAccessPointName" yaml:"objectLambdaAccessPointName"`
	// The name of the IAM policy for the IAM role.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

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

type ComprehendS3olabProps struct {
	// The parameters needed for the `ComprehendPiiAccessControlS3ObjectLambda` function.
	AccessControlLambdaConfig *AccessConrtolLambdaProps `field:"optional" json:"accessControlLambdaConfig" yaml:"accessControlLambdaConfig"`
	// The parameters of the `ComprehendPiiRedactionS3ObjectLambda` function for the `AdminRole`.
	AdminRedactionLambdaConfig *RedactionLambdaProps `field:"optional" json:"adminRedactionLambdaConfig" yaml:"adminRedactionLambdaConfig"`
	// The manageable properties for the administrator IAM role in the redaction case.
	AdminRoleConfig *AdminRoleProps `field:"optional" json:"adminRoleConfig" yaml:"adminRoleConfig"`
	// The parameters of the `ComprehendPiiRedactionS3ObjectLambda` function for the `BillingRole`.
	BillingRedactionLambdaConfig *RedactionLambdaProps `field:"optional" json:"billingRedactionLambdaConfig" yaml:"billingRedactionLambdaConfig"`
	// The manageable properties for the billing IAM role in the redaction case.
	BillingRoleConfig *BillingRoleProps `field:"optional" json:"billingRoleConfig" yaml:"billingRoleConfig"`
	// The parameters of the `ComprehendPiiRedactionS3ObjectLambda` function for the `CustSupportRole`.
	CusrtSupportRedactionLambdaConfig *RedactionLambdaProps `field:"optional" json:"cusrtSupportRedactionLambdaConfig" yaml:"cusrtSupportRedactionLambdaConfig"`
	// The manageable properties for the customer support IAM role in the redaction case.
	CustSupportRoleConfig *CustSupportRoleProps `field:"optional" json:"custSupportRoleConfig" yaml:"custSupportRoleConfig"`
	// The directory path where `files/access_control/*.txt` and `files/redaction/*.txt` will be put.
	//
	// DO NOT INCLUDE `/` in the end.
	ExampleFileDir *string `field:"optional" json:"exampleFileDir" yaml:"exampleFileDir"`
	// The manageable properties for the IAM role used to access the `survey-results.txt` data.
	GeneralRoleConfig *GeneralRoleProps `field:"optional" json:"generalRoleConfig" yaml:"generalRoleConfig"`
	// For distinguish test and normal deployment.
	GenerateRandomCharacters *bool `field:"optional" json:"generateRandomCharacters" yaml:"generateRandomCharacters"`
	// The names of the S3 access points for the access control case and redaction case.
	S3AccessPointNames *S3AccessPointNames `field:"optional" json:"s3AccessPointNames" yaml:"s3AccessPointNames"`
	// The prefix attached to the name of the S3 bucket where you are going to explore the S3 Object Lambda pertaining to the access control case.
	SurveyBucketPrefix *string `field:"optional" json:"surveyBucketPrefix" yaml:"surveyBucketPrefix"`
	// The prefix attached to the name of the S3 bucket where you are going to explore the S3 Object Lambda pertaining to the redaction case.
	TranscriptsBucketPrefix *string `field:"optional" json:"transcriptsBucketPrefix" yaml:"transcriptsBucketPrefix"`
}

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

type CustSupportRoleProps struct {
	// The name of the IAM role.
	IamRoleName *string `field:"optional" json:"iamRoleName" yaml:"iamRoleName"`
	// The name of the object Lambda access point, which will be the same as the S3 acceess point for the S3 bucket in the demostration.
	ObjectLambdaAccessPointName *string `field:"optional" json:"objectLambdaAccessPointName" yaml:"objectLambdaAccessPointName"`
	// The name of the IAM policy for the IAM role.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

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

type GeneralRoleProps struct {
	// The name of the IAM role.
	IamRoleName *string `field:"optional" json:"iamRoleName" yaml:"iamRoleName"`
	// The name of the object Lambda access point, which will be the same as the S3 acceess point for the S3 bucket in the demostration.
	ObjectLambdaAccessPointName *string `field:"optional" json:"objectLambdaAccessPointName" yaml:"objectLambdaAccessPointName"`
	// The name of the IAM policy for the IAM role.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

type IamRoleName string

const (
	IamRoleName_GENERAL IamRoleName = "GENERAL"
	IamRoleName_ADMIN IamRoleName = "ADMIN"
	IamRoleName_BILLING IamRoleName = "BILLING"
	IamRoleName_CUST_SUPPORT IamRoleName = "CUST_SUPPORT"
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

type LambdaArnCaptorResourceProps struct {
	// The partial fixed name of the gemeral Lambda function created from the serverless application.
	PartialLambdaName *string `field:"required" json:"partialLambdaName" yaml:"partialLambdaName"`
	// the name of the corresponding IAM role.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

type RedactionLambda interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The name of the underlying resoure in the serverless application.
	StackName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RedactionLambda
type jsiiProxy_RedactionLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_RedactionLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedactionLambda) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}


func NewRedactionLambda(scope constructs.Construct, id *string, props *RedactionLambdaProps) RedactionLambda {
	_init_.Initialize()

	j := jsiiProxy_RedactionLambda{}

	_jsii_.Create(
		"cdk-comprehend-s3olap.RedactionLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRedactionLambda_Override(r RedactionLambda, scope constructs.Construct, id *string, props *RedactionLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-comprehend-s3olap.RedactionLambda",
		[]interface{}{scope, id, props},
		r,
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
func RedactionLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-comprehend-s3olap.RedactionLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedactionLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RedactionLambdaProps struct {
	// The minimum prediction confidence score above which PII classification and detection would be considered as final answer.
	//
	// Valid range (0.5 to 1.0).
	ConfidenceThreshold *string `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// Number of threads to use for calling Comprehend's ContainsPiiEntities API.
	//
	// This controls the number of simultaneous calls that will be made from this Lambda.
	ContainsPiiEntitiesThreadCount *string `field:"optional" json:"containsPiiEntitiesThreadCount" yaml:"containsPiiEntitiesThreadCount"`
	// Default language of the text to be processed.
	//
	// This code will be used for interacting with Comprehend.
	DefaultLanguageCode *string `field:"optional" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// Number of threads to use for calling Comprehend's DetectPiiEntities API.
	//
	// This controls the number of simultaneous calls that will be made from this Lambda.
	DetectPiiEntitiesThreadCount *string `field:"optional" json:"detectPiiEntitiesThreadCount" yaml:"detectPiiEntitiesThreadCount"`
	// Default maximum document size (in bytes) that this function can process otherwise will throw exception for too large document size.
	DocumentMaxSize *string `field:"optional" json:"documentMaxSize" yaml:"documentMaxSize"`
	// Maximum document size (in bytes) to be used for making calls to Comprehend's ContainsPiiEntities API.
	DocumentMaxSizeContainsPiiEntities *string `field:"optional" json:"documentMaxSizeContainsPiiEntities" yaml:"documentMaxSizeContainsPiiEntities"`
	// Maximum document size (in bytes) to be used for making calls to Comprehend's DetectPiiEntities API.
	DocumentMaxSizeDetectPiiEntities *string `field:"optional" json:"documentMaxSizeDetectPiiEntities" yaml:"documentMaxSizeDetectPiiEntities"`
	// Whether to support partial objects or not.
	//
	// Accessing partial object through http headers such byte-range can corrupt the object and/or affect PII detection accuracy.
	IsPartialObjectSupported *string `field:"optional" json:"isPartialObjectSupported" yaml:"isPartialObjectSupported"`
	// Log level for Lambda function logging, e.g., ERROR, INFO, DEBUG, etc.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// A character that replaces each character in the redacted PII entity.
	MaskCharacter *string `field:"optional" json:"maskCharacter" yaml:"maskCharacter"`
	// Specifies whether the PII entity is redacted with the mask character or the entity type.
	//
	// Valid values - REPLACE_WITH_PII_ENTITY_TYPE and MASK.
	MaskMode *string `field:"optional" json:"maskMode" yaml:"maskMode"`
	// Maximum characters to overlap among segments of a document in case chunking is needed because of maximum document size limit.
	MaxCharsOverlap *string `field:"optional" json:"maxCharsOverlap" yaml:"maxCharsOverlap"`
	// List of comma separated PII entity types to be considered for redaction.
	//
	// Refer Comprehend's documentation page for list of supported PII entity types.
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// True if publish metrics to Cloudwatch, false otherwise.
	//
	// See README.md for details on CloudWatch metrics.
	PublishCloudWatchMetrics *string `field:"optional" json:"publishCloudWatchMetrics" yaml:"publishCloudWatchMetrics"`
	// The version of the serverless application.
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
	// Number of tokens/words to overlap among segments of a document in case chunking is needed because of maximum document size limit.
	SubsegmentOverlappingTokens *string `field:"optional" json:"subsegmentOverlappingTokens" yaml:"subsegmentOverlappingTokens"`
	// Handling logic for Unsupported files.
	//
	// Valid values are PASS and FAIL.
	UnsupportedFileHandling *string `field:"optional" json:"unsupportedFileHandling" yaml:"unsupportedFileHandling"`
}

type S3AccessPointNames struct {
	// The name of the S3 aceess point for the admin role in the redaction case.
	Admin *string `field:"required" json:"admin" yaml:"admin"`
	// The name of the S3 aceess point for the billing role in the redaction case.
	Billing *string `field:"required" json:"billing" yaml:"billing"`
	// The name of the S3 aceess point for the customer support role in the redaction case.
	CustomerSupport *string `field:"required" json:"customerSupport" yaml:"customerSupport"`
	// The name of the S3 aceess point for the general role in the access control case.
	General *string `field:"required" json:"general" yaml:"general"`
}

