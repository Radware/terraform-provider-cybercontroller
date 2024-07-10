package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resource_alteon_real_server() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_alteon_real_server_create,
		ReadContext:   resource_alteon_real_server_read,
		UpdateContext: resource_alteon_real_server_update,
		DeleteContext: resource_alteon_real_server_delete,
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
				Description: "The real server number",
			},
			"last_updated": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Resource last updated time.",
			},
			"elements": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipaddr": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
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
							Description: "List of egress ports attached to the real server (security device), used for SSL inspection WebUI wizard",
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
					},
				},
			},
			"elements_2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"elements_3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Description: "Critical connection threshold.",
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
				},
			},
			/*"elements_4": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"realport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"deletestatus": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"realportfreeidx": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},*/
		},
	}
}

func resource_alteon_real_server_create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	RealServerID := d.Get("index").(string)

	Tables := [3]string{"SlbNewCfgEnhRealServerTable",
		"SlbNewCfgEnhRealServerSecondPartTable",
		"SlbNewCfgEnhRealServerThirdPartTable",
		//"SlbNewCfgEnhRealServPortTable",
	}

	Elements := [3]string{"elements",
		"elements_2",
		"elements_3",
		//"elements_4",
	}

	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)

	for tabl_indx, Table := range Tables {

		if clustername != "" {
			api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + RealServerID + "/"
		} else if alteonip != "" {
			api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + RealServerID + "/"
		}
		/*if tabl_indx == 3 || tabl_indx == 4{
			api = api +
		}*/

		items := d.Get(Elements[tabl_indx]).([]interface{})
		validvals := make(map[string]interface{})

		for _, item := range items {

			i := item.(map[string]interface{})
			for key, conf_val := range i {
				if conf_val != "" && conf_val != 0 {
					validvals[key] = conf_val
				}
			}
		}

		brss, err := json.MarshalIndent(validvals, "", "    ")
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error encoding JSON : " + err.Error(),
				Detail:   err.Error(),
			})
			return diags
		}

		APIBytes := []byte(brss)
		var status int
		var message string

		if tabl_indx == 0 {

			status, message, err = client.CreateItem(api, APIBytes, nil)

		} else if len(APIBytes) == 2 { //empty map will take len of 2 due to braces
			continue
		} else {

			status, message, err = client.UpdateItem(api, APIBytes, nil)
		}

		resp_body := map[string]interface{}{}
		json.Unmarshal([]byte(message), &resp_body)

		detail := "Status Code Received: " + strconv.Itoa(status) + "\nResponse Received: \n" + message

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST CreateItem Failed With Error:" + err.Error() + "\n",
				Detail:   detail + "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		} else if status != 200 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST CreateItem Failed as response code received is: " + strconv.Itoa(status),
				Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
			})
			return diags
		} else if status == 200 && resp_body["status"].(string) != "ok" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST CreateItem Failed as error received in 200 OK Response : " + resp_body["status"].(string),
				Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
			})
			return diags
		} else {
			d.SetId("Resource Create for Real Server")
		}
	}

	resource_alteon_real_server_read(ctx, d, m)

	return diags
}

func resource_alteon_real_server_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	/*var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST ReadItem Failed",
		Detail:   "Read Item not supported for this resource type",
	})
	return diags*/
	return nil
}

func resource_alteon_real_server_update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	RealServerID := d.Get("index").(string)

	Tables := [3]string{"SlbNewCfgEnhRealServerTable",
		"SlbNewCfgEnhRealServerSecondPartTable",
		"SlbNewCfgEnhRealServerThirdPartTable",
	}

	Elements := [3]string{"elements",
		"elements_2",
		"elements_3",
	}

	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)

	for tabl_indx, Table := range Tables {

		if clustername != "" {
			api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + RealServerID + "/"
		} else if alteonip != "" {
			api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + RealServerID + "/"
		}

		items := d.Get(Elements[tabl_indx]).([]interface{})
		validvals := make(map[string]interface{})

		for _, item := range items {

			i := item.(map[string]interface{})
			for key, conf_val := range i {
				if conf_val != "" && conf_val != 0 {
					validvals[key] = conf_val
				}
			}
		}

		brss, err := json.MarshalIndent(validvals, "", "    ")
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error encoding JSON : " + err.Error(),
				Detail:   err.Error(),
			})
			return diags
		}

		APIBytes := []byte(brss)
		var status int
		var message string

		if len(APIBytes) == 2 { //empty map will take len of 2 due to braces
			continue
		} else {

			status, message, err = client.UpdateItem(api, APIBytes, nil)
		}

		//status, message, err = client.UpdateItem(api, APIBytes, nil)

		resp_body := map[string]interface{}{}
		json.Unmarshal([]byte(message), &resp_body)

		detail := "Status Code Received: " + strconv.Itoa(status) + "\nResponse Received: \n" + message

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST CreateItem Failed With Error:" + err.Error() + "\n",
				Detail:   detail + "\nAPI Call Made is:" + api + "\n",
			})
			return diags
		} else if status != 200 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST CreateItem Failed as response code received is: " + strconv.Itoa(status),
				Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
			})
			return diags
		} else if status == 200 && resp_body["status"].(string) != "ok" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "REST CreateItem Failed as error received in 200 OK Response : " + resp_body["status"].(string),
				Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
			})
			return diags
		} else {
			d.Set("last_updated", time.Now().Format(time.RFC3339))
		}
	}

	resource_alteon_real_server_read(ctx, d, m)

	return diags

	/*client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	RealServerID := d.Get("index").(string)
	Table := "SlbNewCfgEnhRealServerTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + RealServerID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + RealServerID + "/"
	}

	items := d.Get("items").([]interface{})
	validvals := make(map[string]interface{})

	for _, item := range items {

		i := item.(map[string]interface{})
		for key, conf_val := range i {
			if conf_val != "" && conf_val != 0 {
				validvals[key] = conf_val
			}
		}
	}

	brss, err := json.MarshalIndent(validvals, "", "    ")
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error encoding JSON : " + err.Error(),
			Detail:   err.Error(),
		})
		return diags
	}
	APIBytes := []byte(brss)
	status, message, err := client.UpdateItem(api, APIBytes, nil)

	resp_body := map[string]interface{}{}
	json.Unmarshal([]byte(message), &resp_body)

	detail := "Status Code Received: " + strconv.Itoa(status) + "\nResponse Received: \n" + message

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST UpdateItem Failed With Error:" + err.Error() + "\n",
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	} else if status != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST UpdateItem Failed as response code received is: " + strconv.Itoa(status),
			Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
		})
		return diags
	} else if status == 200 && resp_body["status"].(string) != "ok" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST UpdateItem Failed as error received in 200 OK Response : " + resp_body["status"].(string),
			Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
		})
		return diags
	} else {
		d.Set("last_updated", time.Now().Format(time.RFC3339))
	}

	resource_alteon_real_server_read(ctx, d, m)

	return diags*/

}

func resource_alteon_real_server_delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	RealServerID := d.Get("index").(string)
	Table := "SlbNewCfgEnhRealServerTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + RealServerID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + RealServerID + "/"
	}
	status, message, err := client.DeleteItem(api, nil, nil)

	resp_body := map[string]interface{}{}
	json.Unmarshal([]byte(message), &resp_body)

	detail := "Status Code Received: " + strconv.Itoa(status) + "\nResponse Received: \n" + message

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST DeleteItem Failed With Error:" + err.Error() + "\n",
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	} else if status != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST DeleteItem Failed as response code received is: " + strconv.Itoa(status),
			Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
		})
		return diags
	} else if status == 200 && resp_body["status"].(string) != "ok" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST DeleteItem Failed as error received in 200 OK Response : " + resp_body["status"].(string),
			Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
		})
		return diags
	} else {
		d.SetId("")
	}

	//resourceRealServerRead(ctx, d, m)

	return diags
}
