package core


type ApiTriggerArgs struct {
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	Method *string `field:"required" json:"method" yaml:"method"`
	Auth *ApiAuth `field:"optional" json:"auth" yaml:"auth"`
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	Cookies *map[string]*string `field:"optional" json:"cookies" yaml:"cookies"`
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	PathParams *map[string]*string `field:"optional" json:"pathParams" yaml:"pathParams"`
	QueryParams *map[string]*string `field:"optional" json:"queryParams" yaml:"queryParams"`
}

