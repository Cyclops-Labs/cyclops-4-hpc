package syncManager

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"gitlab.com/cyclops-utilities/datamodels"
	pmBundle "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/client/bundle_management"
	pmModels "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
	l "gitlab.com/cyclops-utilities/logging"
)

func (m *SyncManager) updateBundles(param *http.Request) (state int, err error) {

	l.Trace.Printf("[LoaderManager] UpdateBundles endpoint invoked.\n")

	var opts gophercloud.AuthOptions
	var bundles []*pmModels.SkuBundle

	opts = gophercloud.AuthOptions{
		IdentityEndpoint: m.db.Configs.OpenStack.Keystone,
		Username:         m.db.Configs.OpenStack.Username,
		Password:         m.db.Configs.OpenStack.Password,
		DomainName:       m.db.Configs.OpenStack.Domain,
	}

	if len(m.db.Configs.OpenStack.Project) > 0 {

		opts.TenantName = m.db.Configs.OpenStack.Project

	}

	provider, e := openstack.AuthenticatedClient(opts)

	if e != nil {

		l.Error.Printf("[SyncManager] Authentication failed. Error: %v", e)

		state = statusFail
		err = e

		return

	}

	pmClient := m.getPMClient(param)
	ctx := context.Background()

	for _, region := range m.db.Configs.OpenStack.Regions {

		l.Info.Printf("[SyncManager] Processing OS region: %s\n", region)

		client, e := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
			Region: region,
		})

		if e != nil {

			l.Warning.Printf("[SyncManager] There was a problem creating compute collector client. Error: %v", e)

			state = statusFail
			err = e

			return

		}

		listOpts := flavors.ListOpts{
			AccessType: flavors.AllAccess,
		}

		allPages, e := flavors.ListDetail(client, listOpts).AllPages()

		if e != nil {

			l.Warning.Printf("[SyncManager] There was a problem while getting flavor list. Error: %v", e)

			state = statusFail
			err = e

			return

		}

		allFlavors, e := flavors.ExtractFlavors(allPages)

		if e != nil {

			l.Warning.Printf("[SyncManager] Could not extract flavors. Error: %v", e)

			state = statusFail
			err = e

			return

		}

	flavorLoop:

		for _, flavor := range allFlavors {

			name := strings.ToLower(flavor.Name)

			for _, filter := range m.db.Configs.OpenStack.Filters {

				if strings.Contains(flavor.Name, filter) {

					continue flavorLoop

				}

			}

			l.Trace.Printf("[SyncManager] Located flavor: VCPUs: %2d, Disk (GB): %3d, RAM (MB): %5d, Name: %15s, ID:%s\n", flavor.VCPUs, flavor.Disk, flavor.RAM, flavor.Name, flavor.ID)

			//here insert or update flavor details in the db
			var bundle pmModels.SkuBundle

			bundle.ID = flavor.ID
			bundle.Name = &flavor.Name
			bundle.SkuPrices = make(datamodels.JSONdb)

			bundle.SkuPrices["vcpu"] = int64(flavor.VCPUs)
			bundle.SkuPrices["ram"] = float64(float64(flavor.RAM) / float64(1024))

			fmt.Printf("RAM: %v\n", bundle.SkuPrices["ram"])

			if strings.Contains(name, "windows") {

				bundle.SkuPrices["license"] = int64(flavor.VCPUs)

			}

			if strings.HasPrefix(name, "g1") {

				l.Trace.Printf("[SyncManager] Found a GPU flavor: %s, proceeding to parse next.\n", name)

				var temp string

				if strings.Contains(name, "t4") {

					temp = "t4"

				}

				if strings.Contains(name, "p100") {

					temp = "p100"

				}

				if strings.Contains(name, "titanxp") {

					temp = "titanxp"

				}

				typeCut := strings.Split(name, temp)
				coresCut := strings.Split(typeCut[0], "-")

				var gpucount int64

				if coresCut[len(coresCut)-1] == "" {

					gpucount = int64(1)

				} else {

					n, e := strconv.ParseInt(coresCut[len(coresCut)-1], 10, 64)

					if e != nil {

						l.Warning.Printf("[SyncManager] Non numeric count detected in gpu flavor name %s\n", name)

					} else {

						gpucount = n

					}

				}

				bundle.SkuPrices[temp] = int64(gpucount)

			}

			if strings.Contains(name, "ssd") {

				bundle.SkuPrices["rootdisk_ssd"] = float64(flavor.Disk)
				bundle.SkuPrices["ephemeraldisk_ssd"] = float64(flavor.Ephemeral)

			} else {

				bundle.SkuPrices["rootdisk"] = float64(flavor.Disk)
				bundle.SkuPrices["ephemeraldisk"] = float64(flavor.Ephemeral)

			}

			params := pmBundle.NewCreateSkuBundleParams().WithBundle(&bundle)

			_, e = pmClient.BundleManagement.CreateSkuBundle(ctx, params)

			if e != nil {

				uparams := pmBundle.NewUpdateSkuBundleParams().WithID(bundle.ID).WithBundle(&bundle)

				_, e := pmClient.BundleManagement.UpdateSkuBundle(ctx, uparams)

				if e != nil {

					err = e
					state = statusFail

				} else {

					state = statusOK

				}

			}

			bundles = append(bundles, &bundle)

		}

	}

	state = statusOK

	if len(bundles) < 1 {

		state = statusFail

	} else if e != nil {

		state = statusMissing

	}

	return

}
