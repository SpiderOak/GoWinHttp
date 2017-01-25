// +build windows

// Implements flags necessary for working with WinHTTP.

package winhttp

import (
	"syscall"
)

const (
	/* WinHTTP error codes */
	WINHTTP_ERROR_BASE                                  = 12000
	ERROR_WINHTTP_OUT_OF_HANDLES                        = syscall.Errno(WINHTTP_ERROR_BASE + 1)
	ERROR_WINHTTP_TIMEOUT                               = syscall.Errno(WINHTTP_ERROR_BASE + 2)
	ERROR_WINHTTP_INTERNAL_ERROR                        = syscall.Errno(WINHTTP_ERROR_BASE + 4)
	ERROR_WINHTTP_INVALID_URL                           = syscall.Errno(WINHTTP_ERROR_BASE + 5)
	ERROR_WINHTTP_UNRECOGNIZED_SCHEME                   = syscall.Errno(WINHTTP_ERROR_BASE + 6)
	ERROR_WINHTTP_NAME_NOT_RESOLVED                     = syscall.Errno(WINHTTP_ERROR_BASE + 7)
	ERROR_WINHTTP_INVALID_OPTION                        = syscall.Errno(WINHTTP_ERROR_BASE + 9)
	ERROR_WINHTTP_OPTION_NOT_SETTABLE                   = syscall.Errno(WINHTTP_ERROR_BASE + 11)
	ERROR_WINHTTP_SHUTDOWN                              = syscall.Errno(WINHTTP_ERROR_BASE + 12)
	ERROR_WINHTTP_LOGIN_FAILURE                         = syscall.Errno(WINHTTP_ERROR_BASE + 15)
	ERROR_WINHTTP_OPERATION_CANCELLED                   = syscall.Errno(WINHTTP_ERROR_BASE + 17)
	ERROR_WINHTTP_INCORRECT_HANDLE_TYPE                 = syscall.Errno(WINHTTP_ERROR_BASE + 18)
	ERROR_WINHTTP_INCORRECT_HANDLE_STATE                = syscall.Errno(WINHTTP_ERROR_BASE + 19)
	ERROR_WINHTTP_CANNOT_CONNECT                        = syscall.Errno(WINHTTP_ERROR_BASE + 29)
	ERROR_WINHTTP_CONNECTION_ERROR                      = syscall.Errno(WINHTTP_ERROR_BASE + 30)
	ERROR_WINHTTP_RESEND_REQUEST                        = syscall.Errno(WINHTTP_ERROR_BASE + 32)
	ERROR_WINHTTP_SECURE_CERT_DATE_INVALID              = syscall.Errno(WINHTTP_ERROR_BASE + 37)
	ERROR_WINHTTP_SECURE_CERT_CN_INVALID                = syscall.Errno(WINHTTP_ERROR_BASE + 38)
	ERROR_WINHTTP_CLIENT_AUTH_CERT_NEEDED               = syscall.Errno(WINHTTP_ERROR_BASE + 44)
	ERROR_WINHTTP_SECURE_INVALID_CA                     = syscall.Errno(WINHTTP_ERROR_BASE + 45)
	ERROR_WINHTTP_SECURE_CERT_REV_FAILED                = syscall.Errno(WINHTTP_ERROR_BASE + 57)
	ERROR_WINHTTP_CANNOT_CALL_BEFORE_OPEN               = syscall.Errno(WINHTTP_ERROR_BASE + 100)
	ERROR_WINHTTP_CANNOT_CALL_BEFORE_SEND               = syscall.Errno(WINHTTP_ERROR_BASE + 101)
	ERROR_WINHTTP_CANNOT_CALL_AFTER_SEND                = syscall.Errno(WINHTTP_ERROR_BASE + 102)
	ERROR_WINHTTP_CANNOT_CALL_AFTER_OPEN                = syscall.Errno(WINHTTP_ERROR_BASE + 103)
	ERROR_WINHTTP_HEADER_NOT_FOUND                      = syscall.Errno(WINHTTP_ERROR_BASE + 150)
	ERROR_WINHTTP_INVALID_SERVER_RESPONSE               = syscall.Errno(WINHTTP_ERROR_BASE + 152)
	ERROR_WINHTTP_INVALID_HEADER                        = syscall.Errno(WINHTTP_ERROR_BASE + 153)
	ERROR_WINHTTP_INVALID_QUERY_REQUEST                 = syscall.Errno(WINHTTP_ERROR_BASE + 154)
	ERROR_WINHTTP_HEADER_ALREADY_EXISTS                 = syscall.Errno(WINHTTP_ERROR_BASE + 155)
	ERROR_WINHTTP_REDIRECT_FAILED                       = syscall.Errno(WINHTTP_ERROR_BASE + 156)
	ERROR_WINHTTP_SECURE_CHANNEL_ERROR                  = syscall.Errno(WINHTTP_ERROR_BASE + 157)
	ERROR_WINHTTP_BAD_AUTO_PROXY_SCRIPT                 = syscall.Errno(WINHTTP_ERROR_BASE + 166)
	ERROR_WINHTTP_UNABLE_TO_DOWNLOAD_SCRIPT             = syscall.Errno(WINHTTP_ERROR_BASE + 167)
	ERROR_WINHTTP_SECURE_INVALID_CERT                   = syscall.Errno(WINHTTP_ERROR_BASE + 169)
	ERROR_WINHTTP_SECURE_CERT_REVOKED                   = syscall.Errno(WINHTTP_ERROR_BASE + 170)
	ERROR_WINHTTP_NOT_INITIALIZED                       = syscall.Errno(WINHTTP_ERROR_BASE + 172)
	ERROR_WINHTTP_SECURE_FAILURE                        = syscall.Errno(WINHTTP_ERROR_BASE + 175)
	ERROR_WINHTTP_AUTO_PROXY_SERVICE_ERROR              = syscall.Errno(WINHTTP_ERROR_BASE + 178)
	ERROR_WINHTTP_SECURE_CERT_WRONG_USAGE               = syscall.Errno(WINHTTP_ERROR_BASE + 179)
	ERROR_WINHTTP_AUTODETECTION_FAILED                  = syscall.Errno(WINHTTP_ERROR_BASE + 180)
	ERROR_WINHTTP_HEADER_COUNT_EXCEEDED                 = syscall.Errno(WINHTTP_ERROR_BASE + 181)
	ERROR_WINHTTP_HEADER_SIZE_OVERFLOW                  = syscall.Errno(WINHTTP_ERROR_BASE + 182)
	ERROR_WINHTTP_CHUNKED_ENCODING_HEADER_SIZE_OVERFLOW = syscall.Errno(WINHTTP_ERROR_BASE + 183)
	ERROR_WINHTTP_RESPONSE_DRAIN_OVERFLOW               = syscall.Errno(WINHTTP_ERROR_BASE + 184)
	ERROR_WINHTTP_CLIENT_CERT_NO_PRIVATE_KEY            = syscall.Errno(WINHTTP_ERROR_BASE + 185)
	ERROR_WINHTTP_CLIENT_CERT_NO_ACCESS_PRIVATE_KEY     = syscall.Errno(WINHTTP_ERROR_BASE + 186)
	WINHTTP_ERROR_LAST                                  = (WINHTTP_ERROR_BASE + 186)
)

const (
	SECURITY_FLAG_IGNORE_UNKNOWN_CA        = 0x00000100
	SECURITY_FLAG_IGNORE_CERT_DATE_INVALID = 0x00002000
	SECURITY_FLAG_IGNORE_CERT_CN_INVALID   = 0x00001000
	SECURITY_FLAG_IGNORE_CERT_WRONG_USAGE  = 0x00000200
	SECURITY_FLAG_SECURE                   = 0x00000001
	SECURITY_FLAG_STRENGTH_WEAK            = 0x10000000
	SECURITY_FLAG_STRENGTH_MEDIUM          = 0x40000000
	SECURITY_FLAG_STRENGTH_STRONG          = 0x20000000
)

const (
	INTERNET_DEFAULT_PORT       = 0
	INTERNET_DEFAULT_HTTP_PORT  = 80
	INTERNET_DEFAULT_HTTPS_PORT = 443

	WINHTTP_FLAG_ASYNC = 0x10000000

	WINHTTP_FLAG_SECURE               = 0x00800000
	WINHTTP_FLAG_ESCAPE_PERCENT       = 0x00000004
	WINHTTP_FLAG_NULL_CODEPAGE        = 0x00000008
	WINHTTP_FLAG_BYPASS_PROXY_CACHE   = 0x00000100
	WINHTTP_FLAG_REFRESH              = 0x00000100
	WINHTTP_FLAG_ESCAPE_DISABLE       = 0x00000040
	WINHTTP_FLAG_ESCAPE_DISABLE_QUERY = 0x00000080

	WINHTTP_AUTOPROXY_AUTO_DETECT         = 0x00000001
	WINHTTP_AUTOPROXY_CONFIG_URL          = 0x00000002
	WINHTTP_AUTOPROXY_HOST_KEEPCASE       = 0x00000004
	WINHTTP_AUTOPROXY_HOST_LOWERCASE      = 0x00000008
	WINHTTP_AUTOPROXY_RUN_INPROCESS       = 0x00010000
	WINHTTP_AUTOPROXY_RUN_OUTPROCESS_ONLY = 0x00020000
	WINHTTP_AUTOPROXY_NO_DIRECTACCESS     = 0x00040000
	WINHTTP_AUTOPROXY_NO_CACHE_CLIENT     = 0x00080000
	WINHTTP_AUTOPROXY_NO_CACHE_SVC        = 0x00100000

	WINHTTP_AUTOPROXY_SORT_RESULTS = 0x00400000

	WINHTTP_AUTO_DETECT_TYPE_DHCP  = 0x00000001
	WINHTTP_AUTO_DETECT_TYPE_DNS_A = 0x00000002

	WINHTTP_ACCESS_TYPE_DEFAULT_PROXY   = 0
	WINHTTP_ACCESS_TYPE_NO_PROXY        = 1
	WINHTTP_ACCESS_TYPE_NAMED_PROXY     = 3
	WINHTTP_ACCESS_TYPE_AUTOMATIC_PROXY = 4

	WINHTTP_NO_PROXY_NAME   = 0
	WINHTTP_NO_PROXY_BYPASS = 0

	WINHTTP_NO_ADDITIONAL_HEADERS = 0
	WINHTTP_NO_REQUEST_DATA       = 0

	WINHTTP_QUERY_MIME_VERSION              = 0
	WINHTTP_QUERY_CONTENT_TYPE              = 1
	WINHTTP_QUERY_CONTENT_TRANSFER_ENCODING = 2
	WINHTTP_QUERY_CONTENT_ID                = 3
	WINHTTP_QUERY_CONTENT_DESCRIPTION       = 4
	WINHTTP_QUERY_CONTENT_LENGTH            = 5
	WINHTTP_QUERY_CONTENT_LANGUAGE          = 6
	WINHTTP_QUERY_ALLOW                     = 7
	WINHTTP_QUERY_PUBLIC                    = 8
	WINHTTP_QUERY_DATE                      = 9
	WINHTTP_QUERY_EXPIRES                   = 10
	WINHTTP_QUERY_LAST_MODIFIED             = 11
	WINHTTP_QUERY_MESSAGE_ID                = 12
	WINHTTP_QUERY_URI                       = 13
	WINHTTP_QUERY_DERIVED_FROM              = 14
	WINHTTP_QUERY_COST                      = 15
	WINHTTP_QUERY_LINK                      = 16
	WINHTTP_QUERY_PRAGMA                    = 17
	WINHTTP_QUERY_VERSION                   = 18 // special: part of status line
	WINHTTP_QUERY_STATUS_CODE               = 19 // special: part of status line
	WINHTTP_QUERY_STATUS_TEXT               = 20 // special: part of status line
	WINHTTP_QUERY_RAW_HEADERS               = 21 // special: all headers as ASCIIZ
	WINHTTP_QUERY_RAW_HEADERS_CRLF          = 22 // special: all headers
	WINHTTP_QUERY_CONNECTION                = 23
	WINHTTP_QUERY_ACCEPT                    = 24
	WINHTTP_QUERY_ACCEPT_CHARSET            = 25
	WINHTTP_QUERY_ACCEPT_ENCODING           = 26
	WINHTTP_QUERY_ACCEPT_LANGUAGE           = 27
	WINHTTP_QUERY_AUTHORIZATION             = 28
	WINHTTP_QUERY_CONTENT_ENCODING          = 29
	WINHTTP_QUERY_FORWARDED                 = 30
	WINHTTP_QUERY_FROM                      = 31
	WINHTTP_QUERY_IF_MODIFIED_SINCE         = 32
	WINHTTP_QUERY_LOCATION                  = 33
	WINHTTP_QUERY_ORIG_URI                  = 34
	WINHTTP_QUERY_REFERER                   = 35
	WINHTTP_QUERY_RETRY_AFTER               = 36
	WINHTTP_QUERY_SERVER                    = 37
	WINHTTP_QUERY_TITLE                     = 38
	WINHTTP_QUERY_USER_AGENT                = 39
	WINHTTP_QUERY_WWW_AUTHENTICATE          = 40
	WINHTTP_QUERY_PROXY_AUTHENTICATE        = 41
	WINHTTP_QUERY_ACCEPT_RANGES             = 42
	WINHTTP_QUERY_SET_COOKIE                = 43
	WINHTTP_QUERY_COOKIE                    = 44
	WINHTTP_QUERY_REQUEST_METHOD            = 45 // special: GET/POST etc.
	WINHTTP_QUERY_REFRESH                   = 46
	WINHTTP_QUERY_CONTENT_DISPOSITION       = 47

	//
	// HTTP 1.1 defined headers
	//

	WINHTTP_QUERY_AGE                   = 48
	WINHTTP_QUERY_CACHE_CONTROL         = 49
	WINHTTP_QUERY_CONTENT_BASE          = 50
	WINHTTP_QUERY_CONTENT_LOCATION      = 51
	WINHTTP_QUERY_CONTENT_MD5           = 52
	WINHTTP_QUERY_CONTENT_RANGE         = 53
	WINHTTP_QUERY_ETAG                  = 54
	WINHTTP_QUERY_HOST                  = 55
	WINHTTP_QUERY_IF_MATCH              = 56
	WINHTTP_QUERY_IF_NONE_MATCH         = 57
	WINHTTP_QUERY_IF_RANGE              = 58
	WINHTTP_QUERY_IF_UNMODIFIED_SINCE   = 59
	WINHTTP_QUERY_MAX_FORWARDS          = 60
	WINHTTP_QUERY_PROXY_AUTHORIZATION   = 61
	WINHTTP_QUERY_RANGE                 = 62
	WINHTTP_QUERY_TRANSFER_ENCODING     = 63
	WINHTTP_QUERY_UPGRADE               = 64
	WINHTTP_QUERY_VARY                  = 65
	WINHTTP_QUERY_VIA                   = 66
	WINHTTP_QUERY_WARNING               = 67
	WINHTTP_QUERY_EXPECT                = 68
	WINHTTP_QUERY_PROXY_CONNECTION      = 69
	WINHTTP_QUERY_UNLESS_MODIFIED_SINCE = 70

	WINHTTP_HEADER_NAME_BY_INDEX = ""
	WINHTTP_NO_OUTPUT_BUFFER     = 0
	WINHTTP_NO_HEADER_INDEX      = 0

	WINHTTP_OPTION_CALLBACK                      = 1
	WINHTTP_OPTION_RESOLVE_TIMEOUT               = 2
	WINHTTP_OPTION_CONNECT_TIMEOUT               = 3
	WINHTTP_OPTION_CONNECT_RETRIES               = 4
	WINHTTP_OPTION_SEND_TIMEOUT                  = 5
	WINHTTP_OPTION_RECEIVE_TIMEOUT               = 6
	WINHTTP_OPTION_RECEIVE_RESPONSE_TIMEOUT      = 7
	WINHTTP_OPTION_HANDLE_TYPE                   = 9
	WINHTTP_OPTION_READ_BUFFER_SIZE              = 12
	WINHTTP_OPTION_WRITE_BUFFER_SIZE             = 13
	WINHTTP_OPTION_PARENT_HANDLE                 = 21
	WINHTTP_OPTION_EXTENDED_ERROR                = 24
	WINHTTP_OPTION_SECURITY_FLAGS                = 31
	WINHTTP_OPTION_SECURITY_CERTIFICATE_STRUCT   = 32
	WINHTTP_OPTION_URL                           = 34
	WINHTTP_OPTION_SECURITY_KEY_BITNESS          = 36
	WINHTTP_OPTION_PROXY                         = 38
	WINHTTP_OPTION_USER_AGENT                    = 41
	WINHTTP_OPTION_CONTEXT_VALUE                 = 45
	WINHTTP_OPTION_CLIENT_CERT_CONTEXT           = 47
	WINHTTP_OPTION_REQUEST_PRIORITY              = 58
	WINHTTP_OPTION_HTTP_VERSION                  = 59
	WINHTTP_OPTION_DISABLE_FEATURE               = 63
	WINHTTP_OPTION_CODEPAGE                      = 68
	WINHTTP_OPTION_MAX_CONNS_PER_SERVER          = 73
	WINHTTP_OPTION_MAX_CONNS_PER_1_0_SERVER      = 74
	WINHTTP_OPTION_AUTOLOGON_POLICY              = 77
	WINHTTP_OPTION_SERVER_CERT_CONTEXT           = 78
	WINHTTP_OPTION_ENABLE_FEATURE                = 79
	WINHTTP_OPTION_WORKER_THREAD_COUNT           = 80
	WINHTTP_OPTION_PASSPORT_COBRANDING_TEXT      = 81
	WINHTTP_OPTION_PASSPORT_COBRANDING_URL       = 82
	WINHTTP_OPTION_CONFIGURE_PASSPORT_AUTH       = 83
	WINHTTP_OPTION_SECURE_PROTOCOLS              = 84
	WINHTTP_OPTION_ENABLETRACING                 = 85
	WINHTTP_OPTION_PASSPORT_SIGN_OUT             = 86
	WINHTTP_OPTION_PASSPORT_RETURN_URL           = 87
	WINHTTP_OPTION_REDIRECT_POLICY               = 88
	WINHTTP_OPTION_MAX_HTTP_AUTOMATIC_REDIRECTS  = 89
	WINHTTP_OPTION_MAX_HTTP_STATUS_CONTINUE      = 90
	WINHTTP_OPTION_MAX_RESPONSE_HEADER_SIZE      = 91
	WINHTTP_OPTION_MAX_RESPONSE_DRAIN_SIZE       = 92
	WINHTTP_OPTION_CONNECTION_INFO               = 93
	WINHTTP_OPTION_CLIENT_CERT_ISSUER_LIST       = 94
	WINHTTP_OPTION_SPN                           = 96
	WINHTTP_OPTION_GLOBAL_PROXY_CREDS            = 97
	WINHTTP_OPTION_GLOBAL_SERVER_CREDS           = 98
	WINHTTP_OPTION_UNLOAD_NOTIFY_EVENT           = 99
	WINHTTP_OPTION_REJECT_USERPWD_IN_URL         = 100
	WINHTTP_OPTION_USE_GLOBAL_SERVER_CREDENTIALS = 101
	WINHTTP_LAST_OPTION                          = WINHTTP_OPTION_USE_GLOBAL_SERVER_CREDENTIALS
	WINHTTP_OPTION_USERNAME                      = 0x1000
	WINHTTP_OPTION_PASSWORD                      = 0x1001
	WINHTTP_OPTION_PROXY_USERNAME                = 0x1002
	WINHTTP_OPTION_PROXY_PASSWORD                = 0x1003
	WINHTTP_ADDREQ_INDEX_MASK                    = 0x0000FFFF
	WINHTTP_ADDREQ_FLAGS_MASK                    = 0xFFFF0000
	WINHTTP_ADDREQ_FLAG_ADD_IF_NEW               = 0x10000000
	WINHTTP_ADDREQ_FLAG_ADD                      = 0x20000000
	WINHTTP_ADDREQ_FLAG_COALESCE_WITH_COMMA      = 0x40000000
	WINHTTP_ADDREQ_FLAG_COALESCE_WITH_SEMICOLON  = 0x01000000
	WINHTTP_ADDREQ_FLAG_COALESCE                 = WINHTTP_ADDREQ_FLAG_COALESCE_WITH_COMMA
	WINHTTP_ADDREQ_FLAG_REPLACE                  = 0x80000000

	WINHTTP_CALLBACK_STATUS_RESOLVING_NAME        = 0x00000001
	WINHTTP_CALLBACK_STATUS_NAME_RESOLVED         = 0x00000002
	WINHTTP_CALLBACK_STATUS_CONNECTING_TO_SERVER  = 0x00000004
	WINHTTP_CALLBACK_STATUS_CONNECTED_TO_SERVER   = 0x00000008
	WINHTTP_CALLBACK_STATUS_SENDING_REQUEST       = 0x00000010
	WINHTTP_CALLBACK_STATUS_REQUEST_SENT          = 0x00000020
	WINHTTP_CALLBACK_STATUS_RECEIVING_RESPONSE    = 0x00000040
	WINHTTP_CALLBACK_STATUS_RESPONSE_RECEIVED     = 0x00000080
	WINHTTP_CALLBACK_STATUS_CLOSING_CONNECTION    = 0x00000100
	WINHTTP_CALLBACK_STATUS_CONNECTION_CLOSED     = 0x00000200
	WINHTTP_CALLBACK_STATUS_HANDLE_CREATED        = 0x00000400
	WINHTTP_CALLBACK_STATUS_HANDLE_CLOSING        = 0x00000800
	WINHTTP_CALLBACK_STATUS_DETECTING_PROXY       = 0x00001000
	WINHTTP_CALLBACK_STATUS_REDIRECT              = 0x00004000
	WINHTTP_CALLBACK_STATUS_INTERMEDIATE_RESPONSE = 0x00008000
	WINHTTP_CALLBACK_STATUS_SECURE_FAILURE        = 0x00010000
	WINHTTP_CALLBACK_STATUS_HEADERS_AVAILABLE     = 0x00020000
	WINHTTP_CALLBACK_STATUS_DATA_AVAILABLE        = 0x00040000
	WINHTTP_CALLBACK_STATUS_READ_COMPLETE         = 0x00080000
	WINHTTP_CALLBACK_STATUS_WRITE_COMPLETE        = 0x00100000
	WINHTTP_CALLBACK_STATUS_REQUEST_ERROR         = 0x00200000
	WINHTTP_CALLBACK_STATUS_SENDREQUEST_COMPLETE  = 0x00400000
	WINHTTP_CALLBACK_FLAG_RESOLVE_NAME            = WINHTTP_CALLBACK_STATUS_RESOLVING_NAME | WINHTTP_CALLBACK_STATUS_NAME_RESOLVED
	WINHTTP_CALLBACK_FLAG_CONNECT_TO_SERVER       = WINHTTP_CALLBACK_STATUS_CONNECTING_TO_SERVER | WINHTTP_CALLBACK_STATUS_CONNECTED_TO_SERVER
)

const (
	WINHTTP_CALLBACK_STATUS_FLAG_CERT_REV_FAILED        = 0x00000001
	WINHTTP_CALLBACK_STATUS_FLAG_INVALID_CERT           = 0x00000002
	WINHTTP_CALLBACK_STATUS_FLAG_CERT_REVOKED           = 0x00000004
	WINHTTP_CALLBACK_STATUS_FLAG_INVALID_CA             = 0x00000008
	WINHTTP_CALLBACK_STATUS_FLAG_CERT_CN_INVALID        = 0x00000010
	WINHTTP_CALLBACK_STATUS_FLAG_CERT_DATE_INVALID      = 0x00000020
	WINHTTP_CALLBACK_STATUS_FLAG_CERT_WRONG_USAGE       = 0x00000040
	WINHTTP_CALLBACK_STATUS_FLAG_SECURITY_CHANNEL_ERROR = 0x80000000
)

// Constants we're missing from the Windows syscall package.
const (
	ERROR_INVALID_PARAMETER syscall.Errno = 0x00000057
)
