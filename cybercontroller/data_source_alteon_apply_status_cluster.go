package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_apply_status_cluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_apply_status_cluster_read,
		Schema: map[string]*schema.Schema{
			"clustername": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the cluster",
			},
			"apply_response": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agapplyconfig": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Last Apply Status-DONE or FAILED",
						},
						"agapplytable": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Error messages if last apply was failed.",
						},
						"last_apply": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time when last apply happened.",
						},
					},
				},
			},
		},
	}
}

func data_source_alteon_apply_status_cluster_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)
	var diags diag.Diagnostics

	var api string
	clustername := d.Get("clustername").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/AgApplyState"
	}

	status, message, err1 := client.GetItem(api, nil, nil)
	detail := "Status Code Received: " + strconv.Itoa(status) + "\nResponse Received: \n" + message

	if err1 != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST Apply Status-Get Failed With Error:" + err1.Error() + "\n",
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	} else if status != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST Apply Status-Get Failed as response code received is: " + strconv.Itoa(status),
			Detail:   detail + "\nAPI Call Made is:\n" + api + "\n",
		})
		return diags
	}
	type ApplyResponse struct {
		AgApplyConfig string `json:"agapplyconfig"`
		AgApplyTable  []struct {
			Index     int    `json:"Index"`
			StringVal string `json:"StringVal"`
		} `json:"agapplytable"`
		LastApply string `json:"last_apply"`
	}
	var msg ApplyResponse
	err := json.Unmarshal([]byte(message), &msg)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST Apply Status-Get Failed With Error:" + err.Error() + "\n",
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}

	var resp_list []interface{}

	if msg.AgApplyConfig == "DONE" { //agApplyTable will be empty if apply is successful
		resp_list = append(resp_list, map[string]interface{}{
			"agapplyconfig": msg.AgApplyConfig,
			//"agapplytable":  msg.AgApplyTable[0].StringVal,
			"last_apply": msg.LastApply,
		})
		/*}else if len(msg.AgApplyTable) == 0 { //if response is empty list for agApplyTable in cases were applyconfig is not "DONE"
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Apply status failed as ApplyTable is empty:" + "\n",
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags*/
	} else if msg.AgApplyConfig == "FAILED" {
		resp_list = append(resp_list, map[string]interface{}{ //ApplyTable has details if apply failed
			"agapplyconfig": msg.AgApplyConfig,
			"agapplytable":  msg.AgApplyTable[0].StringVal,
			"last_apply":    msg.LastApply,
		})
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Apply status response is not as expected :" + "\n",
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}

	d.Set("apply_response", resp_list)
	if msg.AgApplyConfig != "DONE" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Apply was not successful due to below errors:" + "\n",
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}
	d.SetId("GET for Apply Status") //for state machine
	return diags
}
