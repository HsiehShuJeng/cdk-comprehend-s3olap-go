//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ComprehendS3olab) validateGenerateS3PrefixParameters(length *float64) error {
	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}

func validateComprehendS3olab_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewComprehendS3olabParameters(scope constructs.Construct, id *string, props *ComprehendS3olabProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

