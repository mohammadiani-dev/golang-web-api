package logging

type Category string
type subCategory string
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
	//general
	StartUp         subCategory = "StartUp"
	ExternalService subCategory = "ExternalService"

	//postgres
	Select   subCategory = "Select"
	Rollback subCategory = "Rollback"
	Update   subCategory = "Update"
	Delete   subCategory = "Delete"
	Insert   subCategory = "Insert"

	//api
	Api                 subCategory = "Api"
	HashPassword        subCategory = "HashPassword"
	DefaultRoleNotFound subCategory = "DefaultRoleNotFound"

	//validation
	MobileValidation   subCategory = "MobileValidation"
	PasswordValidation subCategory = "PasswordValidation"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIP     ExtraKey = "ClientIP"
	HostIP       ExtraKey = "HostIP"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
	ClientIp     ExtraKey = "ClientIp"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
)
