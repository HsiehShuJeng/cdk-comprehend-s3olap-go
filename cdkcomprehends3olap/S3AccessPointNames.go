package cdkcomprehends3olap


type S3AccessPointNames struct {
	// The name of the S3 aceess point for the admin role in the redaction case.
	// Default: 'admin-s3-access-point-call-transcripts-known-pii'.
	//
	Admin *string `field:"required" json:"admin" yaml:"admin"`
	// The name of the S3 aceess point for the billing role in the redaction case.
	// Default: 'bill-s3-access-point-call-transcripts-known-pii'.
	//
	Billing *string `field:"required" json:"billing" yaml:"billing"`
	// The name of the S3 aceess point for the customer support role in the redaction case.
	// Default: 'cs-s3-access-point-call-transcripts-known-pii'.
	//
	CustomerSupport *string `field:"required" json:"customerSupport" yaml:"customerSupport"`
	// The name of the S3 aceess point for the general role in the access control case.
	// Default: 'accessctl-s3-ap-survey-results-unknown-pii'.
	//
	General *string `field:"required" json:"general" yaml:"general"`
}

