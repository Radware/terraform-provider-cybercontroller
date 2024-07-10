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

func resource_alteon_ssl_policy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_alteon_ssl_policy_create,
		ReadContext:   resource_alteon_ssl_policy_read,
		UpdateContext: resource_alteon_ssl_policy_update,
		DeleteContext: resource_alteon_ssl_policy_delete,
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
			"nameidindex": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The SSL policy name(key id) as an index.",
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
							Description: "SSL policy name,length of the string should be 32 characters.",
						},
						"adminstatus": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     2,
							Description: "Enable or disable ssl policy.",
						},
						"fessl": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
							Description: "Set frontend SSL encryption mode, default value is enabled.",
						},
						"bessl": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable backend SSL encryption.",
						},
						"fesslv3version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable frontend sslv3.",
						},
						"passinfociphername": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The SSL cipher-suite header name.",
						},
						"passinfocipherflag": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable cipher-suite information to backend servers.",
						},
						"passinfoversionname": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "SSL version header name.",
						},
						"passinfoversionflag": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable SSL version information to backend servers.",
						},
						"passinfoheadbitsname": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The passive cipher bits information to backend server.",
						},
						"passinfoheadbitsflag": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable Cipher bits information to backend servers.",
						},
						"passinfofrontend": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable add Front-End-Https: on header.",
						},
						"ciphername": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Allowed cipher-suites in frontend SSL.",
						},
						"cipheruserdef": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cipher-suite allowed for SSL.",
						},
						"intermcachainname": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Intermediate CA certificate chain name.",
						},
						"intermcachaintype": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Intermediate CA certificate chain type certificate=cert,Group=group,None=empty string.",
						},
						"becipher": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Allowed cipher-suites in backend SSL.",
						},
						"authpol": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Client authentication policy.",
						},
						"convuri": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Host regex for HTTP redirection conversion.",
						},
						"convert": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable HTTP redirection conversion.",
						},
						"del": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete SSL policy.",
						},
						"passinfocomply": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable X-SSL header compatible with 2424SSL headers.",
						},
						"fetls10version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable frontend tls1_0.",
						},
						"fetls11version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable frontend tls1_1.",
						},
						"besslv3version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable backend sslv3.",
						},
						"betls10version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable backend tls1_0.",
						},
						"betls11version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable backend tls1_1.",
						},
						"fetls12version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable frontend tls1_2.",
						},
						"betls12version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable backend tls1_2.",
						},
						"cipherexpertuserdef": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Expert-Cipher-suite allowed for SSL.",
						},
						"becipheruserdef": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "BeCipher-suite allowed for SSL.",
						},
						"becipherexpertuserdef": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Expert-BeCipher-suite allowed for SSL.",
						},
						"beclientcertname": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Backend client certificate name.",
						},
						"betrustedcacertname": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Backend trusted CA certificate chain name.",
						},
						"betrustedcacerttype": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Backend trusted CA certificate chain type certificate=cert,Group=group,None=empty string.",
						},
						"secreneg": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secure Renegotiation Frontend and Backend SSL.",
						},
						"dhkey": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "num of bits in Diffie Helman key.",
						},
						"beauthpol": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Backend authentication policy.",
						},
						"hwoffldfersa": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for RSA algorithm.",
						},
						"hwoffldfedh": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for DH algorithm.",
						},
						"hwoffldfeec": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for EC algorithm.",
						},
						"hwoffldfebulk": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for bulk encryption ciphers.",
						},
						"hwoffldfepkey": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for PKEY functionality.",
						},
						"hwoffldfefeatures": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload Features.",
						},
						"hwoffldbersa": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for RSA algorithm.",
						},
						"hwoffldbedh": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for DH algorithm.",
						},
						"hwoffldbeec": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for EC algorithm.",
						},
						"hwoffldbebulk": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for bulk encryption ciphers.",
						},
						"hwoffldbepkey": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload for PKEY functionality.",
						},
						"hwoffldbefeatures": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Disable/Enable HW offload Features.",
						},
						"besni": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/Disable to include SNI.",
						},
						"fetls13version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable frontend tls1_3.",
						},
						"betls13version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable backend tls1_3.",
						},
						"sslpol0rttfedata": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Set maximum allowed early data on frontend connection.",
						},
						"fereusestate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bereusestate": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Set Frontend SSL reuse state",
						},
						"bereusesrcmatch": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/disable reuse for same client only",
						},
						"bereuseticket": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/disable TLS 1.2 session ticket",
						},
						"fereuseticket": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable/disable TLS 1.2 session ticket",
						},
						"fegmsslversion": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable frontend GM SSL.",
						},
						"begmsslversion": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable backend GM SSL.",
						},
						"fegmsslpriority": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Enable or disable frontend GM SSL.",
						},
						"fesslsigs": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Allowed signature algorithms.",
						},
						"besslsigs": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Allowed signature algorithms in backend SSL.",
						},
						"fesslgroups": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Allowed groups.",
						},
						"besslgroups": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Set allowed groups in backend SSL.",
						},
						"hstmout": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Frontend SSL handshake timeout in msec.",
						},
					},
				},
			},
		},
	}
}

func resource_alteon_ssl_policy_create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	SslPolicyID := d.Get("nameidindex").(string)
	Table := "SlbNewSslCfgSSLPolTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + SslPolicyID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + SslPolicyID + "/"
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
		d.SetId("Resource Create for Ssl Policy")
	}

	resource_alteon_ssl_policy_read(ctx, d, m)

	return diags
}

func resource_alteon_ssl_policy_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	/*var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "REST ReadItem Failed",
		Detail:   "Read Item not supported for this resource type",
	})
	return diags*/
	return nil
}

func resource_alteon_ssl_policy_update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	SslPolicyID := d.Get("nameidindex").(string)
	Table := "SlbNewSslCfgSSLPolTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + SslPolicyID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + SslPolicyID + "/"
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

	resource_alteon_ssl_policy_read(ctx, d, m)

	return diags
}

func resource_alteon_ssl_policy_delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*radwaregosdk.New_Client)

	var diags diag.Diagnostics

	SslPolicyID := d.Get("nameidindex").(string)
	Table := "SlbNewSslCfgSSLPolTable"
	var api string
	clustername := d.Get("clustername").(string)
	alteonip := d.Get("alteonip").(string)
	if clustername != "" {
		api = "/acm/main/updateAlteons/" + clustername + "/config/" + Table + "/" + SslPolicyID + "/"
	} else if alteonip != "" {
		api = "/mgmt/device/byip/" + alteonip + "/config/" + Table + "/" + SslPolicyID + "/"
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

	//resourceSslPolicyRead(ctx, d, m)

	return diags
}
