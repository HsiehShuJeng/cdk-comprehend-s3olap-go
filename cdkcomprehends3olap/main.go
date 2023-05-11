// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.AccessConrtolLambda",
		reflect.TypeOf((*AccessConrtolLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AccessConrtolLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.AccessConrtolLambdaProps",
		reflect.TypeOf((*AccessConrtolLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.AdminRole",
		reflect.TypeOf((*AdminRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleId", GoGetter: "RoleId"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdminRole{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.AdminRoleProps",
		reflect.TypeOf((*AdminRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.BillingRole",
		reflect.TypeOf((*BillingRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleId", GoGetter: "RoleId"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BillingRole{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.BillingRoleProps",
		reflect.TypeOf((*BillingRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.ComprehendS3olab",
		reflect.TypeOf((*ComprehendS3olab)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adminLambdaArn", GoGetter: "AdminLambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "billingLambdaArn", GoGetter: "BillingLambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "customerSupportLambdaArn", GoGetter: "CustomerSupportLambdaArn"},
			_jsii_.MemberMethod{JsiiMethod: "generateS3Prefix", GoMethod: "GenerateS3Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "piiAccessConrtolLambdaArn", GoGetter: "PiiAccessConrtolLambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "s3objectLambdaAccessControlArn", GoGetter: "S3objectLambdaAccessControlArn"},
			_jsii_.MemberProperty{JsiiProperty: "s3objectLambdaAdminArn", GoGetter: "S3objectLambdaAdminArn"},
			_jsii_.MemberProperty{JsiiProperty: "s3objectLambdaBillingArn", GoGetter: "S3objectLambdaBillingArn"},
			_jsii_.MemberProperty{JsiiProperty: "s3objectLambdaCustomerSupportArn", GoGetter: "S3objectLambdaCustomerSupportArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ComprehendS3olab{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.ComprehendS3olabProps",
		reflect.TypeOf((*ComprehendS3olabProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.CustSupportRole",
		reflect.TypeOf((*CustSupportRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleId", GoGetter: "RoleId"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustSupportRole{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.CustSupportRoleProps",
		reflect.TypeOf((*CustSupportRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.GeneralRole",
		reflect.TypeOf((*GeneralRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleId", GoGetter: "RoleId"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GeneralRole{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.GeneralRoleProps",
		reflect.TypeOf((*GeneralRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-comprehend-s3olap.IamRoleName",
		reflect.TypeOf((*IamRoleName)(nil)).Elem(),
		map[string]interface{}{
			"GENERAL": IamRoleName_GENERAL,
			"ADMIN": IamRoleName_ADMIN,
			"BILLING": IamRoleName_BILLING,
			"CUST_SUPPORT": IamRoleName_CUST_SUPPORT,
		},
	)
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.LambdaArnCaptorCustomResource",
		reflect.TypeOf((*LambdaArnCaptorCustomResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lambdaArn", GoGetter: "LambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaArnCaptorCustomResource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.LambdaArnCaptorResourceProps",
		reflect.TypeOf((*LambdaArnCaptorResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-comprehend-s3olap.RedactionLambda",
		reflect.TypeOf((*RedactionLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RedactionLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.RedactionLambdaProps",
		reflect.TypeOf((*RedactionLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-comprehend-s3olap.S3AccessPointNames",
		reflect.TypeOf((*S3AccessPointNames)(nil)).Elem(),
	)
}
