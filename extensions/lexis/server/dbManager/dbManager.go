package dbManager

import (
	"context"
	"math"
	"net/http"
	"reflect"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/encoding/json"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/models"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/cacheManager"
	csClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/client"
	csAccount "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/client/account_management"
	csCredit "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/client/credit_management"
	cusClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/client"
	cusReseller "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/client/reseller_management"
	cusModels "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/models"
	pmClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/client"
	udrClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/udr/client"
	l "gitlab.com/cyclops-utilities/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Report status code vars
const (
	scaler           = 1e4
	StatusDuplicated = iota
	StatusFail
	StatusMissing
	StatusOK
)

type Config struct {
	CustomerDB   cusClient.Config
	CreditSystem csClient.Config
	PlanManager  pmClient.Config
	UDR          udrClient.Config
	OpenStack    OpenStackConfig
}

// OpenStackConfig is the struck to hold the configuration parameters needed
// for quering the flavors active in OpenStack.
type OpenStackConfig struct {
	Domain   string
	Filters  []string
	Keystone string
	Password string
	Project  string
	Regions  []string
	Username string
}

// DbParameter is the struct defined to group and contain all the methods
// that interact with the database.
// Parameters:
// - Cache: CacheManager pointer for the cache mechanism.
// - connStr: strings with the connection information to the database
// - Db: a gorm.DB pointer to the db to invoke all the db methods
type DbParameter struct {
	Cache   *cacheManager.CacheManager
	connStr string
	Db      *gorm.DB
	Metrics map[string]*prometheus.GaugeVec
	Configs Config
}

// New is the function to create the struct DbParameter.
// Parameters:
// - dbConn: strings with the connection information to the database
// - tables: array of interfaces that will contains the models to migrate
// to the database on initialization
// Returns:
// - DbParameter: struct to interact with dbManager functionalities
func New(dbConn string, tables ...interface{}) *DbParameter {

	l.Trace.Printf("[DB] Gerenating new DBParameter.\n")

	var (
		dp  DbParameter
		err error
	)

	dp.connStr = dbConn

	dp.Db, err = gorm.Open(postgres.Open(dbConn), &gorm.Config{})

	if err != nil {

		l.Error.Printf("[DB] Error opening connection. Error: %v\n", err)

		panic(err)

	}

	//l.Trace.Printf("[DB] Migrating tables.\n")

	//Database migration, it handles everything
	//dp.Db.AutoMigrate(tables...)

	//l.Trace.Printf("[DB] Generating hypertables.\n")

	// Hypertables creation for timescaledb in case of needed
	//dp.Db.Exec("SELECT create_hypertable('" + dp.Db.NewScope(&models.TABLE).TableName() + "', 'TIMESCALE-ROW-INDEX');")

	return &dp

}

// getCusClient job is to initialize a customerDB client to be able to connect
// to the customerDB service.
// Parameters:
// - http.Request parameter to extract the keycloak bearer token if exists.
// Returns:
// - CustomerDB client struct for connecting to the custoimerDB service.
func (m *DbParameter) getCusClient(param *http.Request) *cusClient.CustomerDatabaseManagement {

	config := m.Configs.CustomerDB

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := cusClient.New(config)

	return r

}

// getCSClient job is to initialize a credit system client to be able to connect
// to the credit system service.
// Parameters:
// - http.Request parameter to extract the keycloak bearer token if exists.
// Returns:
// - CreditSystem client struct for connecting to the creditSystem service.
func (m *DbParameter) getCSClient(param *http.Request) *csClient.CreditManagerManagementAPI {

	config := m.Configs.CreditSystem

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := csClient.New(config)

	return r

}

// getFloat job is to check the type of the interface and properly cast its
// getFloat job is to check the type of the interface and properly cast its
// value into a float64.
// Parameters:
// - i: interface that should contain a float number.
// Returns:
// - f: a float64 with the value contained in the interface provided.
func (m *DbParameter) getFloat(i interface{}) (f float64) {

	var e error

	if i == nil {

		f = float64(0)

		return

	}

	if v := reflect.ValueOf(i); v.Kind() == reflect.Float64 {

		f = i.(float64)

	} else {

		f, e = i.(json.Number).Float64()

		if e != nil {

			l.Trace.Printf("[DB] GetFloat failed to convert [ %v ]. Error: %v\n", i, e)

		}

	}

	return

}

// getNiceFloat job is to turn the ugly float64 provided into a "nice looking"
// one according to the scale set, so we move from a floating coma number into
// a fixed coma number.
// Parameters:
// - i: the ugly floating coma number.
// Returns:
// - o: the "nice looking" fixed coma float.
func (m *DbParameter) getNiceFloat(i float64) (o float64) {

	min := float64(float64(1) / float64(scaler))

	if diff := math.Abs(i - min); diff < min {

		o = float64(0)

	} else {

		o = float64(math.Round(i*scaler) / scaler)

	}

	return

}

// getNiceFloat job is to turn the ugly float64 provided into a "nice looking"
// one according to the scale set, so we move from a floating coma number into
// a fixed coma number.
// Parameters:
// - i: the ugly floating coma number.
// Returns:
// - o: the "nice looking" fixed coma float.
func (m *DbParameter) getNiceFloatAVG(i float64) (o float64) {

	scaler := 1e2

	min := float64(float64(1) / float64(scaler))

	if diff := math.Abs(i - min); diff < min {

		o = float64(0)

	} else {

		o = float64(math.Round(i*scaler) / scaler)

	}

	return

}

func (m *DbParameter) SyncHierarchy(ctx context.Context, param *http.Request) (state int, e error) {

	var resellers []*cusModels.Reseller
	var orgs []*models.Organization
	billingPeriod := "monthly"
	billingCurrency := "EUR"
	billingPlan := "1"

	creditsProjects := make(map[string]float64)

	// get all the organizations
	// - create a reseller entity
	if e = m.Db.Find(&orgs).Error; e != nil {

		l.Warning.Printf("[SYNC] Error retrieving the list of orgs in the system. Error: %v\n", e)

		state = StatusFail

		return

	}

	for _, org := range orgs {

		var customers []*cusModels.Customer
		var prjs []*models.Project

		active := org.OrganizationStatus == "APPROVED"

		reseller := cusModels.Reseller{
			Address:       org.RegisteredAddress1 + "; " + org.RegisteredAddress2 + "; " + org.RegisteredAddress3,
			Billable:      &active,
			BillContact:   org.CreatedBy.String(),
			BillCurrency:  &billingCurrency,
			BillPeriod:    &billingPeriod,
			ContractStart: (strfmt.Date)(org.CreationDate),
			Discount:      float64(0),
			EmailTo:       org.OrganizationEmailAddress,
			IsActive:      &active,
			Name:          org.FormalName,
			PlanID:        billingPlan,
			ResellerID:    org.ID.String(),
		}

		if e = m.Db.Where(models.Project{LinkedOrganization: org.ID}).Find(&prjs).Error; e != nil {

			l.Warning.Printf("[SYNC] Error retrieving the list of prjs for Organization [ %v ] in the system. Error: %v\n", org.ID, e)

			continue

		}

		// for each, get all the project
		// - create a customer entity
		for _, prj := range prjs {

			var products []*cusModels.Product
			var hpcres []*models.HPCResource

			prjActive := prj.ProjectStatus == "ACTIVE"
			customer := cusModels.Customer{
				Billable:      &prjActive,
				BillContact:   prj.ProjectCreatedBy.String(),
				BillCurrency:  &billingCurrency,
				BillPeriod:    &billingPeriod,
				ContractEnd:   (strfmt.Date)(prj.ProjectTerminationDate),
				ContractStart: (strfmt.Date)(prj.ProjectStartDate),
				CustomerID:    prj.ProjectID.String(),
				Discount:      float64(0),
				EmailTo:       prj.ProjectContactEmail,
				IsActive:      &prjActive,
				Name:          prj.ProjectName,
				PlanID:        billingPlan,
				ResellerID:    org.ID.String(),
			}

			creditsProjects[prj.ProjectID.String()] = float64(*prj.NormCoreHours)

			if e = m.Db.Where(models.HPCResource{AssociatedLEXISProject: prj.ProjectID}).Find(&hpcres).Error; e != nil {

				l.Warning.Printf("[SYNC] Error retrieving the list of HPCResources for Project [ %v ] in the system. Error: %v\n", prj.ProjectID, e)

				continue

			}

			// for each get all the HPCRes
			// - create a product entity
			for _, hpc := range hpcres {

				if hpc.ApprovalStatus != "ACCEPTED" {

					l.Info.Printf("[SYNC] HPCResource [ %v ] for Project [ %v ] not accepted yet, skipping...\n", hpc.HPCResourceID, prj.ProjectID)

					continue

				}

				id := hpc.OpenStackProjectID

				if id == "" {

					id = hpc.AssociatedHPCProject

				}

				product := cusModels.Product{
					CustomerID: prj.ProjectID.String(),
					Discount:   float64(0),
					Name:       hpc.HPCProvider + " " + hpc.AssociatedHPCProject,
					ProductID:  id,
					Type:       hpc.HPCProvider + "-" + hpc.ResourceType,
				}

				products = append(products, &product)

			}
			customer.Products = products

			customers = append(customers, &customer)

		}

		reseller.Customers = customers

		resellers = append(resellers, &reseller)

	}

	// For each, first try to create,
	// if fail, then to update
	// if fail, ignore and continue
	customerDB := m.getCusClient(param)
	creditSystem := m.getCSClient(param)

	creditAccounts := make(map[string]bool)

	csParams := csAccount.NewListAccountsParams()
	r, err := creditSystem.AccountManagement.ListAccounts(ctx, csParams)

	if err != nil {

		l.Warning.Printf("[SYNC] Error retrieving the list of Credit accounts in the system. Error: %v\n", e)

	} else {

		acs := r.Payload

		for i := range acs {

			creditAccounts[acs[i].AccountID] = acs[i].Enabled

		}

	}

	for i := range resellers {

		params := cusReseller.NewAddResellerParams().WithReseller(resellers[i])

		_, _, err := customerDB.ResellerManagement.AddReseller(ctx, params)

		if err != nil {

			l.Warning.Printf("[SYNC] Error while adding the data for reseller. Trying with an update... Error: %v", err)

			uParams := cusReseller.NewUpdateResellerParams().WithReseller(resellers[i]).WithID(resellers[i].ResellerID)

			_, _, err := customerDB.ResellerManagement.UpdateReseller(ctx, uParams)

			if err != nil {

				l.Warning.Printf("[SYNC] Error while updating the data for reseller. Error: %v", err)

			}

		}

		// Credit:
		// create account
		// if project active -> credit ac to enable
		// if credit ac previously was diable then add normCorehours as credit amount
		for _, customer := range resellers[i].Customers {

			active, exists := creditAccounts[customer.CustomerID]

			if exists && active {

				continue

			}

			if !exists {

				csParams := csAccount.NewCreateAccountParams().WithID(customer.CustomerID)
				_, err := creditSystem.AccountManagement.CreateAccount(ctx, csParams)

				if err != nil {

					l.Warning.Printf("[SYNC] Error creating the Credit accounts in the system for project [ %v ]. Error: %v\n", customer.CustomerID, err)

					continue

				}

				exists = true
				active = false

			}

			if *customer.Billable {

				if exists && !active {

					csParams := csAccount.NewEnableAccountParams().WithID(customer.CustomerID)
					_, err := creditSystem.AccountManagement.EnableAccount(ctx, csParams)

					if err != nil {

						l.Warning.Printf("[SYNC] Error enabling the Credit accounts in the system for project [ %v ]. Error: %v\n", customer.CustomerID, err)

						continue

					}

					active = true

				}

			} else {

				if exists && active {

					csParams := csAccount.NewDisableAccountParams().WithID(customer.CustomerID)
					_, err := creditSystem.AccountManagement.DisableAccount(ctx, csParams)

					if err != nil {

						l.Warning.Printf("[SYNC] Error disabling the Credit accounts in the system for project [ %v ]. Error: %v\n", customer.CustomerID, err)

						continue

					}

					active = false

				}

			}

			if hours, hasHours := creditsProjects[customer.CustomerID]; exists && active && hasHours {

				csParams := csCredit.NewIncreaseCreditParams().WithID(customer.CustomerID).WithAmount(hours).WithMedium("credit")
				_, err := creditSystem.CreditManagement.IncreaseCredit(ctx, csParams)

				if err != nil {

					l.Warning.Printf("[SYNC] Error setting the Credit for the account for project [ %v ] in the system. Error: %v\n", customer.CustomerID, err)

					continue

				}

			}

		}

	}

	return

}
