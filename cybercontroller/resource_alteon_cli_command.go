package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resource_alteon_cli_command() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_alteon_cli_command_create,
		ReadContext:   resource_alteon_cli_command_read,
		UpdateContext: resource_alteon_cli_command_update,
		DeleteContext: resource_alteon_cli_command_delete,
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
			"last_updated": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Resource last updated time.",
			},
			"elements": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agalteonclicommand": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Alteon cli command to be configured.",
						},
					},
				},
			},
		},
	}
}

func resource_alteon_cli_command_create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//mib := "agAlteonCliCommand"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "//config"
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
		d.SetId("cli command configuration is performed")
	}

	return diags
}

func resource_alteon_cli_command_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST ReadItem not supported",
		Detail:   "Read Item not supported for this resource type",
	})
	return diags
}

func resource_alteon_cli_command_update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST UpdateItem Failed",
		Detail:   "Update Item not supported for this resource type",
	})
	return diags
}

func resource_alteon_cli_command_delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST DeleteItem Failed",
		Detail:   "Delete Item not supported for this resource type",
	})
	return diags
}
