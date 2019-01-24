package gedcom_test

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestNewUUIDFromString(t *testing.T) {
	NewUUIDFromString := tf.Function(t, gedcom.NewUUIDFromString)

	NewUUIDFromString("").Returns("", fmt.Errorf("invalid UUID: "))

	NewUUIDFromString("EE13561DDB204985BFFDEEBF82A5226C5B2E").
		Returns("", fmt.Errorf("invalid UUID: EE13561DDB204985BFFDEEBF82A5226C5B2E"))

	NewUUIDFromString("EE13561DDB204985BFFDEEBF82A5226C").
		Returns("ee13561d-db20-4985-bffd-eebf82a5226c", nil)

	NewUUIDFromString("e0d4d387-618a-4713-ab3b-5fa3500b7a75").
		Returns("e0d4d387-618a-4713-ab3b-5fa3500b7a75", nil)

	NewUUIDFromString("z0d4d387-618a-4713-ab3b-5fa3500b7a75").
		Returns("", fmt.Errorf("invalid UUID: z0d4d387-618a-4713-ab3b-5fa3500b7a75"))

	NewUUIDFromString("e0d4d387618a4713ab3b5fa3500b7a75").
		Returns("e0d4d387-618a-4713-ab3b-5fa3500b7a75", nil)
}
