package framework

import (
	"fmt"
	"strings"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	. "github.com/onsi/ginkgo"
)

const DefaultRegion = "eu-west-1"

type Framework struct {
	TerraformOptions *terraform.Options
}

func NewFramework() *Framework {
	expectedName := fmt.Sprintf("terratest-aws-s3-example-%s", strings.ToLower(random.UniqueId()))

	// Give this S3 Bucket an environment to operate as a part of for the purposes of resource tagging
	expectedEnvironment := "Automated Testing"

	f := &Framework{
		TerraformOptions: &terraform.Options{
			// The path to where our Terraform code is located
			TerraformDir: "../plan",

			// Variables to pass to our Terraform code using -var options
			Vars: map[string]interface{}{
				"tag_bucket_name":        expectedName,
				"tag_bucket_environment": expectedEnvironment,
				"with_policy":            "true",
			},

			// Environment variables to set when running Terraform
			EnvVars: map[string]string{
				"AWS_DEFAULT_REGION": DefaultRegion,
			},
		},
	}

	BeforeSuite(f.BeforeSuite)
	AfterSuite(f.AfterSuite)
	return f
}

func (f *Framework) BeforeSuite() {
	print("InitAndApply Terraform")
	terraform.InitAndApply(GetT(), f.TerraformOptions)
}

func (f *Framework) AfterSuite() {
	print("Destroy Terraform")
	terraform.Destroy(GetT(), f.TerraformOptions)
}

func DescribeExample(text string, body func()) bool {
	return Describe("[example] "+text, body)
}
