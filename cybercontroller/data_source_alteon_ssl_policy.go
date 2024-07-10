package cybercontroller

import (
	"context"
	"encoding/json"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func data_source_alteon_ssl_policy() *schema.Resource {
	return &schema.Resource{
		ReadContext: data_source_alteon_ssl_policy_read,
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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SSL policy name,length of the string should be 32 characters.",
			},
			"adminstatus": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Enable or disable ssl policy.",
			},
			"fessl": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
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
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set Frontend SSL reuse state",
			},
			"bereusestate": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set Backend SSL reuse state",
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
	}
}

func data_source_alteon_ssl_policy_read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	SslPolicy := make(map[string]interface{})
	json.Unmarshal([]byte(message), &SslPolicy)

	if len(SslPolicy) == 0 { //catch for response with 200 ok and not a json body
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem Failed With Response:" + "\n" + message,
			Detail:   "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}

	//catch to handle GET response body with "status" err and "message": "err getting the data obj not found"
	Items1 := SslPolicy[Table]
	if Items1 == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "REST GetItem failed for non-existing ssl policy index :" + "\n" + message,
			Detail:   detail + "\nAPI Call Made is:" + api + "\n",
		})
		return diags
	}
	//Below fix for defect # AAE-930 - GET with empty response
	Items := SslPolicy[Table].([]interface{})
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
	type SslPolicyItem struct {
		Name                  string `json:"name"`
		PassInfoCipherName    string `json:"passinfociphername"`
		PassInfoCipherFlag    int    `json:"passinfocipherflag"`
		PassInfoVersionName   string `json:"passinfoversionname"`
		PassInfoVersionFlag   int    `json:"passinfoversionflag"`
		PassInfoHeadBitsName  string `json:"passinfoheadbitsname"`
		PassInfoHeadBitsFlag  int    `json:"passinfoheadbitsflag"`
		PassInfoFrontend      int    `json:"passinfofrontend"`
		CipherName            int    `json:"ciphername"`
		CipherUserdef         string `json:"cipheruserdef"`
		IntermcaChainName     string `json:"intermcachainname"`
		IntermcaChainType     string `json:"intermcachaintype"`
		Becipher              int    `json:"becipher"`
		Authpol               string `json:"authpol"`
		Convuri               string `json:"convuri"`
		Bessl                 int    `json:"bessl"`
		Convert               int    `json:"convert"`
		AdminStatus           int    `json:"adminstatus"`
		Del                   int    `json:"del"`
		PassInfoComply        int    `json:"passinfocomply"`
		Fessl                 int    `json:"fessl"`
		FESslv3Version        int    `json:"fesslv3version"`
		FETls10Version        int    `json:"fetls10version"`
		FETls11Version        int    `json:"fetls11version"`
		BESslv3Version        int    `json:"besslv3version"`
		BETls10Version        int    `json:"betls10version"`
		FETls12Version        int    `json:"fetls12version"`
		BETls12Version        int    `json:"betls12version"`
		CipherExpertUserdef   string `json:"cipherexpertuserdef"`
		BeCipherUserdef       string `json:"becipheruserdef"`
		BeCipherExpertUserdef string `json:"becipherexpertuserdef"`
		BEClientCertName      string `json:"beclientcertname"`
		BETrustedCAcertName   string `json:"betrustedcacertname"`
		BETrustedCAcertType   string `json:"betrustedcacerttype"`
		Secreneg              string `json:"secreneg"`
		DHkey                 int    `json:"dhkey"`
		BEAuthpol             string `json:"beauthpol"`
		HwoffldFeRsa          int    `json:"hwoffldfersa"`
		HwoffldFeDh           int    `json:"hwoffldfedh"`
		HwoffldFeEc           int    `json:"hwoffldfeec"`
		HwoffldFeBulk         int    `json:"hwoffldfebulk"`
		HwoffldFePkey         int    `json:"hwoffldfepkey"`
		HwoffldFeFeatures     int    `json:"hwoffldFefeatures"`
		HwoffldBeRsa          int    `json:"hwoffldbersa"`
		HwoffldBeDh           int    `json:"hwoffldbedh"`
		HwoffldBeEc           int    `json:"hwoffldbeec"`
		HwoffldBeBulk         int    `json:"hwoffldbebulk"`
		HwoffldBePkey         int    `json:"hwoffldbepkey"`
		HwoffldBeFeatures     int    `json:"hwoffldbefeatures"`
		Besni                 int    `json:"besni"`
		FETls13Version        int    `json:"fetls13version"`
		BETls13Version        int    `json:"betls13version"`
		SSLPol0RTTFEData      int    `json:"sslpol0rttfedata"`
		FeReuseState          int    `json:"fereusestate"`
		BeReuseState          int    `json:"bereusestate"`
		BeReuseSrcMatch       int    `json:"bereusesrcmatch"`
		BeReuseTicket         int    `json:"bereuseticket"`
		FeReuseTicket         int    `json:"fereuseticket"`
		FEGmSslVersion        int    `json:"fegmsslversion"`
		BEGmSslVersion        int    `json:"begmsslversion"`
		FEGmSslPriority       int    `json:"fegmsslpriority"`
		FESslsigs             string `json:"fesslsigs"`
		BESslsigs             string `json:"besslsigs"`
		FESslgroups           string `json:"fesslgroups"`
		BESslgroups           string `json:"besslgroups"`
		HSTmout               int    `json:"hstmout"`
	}

	var List_Item []SslPolicyItem
	json.Unmarshal([]byte(helper), &List_Item)
	data1 := make([]map[string]interface{}, len(List_Item))
	rss, _ := json.Marshal(List_Item)
	json.Unmarshal(rss, &data1)
	for key, _ := range data1[0] {
		d.Set(key, data1[0][key])
	}
	d.SetId("Resource GET for SslPolicy")
	return diags
}
