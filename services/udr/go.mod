module github.com/Cyclops-Labs/cyclops-4-hpc.git/services/udr

go 1.15

require (
	github.com/Nerzal/gocloak/v7 v7.11.0
	github.com/go-openapi/errors v0.20.1
	github.com/go-openapi/loads v0.21.0
	github.com/go-openapi/runtime v0.21.0
	github.com/go-openapi/spec v0.20.4
	github.com/go-openapi/strfmt v0.21.1
	github.com/go-openapi/swag v0.19.15
	github.com/go-openapi/validate v0.20.3
	github.com/prometheus/client_golang v1.11.0
	github.com/remeh/sizedwaitgroup v1.0.0
	github.com/rs/cors v1.8.2
	github.com/segmentio/encoding v0.3.2
	github.com/segmentio/kafka-go v0.4.25
	github.com/spf13/viper v1.10.1
	gitlab.com/cyclops-utilities/datamodels v0.0.0-20191016132854-e9313e683e5b
	github.com/Cyclops-Labs/cyclops-4-hpc.git/services/cdr v0.0.2
	github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb v0.0.2
	github.com/Cyclops-Labs/cyclops-4-hpc.git/services/eventsengine v0.0.2
	github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager v0.0.2
	gitlab.com/cyclops-utilities/logging v0.0.0-20200914110347-ca1d02efd346
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.4
)
