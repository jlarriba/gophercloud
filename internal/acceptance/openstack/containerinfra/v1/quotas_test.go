//go:build acceptance || containerinfra
// +build acceptance containerinfra

package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/internal/acceptance/clients"
	"github.com/gophercloud/gophercloud/internal/acceptance/tools"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestQuotasCRUD(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	th.AssertNoErr(t, err)

	quota, err := CreateQuota(t, client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, quota)
}
