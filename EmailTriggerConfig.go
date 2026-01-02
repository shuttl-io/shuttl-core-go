package core


type EmailTriggerConfig struct {
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	FromAddress *string `field:"required" json:"fromAddress" yaml:"fromAddress"`
	InboxName *string `field:"required" json:"inboxName" yaml:"inboxName"`
	SubjectPattern *string `field:"required" json:"subjectPattern" yaml:"subjectPattern"`
}

