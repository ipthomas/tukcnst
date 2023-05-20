/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE', which is part of this source code package.
 */

package tukcnst

const (
	DEFAULT_TUK_BASEPATH                    = "./config/"
	DEFAULT_TUK_SERVICE_CONFIG_FILE         = "eventsrvc"
	DEFAULT_TUK_SERVICE_LOG_FOLDER          = "logs"
	QUERY_PARAM_INCLUDE                     = "_include"
	QUERY_PARAM_MRN_ID                      = "mrnid"
	QUERY_PARAM_MRN_OID                     = "mrnoid"
	QUERY_PARAM_NHS_ID                      = "nhsid"
	QUERY_PARAM_NHS_OID                     = "nhsoid"
	QUERY_PARAM_REG_ID                      = "regid"
	QUERY_PARAM_REG_OID                     = "regoid"
	QUERY_PARAM_ACTION                      = "action"
	QUERY_PARAM_BROKER_REF                  = "brokerref"
	QUERY_PARAM_PATHWAY                     = "pathway"
	QUERY_PARAM_TOPIC                       = "topic"
	QUERY_PARAM_EXPRESSION                  = "expression"
	QUERY_PARAM_ID                          = "id"
	QUERY_PARAM_TASK                        = "task"
	QUERY_PARAM_USER                        = "user"
	QUERY_PARAM_ORG                         = "org"
	QUERY_PARAM_ROLE                        = "role"
	QUERY_PARAM_VERSION                     = "vers"
	QUERY_PARAM_PDQ_SERVER_TYPE             = "pdqserver"
	QUERY_PARAM_RESPONSE_TYPE               = "rsptype"
	QUERY_PARAM_CACHE                       = "cache"
	QUERY_PARAM_RSP_TYPE                    = "rsptype"
	ENV_TUK_CONFIG                          = "TUK_CONFIG"
	ENV_TUK_CONFIG_FILE                     = "TUK_CONFIG_FILE"
	ENV_RESPONSE_TYPE                       = "RSP_TYPE"
	ENV_RESPONSE_TYPE_HTTP_CODE             = "code"
	ENV_RESPONSE_TYPE_BOOL                  = "bool"
	ENV_PATIENT_CACHE                       = "PATIENT_CACHE"
	ENV_NHS_OID                             = "NHS_OID"
	ENV_REG_OID                             = "REG_OID"
	ENV_IHE_PDQV3_SERVER_URL                = "IHE_PDQV3_SERVER_URL"
	ENV_IHE_PIXV3_SERVER_URL                = "IHE_PIXV3_SERVER_URL"
	ENV_IHE_PIXM_SERVER_URL                 = "IHE_PIXM_SERVER_URL"
	ENV_CGL_SERVER_URL                      = "CGL_SERVER_URL"
	ENV_CGL_X_API_KEY                       = "CGL_API_KEY"
	ENV_CGL_X_API_SECRET                    = "CGL_X_API_SECRET"
	ENV_PDQ_SERVER_TYPE                     = "PDQ_SERVER_TYPE"
	ENV_PDQ_SERVER_URL                      = "PDQ_SERVER_URL"
	ENV_DSUB_BROKER_URL                     = "DSUB_BROKER_URL"
	ENV_DSUB_CONSUMER_URL                   = "DSUB_CONSUMER_URL"
	ENV_HTML_CREATOR_URL                    = "HTML_CREATOR_URL"
	ENV_EVENT_CONSUMER_URL                  = "EVENT_CONSUMER_URL"
	ENV_EVENT_SUBSCRIBER_URL                = "EVENT_SUBSCRIBER_URL"
	ENV_EVENT_NOTIFIER_URL                  = "EVENT_NOTIFIER_URL"
	ENV_XDW_CONSUMER_URL                    = "XDW_CONSUMER_URL"
	ENV_XDW_CREATOR_URL                     = "XDW_CREATOR_URL"
	ENV_XDW_ADMIN_URL                       = "ENV_XDW_ADMIN_URL"
	ENV_TUK_DB_URL                          = "TUK_DB_URL"
	ENV_DB_HOST                             = "DB_HOST"
	ENV_DB_NAME                             = "DB_NAME"
	ENV_DB_PORT                             = "DB_PORT"
	ENV_DB_USER                             = "DB_USER"
	ENV_DB_PASSWORD                         = "DB_PASSWORD"
	ENV_DEBUG_MODE                          = "DEBUG_MODE"
	TUK_HTTP_SERVER_SCHEME                  = "http://"
	TUK_HTTP_SERER_SCHEME_SECURE            = "https://"
	TUK_HTTP_SERVER_DEFAULT_PORT            = ":8080"
	PDQ_SERVER_TYPE_IHE_PIXM                = "pixm"
	PDQ_SERVER_TYPE_IHE_PDQV3               = "pdqv3"
	PDQ_SERVER_TYPE_IHE_PIXV3               = "pixv3"
	PDQ_SERVER_TYPE_CGL                     = "cgl"
	OPEN                                    = "OPEN"
	READY                                   = "READY"
	CLOSED                                  = "CLOSED"
	TASK                                    = "task"
	NO_VALUE                                = "Not Provided"
	TUK_DB_TABLE_SUBSCRIPTIONS              = "subscriptions"
	TUK_DB_TABLE_XDWS                       = "xdws"
	APPLICATION_JSON                        = "application/json"
	APPLICATION_JSON_CHARSET_UTF_8          = APPLICATION_JSON + "; charset=utf-8"
	XDW_DEFINITION_FILE                     = "_xdwdef"
	DASHBOARD                               = "dashboard"
	SPA                                     = "spa"
	TIMELINE                                = "timeline"
	COMPLETE                                = "COMPLETE"
	CREATED                                 = "CREATED"
	XML                                     = "xml"
	JSON                                    = "json"
	HTML                                    = "html"
	PATHWAY                                 = "pathway"
	DEMO                                    = "demo"
	PULLPOINTS                              = "pullpoints"
	CONFIG                                  = "config"
	AUDITS                                  = "audits"
	EVENT_ACKS                              = "eventacks"
	EVENTS                                  = "events"
	ID_MAPS                                 = "idmaps"
	TEMPLATES                               = "templates"
	STATICS                                 = "statics"
	SERVICE_STATES                          = "servicestates"
	SUBSCRIBE                               = "subscribe"
	SUBSCRIBER                              = "subscriber"
	SUBSCRIPTIONS                           = "subscriptions"
	REGISTER                                = "register"
	WIDGET                                  = "widget"
	ADMIN                                   = "admin"
	TEMPLATE                                = "template"
	CREATE                                  = "create"
	CREATOR                                 = "creator"
	CANCEL                                  = "cancel"
	LIST                                    = "list"
	ACK                                     = "ack"
	ATTACHMENT                              = "attachment"
	TOPICS                                  = "topics"
	EXPRESSIONS                             = "expressions"
	PIX_QUERY                               = "pixm"
	STS                                     = "sts"
	STATEID                                 = "stateID"
	PATIENT                                 = "patient"
	SERVICES                                = "services"
	XDW                                     = "xdw"
	XDWS                                    = "xdws"
	WORKFLOW                                = "workflow"
	WORKFLOWS                               = "workflows"
	ALERT                                   = "alert"
	HTTP_POST                               = "POST"
	HTTP_PUT                                = "PUT"
	HTTP_GET                                = "GET"
	HTTP_DELETE                             = "DELETE"
	RST_ISSUE                               = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/RST/Issue"
	SUBSCRIPTION_REFERENCE_PREFIX           = "https://NotificationBrokerServer/Subscription/"
	SOAP_ACTION_UNSUBSCRIBE_REQUEST         = "http://docs.oasis-open.org/wsn/bw2/SubscriptionManager/UnsubscribeRequest"
	SOAP_ACTION_SUBSCRIBE_REQUEST           = "http://docs.oasis-open.org/wsn/bw-2/NotificationProducer/SubscribeRequest"
	SOAP_ACTION_PIXV3_Request               = "urn:hl7-org:v3:PRPA_IN201309UV02"
	SOAP_ACTION_PDQV3_Request               = "urn:hl7-org:v3:PRPA_IN201305UV02"
	SOAP_ACTION                             = "SOAPAction"
	CONTENT_TYPE                            = "Content-Type"
	TEXT_HTML                               = "text/html"
	TEXT_PLAIN                              = "text/plain"
	APPLICATION_XML                         = "application/xml"
	SOAP_XML                                = "application/soap+xml"
	ACCEPT                                  = "Accept"
	AUTHORIZATION                           = "Authorization"
	ALL                                     = "*/*"
	CONNECTION                              = "Connection"
	KEEP_ALIVE                              = "keep-alive"
	UNSUBSCRIBE_RESPONSE                    = "UnsubscribeResponse"
	PULLPOINT                               = "pullpoint"
	IN_PROGRESS                             = "IN_PROGRESS"
	TEXT_XML_CHARSET_UTF_8                  = "text/xml; charset=utf-8"
	A_ADDRESS                               = "a:Address"
	OK                                      = "OK"
	SUBMISSION_TIME                         = "submissionTime"
	CREATION_TIME                           = "creationTime"
	REPOSITORY_UID                          = "repositoryUniqueId"
	SOURCE_PATIENT_ID                       = "sourcePatientId"
	URN_XDS_PID                             = "urn:uuid:58a6f841-87b3-4a3e-92fd-a8ffeff98427"
	URN_XDS_DOCUID                          = "urn:uuid:2e82c1f6-a085-4c72-9da3-8640a32e42ab"
	URN_SUBMISSION_SOURCE_ID                = "urn:uuid:127d4267-6169-4896-b46c-b34d4d1b6d5d"
	URN_SUBMISSION_UID                      = "urn:uuid:96fdda7c-d067-4183-912e-bf5ee74998a8"
	URN_SUBMISSION_PID                      = "urn:uuid:6b5aea1a-874d-4603-a4bc-96a0a7b38446"
	URN_CLASS_CODE                          = "urn:uuid:41a5887f-8865-4c09-adf7-e362475b143a"
	URN_CONF_CODE                           = "urn:uuid:f4f85eac-e6cb-4883-b524-f2705394840f"
	URN_FORMAT_CODE                         = "urn:uuid:a09d5840-386c-46f2-b5ad-9c3699a4309d"
	URN_FACILITY_CODE                       = "urn:uuid:f33fb8ac-18af-42cc-ae0e-ed0b0bdb91e1"
	URN_PRACTICE_CODE                       = "urn:uuid:cccf5598-8b07-4b77-a05e-ae952c785ead"
	URN_TYPE_CODE                           = "urn:uuid:f0306f51-975f-434e-a61c-c59651d33983"
	URN_AUTHOR                              = "urn:uuid:93606bcf-9494-43ec-9b4e-a7748d1a838d"
	URN_EVENT_LIST                          = "urn:uuid:2c6b8cb7-8b2a-4051-b291-b1ae6a575ef4"
	AUTHOR_PERSON                           = "authorPerson"
	AUTHOR_INSTITUTION                      = "authorInstitution"
	AUTHOR_SPECIALITY                       = "authorSpecialty"
	AUTHOR_ROLE                             = "authorRole"
	XMLNS_XSI                               = "http://www.w3.org/2001/XMLSchema-instance"
	XMLNS_XSD                               = "http://www.w3.org/2001/XMLSchema"
	XMLNS                                   = "http://docs.oasis-open.org/wsn/b-2"
	MYSQL                                   = "mysql"
	INSERT                                  = "insert"
	SELECT                                  = "select"
	DELETE                                  = "delete"
	DEPRECATE                               = "deprecate"
	REPLACE                                 = "replace"
	UPDATE                                  = "update"
	ISPUBLISHED                             = "ispublished"
	APPEND                                  = "append"
	XDSDOMAIN                               = "XDSDOMAIN"
	WorkflowInstanceId                      = "^^^^urn:ihe:iti:xdw:2013:workflowInstanceId"
	XDWNameSpace                            = "urn:ihe:iti:xdw:2011"
	XDWNameLocal                            = "xdw"
	HL7NameSpace                            = "urn:hl7-org:v3"
	HL7NameLocal                            = "hl7"
	WHTNameSpace                            = "http://docs.oasis-open.org/ns/bpel4people/ws-humantask/types/200803"
	WHTNameLocal                            = "ws-ht"
	WorkflowDocumentXsi                     = "http://www.w3.org/2001/XMLSchema-instance"
	WorkflowDocumentSchemaLocation          = "urn:ihe:iti:xdw:2011 XDW-2014-12-23.xsd"
	XDS_REGISTERED                          = "urn:ihe:iti:xdw:2011:XDSregistered"
	MEDIA_TYPES                             = "http://www.iana.org/assignments/media-types"
	ASSERTION_SUBJECT_ID                    = "urn:oasis:names:tc:xspa:1.0:subject:subject-id"
	ASSERTION_ORGANISATION                  = "urn:oasis:names:tc:xspa:1.0:subject:organization"
	ASSERTION_ROLE                          = "urn:oasis:names:tc:xacml:2.0:subject:role"
	ASSERTION_POU                           = "urn:oasis:names:tc:xspa:1.0:subject:purposeofuse"
	ACT                                     = "act"
	RETRIEVE_DOCUMENT_SET                   = "urn:ihe:iti:2007:RetrieveDocumentSet"
	NOTIFICATION_ELEMENT                    = "<wsnt:NotificationMessage>"
	DSUB_NOTIFY_ELEMENT                     = "wsnt:Notify"
	EXPRESSION_EQUALS                       = "expression ="
	SQL_DEFAULT_EVENT_ACKS                  = "SELECT * FROM eventacks"
	SQL_DEFAULT_IDMAPS                      = "SELECT * FROM idmaps"
	SQL_DEFAULT_SERVICESTATES               = "SELECT * FROM servicestates"
	SQL_DEFAULT_XDWS                        = "SELECT * FROM xdws"
	SQL_DEFAULT_EVENTS                      = "SELECT * FROM events"
	SQL_DEFAULT_TEMPLATES                   = "SELECT * FROM templates"
	SQL_DEFAULT_STATICS                     = "SELECT * FROM statics"
	SQL_DEFAULT_WORKFLOWS                   = "SELECT * FROM workflows"
	SQL_DEFAULT_SUBSCRIPTIONS               = "SELECT * FROM subscriptions"
	DSUB_TOPIC_TYPE_CODE                    = "$XDSDocumentEntryTypeCode"
	DSUB_BROKER_URL                         = "DSUB_Broker_URL"
	PIX_URL                                 = "PIX_URL"
	NHS_OID                                 = "NHS_OID"
	NHS_OID_DEFAULT                         = "2.16.840.1.113883.2.1.4.1"
	REGION_OID                              = "Region_OID"
	HOME_COMMUNITY_OID                      = "Home_Community_OID"
	FORMAT_JSON_PRETTY                      = "&_format=json&_pretty=true"
	URN_OID_PREFIX                          = "urn:oid:"
	DSUB_ACK_TEMPLATE                       = "DSUB_ACK_TEMPLATE"
	DSUB_SUBSCRIBE_TEMPLATE                 = "DSUB_SUBSCRIBE_TEMPLATE"
	DSUB_CANCEL_TEMPLATE                    = "DSUB_CANCEL_TEMPLATE"
	GO_Template_PDQ_V3_Request              = "{{define \"pdqv3\"}}<S:Envelope xmlns:S='http://www.w3.org/2003/05/soap-envelope' xmlns:env='http://www.w3.org/2003/05/soap-envelope'><S:Header><To xmlns='http://www.w3.org/2005/08/addressing'>{{.Server_URL}}</To><Action xmlns='http://www.w3.org/2005/08/addressing' S:mustUnderstand='true' xmlns:S='http://www.w3.org/2003/05/soap-envelope'>urn:hl7-org:v3:PRPA_IN201305UV02</Action><ReplyTo xmlns='http://www.w3.org/2005/08/addressing'><Address>http://www.w3.org/2005/08/addressing/anonymous</Address></ReplyTo><FaultTo xmlns='http://www.w3.org/2005/08/addressing'><Address>http://www.w3.org/2005/08/addressing/anonymous</Address></FaultTo><MessageID xmlns='http://www.w3.org/2005/08/addressing'>uuid:{{newuuid}}</MessageID></S:Header><S:Body><PRPA_IN201305UV02 xmlns='urn:hl7-org:v3' ITSVersion='XML_1.0'><id extension='1663079209882' root='1.3.6.1.4.1.21998.2.1.10.15'/><creationTime value='{{simpledatetime}}'/><versionCode code='V3PR1'/><interactionId extension='PRPA_IN201305UV02' root='2.16.840.1.113883.1.6'/><processingCode code='P'/><processingModeCode code='T'/><acceptAckCode code='AL'/><receiver typeCode='RCV'><device classCode='DEV' determinerCode='INSTANCE'><id root='1.3.6.1.4.1.21367.2009.2.2.795'/><asAgent classCode='AGNT'><representedOrganization classCode='ORG' determinerCode='INSTANCE'><id root='1.3.6.1.4.1.21367.2009.2.2.1'/></representedOrganization></asAgent></device></receiver><sender typeCode='SND'><device classCode='DEV' determinerCode='INSTANCE'><id assigningAuthorityName='EHR_TIANI-SPIRIT' root='1.3.6.1.4.1.21367.2011.2.2.7919'/><asAgent classCode='AGNT'><representedOrganization classCode='ORG' determinerCode='INSTANCE'><id assigningAuthorityName='Tiani-Cisco' root='1.3.6.1.4.1.21367.2011.2.7.5572'/></representedOrganization></asAgent></device></sender><controlActProcess classCode='CACT' moodCode='EVN'><code code='PRPA_TE201305UV02' codeSystem='2.16.840.1.113883.1.6'/><queryByParameter><queryId extension='1663079209880' root='1.3.6.1.4.1.21998.2.1.10.15'/><statusCode code='new'/><responseModalityCode code='R'/><responsePriorityCode code='I'/><matchCriterionList/><parameterList><livingSubjectId><value extension='{{.Used_PID}}'/><semanticsText>LivingSubject.id</semanticsText></livingSubjectId></parameterList></queryByParameter></controlActProcess></PRPA_IN201305UV02></S:Body></S:Envelope>{{end}}"
	GO_Template_PIX_V3_Request              = "{{define \"pixv3\"}}<S:Envelope xmlns:S='http://www.w3.org/2003/05/soap-envelope' xmlns:env='http://www.w3.org/2003/05/soap-envelope'><S:Header><To xmlns='http://www.w3.org/2005/08/addressing'>{{.Server_URL}}</To><Action xmlns='http://www.w3.org/2005/08/addressing' S:mustUnderstand='true' xmlns:S='http://www.w3.org/2003/05/soap-envelope'>urn:hl7-org:v3:PRPA_IN201309UV02</Action><ReplyTo xmlns='http://www.w3.org/2005/08/addressing'><Address>http://www.w3.org/2005/08/addressing/anonymous</Address></ReplyTo><FaultTo xmlns='http://www.w3.org/2005/08/addressing'><Address>http://www.w3.org/2005/08/addressing/anonymous</Address></FaultTo><MessageID xmlns='http://www.w3.org/2005/08/addressing'>uuid:{{newuuid}}</MessageID></S:Header><S:Body><PRPA_IN201309UV02 xmlns='urn:hl7-org:v3' ITSVersion='XML_1.0'><id extension='1663059665645' root='1.3.6.1.4.1.21998.2.1.10.12'/><creationTime value='{{simpledatetime}}'/><versionCode code='V3PR1'/><interactionId extension='PRPA_IN201309UV02' root='2.16.840.1.113883.1.6'/><processingCode code='P'/><processingModeCode code='T'/><acceptAckCode code='AL'/><receiver typeCode='RCV'><device classCode='DEV' determinerCode='INSTANCE'><id root='1.3.6.1.4.1.21367.2009.2.2.795'/><asAgent classCode='AGNT'><representedOrganization classCode='ORG' determinerCode='INSTANCE'><id root='1.3.6.1.4.1.21367.2009.2.2.1'/></representedOrganization></asAgent></device></receiver><sender typeCode='SND'><device classCode='DEV' determinerCode='INSTANCE'><id assigningAuthorityName='NHS' root='1.3.6.1.4.1.21367.2011.2.2.7919'/><asAgent classCode='AGNT'><representedOrganization classCode='ORG' determinerCode='INSTANCE'><id assigningAuthorityName='ICB' root='1.3.6.1.4.1.21367.2011.2.7.5572'/></representedOrganization></asAgent></device></sender><controlActProcess classCode='CACT' moodCode='EVN'><code code='PRPA_TE201309UV02' codeSystem='2.16.840.1.113883.1.6'/><queryByParameter><queryId extension='1663059665645' root='1.3.6.1.4.1.21998.2.1.10.12'/><statusCode code='new'/><responsePriorityCode code='I'/><parameterList><patientIdentifier><value assigningAuthorityName='{{.Used_PID_OID}}' extension='{{.Used_PID}}' root='{{.Used_PID_OID}}'/><semanticsText>Patient.id</semanticsText></patientIdentifier></parameterList></queryByParameter></controlActProcess></PRPA_IN201309UV02></S:Body></S:Envelope>{{end}}"
	GO_TEMPLATE_DSUB_ACK                    = "<SOAP-ENV:Envelope xmlns:SOAP-ENV='http://www.w3.org/2003/05/soap-envelope' xmlns:s='http://www.w3.org/2001/XMLSchema' xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'><SOAP-ENV:Body/></SOAP-ENV:Envelope>"
	GO_TEMPLATE_DSUB_CANCEL                 = "{{define \"cancel\"}}<soap:Envelope xmlns:soap='http://www.w3.org/2003/05/soap-envelope'><soap:Header><Action xmlns='http://www.w3.org/2005/08/addressing' soap:mustUnderstand='true'>http://docs.oasis-open.org/wsn/bw-2/SubscriptionManager/UnsubscribeRequest</Action><MessageID xmlns='http://www.w3.org/2005/08/addressing' soap:mustUnderstand='true'>urn:uuid:{{.UUID}}</MessageID><To xmlns='http://www.w3.org/2005/08/addressing' soap:mustUnderstand='true'>{{.BrokerRef}}</To><ReplyTo xmlns='http://www.w3.org/2005/08/addressing' soap:mustUnderstand='true'><Address>http://www.w3.org/2005/08/addressing/anonymous</Address></ReplyTo></soap:Header><soap:Body><Unsubscribe xmlns='http://docs.oasis-open.org/wsn/b-2' xmlns:ns2='http://www.w3.org/2005/08/addressing' xmlns:ns3='http://docs.oasis-open.org/wsrf/bf-2' xmlns:ns4='urn:oasis:names:tc:ebxml-regrep:xsd:rim:3.0' xmlns:ns5='urn:oasis:names:tc:ebxml-regrep:xsd:rs:3.0' xmlns:ns6='urn:oasis:names:tc:ebxml-regrep:xsd:lcm:3.0' xmlns:ns7='http://docs.oasis-open.org/wsn/t-1' xmlns:ns8='http://docs.oasis-open.org/wsrf/r-2'/></soap:Body></soap:Envelope>{{end}}"
	GO_TEMPLATE_DSUB_SUBSCRIBE              = "{{define \"subscribe\"}}<SOAP-ENV:Envelope xmlns:SOAP-ENV='http://www.w3.org/2003/05/soap-envelope' xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance' xmlns:s='http://www.w3.org/2001/XMLSchema' xmlns:wsa='http://www.w3.org/2005/08/addressing'><SOAP-ENV:Header><wsa:Action SOAP-ENV:mustUnderstand='true'>http://docs.oasis-open.org/wsn/bw-2/NotificationProducer/SubscribeRequest</wsa:Action><wsa:MessageID>urn:uuid:{{newuuid}}</wsa:MessageID><wsa:ReplyTo SOAP-ENV:mustUnderstand='true'><wsa:Address>http://www.w3.org/2005/08/addressing/anonymous</wsa:Address></wsa:ReplyTo><wsa:To>{{.BrokerURL}}</wsa:To></SOAP-ENV:Header><SOAP-ENV:Body><wsnt:Subscribe xmlns:wsnt='http://docs.oasis-open.org/wsn/b-2' xmlns:a='http://www.w3.org/2005/08/addressing' xmlns:rim='urn:oasis:names:tc:ebxml-regrep:xsd:rim:3.0' xmlns:wsa='http://www.w3.org/2005/08/addressing'><wsnt:ConsumerReference><wsa:Address>{{.ConsumerURL}}</wsa:Address></wsnt:ConsumerReference><wsnt:Filter><wsnt:TopicExpression Dialect='http://docs.oasis-open.org/wsn/t-1/TopicExpression/Simple'>ihe:FullDocumentEntry</wsnt:TopicExpression><rim:AdhocQuery id='urn:uuid:742790e0-aba6-43d6-9f1f-e43ed9790b79'><rim:Slot name='{{.Topic}}'><rim:ValueList><rim:Value>('{{.Expression}}')</rim:Value></rim:ValueList></rim:Slot></rim:AdhocQuery></wsnt:Filter></wsnt:Subscribe></SOAP-ENV:Body></SOAP-ENV:Envelope>{{end}}"
	XDW_ACTOR_CONTENT_CONSUMER              = "XDW_Consumer"
	XDW_ACTOR_CONTENT_CREATOR               = "XDW_Creator"
	XDW_ACTOR_CONTENT_UPDATER               = "XDW_Updater"
	XDW_ADMIN_REGISTER_DEFINITION           = "XDW_Register_Definition"
	XDW_ADMIN_Register_Template             = "HTML_Register_Template"
	XDW_ADMIN_REGISTER_XDS_META             = "XDW_Register_XDS_Meta"
	XDW_DOCEVENTTYPE_CREATE_WORKFLOW        = "CREATE_WORKFLOW"
	XDW_DOCEVENTTYPE_COMPLETED_WORKFLOW     = "COMPLETED_WORKFLOW"
	XDW_TASKEVENTTYPE_CREATE_TASK           = "CREATE_TASK"
	XDW_TASKEVENTTYPE_WORKFLOW_COMPLETED    = "WORKFLOW_COMPLETED"
	XDW_TASKEVENTTYPE_CREATED               = "created"
	XDW_TASKEVENTTYPE_CLAIM                 = "claim"
	XDW_TASKEVENTTYPE_START                 = "start"
	XDW_TASKEVENTTYPE_COMPLETE              = "complete"
	XDW_TASKEVENTTYPE_ATTACHMENT            = "addAttachment"
	XDW_TASKEVENTTYPE_COMMENT               = "addComment"
	XDW_TASKEVENTTYPE_ESCALATED             = "escalated"
	XDW_TASKEVENTTYPE_RESERVED              = "reserved"
	XDW_OPERATION_GET_TASKS                 = "getTasks"
	XDW_OPERATION_GET_PATHWAYS              = "getPathways"
	XDW_OPERATION_GET_ICB_PATIENT           = "getICBPatient"
	XDW_OPERATION_GET_CGL_PATIENT           = "getCGLPatient"
	XDW_OPERATION_GET_ALL_PATIENT           = "getAllPatientSrvs"
	XDW_OPERATION_ADD_ATTACHMENT            = "addAttachment"
	XDW_OPERATION_ADD_COMMENT               = "addComment"
	XDW_OPERATION_CLAIM                     = "claim"
	XDW_OPERATION_COMPLETE                  = "complete"
	XDW_OPERATION_DELEGATE                  = "delegate"
	XDW_OPERATION_GET_TASK_ATTACHMENT       = "getAttachment"
	XDW_OPERATION_GET_TASK_ATTACHMENT_INFOS = "getAttachmentInfos"
	XDW_OPERATION_GET_TASK_COMMENTS         = "getComments"
	XDW_OPERATION_GET_TASK_DESCRIPTION      = "getTaskDescription"
	XDW_OPERATION_GET_TASK_DETAILS          = "getTaskDetails"
	XDW_OPERATION_GET_TASK_HISTORY          = "getTaskHistory"
	XDW_OPERATION_GET_MYTASK_DETAILS        = "getMyTaskDetails"
	XDW_OPERATION_TASK_START                = "start"
	XDW_OPERATION_TASK_QUERY                = "query"
	TUK_EVENT_QUERY_PARAM_ID                = "id"
	TUK_EVENT_QUERY_PARAM_SAML              = "saml"
	TUK_EVENT_QUERY_PARAM_ACT               = "act"
	TUK_EVENT_QUERY_PARAM_TASK              = "task"
	TUK_EVENT_QUERY_PARAM_OP                = "op"
	TUK_EVENT_QUERY_PARAM_VERSION           = "vers"
	TUK_EVENT_QUERY_PARAM_STATUS            = "status"
	TUK_EVENT_QUERY_PARAM_TASK_ID           = "taskid"
	TUK_EVENT_QUERY_PARAM_PASSWORD          = "pwd"
	TUK_EVENT_QUERY_PARAM_ROLE              = "role"
	TUK_EVENT_QUERY_PARAM_ORG               = "org"
	TUK_EVENT_QUERY_PARAM_USER              = "user"
	TUK_EVENT_QUERY_PARAM_PATHWAY           = "pathway"
	TUK_EVENT_QUERY_PARAM_TOPIC             = "topic"
	TUK_EVENT_QUERY_PARAM_EXPRESSION        = "expression"
	TUK_EVENT_QUERY_PARAM_NOTES             = "notes"
	TUK_EVENT_QUERY_PARAM_NHS               = "nhs"
	TUK_EVENT_QUERY_PARAM_PID               = "pid"
	TUK_EVENT_QUERY_PARAM_PID_ORG           = "pidorg"
	TUK_EVENT_QUERY_PARAM_GIVEN_NAME        = "givenname"
	TUK_EVENT_QUERY_PARAM_FAMILY_NAME       = "familyname"
	TUK_EVENT_QUERY_PARAM_DOB               = "dob"
	TUK_EVENT_QUERY_PARAM_ZIP               = "zip"
	TUK_EVENT_QUERY_PARAM_GENDER            = "gender"
	TUK_EVENT_QUERY_PARAM_AUDIEANCE         = "audience"
	TUK_EVENT_QUERY_PARAM_CONFIG            = "config"
	TUK_EVENT_QUERY_PARAM_DOCREF            = "docref"
	TUK_EVENT_QUERY_PARAM_FORMAT            = "_format"
	TUK_EVENT_QUERY_PARAM_FORMAT_CODE       = "formatcode"
	TUK_EVENT_QUERY_PARAM_EVENT_TYPE        = "eventtype"
	TUK_EVENT_QUERY_PARAM_DOC_NAME          = "docname"
	TUK_EVENT_QUERY_PARAM_NAME              = "name"
	TUK_TEMPLATE_TEST                       = "test"
	TUK_TEMPLATE_EVENTS                     = "events"
	TUK_TEMPLATE_NOTIFICATION_WIDGET        = "notificationwidget"
	TUK_TEMPLATE_WORKFLOW                   = "workflow"
	TUK_TEMPLATE_WORKFLOW_TASKS             = "workflowtasks"
	TUK_TEMPLATE_DASHBOARD_WIDGET           = "dashboardwidget"
	TUK_TEMPLATE_TIMELINE_WIDGET            = "timelinewidget"
	TUK_TEMPLATE_ADMIN_SPA_WIDGET           = "adminspawidget"
	TUK_TEMPLATE_SPA_WIDGET                 = "spawidget"
	TUK_TEMPLATE_EVENTS_WIDGET              = "events"
	TUK_TEMPLATE_CONFIG_WIDGET              = "wfconfig"
	TUK_TEMPLATE_WORKFLOWS_WIDGET           = "workflows"
	TUK_TEMPLATE_SUBSCRIPTIONS_WIDGET       = "subscriptions"
	TUK_TASK_RESTART                        = "restart"
	TUK_TASK_INIT_SERVICES                  = "initsrvcs"
	TUK_TASK_INIT_TEMPLATES                 = "inittmplts"
	TUK_TASK_INIT_XDWS                      = "initxdws"
	TUK_TASK_SPA                            = "spa"
	TUK_TASK_GET                            = "get"
	TUK_TASK_SET                            = "set"
	TUK_TASK_GET_META                       = "getmeta"
	TUK_TASK_SET_META                       = "setmeta"
	TUK_TASK_GET_XDW                        = "getxdw"
	TUK_TASK_SET_XDW                        = "setxdw"
	TUK_TASK_GET_HTML                       = "gethtml"
	TUK_TASK_SET_HTML                       = "sethtml"
	TUK_TASK_GET_XML                        = "getxml"
	TUK_TASK_SET_XML                        = "setxml"
	TUK_STATUS_MET                          = "met"
	TUK_STATUS_MISSED                       = "missed"
	TUK_STATUS_ESCALATED                    = "escalated"
	TUK_STATUS_OPEN                         = "OPEN"
	TUK_STATUS_CLOSED                       = "CLOSED"
)
