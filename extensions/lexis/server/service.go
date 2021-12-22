package main

import (
	"net/http"
	"net/url"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/restapi"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/dbManager"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/statusManager"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/syncManager"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/triggerManager"
	csClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/client"
	cusClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/client"
	pmClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/client"
	udrClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/udr/client"
	l "gitlab.com/cyclops-utilities/logging"
)

// getService is the uService instantiation function. A sample of the elements
// that need to be personalized for a new uService are provided.
// Returns:
// - hadler: return the http.handler with the swagger and (optionally) the CORS
// configured according to the config provided
// - e: in case of any error it's propagated to the invoker
func getService() (handler http.Handler, register *prometheus.Registry, e error) {

	l.Trace.Printf("[MAIN] Intializing service [ %v ] handler\n", strings.Title(service))

	// TABLE0,1,n have to be customized
	db := dbStart()
	mon := statusManager.New(db)

	// Prometheus Metrics linked to dbParameter
	db.Metrics, register = prometheusStart()

	// cache linked to the dbParameter
	db.Cache = cacheStart(db.Metrics["cache"])

	// In case of needed, here is where kafka support is started
	// The functions needs to be customized accordingly to the needs of the service
	// And the parts of the service that might need to send messages need to get
	// the channel
	// _ = kafkaStart(db, mon)

	//bp := getBasePath()

	db.Configs.CreditSystem = csClient.Config{
		URL: &url.URL{
			Host:   cfg.General.Services["creditsystem"],
			Path:   cusClient.DefaultBasePath,
			Scheme: "http",
		},
		AuthInfo: httptransport.APIKeyAuth(cfg.APIKey.Key, cfg.APIKey.Place, cfg.APIKey.Token),
	}

	db.Configs.CustomerDB = cusClient.Config{
		URL: &url.URL{
			Host:   cfg.General.Services["customerdb"],
			Path:   cusClient.DefaultBasePath,
			Scheme: "http",
		},
		AuthInfo: httptransport.APIKeyAuth(cfg.APIKey.Key, cfg.APIKey.Place, cfg.APIKey.Token),
	}

	db.Configs.PlanManager = pmClient.Config{
		URL: &url.URL{
			Host:   cfg.General.Services["planmanager"],
			Path:   pmClient.DefaultBasePath,
			Scheme: "http",
		},
		AuthInfo: httptransport.APIKeyAuth(cfg.APIKey.Key, cfg.APIKey.Place, cfg.APIKey.Token),
	}

	db.Configs.UDR = udrClient.Config{
		URL: &url.URL{
			Host:   cfg.General.Services["udr"],
			Path:   pmClient.DefaultBasePath,
			Scheme: "http",
		},
		AuthInfo: httptransport.APIKeyAuth(cfg.APIKey.Key, cfg.APIKey.Place, cfg.APIKey.Token),
	}

	db.Configs.OpenStack = dbManager.OpenStackConfig{
		Domain:   cfg.OpenStack.Domain,
		Filters:  cfg.OpenStack.Filters,
		Keystone: cfg.OpenStack.Keystone,
		Password: cfg.OpenStack.Password,
		Project:  cfg.OpenStack.Project,
		Regions:  cfg.OpenStack.Regions,
		Username: cfg.OpenStack.User,
	}

	// Parts of the service HERE
	syncer := syncManager.New(db, mon)
	trigger := triggerManager.New(db, mon)

	// Initiate the http handler, with the objects that are implementing the business logic.
	h, e := restapi.Handler(restapi.Config{
		StatusManagementAPI:  mon,
		SyncManagementAPI:    syncer,
		TriggerManagementAPI: trigger,
		Logger:               l.Info.Printf,
		AuthKeycloak:         AuthKeycloak,
		AuthAPIKeyHeader:     AuthAPIKey,
		AuthAPIKeyParam:      AuthAPIKey,
	})

	if e != nil {

		return

	}

	handler = h

	// CORS
	if cfg.General.CORSEnabled {

		l.Trace.Printf("[MAIN] Enabling CORS for the service [ %v ]\n", strings.Title(service))

		handler = cors.New(
			cors.Options{
				Debug:          (cfg.General.LogLevel == "DEBUG") || (cfg.General.LogLevel == "TRACE"),
				AllowedOrigins: cfg.General.CORSOrigins,
				AllowedHeaders: cfg.General.CORSHeaders,
				AllowedMethods: cfg.General.CORSMethods,
			}).Handler(h)

	}

	return

}

// kafkaStart handles the initialization of the kafka service.
// This is a sample function with the most basic usage of the kafka service, it
// should be redefined to match the needs of the service.
// Parameters:
// - db: DbPAramenter reference for interaction with the db.
// - mon: statusManager parameter reference to interact with the statusManager
// subsystem
// Returns:
// - ch: a interface{} channel to be able to send thing through the kafka topic generated.
func kafkaStart(db *dbManager.DbParameter, mon *statusManager.StatusManager) (ch chan interface{}) {

	l.Trace.Printf("[MAIN] Intializing Kafka\n")

	f := func(db *dbManager.DbParameter, model interface{}) (e error) {

		////if e = db.ProcessCDR(*model.(*cdrModels.CReport), ""); e != nil {

		////	l.Warning.Printf("[KAFKA] There was an error while apliying the linked function. Error: %v\n", e)

		////} else {

		////	l.Trace.Printf("[KAFKA] Aplication of the linked function done succesfully.\n")

		////}

		return

	}

	ch = make(chan interface{})

	handler := kafkaHandlerConf{
		in: []kafkaPackage{
			{
				topic:    cfg.Kafka.TopicsIn[0],
				model:    nil,
				function: f,
				saveDB:   false,
			},
		},
	}

	kafkaHandler(db, mon, handler)

	return

}
