package cdkcomprehends3olap


type GeneralRoleProps struct {
	// The name of the IAM role.
	// Default: 'GeneralRole'.
	//
	IamRoleName *string `field:"optional" json:"iamRoleName" yaml:"iamRoleName"`
	// The name of the object Lambda access point, which will be the same as the S3 acceess point for the S3 bucket in the demostration.
	// Default: 'accessctl-s3olap-survey-results-unknown-pii'.
	//
	ObjectLambdaAccessPointName *string `field:"optional" json:"objectLambdaAccessPointName" yaml:"objectLambdaAccessPointName"`
	// The name of the IAM policy for the IAM role.
	// Default: 'general-role-s3olap-policy'.
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

