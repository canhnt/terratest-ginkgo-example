package framework

import (
	"github.com/gruntwork-io/terratest/modules/testing"
	. "github.com/onsi/ginkgo"
)

func GetT() testing.TestingT {
	return &myTestingT{
		GinkgoT(),
	}
}

type myTestingT struct {
	GinkgoTInterface
}

// Extends GinkgoTInterface to have #Name() method, that is compatible with testing.TestingT
func (mt *myTestingT) Name() string {
	return "[TerraTest+Ginkgo]"
}
