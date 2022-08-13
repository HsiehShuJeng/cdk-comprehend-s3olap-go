// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap


type IamRoleName string

const (
	IamRoleName_GENERAL IamRoleName = "GENERAL"
	IamRoleName_ADMIN IamRoleName = "ADMIN"
	IamRoleName_BILLING IamRoleName = "BILLING"
	IamRoleName_CUST_SUPPORT IamRoleName = "CUST_SUPPORT"
)

