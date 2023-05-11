package cdkcomprehends3olap


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

