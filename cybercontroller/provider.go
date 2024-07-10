package cybercontroller

import (
	"context"
	"strconv"

	radwaregosdk "github.com/Radware/radware_go_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CC_USERNAME", nil),
				Description: "Cyber Controller Username.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CC_PASSWORD", nil),
				Description: "Cyber Controller Password.",
			},
			"ip": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CC_IP", nil),
				Description: "Management IP of Cyber Controller.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cybercontroller_alteon_real_server":        resource_alteon_real_server(),
			"cybercontroller_alteon_server_group":       resource_alteon_server_group(),
			"cybercontroller_alteon_cli_command":        resource_alteon_cli_command(),
			"cybercontroller_alteon_apply":              resource_alteon_apply(),
			"cybercontroller_alteon_save":               resource_alteon_save(),
			"cybercontroller_alteon_revert":             resource_alteon_revert(),
			"cybercontroller_alteon_revert_apply":       resource_alteon_revert_apply(),
			"cybercontroller_alteon_virtual_server":     resource_alteon_virtual_server(),
			"cybercontroller_alteon_virtual_service":    resource_alteon_virtual_service(),
			"cybercontroller_alteon_ssl_policy":         resource_alteon_ssl_policy(),
			"cybercontroller_alteon_http2_policy":       resource_alteon_http2_policy(),
			"cybercontroller_alteon_https_health_check": resource_alteon_https_health_check(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			//"cybercontroller_real_server": dataSourceRealServer(),
			"cybercontroller_alteon_real_server_data":          data_source_alteon_real_server(),
			"cybercontroller_alteon_apply_status_cluster_data": data_source_alteon_apply_status_cluster(),
			"cybercontroller_alteon_server_group_data":         data_source_alteon_server_group(),
			"cybercontroller_alteon_virtual_server_data":       data_source_alteon_virtual_server(),
			"cybercontroller_alteon_https_health_check_data":   data_source_alteon_https_health_check(),
			"cybercontroller_alteon_ssl_policy_data":           data_source_alteon_ssl_policy(),
			"cybercontroller_alteon_http2_policy_data":         data_source_alteon_http2_policy(),
			"cybercontroller_alteon_virtual_service_data":      data_source_alteon_virtual_service(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	host := d.Get("ip").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	/*
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Warning Message Summary",
			Detail:   "This is the detailed warning message from providerConfigure",
		})*/

	if (username != "") && (password != "") {
		client, status, message, err := radwaregosdk.Login("CYBERCONTROLLER", host, username, password)
		//client.HostIP = host
		detail := strconv.Itoa(status) + message

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to connect to CyberController." + detail + err.Error(),
				Detail:   detail + err.Error(),
			})
			return nil, diags
		}

		if client == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Authentication Failed" + detail,
				Detail:   detail,
			})
			return nil, diags
		}

		return client, diags
	}

	return 1, diags
}
