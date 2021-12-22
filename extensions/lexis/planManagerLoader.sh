#!/bin/bash

SERVER="http://localhost"
API_KEY="1234567890abcdefghi"
API_VERSION="v1.0"

# We load the plans
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/plan" -X POST -d '{ "id":"1", "name": "LEXIS_1", "offeredstartdate":"2019-01-01", "offeredenddate":"2040-12-31" }'

# We load the skus
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "1", "name":              "vcpu", "unit":         "Core" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "2", "name":               "ram", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "3", "name":          "rootdisk", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "4", "name":     "ephemeraldisk", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "5", "name":        "floatingip", "unit":           "IP" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "6", "name":      "blockstorage", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "7", "name":     "objectstorage", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "8", "name":           "license", "unit": "License*Core" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id":  "9", "name":           "titanxp", "unit":         "Core" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "10", "name":                "t4", "unit":         "Core" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "11", "name":              "p100", "unit":         "Core" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "12", "name":      "rootdisk_ssd", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "13", "name": "ephemeraldisk_ssd", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "14", "name":  "blockstorage_ssd", "unit":           "GB" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "15", "name":           "salomon", "unit":    "core-hour" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "16", "name":           "barbora", "unit":    "core-hour" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "17", "name":       "barbora-gpu", "unit":    "core-hour" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku" -X POST -d '{ "id": "18", "name":               "dgx", "unit":    "core-hour" }'

# We load the sku prices
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "1", "skuname":              "vcpu", "unitprice":  0.000011574, "unit":         "Core", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "2", "skuname":               "ram", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "3", "skuname":          "rootdisk", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "4", "skuname":     "ephemeraldisk", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "5", "skuname":        "floatingip", "unitprice":  0.000011574, "unit":           "IP", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "6", "skuname":      "blockstorage", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "7", "skuname":     "objectstorage", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "8", "skuname":           "license", "unitprice":  0.000011574, "unit": "License*Core", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid":  "9", "skuname":           "titanxp", "unitprice":  0.000011574, "unit":         "Core", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "10", "skuname":                "t4", "unitprice":  0.000011574, "unit":         "Core", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "11", "skuname":              "p100", "unitprice":  0.000011574, "unit":         "Core", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "12", "skuname":      "rootdisk_ssd", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "13", "skuname": "ephemeraldisk_ssd", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "14", "skuname":  "blockstorage_ssd", "unitprice":  0.000011574, "unit":           "GB", "planid": "1", "UnitCreditPrice": 0.000011574, "AccountingMode": "CASH" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "15", "skuname":           "salomon", "unitprice":  0.000013889, "unit":    "core-hour", "planid": "1", "UnitCreditPrice": 0.000013889, "AccountingMode": "CREDIT" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "16", "skuname":           "barbora", "unitprice":  0.000020694, "unit":    "core-hour", "planid": "1", "UnitCreditPrice": 0.000020694, "AccountingMode": "CREDIT" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "17", "skuname":       "barbora-gpu", "unitprice":  0.000066111, "unit":    "core-hour", "planid": "1", "UnitCreditPrice": 0.000066111, "AccountingMode": "CREDIT" }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/sku/price" -X POST -d '{ "skuid": "18", "skuname":               "dgx", "unitprice":  0.000161667, "unit":    "core-hour", "planid": "1", "UnitCreditPrice": 0.000161667, "AccountingMode": "CREDIT" }'


# We load the life cycles
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":     "active", "resourceType":  "blockstorage", "skuList":{ "blockstorage":  1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":   "inactive", "resourceType":  "blockstorage", "skuList":{ "blockstorage":  1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":     "active", "resourceType":    "floatingip", "skuList":{ "floatingip":    1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":   "inactive", "resourceType":    "floatingip", "skuList":{ "floatingip":    1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":       "used", "resourceType": "objectstorage", "skuList":{ "objectstorage": 1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":       "used", "resourceType":       "salomon", "skuList":{ "salomon":       1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":       "used", "resourceType":       "barbora", "skuList":{ "barbora":       1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":       "used", "resourceType":   "barbora-gpu", "skuList":{ "barbora-gpu":   1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":       "used", "resourceType":           "dgx", "skuList":{ "dgx":           1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":     "active", "resourceType":        "server", "skuList":{ "vcpu":    1, "ram":     1, "titanxp": 1, "t4":       1, "p100":         1, "rootdisk":      1, "rootdisk_ssd":      1, "ephemeraldisk":     1, "ephemeraldisk_ssd": 1, "license": 1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":   "inactive", "resourceType":        "server", "skuList":{ "titanxp": 1, "t4":      1, "p100":    1, "rootdisk": 1, "rootdisk_ssd": 1, "ephemeraldisk": 1, "ephemeraldisk_ssd": 1, "license":           1 } }'
curl --silent -H "X-API-Key: ${API_KEY}" -H "Content-Type: application/json" "${SERVER}:8600/api/${API_VERSION}/cycle" -X POST -d  '{ "state":  "suspended", "resourceType":        "server", "skuList":{ "ram":     1, "titanxp": 1, "t4":      1, "p100":     1, "rootdisk":     1, "rootdisk_ssd":  1, "ephemeraldisk":     1, "ephemeraldisk_ssd": 1, "license":           1 } }'

