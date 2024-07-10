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

func resource_alteon_http2_policy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_alteon_http2_policy_create,
		ReadContext:   resource_alteon_http2_policy_read,
		UpdateContext: resource_alteon_http2_policy_update,
		DeleteContext: resource_alteon_http2_policy_delete,
		Schema: map[string]*schema.Schema{
			"clustername": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of cluster.",
			},
			"alteonip": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP Address of the Alteon managed by the cybercontroller.",
			},
			"nameidindex": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The http2 policy ID(key id) as an index.",
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
							Description: "Http2 policy name.",
						},
						"adminstatus": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     2,
							Description: "Status (enable/disable) of http2 policy.",
						},
						"streams": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     100,
							Description: "Defines the maximum concurrent streams per connection.",
						},
						"idle": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Defines the number of seconds an HTTP/2 connection is left open idly before it is closed.",
						},
						"enainsert": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Define (enable/disable) of insert private http2 header.",
						},
						"header": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Http2 policy header string.",
						},
						"enaserverpush": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Define (enable/disable) of http2 server push.",
						},
						"hpacksize": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Set HTTP2 policy HPACK table size (small / medium / large).",
						},
						"deletestatus": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete Http2 policy.",
						},
						"backendstatus": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: " Backend Status (enable/disable) of backend http2.",
						},
						"backendstreams": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "maximum concurrent streams per backend connection.",
						},
						"backendhpacksize": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTP2 policy backend HPACK table size (small / medium / large).",
						},
						"backendserverpush": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "maximum concurrent server push streams per backend connection.",
						},
					},
				},
			},
		},
	}
}

func resource_alteon_http2_policy_create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	Http2PolicyID := d.Get("nameidindex").(string)
	Table := "SlbNewAcclCfgHttp2PolTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + Http2PolicyID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + Http2PolicyID + "/"
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
		d.SetId("Resource Create for Http2 Policy")
	}

	resource_alteon_http2_policy_read(ctx, d, m)

	return diags
}

func resource_alteon_http2_policy_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	/*var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST ReadItem Failed",
		Detail:   "Read Item not supported for this resource type",
	})
	return diags*/
	return nil
}

func resource_alteon_http2_policy_update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	Http2PolicyID := d.Get("nameidindex").(string)
	Table := "SlbNewAcclCfgHttp2PolTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + Http2PolicyID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + Http2PolicyID + "/"
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

	resource_alteon_http2_policy_read(ctx, d, m)

	return diags
}

func resource_alteon_http2_policy_delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	Http2PolicyID := d.Get("nameidindex").(string)
	Table := "SlbNewAcclCfgHttp2PolTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + Http2PolicyID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + Http2PolicyID + "/"
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
