package healthprobe_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHealthprobe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Healthprobe Suite")
}
