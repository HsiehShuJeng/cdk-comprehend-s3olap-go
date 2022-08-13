// A constrcut for PII and redaction scenarios with Amazon Comprehend and S3 Object Lambda
package cdkcomprehends3olap


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

