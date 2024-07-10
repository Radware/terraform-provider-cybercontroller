package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_server_group() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_server_group_read,
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
			/*"state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},*/
		},
	}
}

func data_source_alteon_server_group_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)
	var diags diag.Diagnostics

	ServerGroupID := d.Get("index").(string)
	//RealServerID := d.Get("realserverindex").(string)

	/*Tables := [3]string{"SlbNewCfgEnhGroupTable",
		"SlbEnhGroupRealServersTable",
		"SlbNewCfgEnhGroupRealServerTable",
	}*/
	Table := "SlbNewCfgEnhGroupTable"

	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)

	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + ServerGroupID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + ServerGroupID + "/"
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

	ServerGroup := make(map[string]interface{})
	json.Unmarshal([]byte(message), &ServerGroup)

	if len(ServerGroup) == 0 { //catch for response with 200 ok and not a json body
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem Failed With Response:" + "\n" + message,
			Detail:   "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}

	//catch to handle GET response body with "status" err and "message": "err getting the data obj not found"
	Items1 := ServerGroup[Table]
	if Items1 == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem failed for non-existing Server Group index :" + "\n" + message,
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}

	//Below fix for defect # AAE-930 - GET with empty response
	Items := ServerGroup[Table].([]interface{})
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

	type ServerGroupItem struct {
		AddServer          string `json:"addserver"`
		RemoveServer       string `json:"removeserver"`
		Metric             int    `json:"metric"`
		BackupServer       string `json:"backupserver"`
		BackupGroup        string `json:"backupgroup"`
		HealthCheckUrl     string `json:"healthcheckurl"`
		HealthCheckLayer   int    `json:"healthchecklayer"`
		Name               string `json:"name"`
		RealThreshold      int    `json:"realthreshold"`
		VipHealthCheck     int    `json:"viphealthcheck"`
		IdsState           int    `json:"idsstate"`
		IdsPort            int    `json:"idsport"`
		DeleteStatus       int    `json:"deletestatus"`
		IdsFlood           int    `json:"idsflood"`
		MinmissHash        int    `json:"minmisshash"`
		PhashMask          string `json:"phashmask"`
		Rmetric            int    `json:"rmetric"`
		HealthCheckFormula string `json:"healthcheckformula"`
		OperatorAccess     int    `json:"operatoraccess"`
		Wlm                int    `json:"wlm"`
		RadiusAuthenString string `json:"radiusauthenstring"`
		SecBackupGroup     string `json:"secbackupgroup"`
		Slowstart          int    `json:"slowstart"`
		MinThreshold       int    `json:"minthreshold"`
		MaxThreshold       int    `json:"maxthreshold"`
		IpVer              int    `json:"ipver"`
		Backup             string `json:"backup"`
		BackupType         int    `json:"backuptype"`
		HealthID           string `json:"healthid"`
		PhashPrefixLength  int    `json:"phashprefixlength"`
		Type               int    `json:"type"`
		Copy               string `json:"copy"`
		IdsChain           int    `json:"idschain"`
		SecType            int    `json:"sectype"`
		SecDeviceFlag      int    `json:"secdeviceflag"`
		MaxConEx           int    `json:"maxconex"`
	}

	var List_Item []ServerGroupItem
	json.Unmarshal([]byte(helper), &List_Item)
	data1 := make([]map[string]interface{}, len(List_Item))
	rss, _ := json.Marshal(List_Item)
	json.Unmarshal(rss, &data1)
	for key, _ := range data1[0] {
		d.Set(key, data1[0][key])
	}

	d.SetId("Resource GET for Server Group")

	return diags
}
