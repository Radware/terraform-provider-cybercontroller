package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_http2_policy() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_http2_policy_read,
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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Http2 policy name.",
			},
			"adminstatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Status (enable/disable) of http2 policy.",
			},
			"streams": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
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
	}
}

func data_source_alteon_http2_policy_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	Http2Policy := make(map[string]interface{})
	json.Unmarshal([]byte(message), &Http2Policy)

	if len(Http2Policy) == 0 { //catch for response with 200 ok and not a json body
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem Failed With Response:" + "\n" + message,
			Detail:   "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}
	//catch to handle GET response body with "status" err and "message": "err getting the data obj not found"
	Items1 := Http2Policy[Table]
	if Items1 == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem failed for non-existing Http2Policy index :" + "\n" + message,
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}

	//Below fix for defect # AAE-930 - GET with empty response
	Items := Http2Policy[Table].([]interface{})
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
	type Http2PolicyItem struct {
		Name              string `json:"name"`
		AdminStatus       int    `json:"adminstatus"`
		Streams           int    `json:"streams"`
		Idle              int    `json:"idle"`
		EnaInsert         int    `json:"enainsert"`
		Header            string `json:"header"`
		EnaServerPush     int    `json:"enaserverpush"`
		HpackSize         string `json:"hpacksize"`
		DeleteStatus      int    `json:"deletestatus"`
		BackendStatus     int    `json:"backendstatus"`
		BackendStreams    int    `json:"backendstreams"`
		BackendHpackSize  string `json:"backendhpacksize"`
		BackendServerPush int    `json:"backendserverpush"`
	}

	var List_Item []Http2PolicyItem
	json.Unmarshal([]byte(helper), &List_Item)
	data1 := make([]map[string]interface{}, len(List_Item))
	rss, _ := json.Marshal(List_Item)
	json.Unmarshal(rss, &data1)
	for key, _ := range data1[0] {
		d.Set(key, data1[0][key])
	}
	d.SetId("Resource GET for Http2 Policy")
	return diags
}
