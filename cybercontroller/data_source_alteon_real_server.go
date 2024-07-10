package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_real_server() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_real_server_read,
		Schema: map[string]*schema.Schema{
			"clustername": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the cluster",
			},
			"alteonip": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP Address of the Alteon managed by the cybercontroller",
			},
			"index": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The real server number.",
			},
			"ipaddr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of the real server identified by the instance of slbRealServerIndex.",
			},
			"weight": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The server weight.",
			},
			"maxconns": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum number of connections that are allowed.",
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum number of minutes an inactive connection remains open.",
			},
			"pinginterval": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The interval between keep-alive (ping) attempts in number of seconds. Zero means disabling ping attempt.",
			},
			"failretry": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of failed attempts to declare this server DOWN.",
			},
			"succretry": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of successful attempts to declare a server UP.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable the server and remove the existing sessions using disabled-with-fastage option.",
			},
			"deletestatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "By setting the value to delete(2), the entire row is deleted.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The server type. It participates in global server load balancing when it is configured as remote-server.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the real server.",
			},
			"addurl": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The URL Path (slbCurCfgUrlLbPathIndex) to be added to the real server. A zero is returned when read.",
			},
			"remurl": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The URL Path (slbCurCfgUrlLbPathIndex) to be removed from the real server. A zero is returned when read.",
			},
			"cookie": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable real server to handle client requests that don't contain a cookie if cookie loadbalance is enabled.",
			},
			"excludestr": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable exclusionary matching string on real server.",
			},
			"submac": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The real server config to enable/disable MAC SA substitution for L4 traffic.If disabled will not substitute the MAC SA of client-to-server frames, if enabled will substitute the MAC SA. ",
			},
			"idsport": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port to be connected to IDS server.",
			},
			"ipver": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The type of IP address.",
			},
			"ipv6addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPV6 address of the real server identified by the instance of the slbRealServerIndex.",
			},
			"nxtrportidx": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The next available free slot index number, to add the real port to the server. Value 0 will be returned if no free slot available.",
			},
			"nxtbuddyidx": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The next available free slot Buddy index number, to add the Buddy Server to the Real server. Value 0 will be returned if no free slot available.",
			},
			"llbtype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The server type.",
			},
			"copy": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The alphanumeric index of the new copy to be created.",
			},
			"portsingress": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of ingress ports attached to the real server (security device), used for SSL inspection WebUI wizard.",
			},
			"portsegress": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of egress ports attached to the real server (security device), used for SSL inspection WebUI wizard.",
			},
			"addportsingress": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ingress port to be added to the specified security device. A '0' value is returned when read.",
			},
			"remportsingress": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ingress port to be removed to the specified security device. A zero is returned when read.",
			},
			"addportsegress": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The egress port to be added to the specified security device. A zero is returned when read.",
			},
			"remportsegress": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The egress port to be removed to the specified security device. A zero is returned when read.",
			},
			"vlaningress": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ingress Vlan specified on security device. Used for SSL wizard",
			},
			"vlanegress": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The egress Vlan specified on security device. Used for SSL wizard",
			},
			"egressif": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The egress interface specified on security device. Used for SSL wizard",
			},
			"sectype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The security device type.",
			},
			"ingressif": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ingress interface specified on security device. Used for SSL wizard",
			},
			"secdeviceflag": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Real security device flag.",
			},
			"ingport": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ingress port to be connected to IDS server.",
			},
			"urlbmap": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The URL Paths selected for URL load balancing for by the real server. The selected URL Paths are presented in a bitmap format.",
			},
			"proxy": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The real server config to enable/disable client proxy operation.",
			},
			"ldapwr": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The real server config to enable/disable LDAP write server.",
			},
			"idsvlan": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The VLAN to be associated with IDS server.",
			},
			"avail": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The remote real server Global SLB availability.",
			},
			"fasthealthcheck": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The real server config to enable/disable Fast Health Check Operation.",
			},
			"subdmac": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The real server config to enable/disable MAC DA substitution for L4 traffic.If disabled will not substitute the MAC DA of client-to-server frames,if enabled will substitute the MAC DA.",
			},
			"overflow": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The real server config to enable/disable Overflow. If enabled(default) allows Backup server to kick in if real server reaches maximum connections.",
			},
			"bkppreempt": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The real server config to enable/disable backup preemption. If enabled (default)allows to preempt the backup server when the primary server comes up.",
			},
			"mode": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the mode of the real server. By default the mode is set to physical.",
			},
			"updateallrealservers": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set all the real servers having the same RIP address as this real server with the mode and/or maximum connection value (if mode is physical) that is set in this real server.",
			},
			"proxyipmode": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the real server Proxy IP mode.Changing from address(2) to any other mode will clear the configured IPv4/IPv6 address,prefix and persistancy.Changing from nwclass(3) to any other mode will clear the configured NWclass and NWpersistancy.",
			},
			"proxyipaddress": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allows configuration of real server proxy IP v4 address . The IP version for addr must be the same as the real server IP version.",
			},
			"proxyipmask": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allows configuration of real server proxy IP Mask. The IP version for addr must be the same as the real server IP version.",
			},
			"proxyipv6address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allows configuration of real server proxy IPv6 address . The IP version for addr must be the same as the real server IP version.Returns emply if IP version is IPv4 or slbNewCfgEnhRealServerProxyIpMode is not set to address. ",
			},
			"proxyipv6prefix": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Allows configuration of real server proxy IPv6 Mask. The IP version for addr must be the same as the real server IP version.",
			},
			"proxyippersistency": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "When a subnet is configured user has the ability to select PIP persistency mode.can be set only if slbNewCfgEnhRealServerProxyIpMode is address else return failure.If PIP is not configured the persistency configuration is disable.",
			},
			"proxyipnwclass": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allows configuration of real server proxy IP IPv4 or IPv6 Network Class as PIP. It can be set only if slbNewCfgEnhRealServerProxyIpMode is nwclss else return failure. ",
			},
			"proxyipnwclasspersistency": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Allows configuration of real server Network Class PIP persistency mode.It can be set only if slbNewCfgEnhRealServerProxyIpMode is nwclss else return failure.",
			},
			"ingvlan": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ingress VLAN to be associated with IDS server.",
			},
			"oid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The OID to be sent in the SNMP get request packet.",
			},
			"commstring": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The community string to be used in the SNMP get request packet.",
			},
			"backup": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The backup server number for this server.",
			},
			"healthid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Advanced HC ID.",
			},
			"criticalconnthrsh": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: " Critical connection threshold.",
			},
			"highconnthrsh": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "High connection threshold.",
			},
			"uploadbandwidth": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Upload bandwidth limit for WAN Link real server in Mbps.",
			},
			"downloadbandwidth": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Download bandwidth limit for WAN Link real server in Mbps.",
			},
		},
	}
}

func data_source_alteon_real_server_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)
	var diags diag.Diagnostics

	RealServerID := d.Get("index").(string)
	Tables := [3]string{"SlbNewCfgEnhRealServerTable",
		"SlbNewCfgEnhRealServerSecondPartTable",
		"SlbNewCfgEnhRealServerThirdPartTable",
	}
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	for _, Table := range Tables {
		if clustername != "" {
			api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + RealServerID + "/"
		} else if alteonip != "" {
			api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + RealServerID + "/"
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

		RealServer := make(map[string]interface{})
		json.Unmarshal([]byte(message), &RealServer)
		if len(RealServer) == 0 { //catch for response with 200 ok and not a json body
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST GetItem Failed With Response:" + "\n" + message,
				Detail:   "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		}
		//catch to handle GET response body with "status" err and "message": "err getting the data obj not found"
		Items1 := RealServer[Table]
		if Items1 == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST GetItem failed for non-existing real server index :" + "\n" + message,
				Detail:   detail + "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		}
		//Below fix for defect # AAE-930 - GET with empty response
		Items := RealServer[Table].([]interface{})
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
		type RealServerItem1 struct {
			IpAddr          string `json:"ipaddr"`
			Name            string `json:"name"`
			Weight          int    `json:"weight"`
			MaxConns        int    `json:"maxconns"`
			TimeOut         int    `json:"timeout"`
			PingInterval    int    `json:"pinginterval"`
			FailRetry       int    `json:"failretry"`
			SuccRetry       int    `json:"succretry"`
			State           int    `json:"state"`
			DeleteStatus    int    `json:"deletestatus"`
			NxtRportIdx     int    `json:"nxtrportidx"`
			NxtBuddyIdx     int    `json:"nxtbuddyidx"`
			LLBType         int    `json:"llbtype"`
			Copy            string `json:"copy"`
			PortsIngress    string `json:"portsingress"`
			PortsEgress     string `json:"portsegress"`
			AddPortsIngress int    `json:"addportsingress"`
			RemPortsIngress int    `json:"remportsingress"`
			AddPortsEgress  int    `json:"addportsegress"`
			RemPortsEgress  int    `json:"remportsegress"`
			Type            int    `json:"type"`
			AddUrl          int    `json:"addurl"`
			RemUrl          int    `json:"remurl"`
			Cookie          int    `json:"cookie"`
			ExcludeStr      int    `json:"excludestr"`
			Submac          int    `json:"submac"`
			Idsport         int    `json:"idsport"`
			IpVer           int    `json:"ipver"`
			Ipv6Addr        string `json:"ipv6addr"`
			VlanIngress     int    `json:"vlaningress"`
			VlanEgress      int    `json:"vlanegress"`
			EgressIf        int    `json:"egressif"`
			SecType         int    `json:"sectype"`
			IngressIf       int    `json:"ingressif"`
			SecDeviceFlag   int    `json:"secdeviceflag"`
			Ingport         int    `json:"ingport"`
		}

		type RealServerItem2 struct {
			UrlBmap                   string `json:"urlbmap"`
			Proxy                     int    `json:"proxy"`
			Ldapwr                    int    `json:"ldapwr"`
			Idsvlan                   int    `json:"idsvlan"`
			Avail                     int    `json:"avail"`
			FastHealthCheck           int    `json:"fasthealthcheck"`
			Subdmac                   int    `json:"subdmac"`
			Overflow                  int    `json:"overflow"`
			BkpPreempt                int    `json:"bkppreempt"`
			Mode                      int    `json:"mode"`
			UpdateAllRealServers      int    `json:"updateallrealservers"`
			ProxyIpMode               int    `json:"proxyipmode"`
			ProxyIpAddr               string `json:"proxyipaddr"`
			ProxyIpMask               string `json:"proxyipmask"`
			ProxyIpv6Addr             string `json:"proxyipv6addr"`
			ProxyIpv6Prefix           int    `json:"proxyipv6prefix"`
			ProxyIpPersistency        int    `json:"proxyippersistency"`
			ProxyIpNWclass            string `json:"proxyipnwclass"`
			ProxyIpNWclassPersistency int    `json:"proxyipnwclasspersistency"`
			Ingvlan                   int    `json:"ingvlan"`
		}

		type RealServerItem3 struct {
			Oid               string `json:"oid"`
			CommString        string `json:"commstring"`
			BackUp            string `json:"backup"`
			HealthID          string `json:"healthid"`
			Criticalconnthrsh int    `json:"criticalconnthrsh"`
			HighConnThrsh     int    `json:"highconnthrsh"`
			UploadBandWidth   int    `json:"uploadbandwidth"`
			DownloadBandWidth int    `json:"downloadbandwidth"`
		}

		var List_Item1 []RealServerItem1
		var List_Item2 []RealServerItem2
		var List_Item3 []RealServerItem3
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
		}

	}
	d.SetId("Resource GET for Real Server")
	return diags
}
