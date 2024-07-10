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

func resource_alteon_virtual_server() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_alteon_virtual_server_create,
		ReadContext:   resource_alteon_virtual_server_read,
		UpdateContext: resource_alteon_virtual_server_update,
		DeleteContext: resource_alteon_virtual_server_delete,
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
			"index": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Virtual Server Number.",
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
						"virtserveripaddress": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address of the virtual server.",
						},
						"virtserverstate": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3,
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
				},
			},
		},
	}
}

func resource_alteon_virtual_server_create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	items := d.Get("elements").([]interface{})
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
	status, message, err := client.CreateItem(api, APIBytes, nil)

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
		d.SetId("Resource Create for Virtual Server")
	}

	resource_alteon_virtual_server_read(ctx, d, m)

	return diags
}

func resource_alteon_virtual_server_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	/*var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST ReadItem Failed",
		Detail:   "Read Item not supported for this resource type",
	})
	return diags*/
	return nil
}

func resource_alteon_virtual_server_update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	items := d.Get("elements").([]interface{})
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

	resource_alteon_virtual_server_read(ctx, d, m)

	return diags
}

func resource_alteon_virtual_server_delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	//resourceVirtualServerRead(ctx, d, m)

	return diags
}
