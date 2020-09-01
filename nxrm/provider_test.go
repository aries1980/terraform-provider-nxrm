package nxrm

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cloudflare": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

type preCheckFunc = func(*testing.T)

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("NXRM_USERNAME"); v == "" {
		t.Fatal("NXRM_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("NXRM_PASSWORD"); v == "" {
		t.Fatal("NXRM_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("NXRM_ENDPOINT"); v == "" {
		t.Fatal("NXRM_ENDPOINT must be set for acceptance tests. The domain is used to create and destroy record against.")
	}
}

func generateRandomResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}
