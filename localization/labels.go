package labels

const (

	// General labels
	Code       = "code"
	RequestID  = "request-id"
	HTTPScheme = "http-scheme"
	HTTPProto  = "http-proto"
	HTTPMethod = "http-method"
	RemoteAddr = "remote-addr"
	UserAgent  = "user-agent"
	URI        = "uri"

	// HTTPContentType label
	HTTPContentTypeKey   = "Content-Type"
	HTTPContentTypeValue = "application/json; charset=utf-8"
	HTTPUSERAGENTKey     = "User-Agent"
	HTTPUSERAGENTValue   = "TriDubai-Api/1.0"

	// logger levels
	Info  = "info"
	Debug = "debug"
	Error = "error"

	// Custom Millisecond for RC3339
	RFC3339Milli = "2006-01-02T15:04:05.999Z07:00"
)