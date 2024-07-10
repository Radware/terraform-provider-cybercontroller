package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_virtual_server() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_virtual_server_read,
		Schema: map[string]*schema.Schema{
			"clustername": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the cluster.",
			},
			"alteonip": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP Address of the Alteon managed by the cybercontroller",
			},
			"index": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Virtual Server Number",
			},
			"virtserveripaddress": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of the virtual server.",
			},
			"virtserverstate": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable the virtual server.",
			},
			"virtserverlayer3only": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable layer3 only balancing.",
			},
			"virtserverdname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain name of the virtual server.",
			},
			"virtserverbwmcontract": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The default BW contract number of the virtual server.",
			},
			"virtserverdelete": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "By setting the value to delete(2), the entire row is deleted.",
			},
			"virtserverweight": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The virtual server Global SLB weight.",
			},
			"virtserveravail": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The virtual server Global SLB availability.",
			},
			"virtserverrule": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Global SLB rules for the domain. The rules are presented in bitmap format. in receiving order: OCTET 1 OCTET 2 ..... xxxxxxxx xxxxxxxx ..... | || |_ server 9 | || | ||___ server 8 | |____ server 7 | . . . |__________ server 1 where x : 1 - The represented rule belongs to the domain 0 - The represented rule does not belong to the domain",
			},
			"virtserveraddrule": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The rule to be added to the domain. When read, 0 is returned.",
			},
			"virtserverremoverule": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The rule to be removed from the domain. When read, 0 is returned.",
			},
			"virtservervname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the virtual server.",
			},
			"virtserveripver": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The type of IP address.",
			},
			"virtserveripv6addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv6 address of the virtual server. Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx",
			},
			"virtserverfreeserviceidx": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The first free service index number of the virtual server. Value 0 will be returned when all 8 virtual services are configured for a virtual server.",
			},
			"virtservercreset": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable client connection reset for invalid VPORT.",
			},
			"virtserversrcnetwork": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source Network Classifier of the virtual server.",
			},
			"virtservernat": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of the NAT.",
			},
			"virtservernat6": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv6 address of the NAT. Address should be 4-byte hexadecimal colon notation. Valid IPv6 address should be in any of the following forms xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx or xxxx::xxxx:xxxx:xxxx:xxxx or ::xxxx",
			},
			"virtserverisdnssecvip": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "This mib returns Yes(1) if virtual server is a DNS Responder VIP, else returns no(0)",
			},
			"virtserveravailpersist": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable GSLB availability persistence.",
			},
			"virtserverwanlink": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Associate a real wanlink server to virtual server.",
			},
			"virtserverrtsrcmac": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable/disable return to source mac address.",
			},
			"virtservercreationtype": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Virtual Server Creation Type.",
			},
		},
	}
}

func data_source_alteon_virtual_server_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)
	var diags diag.Diagnostics

	VirtualServerID := d.Get("index").(string)
	Table := "SlbNewCfgEnhVirtServerTable"

	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)

	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + VirtualServerID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + VirtualServerID + "/"
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

	VirtualServer := make(map[string]interface{})
	json.Unmarshal([]byte(message), &VirtualServer)
	if len(VirtualServer) == 0 { //catch for response with 200 ok and not a json body
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem Failed With Response:" + "\n" + message,
			Detail:   "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}

	//catch to handle GET response body with "status" err and "message": "err getting the data obj not found"
	Items1 := VirtualServer[Table]
	if Items1 == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem failed for non-existing virtual server index :" + "\n" + message,
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}
	//Below fix for defect # AAE-930 - GET with empty response
	Items := VirtualServer[Table].([]interface{})
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
	type VirtualServerItem struct {
		VirtServerIpAddress      string `json:"virtserveripaddress"`
		VirtServerLayer3Only     int    `json:"virtserverlayer3only"`
		VirtServerState          int    `json:"virtserverstate"`
		VirtServerDname          string `json:"virtserverdname"`
		VirtServerBwmContract    int    `json:"virtserverbwmcontract"`
		VirtServerDelete         int    `json:"virtserverdelete"`
		VirtServerWeight         int    `json:"virtserverweight"`
		VirtServerAvail          int    `json:"virtserveravail"`
		VirtServerRule           string `json:"virtserverrule"`
		VirtServerAddRule        int    `json:"virtserveraddrule"`
		VirtServerRemoveRule     int    `json:"virtserverremoverule"`
		VirtServerVname          string `json:"virtservervname"`
		VirtServerIpVer          int    `json:"virtserveripver"`
		VirtServerIpv6Addr       string `json:"virtserveripv6addr"`
		VirtServerFreeServiceIdx int    `json:"virtserverfreeserviceidx"`
		VirtServerCReset         int    `json:"virtservercreset"`
		VirtServerSrcNetwork     string `json:"virtserversrcnetwork"`
		VirtServerNat            string `json:"virtservernat"`
		VirtServerNat6           string `json:"virtservernat6"`
		VirtServerIsDnsSecVip    int    `json:"virtserverisdnssecvip"`
		VirtServerAvailPersist   int    `json:"virtserveravailpersist"`
		VirtServerWanlink        string `json:"virtserverwanlink"`
		VirtServerRtSrcMac       int    `json:"virtserverrtsrcmac"`
		VirtServerCreationType   int    `json:"virtservercreationtype"`
	}

	var List_Item []VirtualServerItem
	json.Unmarshal([]byte(helper), &List_Item)
	data1 := make([]map[string]interface{}, len(List_Item))
	rss, _ := json.Marshal(List_Item)
	json.Unmarshal(rss, &data1)
	for key, _ := range data1[0] {
		d.Set(key, data1[0][key])
	}
	d.SetId("Resource GET for Virtual Server")
	return diags
}
