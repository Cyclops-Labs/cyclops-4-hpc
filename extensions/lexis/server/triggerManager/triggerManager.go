package triggerManager

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/models"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/restapi/operations/trigger_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/dbManager"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/server/statusManager"
	udrClient "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/udr/client"
	udrTrigger "github.com/Cyclops-Labs/cyclops-4-hpc.git/services/udr/client/trigger_management"
	l "gitlab.com/cyclops-utilities/logging"
)

const (
	scaler           = 1e4
	statusDuplicated = iota
	statusFail
	statusMissing
	statusOK
)

type TriggerManager struct {
	db    *dbManager.DbParameter
	monit *statusManager.StatusManager
}

var returnValue models.ErrorResponse

// New is the function to create the struct TriggerManager.
// Returns:
// - TriggerManager: *struct to interact with LoaderManager subsystem functionalities.
func New(db *dbManager.DbParameter, monit *statusManager.StatusManager) *TriggerManager {

	l.Trace.Printf("[TriggerManager] Generating new triggerManager.\n")

	monit.InitEndpoint("trigger")

	// Default return string in case something weird happens.
	// It usually means that something went wrong in the dbManager.
	s := "Something unexpected happened, check with the administrator."

	returnValue = models.ErrorResponse{
		ErrorString: &s,
	}

	return &TriggerManager{
		db:    db,
		monit: monit,
	}

}

// getToken job is to retrieve the keycloak bearer token from the Http.Request.
// Parameters:
// - http.Request parameter to extract the keycloak bearer token if exists.
// Returns:
// - string with the optional Keycloak bearer token.
func (m *TriggerManager) getToken(param *http.Request) (token string) {

	if len(param.Header.Get("Authorization")) > 0 {

		token = strings.Fields(param.Header.Get("Authorization"))[1]

	}

	return

}

func (m *TriggerManager) UDRRedo(ctx context.Context, params trigger_management.UDRRedoParams) middleware.Responder {

	l.Trace.Printf("[Trigger] Generate Reports endpoint invoked.\n")

	callTime := time.Now()
	m.monit.APIHit("trigger", callTime)

	go func() {

		client := udrClient.New(m.db.Configs.UDR)

		increment, _ := time.ParseDuration(*params.Interval)

		f := (time.Time)(*params.From)
		end := (time.Time)(*params.To)

		from := strfmt.DateTime(f)
		f = f.Add(increment)
		to := strfmt.DateTime(f)

		for end.Sub(f.Add(-increment)) >= increment {

			l.Trace.Printf("UDR Compact requested for %v to %v", from, to)

			params := udrTrigger.NewExecCompactationParams().WithFrom(&from).WithTo(&to)
			ctx = context.Background()

			_, e := client.TriggerManagement.ExecCompactation(ctx, params)

			if e != nil {

				l.Warning.Printf("There was a problem while compacting UDR from [ %v ], to [ %v ]. Error: %v", from, to, e)

			} else {

				l.Trace.Printf("UDR Compact completed for %v to %v", from, to)

			}

			from = to
			f = f.Add(increment)
			to = strfmt.DateTime(f)

			time.Sleep(1 * time.Second)

		}

		m.monit.APIHitDone("trigger", callTime)

	}()

	returnValue := models.ItemCreatedResponse{
		Message: "Report generation started",
	}

	return trigger_management.NewUDRRedoOK().WithPayload(&returnValue)

}
