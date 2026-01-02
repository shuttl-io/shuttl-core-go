// The JSII library for Shuttl AI models
package core

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@shuttl-io/core.Agent",
		reflect.TypeOf((*Agent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getTool", GoMethod: "GetTool"},
			_jsii_.MemberMethod{JsiiMethod: "getToolCallResult", GoMethod: "GetToolCallResult"},
			_jsii_.MemberMethod{JsiiMethod: "invoke", GoMethod: "Invoke"},
			_jsii_.MemberProperty{JsiiProperty: "model", GoGetter: "Model"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outcomes", GoGetter: "Outcomes"},
			_jsii_.MemberMethod{JsiiMethod: "respondWithToolCall", GoMethod: "RespondWithToolCall"},
			_jsii_.MemberProperty{JsiiProperty: "systemPrompt", GoGetter: "SystemPrompt"},
			_jsii_.MemberProperty{JsiiProperty: "toolkits", GoGetter: "Toolkits"},
			_jsii_.MemberProperty{JsiiProperty: "tools", GoGetter: "Tools"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
		},
		func() interface{} {
			return &jsiiProxy_Agent{}
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.AgentProps",
		reflect.TypeOf((*AgentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.AgentStreamer",
		reflect.TypeOf((*AgentStreamer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "recieve", GoMethod: "Recieve"},
		},
		func() interface{} {
			j := jsiiProxy_AgentStreamer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IModelStreamer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ApiAuth",
		reflect.TypeOf((*ApiAuth)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.ApiTrigger",
		reflect.TypeOf((*ApiTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activate", GoMethod: "Activate"},
			_jsii_.MemberMethod{JsiiMethod: "bindOutcome", GoMethod: "BindOutcome"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outcome", GoGetter: "Outcome"},
			_jsii_.MemberMethod{JsiiMethod: "parseArgs", GoMethod: "ParseArgs"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "triggerType", GoGetter: "TriggerType"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "withName", GoMethod: "WithName"},
		},
		func() interface{} {
			j := jsiiProxy_ApiTrigger{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseTrigger)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ApiTriggerArgs",
		reflect.TypeOf((*ApiTriggerArgs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ApiTriggerConfig",
		reflect.TypeOf((*ApiTriggerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.App",
		reflect.TypeOf((*App)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAgent", GoMethod: "AddAgent"},
			_jsii_.MemberMethod{JsiiMethod: "addToolkit", GoMethod: "AddToolkit"},
			_jsii_.MemberProperty{JsiiProperty: "agents", GoGetter: "Agents"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "serve", GoMethod: "Serve"},
			_jsii_.MemberProperty{JsiiProperty: "server", GoGetter: "Server"},
		},
		func() interface{} {
			return &jsiiProxy_App{}
		},
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.BaseTrigger",
		reflect.TypeOf((*BaseTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activate", GoMethod: "Activate"},
			_jsii_.MemberMethod{JsiiMethod: "bindOutcome", GoMethod: "BindOutcome"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outcome", GoGetter: "Outcome"},
			_jsii_.MemberMethod{JsiiMethod: "parseArgs", GoMethod: "ParseArgs"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "triggerType", GoGetter: "TriggerType"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "withName", GoMethod: "WithName"},
		},
		func() interface{} {
			j := jsiiProxy_BaseTrigger{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITrigger)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.EmailTrigger",
		reflect.TypeOf((*EmailTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activate", GoMethod: "Activate"},
			_jsii_.MemberMethod{JsiiMethod: "bindOutcome", GoMethod: "BindOutcome"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outcome", GoGetter: "Outcome"},
			_jsii_.MemberMethod{JsiiMethod: "parseArgs", GoMethod: "ParseArgs"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "triggerType", GoGetter: "TriggerType"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "withName", GoMethod: "WithName"},
		},
		func() interface{} {
			j := jsiiProxy_EmailTrigger{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseTrigger)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.EmailTriggerConfig",
		reflect.TypeOf((*EmailTriggerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.EnvSecret",
		reflect.TypeOf((*EnvSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "envVarName", GoGetter: "EnvVarName"},
			_jsii_.MemberMethod{JsiiMethod: "resolveSecret", GoMethod: "ResolveSecret"},
		},
		func() interface{} {
			j := jsiiProxy_EnvSecret{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISecret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.FileAttachment",
		reflect.TypeOf((*FileAttachment)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.FileTrigger",
		reflect.TypeOf((*FileTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activate", GoMethod: "Activate"},
			_jsii_.MemberMethod{JsiiMethod: "bindOutcome", GoMethod: "BindOutcome"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outcome", GoGetter: "Outcome"},
			_jsii_.MemberMethod{JsiiMethod: "parseArgs", GoMethod: "ParseArgs"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "triggerType", GoGetter: "TriggerType"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "withName", GoMethod: "WithName"},
		},
		func() interface{} {
			j := jsiiProxy_FileTrigger{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseTrigger)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.FileTriggerConfig",
		reflect.TypeOf((*FileTriggerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IAgentStreamerWriter",
		reflect.TypeOf((*IAgentStreamerWriter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "write", GoMethod: "Write"},
			_jsii_.MemberMethod{JsiiMethod: "writeObject", GoMethod: "WriteObject"},
		},
		func() interface{} {
			return &jsiiProxy_IAgentStreamerWriter{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IApiAuthenticator",
		reflect.TypeOf((*IApiAuthenticator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "authenticate", GoMethod: "Authenticate"},
		},
		func() interface{} {
			return &jsiiProxy_IApiAuthenticator{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IModel",
		reflect.TypeOf((*IModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "invoke", GoMethod: "Invoke"},
			_jsii_.MemberProperty{JsiiProperty: "threadId", GoGetter: "ThreadId"},
		},
		func() interface{} {
			return &jsiiProxy_IModel{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IModelFactory",
		reflect.TypeOf((*IModelFactory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "create", GoMethod: "Create"},
		},
		func() interface{} {
			return &jsiiProxy_IModelFactory{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IModelFactoryProps",
		reflect.TypeOf((*IModelFactoryProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "systemPrompt", GoGetter: "SystemPrompt"},
			_jsii_.MemberProperty{JsiiProperty: "tools", GoGetter: "Tools"},
		},
		func() interface{} {
			return &jsiiProxy_IModelFactoryProps{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IModelResponseStream",
		reflect.TypeOf((*IModelResponseStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
		},
		func() interface{} {
			return &jsiiProxy_IModelResponseStream{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IModelStreamer",
		reflect.TypeOf((*IModelStreamer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "recieve", GoMethod: "Recieve"},
		},
		func() interface{} {
			return &jsiiProxy_IModelStreamer{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IOutcome",
		reflect.TypeOf((*IOutcome)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bindToRequest", GoMethod: "BindToRequest"},
			_jsii_.MemberMethod{JsiiMethod: "send", GoMethod: "Send"},
		},
		func() interface{} {
			return &jsiiProxy_IOutcome{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IPCRequest",
		reflect.TypeOf((*IPCRequest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "body", GoGetter: "Body"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
		},
		func() interface{} {
			return &jsiiProxy_IPCRequest{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IPCResponse",
		reflect.TypeOf((*IPCResponse)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "errorObj", GoGetter: "ErrorObj"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
			_jsii_.MemberProperty{JsiiProperty: "success", GoGetter: "Success"},
		},
		func() interface{} {
			return &jsiiProxy_IPCResponse{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IPCResponseError",
		reflect.TypeOf((*IPCResponseError)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
		},
		func() interface{} {
			return &jsiiProxy_IPCResponseError{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IRateTriggerOnTrigger",
		reflect.TypeOf((*IRateTriggerOnTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "onTrigger", GoMethod: "OnTrigger"},
		},
		func() interface{} {
			return &jsiiProxy_IRateTriggerOnTrigger{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.ISecret",
		reflect.TypeOf((*ISecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolveSecret", GoMethod: "ResolveSecret"},
		},
		func() interface{} {
			return &jsiiProxy_ISecret{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.IServer",
		reflect.TypeOf((*IServer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "accept", GoMethod: "Accept"},
			_jsii_.MemberMethod{JsiiMethod: "start", GoMethod: "Start"},
			_jsii_.MemberMethod{JsiiMethod: "stop", GoMethod: "Stop"},
		},
		func() interface{} {
			return &jsiiProxy_IServer{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.ITool",
		reflect.TypeOf((*ITool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "execute", GoMethod: "Execute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
		},
		func() interface{} {
			return &jsiiProxy_ITool{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.ITrigger",
		reflect.TypeOf((*ITrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activate", GoMethod: "Activate"},
			_jsii_.MemberMethod{JsiiMethod: "bindOutcome", GoMethod: "BindOutcome"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outcome", GoGetter: "Outcome"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "triggerType", GoGetter: "TriggerType"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "withName", GoMethod: "WithName"},
		},
		func() interface{} {
			return &jsiiProxy_ITrigger{}
		},
	)
	_jsii_.RegisterInterface(
		"@shuttl-io/core.ITriggerInvoker",
		reflect.TypeOf((*ITriggerInvoker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "defaultOutcome", GoMethod: "DefaultOutcome"},
			_jsii_.MemberMethod{JsiiMethod: "invoke", GoMethod: "Invoke"},
		},
		func() interface{} {
			return &jsiiProxy_ITriggerInvoker{}
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.InputContent",
		reflect.TypeOf((*InputContent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.InputTokensDetails",
		reflect.TypeOf((*InputTokensDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.Model",
		reflect.TypeOf((*Model)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Model{}
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ModelContent",
		reflect.TypeOf((*ModelContent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ModelDeltaOutput",
		reflect.TypeOf((*ModelDeltaOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ModelResponse",
		reflect.TypeOf((*ModelResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ModelResponseData",
		reflect.TypeOf((*ModelResponseData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ModelResponseStreamValue",
		reflect.TypeOf((*ModelResponseStreamValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ModelTextOutput",
		reflect.TypeOf((*ModelTextOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ModelToolOutput",
		reflect.TypeOf((*ModelToolOutput)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.Outcomes",
		reflect.TypeOf((*Outcomes)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bindToRequest", GoMethod: "BindToRequest"},
			_jsii_.MemberMethod{JsiiMethod: "send", GoMethod: "Send"},
		},
		func() interface{} {
			j := jsiiProxy_Outcomes{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOutcome)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.OutputTokensDetails",
		reflect.TypeOf((*OutputTokensDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.Rate",
		reflect.TypeOf((*Rate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activate", GoMethod: "Activate"},
			_jsii_.MemberMethod{JsiiMethod: "bindOutcome", GoMethod: "BindOutcome"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outcome", GoGetter: "Outcome"},
			_jsii_.MemberMethod{JsiiMethod: "parseArgs", GoMethod: "ParseArgs"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "triggerType", GoGetter: "TriggerType"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "withName", GoMethod: "WithName"},
			_jsii_.MemberMethod{JsiiMethod: "withOnTrigger", GoMethod: "WithOnTrigger"},
		},
		func() interface{} {
			j := jsiiProxy_Rate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseTrigger)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.Schema",
		reflect.TypeOf((*Schema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
		},
		func() interface{} {
			return &jsiiProxy_Schema{}
		},
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
		},
		func() interface{} {
			return &jsiiProxy_Secret{}
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.SerializedHTTPRequest",
		reflect.TypeOf((*SerializedHTTPRequest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.SlackOutcome",
		reflect.TypeOf((*SlackOutcome)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bindToRequest", GoMethod: "BindToRequest"},
			_jsii_.MemberMethod{JsiiMethod: "send", GoMethod: "Send"},
		},
		func() interface{} {
			j := jsiiProxy_SlackOutcome{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOutcome)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.StdInServer",
		reflect.TypeOf((*StdInServer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "accept", GoMethod: "Accept"},
			_jsii_.MemberMethod{JsiiMethod: "start", GoMethod: "Start"},
			_jsii_.MemberMethod{JsiiMethod: "stop", GoMethod: "Stop"},
		},
		func() interface{} {
			j := jsiiProxy_StdInServer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.StreamingOutcome",
		reflect.TypeOf((*StreamingOutcome)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bindToRequest", GoMethod: "BindToRequest"},
			_jsii_.MemberMethod{JsiiMethod: "send", GoMethod: "Send"},
		},
		func() interface{} {
			j := jsiiProxy_StreamingOutcome{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOutcome)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ToolArg",
		reflect.TypeOf((*ToolArg)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.ToolArgBuilder",
		reflect.TypeOf((*ToolArgBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "argType", GoGetter: "ArgType"},
			_jsii_.MemberMethod{JsiiMethod: "defaultTo", GoMethod: "DefaultTo"},
			_jsii_.MemberProperty{JsiiProperty: "defaultValue", GoGetter: "DefaultValue"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enumValues", GoGetter: "EnumValues"},
			_jsii_.MemberMethod{JsiiMethod: "isRequired", GoMethod: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "required", GoGetter: "Required"},
		},
		func() interface{} {
			return &jsiiProxy_ToolArgBuilder{}
		},
	)
	_jsii_.RegisterClass(
		"@shuttl-io/core.Toolkit",
		reflect.TypeOf((*Toolkit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTool", GoMethod: "AddTool"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "tools", GoGetter: "Tools"},
		},
		func() interface{} {
			return &jsiiProxy_Toolkit{}
		},
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.ToolkitProps",
		reflect.TypeOf((*ToolkitProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.TriggerOutput",
		reflect.TypeOf((*TriggerOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@shuttl-io/core.Usage",
		reflect.TypeOf((*Usage)(nil)).Elem(),
	)
}
