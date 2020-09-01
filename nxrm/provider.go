package nxrm

import (
	"fmt"
	"log"
	"os"

	"github.com/aries1980/terraform-provider-nxrm/version"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/httpclient"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_client_username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NXRM_USERNAME", nil),
				Description: "The Nexus user's username for operations.",
			},

			"api_client_password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NXRM_PASSWORD", nil),
				Description: "The Nexus user's password for operations.",
			},

			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NXRM_ENDPOINT", nil),
				Description: "URL of the Nexus Repository Manager.",
			},

			"api_client_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NXRM_API_CLIENT_LOGGING", false),
				Description: "Whether to print logs from the API client (using the default log library logger)",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"nxrm_asset": dataSourceNxrmAsset(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"nxrm_asset":     resourceNxrmAsset(),
			"nxrm_blobstore": resourceNxrmBlobstore(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	if d.Get("api_client_logging").(bool) {
		options = append(options, cloudflare.UsingLogger(log.New(os.Stderr, "", log.LstdFlags)))
	}

	c := cleanhttp.DefaultClient()
	c.Transport = logging.NewTransport("NXRM", c.Transport)

	tfUserAgent := httpclient.TerraformUserAgent(terraformVersion)
	providerUserAgent := fmt.Sprintf("terraform-provider-nxrm/%s", version.ProviderVersion)
	ua := fmt.Sprintf("%s %s", tfUserAgent, providerUserAgent)
	options = append(options, nxrm.UserAgent(ua))

	config := Config{Options: options}

	if v, ok := d.GetOk("api_client_username"); ok {
		config.APIUsername = v.(string)
	} else if v, ok := d.GetOk("api_client_password"); ok {
		config.APIPassword = v.(string)
	} else {
		return nil, fmt.Errorf("The credentials are not set correctly.")
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	if apiEndpoint, ok := d.GetOk("api_endpoint"); ok {
		log.Printf("[INFO] Using specified %s Nexus Repository URL", apiEndpoint.(string)
		client, err := config.Client()
	} else {
		return client, err
	}

	config.Options = options

	client, err = config.Client()
	if err != nil {
		return nil, err
	}

	return client, err
}
