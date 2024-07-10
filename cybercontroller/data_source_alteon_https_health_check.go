package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_https_health_check() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_https_health_check_read,
		Schema: map[string]*schema.Schema{
			"clustername": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name Of The Cluster.",
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
	}
}

func data_source_alteon_https_health_check_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	HttpsHealthCheck := make(map[string]interface{})
	json.Unmarshal([]byte(message), &HttpsHealthCheck)

	if len(HttpsHealthCheck) == 0 { //catch for response with 200 ok and not a json body
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem Failed With Response:" + "\n" + message,
			Detail:   "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}
	//catch to handle GET response body with "status" err and "message": "err getting the data obj not found"
	Items1 := HttpsHealthCheck[Table]
	if Items1 == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem failed for non-existing HttpsHealth Check  index :" + "\n" + message,
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}
	//Below fix for defect # AAE-930 - GET with empty response
	Items := HttpsHealthCheck[Table].([]interface{})
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
	type HttpsHealthCheckItem struct {
		Name                 string `json:"name"`
		DPort                int    `json:"dport"`
		IPVer                int    `json:"ipver"`
		HostName             string `json:"hostname"`
		Transparent          int    `json:"transparent"`
		Interval             int    `json:"interval"`
		Retries              int    `json:"retries"`
		RestoreRetries       int    `json:"restoreretries"`
		Timeout              int    `json:"timeout"`
		Overflow             int    `json:"overflow"`
		DownInterval         int    `json:"downinterval"`
		Invert               int    `json:"invert"`
		Https                int    `json:"https"`
		Host                 string `json:"host"`
		Path                 string `json:"path"`
		Method               int    `json:"method"`
		Headers              string `json:"headers"`
		Body                 string `json:"body"`
		AuthLevel            int    `json:"authlevel"`
		UserName             string `json:"username"`
		Password             string `json:"password"`
		ResponseType         int    `json:"responsetype"`
		OverloadType         int    `json:"overloadtype"`
		ResponseCode         string `json:"responsecode"`
		ReceiveString        string `json:"receivestring"`
		ResponseCodeOverload string `json:"responsecodeoverload"`
		OverloadString       string `json:"overloadstring"`
		Copy                 string `json:"copy"`
		DeleteStatus         int    `json:"deletestatus"`
		Proxy                int    `json:"proxy"`
		ConnTerm             int    `json:"connterm"`
		HttpsCipherName      int    `json:"httpsciphername"`
		HttpsCipherUserdef   string `json:"httpscipheruserdef"`
		Http2                int    `json:"http2"`
		Always               int    `json:"always"`
		Snat                 int    `json:"snat"`
		ConnTout             int    `json:"conntout"`
	}

	var List_Item []HttpsHealthCheckItem
	json.Unmarshal([]byte(helper), &List_Item)
	data1 := make([]map[string]interface{}, len(List_Item))
	rss, _ := json.Marshal(List_Item)
	json.Unmarshal(rss, &data1)
	for key, _ := range data1[0] {
		d.Set(key, data1[0][key])
	}

	d.SetId("Resource GET for Https Health Check")
	return diags
}
