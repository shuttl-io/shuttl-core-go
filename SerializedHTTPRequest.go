package core


// Represents a serialized HTTP request from the serve command.
type SerializedHTTPRequest struct {
	// The Content-Type header value.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// HTTP headers as key-value pairs with array values.
	Headers *map[string]*[]*string `field:"required" json:"headers" yaml:"headers"`
	// The host header value.
	Host *string `field:"required" json:"host" yaml:"host"`
	// The HTTP method (POST, GET, etc.).
	Method *string `field:"required" json:"method" yaml:"method"`
	// The request path.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The HTTP protocol version.
	Proto *string `field:"required" json:"proto" yaml:"proto"`
	// Query parameters as key-value pairs with array values.
	Query *map[string]*[]*string `field:"required" json:"query" yaml:"query"`
	// The remote address of the client.
	RemoteAddr *string `field:"required" json:"remoteAddr" yaml:"remoteAddr"`
	// Timestamp of when the request was received.
	Timestamp *string `field:"required" json:"timestamp" yaml:"timestamp"`
	// The request body (parsed JSON or raw).
	Body interface{} `field:"optional" json:"body" yaml:"body"`
}

