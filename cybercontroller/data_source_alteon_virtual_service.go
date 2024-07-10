package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_virtual_service() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_virtual_service_read,
		Schema: map[string]*schema.Schema{
			"clustername": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the cluster.",
			},
			"alteonip": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP Address of the Alteon managed by the cybercontroller.",
			},
			"servindex": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The number of the virtual server.",
			},
			"index": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The service index. This has no external meaning",
			},
			"virtport": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The layer4 virtual port number of the service. it can be either 1 for ip or between 9 to 65534, virt port no. 2 to 9 are invalid",
			},
			"realport": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The layer4 real port number of the service, it can be either 0 for multiple real ports or 1 for ip service or between 5 to 65534. (2 to 5 are invalid)",
			},
			"udpbalance": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set protocol for the virtual service to UDP or TCP or SCTP or tcpAndUdp or stateless. tcpAndUdp is applicable only to ip service.",
			},
			"bwmcontract": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The BWM contract number for this service.",
			},
			"dirserverrtn": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable direct server return feature. To translate only MAC addresses in performing server load balancing when enabled. This allow servers to return directly to client since IP addresses have not been changed.",
			},
			"rtspurlparse": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Select RTSP URL load balancing type.",
			},
			"dbind": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable/forceproxy delayed binding.",
			},
			"ftpparsing": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or Disable the ftp parsing for the virtual service.",
			},
			"remapudpfrags": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable remapping UDP server fragments",
			},
			"dnsslb": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable DNS query load balancing.",
			},
			"responsecount": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of cookie search response count.",
			},
			"pbind": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable persistent bindings for the virtual port.",
			},
			"coffset": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The starting byte offset of the cookie value.",
			},
			"clength": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of bytes to extract from the cookie value.",
			},
			"uricookie": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable cookie search in URI",
			},
			"cookiemode": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Select cookie persistence mode. Mode disabled(4) not supported on Alteon",
			},
			"httpslb": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Select HTTP server loadbalancing for the virtual port.",
			},
			"httpslboption": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Select HTTP server loadbalancing for the virtual port.",
			},
			"httpslb2": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Select HTTP server loadbalancing for the virtual port.",
			},
			"deletestatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "When set to the value of 2 (delete), the entire row is deleted. When read, other(1) is returned. Setting the value to anything other than 2(delete) has no effect on the state of the row.Apm - Enable/disable apm.",
			},
			"apm": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable apm.",
			},
			"nonhttp": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/Disable to send non-HTTP traffic.",
			},
			"iprep": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/Disable IP reputation.",
			},
			"cdnproxy": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/Disable CDN proxy.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The service status: up, down, admin down, warning, shutdown, error.",
			},
			"rtsrctnl": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/Disable Return to Source Tunnel.",
			},
			"sideband": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set Sideband policy.",
			},
			"connmgtstatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Connection management configuration for HTTP traffic(Enable/disable/pooling) [Default: Disable].",
			},
			"connmgttimeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Connection management server side connection idle timeout in minutes [0-32768] [Default: 10].",
			},
			"cachepol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cache policy name associated with this virtual service.Set none to delete entry",
			},
			"comppol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Compression policy name associated with this virtual service.Set none to delete entry",
			},
			"sslpol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SSL policy name associated with this virtual service.Set none to delete entry",
			},
			"servcert": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server Certificate name associated with this virtual service.",
			},
			"httpmodlist": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTTP Content Modifications Rule-list associated with this virtual service.Set none to delete entry",
			},
			"cloaksrv": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable server cloaking.",
			},
			"serverrcodestatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable/clear error-code configuration.",
			},
			"serverrcodematch": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Match error-code(s), e.g 203,204 .",
			},
			"serverrcodehttpredir": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Use http redirection [yes/no] [Default: yes].",
			},
			"serverrcodeurl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL for redirection.",
			},
			"serverrcode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "set error code [Default: 302].",
			},
			"serverrcodenew": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter new error code [Default: 302].",
			},
			"serverrcodereason": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter error reason.",
			},
			"servurlchangstatus": &schema.Schema{ //required for third table elements config
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter enabled/disabled/clear [Default: clear].",
			},
			"servurlchanghosttype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter hostname match type [sufx|prefx|eq|incl|any] [Default: any]",
			},
			"fetcppolid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Frontend TCP optimization policy.",
			},
			"betcppolid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Backend TCP optimization policy.",
			},
			"basicconnmgtstatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Connection management configuration for Tcp traffic(Enable/disable) [Default: Disable].",
			},
			"servcertenc": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "GM SSL Server encryption Certificate name associated with this virtual service.",
			},
			"servcertsign": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "GM SSL Server sign Certificate name associated with this virtual service.",
			},
			"servurlchanghostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter hostname to match.",
			},
			"servurlchangpathtype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter path match type [sufx|prefx|eq|incl|any|none].",
			},
			"servurlchangpathmatch": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter path to match.",
			},
			"servurlchangpagename": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter page name to match or none.",
			},
			"servurlchangpagetype": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter page type to match or none.",
			},
			"servurlchangactntype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter path action type.",
			},
			"servurlchangpathinsrt": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: " Enter path to insert.",
			},
			"servurlchanginsrtpostn": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Insert the specified path before or after the matched section",
			},
			"servurlchangnewpgname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter new page name or none.",
			},
			"servurlchangnewpgtype": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter new page type or none.",
			},
			"servpathhidestatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter enabled/disabled/clear [Default: clear].",
			},
			"servpathhidehosttype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter hostname type [sufx|prefx|eq|incl|any] [Default: any].",
			},
			"servpathhidehostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter hostname to match.",
			},
			"servpathhidepathtype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter path match type [sufx|prefx|eq|none].",
			},
			"servpathhidepathname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter path to remove.",
			},
			"servtextrepstatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter enabled/disabled/clear [Default: clear].",
			},
			"servtextrepaction": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enter action [replace|remove|none].",
			},
			"servtextrepmatchtext": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter text to be replaced.",
			},
			"servtextrepreplactxt": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enter new text.",
			},
			"servapplicationtype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Application Type for virtual service.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the virtual service.",
			},
			"action": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Action type of the service.If we set value as group(1) it will Load balance the traffic between the servers defined in the group field after performing all other services actions.when set to a value of redirect(2) for http/https services, an http/s redirection will be performed with the values set in the application redirection.If we set value as discard(3) it will drop the session.",
			},
			"redirect": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Application redirection location.We need to provide this value When action type is set to redirect(2).",
			},
			"servcertgrpmark": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Group Mark for Server Certificate. If we set value as cert(0) It will denote that the server certificate (name) associated with this virtual service, represents a certificate. Otherwise, a value of group(1), denotes that the server certificate (name) represents a group.",
			},
			"dnstype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set DNS type for this service (DNS, DNSSEC).",
			},
			"clntproxtype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set client proximity type for this service.",
			},
			"zerowinsize": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable zero window size in SYN+ACK for this service.",
			},
			"cookiepath": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cookie path name of the virtual server used for cookie load balance.",
			},
			"cookiesecure": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Is cookie secure [yes/no] [Default: no].",
			},
			"nortsp": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable only rtsp SLB for this service.",
			},
			"ckrebind": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable server rebalancing when cookie is absent. When enabled, server load balancing will happen for subsequent request comes without cookie.",
			},
			"parselimit": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable buffer limit for content based selection.",
			},
			"parselength": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The buffer length for content based selection.",
			},
			"urinorm": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable URI normalization for HTTP content matching.",
			},
			"granularity": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: " Sets the Granularity for this service, for statistics report-protocol information. Group(1) - for group level, or GroupAndServers(2) - for server level.",
			},
			"sesslog": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable Session Logging.",
			},
			"udpage": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Fast aging of UDP sessions.",
			},
			"sessentrymode": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Session entry mode.",
			},
			"secpol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Security policy name associated with this virtual service. Set none to delete entry",
			},
			"alwayson": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "service always on when AS++ script attached.",
			},
			"sendrst": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/Disable sending reset/icmp-err when the service is down.",
			},
			"clsonslowage": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set close connection on aging treatment.",
			},
			"cookiesamesite": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The cookie samesite attribute of the virtual server used for cookie load balance.",
			},
			"clsonfastage": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set close connection on fastaging treatment.",
			},
			"cookiehttponly": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Is cookie http only [yes/no] [Default: no].",
			},
			"hname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host name of the virtual service.",
			},
			"cname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cookie name of the virtual server used for cookie load balance.",
			},
			"cexpire": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cookie expire of the virtual server used for insert cookie load balance depending on the mode it has the following format absolute mode or for relative mode.",
			},
			"urlhashlen": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of bytes used to hash onto server, A zero means URL hashing disabled.",
			},
			"dummydelete": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "This mib is dummy,the main Delete mib is in slbNewCfgEnhVirtServicesTable When read, other(1) is returned.",
			},
			"direct": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable DAM for this service.",
			},
			"thash": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set hash parameter.",
			},
			"ldapreset": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable LDAP Server Reset",
			},
			"ldapslb": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable LDAP Server load balancing",
			},
			"sip": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable SIP load balancing.",
			},
			"xforwardedfor": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable X-Forwarded-For for proxy mode.",
			},
			"httpredir": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable HTTP/HTTPS redirect for GSLB.",
			},
			"pbindrport": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable use of rport in the session lookup for a persistent session.",
			},
			"egresspip": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: " Enable/disable pip selection based on egress port/vlan.",
			},
			"cookiedname": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Select dname for insert cookie persistence mode.",
			},
			"wts": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable WTS loadbalancing and persistence.",
			},
			"uhash": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable when there is no Session Directory server.",
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum number of minutes an inactive connection remains open.",
			},
			"sdpnat": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable SIP Media portal NAT.",
			},
			"ipheader": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set ip address header.",
			},
			"userdefinedipheader": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set ip address header set by the user.",
			},
			"realgroup": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The real server group number for this service.",
			},
			"sessionmirror": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable session mirroring.",
			},
			"softgrid": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable softgrid load balancing.",
			},
			"connpooling": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable connection pooling for HTTP traffic.",
			},
			"persistenttimeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum number of minutes a persistent session should exist.",
			},
			"proxyipmode": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the Proxy IP mode, default is ingress(0). Changing from address(2) to any other mode will clear the configured IPv4/IPv6 address,prefix and persistancy. Changing from nwclass(3) to any other mode will clear the configured NWclass and NWpersistancy.",
			},
			"proxyipaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxyipmask": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This object ID allows configuring IPv4 PIP Mask. Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address. When a subnet is configured user has the ability to select PIP persistency mode. Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as subnet. If neither of the addresses (v4 & v6) are configured or are subnets, the persistency value is disable.",
			},
			"proxyipv6addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This object ID allows configuring IPv6 PIP address. Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address. When a subnet is configured user has the ability to select PIP persistency mode. Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as subnet. Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx. If neither of the addresses (v4 & v6) are configured or are subnets, the persistency value is disable",
			},
			"proxyipv6prefix": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "This object ID allows configuring IPv6 PIP Mask. Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address. When a subnet is configured user has the ability to select PIP persistency mode. Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as subnet. If neither of the addresses (v4 & v6) are configured or are subnets, the persistency value is disable.",
			},
			"proxyippersistency": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "When a subnet is configured user has the ability to select PIP persistency mode. Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to address. Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as subnet. If neither of the addresses (v4 & v6) are configured or are subnets, the persistency value is disable.",
			},
			"proxyipnwsclass": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This object ID allows configuring IPv4 Network Class as PIP and PIP persistency mode. Returns empty string when slbNewCfgVirtServiceProxyIpMode is not set to nwclass. Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured. If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.",
			},
			"proxyipv6nwclass": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This object ID allows configuring IPv6 Network Class as PIP and PIP persistency mode. Returns empty string when slbNewCfgVirtServiceProxyIpMode is not set to nwclass. Persistency is relevant only if either IPv4 or IPv6 class (or both) are configured. If neither of the classes (v4 & v6) are configured, the persistency configuration value is disable.",
			},
			"proxyipnwclasspersistency": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "This object ID allows configuring Network Class PIP persistency mode. Returns 0 when slbNewCfgVirtServiceProxyIpMode is not set to nwclass. Persistency is relevant only if either IPv4 PIP or IPv6 PIP (or both) are configured as subnet. If neither of the addresses (v4 & v6) are configured or are subnets, the persistency value is disable.",
			},
			"hashlen": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set length for slb service sip hashing (4- 256 bytes)",
			},
			"clsrst": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable Send RST on connection close.",
			},
			"httphdrname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The HTTP header name of the virtual server.",
			},
			"servfastwa": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Fastview web application name associated with this virtual service.Set none to delete entry",
			},
			"appwallwebappid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This object ID allows configuring the web security ID",
			},
			"http2": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Http2 policy name associated with this virtual service.",
			},
			"cluster": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable Cluster Updates for the service.",
			},
			"dataport": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ftp control service data port",
			},
			"applicname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Correlate several services into one application at the visualization.",
			},
			"report": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable counter based reporting for service.",
			},
			"trevpol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set Traffic Event Log Policy.",
			},
			"satisrt": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Application satisfied response time threshold, inherits the value from the global satisfied value or set with different value 1-999999 ms.",
			},
			"botpol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set Bot Manager Policy.",
			},
			"namesrvr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set DNS Nameserver group.",
			},
			"isdnssecvip": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "It returns Yes(1) if virtual service is configure auto with a DNS Responder VIP, else returns no(0).Http3 - Http3 policy name associated with this virtual service.",
			},
			"http3": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Http3 policy name associated with this virtual service.",
			},
			"quic": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Quic policy name associated with this virtual service.",
			},
			"awinflow": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set if AW processing comes before or after Alteon HTTP parsing.",
			},
			"fallbackuseaw": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable security web application processing when no content rule matches.",
			},
			"http3port": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set HTTP3 port for this virtual service",
			},
			"securepathpol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set SecurePath Policy.",
			},
			"jsinject": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set JS inject mode.",
			},
			"dohtype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set DOH mode.",
			},
			"awthresholdop": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set overload protection threshold",
			},
			"awmintransop": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set overload protection mintrans",
			},
			"reqheaderstimeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set request headers timeout in msec",
			},
		},
	}
}

func data_source_alteon_virtual_service_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)
	var diags diag.Diagnostics

	VirtualServerID := d.Get("servindex").(string)
	VirtualServiceID := strconv.Itoa(d.Get("index").(int))

	Tables := [7]string{"SlbNewCfgEnhVirtServicesTable",
		"SlbNewCfgEnhVirtServicesSecondPartTable",
		"SlbNewCfgEnhVirtServicesThirdPartTable",
		"SlbNewCfgEnhVirtServicesFourthPartTable",
		"SlbNewCfgEnhVirtServicesFifthPartTable",
		"SlbNewCfgEnhVirtServicesSixthPartTable",
		"SlbNewCfgEnhVirtServicesSeventhPartTable",
	}

	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	for _, Table := range Tables {
		if clustername != "" {
			api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + VirtualServerID + "/" + VirtualServiceID + "/"
		} else if alteonip != "" {
			api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + VirtualServerID + "/" + VirtualServiceID + "/"
		}

		status, message, err1 := client.GetItem(api, nil, nil)

		detail := "Status Code Received: " + strconv.Itoa(status) + "\nResponse Received: \n" + message

		if err1 != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST GetItem Failed With Error:" + err1.Error() + "\n",
				Detail:   detail + "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		} else if status != 200 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST GetItem Failed as response code received is: " + strconv.Itoa(status),
				Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
			})
			return diags
		}

		VirtualService := make(map[string]interface{})
		json.Unmarshal([]byte(message), &VirtualService)

		if len(VirtualService) == 0 { //catch for response with 200 ok and not a json body
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST GetItem Failed With Response:" + "\n" + message,
				Detail:   "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		}
		//catch to handle GET response body with "status" err and "message": "err getting the data obj not found"
		Items1 := VirtualService[Table]
		if Items1 == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST GetItem failed for non-existing virtual service index :" + "\n" + message,
				Detail:   detail + "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		}
		//Below fix for defect # AAE-930 - GET with empty response
		Items := VirtualService[Table].([]interface{})
		if len(Items) == 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "REST GetItem retrieved with an empty Response:" + "\n" + message,
				Detail:   "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		}

		helper, err := json.Marshal(Items)
		if err != nil {
			return diag.FromErr(err)
		}
		type VirtualServiceItem1 struct {
			VirtPort      int    `json:"virtport"`
			RealPort      int    `json:"realport"`
			UDPBalance    int    `json:"udpbalance"`
			BwmContract   int    `json:"bwmcontract"`
			DirServerRtn  int    `json:"dirserverrtn"`
			RtspUrlParse  int    `json:"rtspurlparse"`
			DBind         int    `json:"dbind"`
			FtpParsing    int    `json:"ftpparsing"`
			RemapUDPFrags int    `json:"remapudpfrags"`
			DnsSlb        int    `json:"dnsslb"`
			ResponseCount int    `json:"responsecount"`
			PBind         int    `json:"pbind"`
			Coffset       int    `json:"coffset"`
			Clength       int    `json:"clength"`
			UriCookie     int    `json:"uricookie"`
			CookieMode    int    `json:"cookiemode"`
			HttpSlb       int    `json:"httpslb"`
			HttpSlbOption int    `json:"httpslboption"`
			HttpSlb2      int    `json:"httpslb2"`
			DeleteStatus  int    `json:"deletestatus"`
			Apm           int    `json:"apm"`
			NonHTTP       int    `json:"nonhttp"`
			IpRep         int    `json:"iprep"`
			CdnProxy      int    `json:"cdnproxy"`
			Status        int    `json:"status"`
			RtSrcTnl      int    `json:"rtsrctnl"`
			Sideband      string `json:"sideband"`
		}

		type VirtualServiceItem2 struct {
			ConnmgtStatus        int    `json:"connmgtstatus"`
			ConnmgtTimeout       int    `json:"connmgttimeout"`
			Cachepol             string `json:"cachepol"`
			Comppol              string `json:"comppol"`
			SSLpol               string `json:"sslpol"`
			ServCert             string `json:"servcert"`
			HttpmodList          string `json:"httpmodlist"`
			Cloaksrv             int    `json:"cloaksrv"`
			ServErrcodeStatus    int    `json:"serverrcodestatus"`
			ServErrcodeMatch     string `json:"serverrcodematch"`
			ServErrcodeHttpRedir int    `json:"serverrcodehttpredir"`
			ServErrcodeUrl       string `json:"serverrcodeurl"`
			ServErrcode          string `json:"serverrcode"`
			ServErrcodeNew       string `json:"serverrcodenew"`
			ServErrcodeReason    string `json:"serverrcodereason"`
			ServUrlchangStatus   int    `json:"servurlchangstatus"`
			ServUrlchangHostType int    `json:"servurlchanghosttype"`
			FeTcpPolId           string `json:"fetcppolid"`
			BeTcpPolId           string `json:"betcppolid"`
			BasicConnmgtStatus   int    `json:"basicconnmgtstatus"`
			ServCertEnc          string `json:"servcertenc"`
			ServCertSign         string `json:"servcertsign"`
		}

		type VirtualServiceItem3 struct {
			ServUrlchangHostName   string `json:"servurlchanghostname"`
			ServUrlchangPathType   int    `json:"servurlchangpathtype"`
			ServUrlchangPathMatch  string `json:"servurlchangpathmatch"`
			ServUrlchangPageName   string `json:"servurlchangpagename"`
			ServUrlchangPageType   string `json:"servurlchangpagetype"`
			ServUrlchangActnType   int    `json:"servurlchangactntype"`
			ServUrlchangPathInsrt  string `json:"servurlchangpathinsrt"`
			ServUrlchangInsrtPostn int    `json:"servurlchanginsrtpostn"`
		}

		type VirtualServiceItem4 struct {
			ServUrlchangNewPgName string `json:"servurlchangnewpgname"`
			ServUrlchangNewPgType string `json:"servurlchangnewpgtype"`
			ServPathHideStatus    int    `json:"servpathhidestatus"`
			ServPathHideHostType  int    `json:"servpathhidehosttype"`
			ServPathHideHostName  string `json:"servpathhidehostname"`
			ServPathHidePathType  int    `json:"servpathhidepathtype"`
			ServPathHidePathName  string `json:"servpathhidepathname"`
			ServTextrepStatus     int    `json:"servtextrepstatus"`
			ServTextrepAction     int    `json:"servtextrepaction"`
		}

		type VirtualServiceItem5 struct {
			ServTextrepMatchText string `json:"servtextrepmatchtext"`
			ServTextrepReplacTxt string `json:"servtextrepreplactxt"`
			ServApplicationType  int    `json:"servapplicationtype"`
			Name                 string `json:"name"`
			Action               int    `json:"action"`
			Redirect             string `json:"redirect"`
			ServCertGrpMark      int    `json:"servcertgrpmark"`
			DnsType              int    `json:"dnstype"`
			ClntproxType         int    `json:"clntproxtype"`
			ZerowinSize          int    `json:"zerowinsize"`
			CookiePath           string `json:"cookiepath"`
			CookieSecure         int    `json:"cookiesecure"`
			NoRtsp               int    `json:"nortsp"`
			CkRebind             int    `json:"ckrebind"`
			ParseLimit           int    `json:"parselimit"`
			ParseLength          int    `json:"parselength"`
			UriNorm              int    `json:"urinorm"`
			Granularity          int    `json:"granularity"`
			SessLog              int    `json:"sesslog"`
			UdpAge               int    `json:"udpage"`
			SessEntryMode        int    `json:"sessentrymode"`
			SecPol               string `json:"secpol"`
			AlwaysOn             int    `json:"alwayson"`
			SendRST              int    `json:"sendrst"`
			ClsOnSlowage         int    `json:"clsonslowage"`
			CookieSameSite       int    `json:"cookiesamesite"`
			ClsOnFastage         int    `json:"clsonfastage"`
			CookieHttponly       int    `json:"cookiehttponly"`
		}

		type VirtualServiceItem6 struct {
			Hname               string `json:"hname"`
			Cname               string `json:"cname"`
			CExpire             string `json:"cexpire"`
			UrlHashLen          int    `json:"urlhashlen"`
			DummyDelete         int    `json:"dummydelete"`
			Direct              int    `json:"direct"`
			Thash               int    `json:"thash"`
			Ldapreset           int    `json:"ldapreset"`
			Ldapslb             int    `json:"ldapslb"`
			Sip                 int    `json:"sip"`
			XForwardedFor       int    `json:"xforwardedfor"`
			HttpRedir           int    `json:"httpredir"`
			PbindRport          int    `json:"pbindrport"`
			EgressPip           int    `json:"egresspip"`
			CookieDname         int    `json:"cookiedname"`
			Wts                 int    `json:"wts"`
			Uhash               int    `json:"uhash"`
			TimeOut             int    `json:"timeout"`
			SdpNat              int    `json:"sdpnat"`
			IpHeader            string `json:"ipheader"`
			UserDefinedIpHeader string `json:"userdefinedipheader"`
		}

		type VirtualServiceItem7 struct {
			RealGroup                 string `json:"realgroup"`
			SessionMirror             int    `json:"sessionmirror"`
			SoftGrid                  int    `json:"softgrid"`
			ConnPooling               int    `json:"connpooling"`
			PersistentTimeOut         int    `json:"persistenttimeout"`
			ProxyIpMode               int    `json:"proxyipmode"`
			ProxyIpAddr               string `json:"ProxyIpAddr"`
			ProxyIpMask               string `json:"proxyipmask"`
			ProxyIpv6Addr             string `json:"proxyipv6addr"`
			ProxyIpv6Prefix           int    `json:"proxyipv6prefix"`
			ProxyIpPersistency        int    `json:"proxyippersistency"`
			ProxyIpNWclass            string `json:"proxyipnwclass"`
			ProxyIpv6NWclass          string `json:"proxyipv6nwclass"`
			ProxyIpNWclassPersistency string `json:"proxyipnwclasspersistency"`
			HashLen                   int    `json:"hashlen"`
			ClsRST                    int    `json:"clsrst"`
			HttpHdrName               string `json:"httphdrname"`
			ServFastWa                string `json:"servfastwa"`
			AppwallWebappId           string `json:"appwallwebappid"`
			Http2                     string `json:"http2"`
			Cluster                   int    `json:"cluster"`
			DataPort                  int    `json:"dataport"`
			ApplicName                string `json:"applicname"`
			Report                    int    `json:"report"`
			Trevpol                   string `json:"trevpol"`
			Satisrt                   int    `json:"satisrt"`
			Botpol                    string `json:"botpol"`
			Namesrvr                  string `json:"namesrvr"`
			IsDnsSecVip               int    `json:"isdnssecvip"`
			Http3                     string `json:"http3"`
			Quic                      string `json:"quic"`
			Awinflow                  int    `json:"awinflow"`
			FallbackUseAW             int    `json:"fallbackuseaw"`
			Http3port                 int    `json:"http3port"`
			SecurePathpol             string `json:"securepathpol"`
			Jsinject                  int    `json:"jsinject"`
			DohType                   int    `json:"dohtype"`
			AwThresholdOp             int    `json:"awthresholdop"`
			AwMintransOp              int    `json:"awmintransop"`
			ReqHeadersTimeout         int    `json:"reqheaderstimeout"`
		}

		var List_Item1 []VirtualServiceItem1
		var List_Item2 []VirtualServiceItem2
		var List_Item3 []VirtualServiceItem3
		var List_Item4 []VirtualServiceItem4
		var List_Item5 []VirtualServiceItem5
		var List_Item6 []VirtualServiceItem6
		var List_Item7 []VirtualServiceItem7

		if Table == Tables[0] {
			json.Unmarshal([]byte(helper), &List_Item1)
			data1 := make([]map[string]interface{}, len(List_Item1))
			rss, _ := json.Marshal(List_Item1)
			json.Unmarshal(rss, &data1)
			for key, _ := range data1[0] {
				d.Set(key, data1[0][key])
			}
		} else if Table == Tables[1] {
			json.Unmarshal([]byte(helper), &List_Item2)
			data1 := make([]map[string]interface{}, len(List_Item2))
			rss, _ := json.Marshal(List_Item2)
			json.Unmarshal(rss, &data1)
			for key, _ := range data1[0] {
				d.Set(key, data1[0][key])
			}

		} else if Table == Tables[2] {
			json.Unmarshal([]byte(helper), &List_Item3)
			data1 := make([]map[string]interface{}, len(List_Item3))
			rss, _ := json.Marshal(List_Item3)
			json.Unmarshal(rss, &data1)
			for key, _ := range data1[0] {
				d.Set(key, data1[0][key])
			}
		} else if Table == Tables[3] {
			json.Unmarshal([]byte(helper), &List_Item4)
			data1 := make([]map[string]interface{}, len(List_Item4))
			rss, _ := json.Marshal(List_Item4)
			json.Unmarshal(rss, &data1)
			for key, _ := range data1[0] {
				d.Set(key, data1[0][key])
			}
		} else if Table == Tables[4] {
			json.Unmarshal([]byte(helper), &List_Item5)
			data1 := make([]map[string]interface{}, len(List_Item5))
			rss, _ := json.Marshal(List_Item5)
			json.Unmarshal(rss, &data1)
			for key, _ := range data1[0] {
				d.Set(key, data1[0][key])
			}
		} else if Table == Tables[5] {
			json.Unmarshal([]byte(helper), &List_Item6)
			data1 := make([]map[string]interface{}, len(List_Item6))
			rss, _ := json.Marshal(List_Item6)
			json.Unmarshal(rss, &data1)
			for key, _ := range data1[0] {
				d.Set(key, data1[0][key])
			}
		} else if Table == Tables[6] {
			json.Unmarshal([]byte(helper), &List_Item7)
			data1 := make([]map[string]interface{}, len(List_Item7))
			rss, _ := json.Marshal(List_Item7)
			json.Unmarshal(rss, &data1)
			for key, _ := range data1[0] {
				d.Set(key, data1[0][key])
			}
		}

	}
	d.SetId("Resource GET for Virtual Service")
	return diags
}
