package cdkcomprehends3olap


type LambdaArnCaptorResourceProps struct {
	// The partial fixed name of the gemeral Lambda function created from the serverless application.
	PartialLambdaName *string `field:"required" json:"partialLambdaName" yaml:"partialLambdaName"`
	// the name of the corresponding IAM role.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

