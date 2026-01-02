package core


type ApiTriggerConfig struct {
	Authenticator IApiAuthenticator `field:"optional" json:"authenticator" yaml:"authenticator"`
	Cors *[]*string `field:"optional" json:"cors" yaml:"cors"`
}

