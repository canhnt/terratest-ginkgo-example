package test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"

	// tests to run
	_ "github.com/canhnt/terratest-ginkgo-example/test/s3"
)

func RunE2ETests(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "terratest e2e suite", []ginkgo.Reporter{reporters.NewJUnitReporter("it_cov.xml")})
}
