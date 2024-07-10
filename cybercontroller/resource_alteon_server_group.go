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

func resource_alteon_server_group() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_alteon_server_group_create,
		ReadContext:   resource_alteon_server_group_read,
		UpdateContext: resource_alteon_server_group_update,
		DeleteContext: resource_alteon_server_group_delete,
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
				Optional:    true,
				Description: "The group alphanumeric index for which the information pertains.",
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
						"addserver": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The real server to be added to the group. When read, 0 is returned.",
						},
						"removeserver": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The real server to be removed from the group. When read, 0 is returned.",
						},
						"healthcheckurl": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The specific content which is examined during health checks. The content depends on the type of health check.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the real server group.",
						},
						"healthchecklayer": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The OSI layer at which servers are health checked. From version 29.0.0.0 the following values are not supported: snmp2-snmp5, script1-script64.",
						},
						"metric": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The metric used to select next server in group.",
						},
						"backupserver": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The backup real server for this group.",
						},
						"backupgroup": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The backup real server group for this group.",
						},
						"realthreshold": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The minimum number of real servers available.If it reaches the minimum limit a SYSLOG ALERT message is send to to the configured syslog servers stating that the real server threshold has been reached for the concerned group.",
						},
						"viphealthcheck": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable VIP health checking in DSR mode.",
						},
						"idsstate": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable intrusion detection.",
						},
						"idsport": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The intrusion detection port. A value of 1 is invalid.",
						},
						"deletestatus": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "By setting the value to delete(2), the entire group is deleted.",
						},
						"idsflood": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable intrusion detection group flood.",
						},
						"minmisshash": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "24|32 number of sip bits used for minmisses hash in the new_configuration block.",
						},
						"phashmask": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address mask used by the persistent hash metric.",
						},
						"rmetric": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The metric used to select next rport in server.",
						},
						"healthcheckformula": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The formula used to state the actual health of a virtual service. It allows user to use the symbols of '(', ')', '|', '&' to construct a formula to state the health of the server group.This string can take the following formats : '(1&2|3..)', '128' or 'none'",
						},
						"operatoraccess": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable access to this group for operator.",
						},
						"wlm": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The Workload Manager for this Group.",
						},
						"radiusauthenstring": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The group RADIUS authentication string. The string is used for generating encrypted authentication string while doing RADIUS health check for this group radius servers.",
						},
						"secbackupgroup": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Secondary backup real server group for this group.",
						},
						"slowstart": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The slow-start time for this group.",
						},
						"minthreshold": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The minimum threshold value for this group.",
						},
						"maxthreshold": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The maximum threshold value for this group.",
						},
						"ipver": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The type of real server group IP address.",
						},
						"backup": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The backup real group or real server for this group.",
						},
						"backuptype": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Backup type of the real server group.",
						},
						"healthid": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Advanced HC ID.",
						},
						"phashprefixlength": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Prefix length used by the persistent hash metric.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Group type.",
						},
						"copy": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The alphanumeric index of the new copy to be created.",
						},
						"idschain": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable IDS group participation in inspection chain.",
						},
						"sectype": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The Group security device type.",
						},
						"secdeviceflag": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The Group security device flag",
						},
						"maxconex": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or Disable override maximum connections limit.",
						},
					},
				},
			},
		},
	}
}

func resource_alteon_server_group_create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ServerGroupID := d.Get("index").(string)
	Table := "SlbNewCfgEnhGroupTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + ServerGroupID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + ServerGroupID + "/"
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
		d.SetId("Resource Create for Server Group")
	}

	resource_alteon_server_group_read(ctx, d, m)

	return diags
}

func resource_alteon_server_group_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resource_alteon_server_group_update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ServerGroupID := d.Get("index").(string)
	Table := "SlbNewCfgEnhGroupTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + ServerGroupID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + ServerGroupID + "/"
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

	resource_alteon_server_group_read(ctx, d, m)

	return diags
}

func resource_alteon_server_group_delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ServerGroupID := d.Get("index").(string)
	Table := "SlbNewCfgEnhGroupTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + ServerGroupID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + ServerGroupID + "/"
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

	//resourceServerGroupRead(ctx, d, m)

	return diags
}
