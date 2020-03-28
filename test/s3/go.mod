module example.com/test/s3

go 1.14

require (
	example.com/test/framework v0.0.0

	github.com/gruntwork-io/terratest v0.26.1
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.7.1
	github.com/stretchr/testify v1.5.1
)

replace example.com/test/framework => ../framework
