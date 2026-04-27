package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)

const (
	StartUp         SubCategory = "StartUp"
	Get             SubCategory = "Get"
	Shutdown        SubCategory = "Shutdown"
	ExternalService SubCategory = "ExternalService"

	Select   SubCategory = "Select"
	Insert   SubCategory = "Insert"
	Update   SubCategory = "Update"
	Delete   SubCategory = "Delete"
	Rollback SubCategory = "Rollback"

	Api          SubCategory = "Api"
	HashPassword SubCategory = "HashPassword"

	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
)
