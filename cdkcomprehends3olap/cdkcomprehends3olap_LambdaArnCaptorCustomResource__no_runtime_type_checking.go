//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap

// Building without runtime type checking enabled, so all the below just return nil

func validateLambdaArnCaptorCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaArnCaptorCustomResourceParameters(scope constructs.Construct, id *string, props *LambdaArnCaptorResourceProps) error {
	return nil
}
