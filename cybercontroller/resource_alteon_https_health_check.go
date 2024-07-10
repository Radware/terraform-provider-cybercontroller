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

func resource_alteon_https_health_check() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_alteon_https_health_check_create,
		ReadContext:   resource_alteon_https_health_check_read,
		UpdateContext: resource_alteon_https_health_check_update,
		DeleteContext: resource_alteon_https_health_check_delete,
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
				Description: "HTTP Health check id.",
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
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check name.",
						},
						"dport": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check destination port.",
						},
						"ipver": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check destination IP version.",
						},
						"hostname": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check destination hostname.",
						},
						"transparent": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check transparent flag.",
						},
						"interval": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check interval.",
						},
						"retries": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check retries counter.",
						},
						"restoreretries": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check retries in down state counter.",
						},
						"timeout": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check timeout.",
						},
						"overflow": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check overflow flag.",
						},
						"downinterval": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check interval in down state.",
						},
						"invert": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check invert flag.",
						},
						"https": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check HTTPS enable/disable flag.",
						},
						"host": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check host field.",
						},
						"path": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check path field.",
						},
						"method": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check HTTP method.",
						},
						"headers": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check headers list.",
						},
						"body": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check body field.",
						},
						"authlevel": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check authentication method.",
						},
						"username": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check user name.",
						},
						"password": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check password.",
						},
						"responsetype": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check response handling method.",
						},
						"overloadtype": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Overload string is included or not included.",
						},
						"responsecode": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check expected response code.",
						},
						"receivestring": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check expected response string.",
						},
						"responsecodeoverload": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check expected code for overflow state.",
						},
						"overloadstring": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Expected response for server overload.",
						},
						"copy": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP Health check copy action trigger.",
						},
						"deletestatus": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "When set to the value of 2 (delete), the entire row is deleted. When read, other(1) is returned. Setting the value to anything other than 2(delete) has no effect on the state of the row.",
						},
						"proxy": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/disable HTTP health check proxy request.",
						},
						"connterm": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Connection termination type.",
						},
						"httpsciphername": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Cipher name for SSL for HTTPS HC Context.",
						},
						"httpscipheruserdef": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cipher-suite allowed for SSL for HTTPS HC Context.",
						},
						"http2": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health check HTTP2 flag.",
						},
						"always": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "This flag determines whether HC is allowed for standalone real.",
						},
						"snat": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "HTTP Health Check src NAT (PIP) flag.",
						},
						"conntout": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Connection termination on timeout type.",
						},
					},
				},
			},
		},
	}
}

func resource_alteon_https_health_check_create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	AdvhcHttpID := d.Get("index").(string)
	Table := "SlbNewAdvhcHttpTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + AdvhcHttpID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + AdvhcHttpID + "/"
	}

	items := d.Get("elements").([]interface{})
	//rss := []radwaregosdk.Http2PolicyItem{}
	validvals := make(map[string]interface{})

	for _, item := range items {
		i := item.(map[string]interface{})
		for key, conf_val := range i {
			if conf_val != "" && conf_val != 0 {
				validvals[key] = conf_val
			}
		}
		/*rsi := radwaregosdk.Http2PolicyItem{
			Name:        i["name"].(string),
			AdminStatus: i["adminstatus"].(int),
		}
		rss = append(rss, rsi)*/
	}

	//brss, err := json.MarshalIndent(rss[0], "", "    ")
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
		d.SetId("Resource Create for Https Health Check")
	}

	resource_alteon_https_health_check_read(ctx, d, m)

	return diags
}

func resource_alteon_https_health_check_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	/*var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST ReadItem Failed",
		Detail:   "Read Item not supported for this resource type",
	})
	return diags*/
	return nil
}

func resource_alteon_https_health_check_update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	AdvhcHttpID := d.Get("index").(string)
	Table := "SlbNewAdvhcHttpTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + AdvhcHttpID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + AdvhcHttpID + "/"
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

	resource_alteon_https_health_check_read(ctx, d, m)

	return diags
}

func resource_alteon_https_health_check_delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	AdvhcHttpID := d.Get("index").(string)
	Table := "SlbNewAdvhcHttpTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + AdvhcHttpID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + AdvhcHttpID + "/"
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

	//resourceHttp2PolicyRead(ctx, d, m)

	return diags
}
