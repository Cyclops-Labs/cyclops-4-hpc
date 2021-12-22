package syncManager

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/models"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/restapi/operations/sync_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/dbManager"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/statusManager"
	cusClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/client"
	pmClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/client"
	l "gitlab.com/cyclops-utilities/logging"
)

const (
	statusDuplicated = iota
	statusFail
	statusMissing
	statusOK
)

// SyncManager is the struct defined to group and contain all the methods
// that interact with the sync subsystem.
type SyncManager struct {
	db    *dbManager.DbParameter
	monit *statusManager.StatusManager
}

var returnValue models.ErrorResponse

// New is the function to create the struct SyncManager.
// Returns:
// - SyncManager: *struct to interact with SyncManager subsystem functionalities.
func New(db *dbManager.DbParameter, monit *statusManager.StatusManager) *SyncManager {

	l.Trace.Printf("[SyncManager] Generating new SyncManager.\n")

	monit.InitEndpoint("sync")

	// Default return string in case something weird happens.
	// It usually means that something went wrong in the dbManager.
	s := "Something unexpected happened, check with the administrator."

	returnValue = models.ErrorResponse{
		ErrorString: &s,
	}

	return &SyncManager{
		db:    db,
		monit: monit,
	}

}

// getCusClient job is to initialize a customerDB client to be able to connect
// to the customerDB service.
// Parameters:
// - http.Request parameter to extract the keycloak bearer token if exists.
// Returns:
// - CustomerDB client struct for connecting to the custoimerDB service.
func (m *SyncManager) getCusClient(param *http.Request) *cusClient.CustomerDatabaseManagement {

	config := m.db.Configs.CustomerDB

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := cusClient.New(config)

	return r

}

// getPMClient job is to initialize a planManager client to be able to connect
// to the planManager service.
// Parameters:
// - http.Request parameter to extract the keycloak bearer token if exists.
// Returns:
// - PlanManager client struct for connecting to the PlanManager service.
func (m *SyncManager) getPMClient(param *http.Request) *pmClient.PlanManagerManagementAPI {

	config := m.db.Configs.PlanManager

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := pmClient.New(config)

	return r

}

// SyncHierarchy (Swagger func) is the function behind the (GET) endpoint
// /load/customers
// Its job is to process the data from SWITCH coscenters endpoint and load
// them into Cyclops.
func (m *SyncManager) SyncHierarchy(ctx context.Context, params sync_management.SyncHierarchyParams) middleware.Responder {

	l.Trace.Printf("[SyncManager] Starting the loading of customers.\n")

	callTime := time.Now()
	m.monit.APIHit("sync", callTime)

	state, e := m.db.SyncHierarchy(ctx, params.HTTPRequest)

	switch state {

	case statusOK:

		m.db.Metrics["api"].With(prometheus.Labels{"code": fmt.Sprint(http.StatusOK), "method": params.HTTPRequest.Method, "route": params.HTTPRequest.URL.Path}).Inc()

		m.monit.APIHitDone("sync", callTime)

		return sync_management.NewSyncHierarchyOK()

	case statusFail:

		err := fmt.Sprintf("Something went wrong with the Hierarchy processing. Error: %v", e)

		returnValueError := models.ErrorResponse{
			ErrorString: &err,
		}

		m.db.Metrics["api"].With(prometheus.Labels{"code": fmt.Sprint(http.StatusInternalServerError), "method": params.HTTPRequest.Method, "route": params.HTTPRequest.URL.Path}).Inc()

		m.monit.APIHitDone("sync", callTime)

		return sync_management.NewSyncHierarchyInternalServerError().WithPayload(&returnValueError)

	}

	m.db.Metrics["api"].With(prometheus.Labels{"code": fmt.Sprint(http.StatusAccepted), "method": params.HTTPRequest.Method, "route": params.HTTPRequest.URL.Path}).Inc()

	m.monit.APIHitDone("sync", callTime)

	return sync_management.NewSyncHierarchyAccepted()

}

// SyncFlavors (Swagger func) is the function behind the (GET) endpoint
// /load/flavors
// Its job is to process the data from OppenStack flavors in SWITCH and load
// them into Cyclops.
func (m *SyncManager) SyncFlavors(ctx context.Context, params sync_management.SyncFlavorsParams) middleware.Responder {

	l.Trace.Printf("[SyncManager] Starting the loading of flavors.\n")

	callTime := time.Now()
	m.monit.APIHit("sync", callTime)

	state, e := m.updateBundles(params.HTTPRequest)

	switch state {

	case statusOK:

		m.db.Metrics["api"].With(prometheus.Labels{"code": fmt.Sprint(http.StatusOK), "method": params.HTTPRequest.Method, "route": params.HTTPRequest.URL.Path}).Inc()

		m.monit.APIHitDone("sync", callTime)

		return sync_management.NewSyncFlavorsOK()

	case statusFail:

		err := fmt.Sprintf("Something went wrong with the Flavors processing. Error: %v", e)

		returnValueError := models.ErrorResponse{
			ErrorString: &err,
		}

		m.db.Metrics["api"].With(prometheus.Labels{"code": fmt.Sprint(http.StatusInternalServerError), "method": params.HTTPRequest.Method, "route": params.HTTPRequest.URL.Path}).Inc()

		m.monit.APIHitDone("sync", callTime)

		return sync_management.NewSyncFlavorsInternalServerError().WithPayload(&returnValueError)

	}

	m.db.Metrics["api"].With(prometheus.Labels{"code": fmt.Sprint(http.StatusAccepted), "method": params.HTTPRequest.Method, "route": params.HTTPRequest.URL.Path}).Inc()

	m.monit.APIHitDone("sync", callTime)

	return sync_management.NewSyncFlavorsAccepted()

}
