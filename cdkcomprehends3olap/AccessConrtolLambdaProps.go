package cdkcomprehends3olap


type AccessConrtolLambdaProps struct {
	// The minimum prediction confidence score above which PII classification and detection would be considered as final answer.
	//
	// Valid range (0.5 to 1.0).
	// Default: '0.5'
	//
	ConfidenceThreshold *string `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// Number of threads to use for calling Comprehend's ContainsPiiEntities API.
	//
	// This controls the number of simultaneous calls that will be made from this Lambda.
	// Default: '20'.
	//
	ContainsPiiEntitiesThreadCount *string `field:"optional" json:"containsPiiEntitiesThreadCount" yaml:"containsPiiEntitiesThreadCount"`
	// Default language of the text to be processed.
	//
	// This code will be used for interacting with Comprehend.
	// Default: 'en'.
	//
	DefaultLanguageCode *string `field:"optional" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// Default maximum document size (in bytes) that this function can process otherwise will throw exception for too large document size.
	// Default: '102400'.
	//
	DocumentMaxSize *string `field:"optional" json:"documentMaxSize" yaml:"documentMaxSize"`
	// Maximum document size (in bytes) to be used for making calls to Comprehend's ContainsPiiEntities API.
	// Default: '50000'.
	//
	DocumentMaxSizeContainsPiiEntities *string `field:"optional" json:"documentMaxSizeContainsPiiEntities" yaml:"documentMaxSizeContainsPiiEntities"`
	// Whether to support partial objects or not.
	//
	// Accessing partial object through http headers such byte-range can corrupt the object and/or affect PII detection accuracy.
	// Default: 'false'.
	//
	IsPartialObjectSupported *string `field:"optional" json:"isPartialObjectSupported" yaml:"isPartialObjectSupported"`
	// Log level for Lambda function logging, e.g., ERROR, INFO, DEBUG, etc.
	// Default: 'INFO'.
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Maximum characters to overlap among segments of a document in case chunking is needed because of maximum document size limit.
	// Default: '200'.
	//
	MaxCharsOverlap *string `field:"optional" json:"maxCharsOverlap" yaml:"maxCharsOverlap"`
	// List of comma separated PII entity types to be considered for access control.
	//
	// Refer Comprehend's documentation page for list of supported PII entity types.
	// Default: 'ALL'.
	//
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// True if publish metrics to Cloudwatch, false otherwise.
	//
	// See README.md for details on CloudWatch metrics.
	// Default: 'true'.
	//
	PublishCloudWatchMetrics *string `field:"optional" json:"publishCloudWatchMetrics" yaml:"publishCloudWatchMetrics"`
	// The version of the serverless application.
	// Default: '1.0.2'
	//
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
	// Number of tokens/words to overlap among segments of a document in case chunking is needed because of maximum document size limit.
	// Default: '20'.
	//
	SubsegmentOverlappingTokens *string `field:"optional" json:"subsegmentOverlappingTokens" yaml:"subsegmentOverlappingTokens"`
	// Handling logic for Unsupported files.
	//
	// Valid values are PASS and FAIL.
	// Default: 'FAIL'.
	//
	UnsupportedFileHandling *string `field:"optional" json:"unsupportedFileHandling" yaml:"unsupportedFileHandling"`
}

