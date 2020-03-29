package s3

import (
	"github.com/canhnt/terratest-ginkgo-example/test/framework"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = framework.DescribeExample("test s3", func() {
	f := framework.NewFramework()
	t := framework.GetT()
	var bucketID string

	ginkgo.BeforeEach(func() {
		bucketID = terraform.Output(t, f.TerraformOptions, "bucket_id")
	})

	ginkgo.It("bucket should exist", func() {
		aws.AssertS3BucketExists(t, framework.DefaultRegion, bucketID)
	})

	ginkgo.It("should have versioning", func() {
		actualStatus := aws.GetS3BucketVersioning(t, framework.DefaultRegion, bucketID)
		assert.Equal(t, "Enabled", actualStatus)
	})

	ginkgo.It("should have bucket policy", func() {
		aws.AssertS3BucketPolicyExists(t, framework.DefaultRegion, bucketID)
	})

})
