package action_test

import (
	"github.com/cloudfoundry/bosh-cpi-go/apiv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/orange-cloudfoundry/bosh-cpi-cloudstack/config"
)

var _ = Describe("FactoryOpts", func() {
	var (
		opts FactoryOpts

		validOptions = FactoryOpts{
			Agent: apiv1.AgentOptions{
				Mbus: "fake-mbus",
				NTP:  []string{},
			},
		}
	)

	Describe("Validate", func() {
		BeforeEach(func() {
			opts = validOptions
		})

		It("does not return error if all fields are valid", func() {
			err := opts.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if agent section is not valid", func() {
			opts.Agent.Mbus = ""

			err := opts.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Validating Agent configuration"))
		})
	})
})
