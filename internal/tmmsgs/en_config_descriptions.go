// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tmmsgs

import (
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"golang.org/x/text/language"
)

var ffc = func(key, translation, fieldType string) i18n.ConfigMessageKey {
	return i18n.FFC(language.AmericanEnglish, key, translation, fieldType)
}

//revive:disable
var (
	ConfigAPIDefaultRequestTimeout = ffc("config.api.defaultRequestTimeout", "Default server-side request timeout for API calls", i18n.TimeDurationType)
	ConfigAPIMaxRequestTimeout     = ffc("config.api.maxRequestTimeout", "Maximum server-side request timeout a caller can request with a Request-Timeout header", i18n.TimeDurationType)
	ConfigAPIAddress               = ffc("config.api.address", "Listener address for API", i18n.StringType)
	ConfigAPIPort                  = ffc("config.api.port", "Listener port for API", i18n.IntType)
	ConfigAPIPublicURL             = ffc("config.api.publicURL", "External address callers should access API over", i18n.StringType)
	ConfigAPIReadTimeout           = ffc("config.api.readTimeout", "The maximum time to wait when reading from an HTTP connection", i18n.TimeDurationType)
	ConfigAPIWriteTimeout          = ffc("config.api.writeTimeout", "The maximum time to wait when writing to a HTTP connection", i18n.TimeDurationType)
	ConfigAPIShutdownTimeout       = ffc("config.api.shutdownTimeout", "The maximum amount of time to wait for any open HTTP requests to finish before shutting down the HTTP server", i18n.TimeDurationType)

	ConfigConfirmationsBlockCacheSize           = ffc("config.confirmations.blockCacheSize", "The maximum number of block headers to keep in the cache", i18n.IntType)
	ConfigConfirmationsBlockPollingInterval     = ffc("config.confirmations.blockPollingInterval", "How often to poll for new block headers", i18n.TimeDurationType)
	ConfigConfirmationsNotificationsQueueLength = ffc("config.confirmations.notificationQueueLength", "Internal queue length for notifying the confirmations manager of new transactions/events", i18n.IntType)
	ConfigConfirmationsRequired                 = ffc("config.confirmations.required", "Number of confirmations required to consider a transaction/event final", i18n.IntType)
	ConfigConfirmationsStaleReceiptTimeout      = ffc("config.confirmations.staleReceiptTimeout", "Duration after which to force a receipt check for a pending transaction", i18n.TimeDurationType)

	ConfigFFCoreURL      = ffc("config.ffcore.url", "The URL of the FireFly core admin API server to connect to", i18n.StringType)
	ConfigFFCoreProxyURL = ffc("config.ffcore.proxy.url", "Optional HTTP proxy URL to use for the FireFly core admin API server", i18n.StringType)

	ConfigManagerName = ffc("config.manager.name", "The name of this Transaction Manager, used in operation metadata to track which operations are to be updated", i18n.StringType)

	ConfigOperationsTypes                     = ffc("config.operations.types", "The operation types to query in FireFly core, that might have been submitted via this Transaction Manager", "string[]")
	ConfigOperationsFullScanMinimumDelay      = ffc("config.operations.fullScan.minimumDelay", "The minimum delay between full scans of the FireFly core API, when reconnecting, or recovering from missed events / errors", i18n.TimeDurationType)
	ConfigOperationsFullScanPageSize          = ffc("config.operations.fullScan.pageSize", "The page size to use when performing a full scan of the ForeFly core API on startup, or recovery", i18n.IntType)
	ConfigOperationsFullScanStartupMaxRetries = ffc("config.operations.fullScan.startupMaxRetries", "The page size to use when performing a full scan of the ForeFly core API on startup, or recovery", i18n.IntType)
	ConfigOperationsErrorHistoryCount         = ffc("config.operations.errorHistoryCount", "The number of historical errors to retain in the operation", i18n.IntType)
	ConfigOperationsChangeListenerEnabled     = ffc("config.operations.changeListener.enabled", "Whether to enable the change event listener to detect updates made to operations outside of the FFTM", i18n.BooleanType)

	ConfigPolicyEngineName = ffc("config.policyengine.name", "The name of the policy engine to use", i18n.StringType)

	ConfigLoopInterval = ffc("config.policyloop.interval", "Interval at which to invoke the policy engine to evaluate outstanding transactions", i18n.TimeDurationType)

	ConfigPolicyEngineSimpleFixedGasPrice          = ffc("config.policyengine.simple.fixedGasPrice", "A fixed gasPrice value/structure to pass to the connector", "Raw JSON")
	ConfigPolicyEngineSimpleWarnInterval           = ffc("config.policyengine.simple.warnInterval", "The time between warnings when a blockchain transaction has not been allocated a receipt", i18n.TimeDurationType)
	ConfigPolicyEngineSimpleGasOracleEnabled       = ffc("config.policyengine.simple.gasOracle.mode", "The gas oracle mode", "connector | restapi | disabled")
	ConfigPolicyEngineSimpleGasOracleGoTemplate    = ffc("config.policyengine.simple.gasOracle.template", "REST API Gas Oracle: A go template to execute against the result from the Gas Oracle, to create a JSON block that will be passed as the gas price to the connector", i18n.GoTemplateType)
	ConfigPolicyEngineSimpleGasOracleURL           = ffc("config.policyengine.simple.gasOracle.url", "REST API Gas Oracle: The URL of a Gas Oracle REST API to call", i18n.StringType)
	ConfigPolicyEngineSimpleGasOracleProxyURL      = ffc("config.policyengine.simple.gasOracle.proxy.url", "Optional HTTP proxy URL to use for the Gas Oracle REST API", i18n.StringType)
	ConfigPolicyEngineSimpleGasOracleMethod        = ffc("config.policyengine.simple.gasOracle.method", "The HTTP Method to use when invoking the Gas Oracle REST API", i18n.StringType)
	ConfigPolicyEngineSimpleGasOracleQueryInterval = ffc("config.policyengine.simple.gasOracle.queryInterval", "The minimum interval between queries to the Gas Oracle", i18n.TimeDurationType)

	ConfigEventStreamsDefaultsBatchSize                 = ffc("config.eventstreams.defaults.batchSize", "Default batch size for newly created event streams", i18n.IntType)
	ConfigEventStreamsDefaultsBatchTimeout              = ffc("config.eventstreams.defaults.batchTimeout", "Default batch timeout for newly created event streams", i18n.TimeDurationType)
	ConfigEventStreamsDefaultsErrorHandling             = ffc("config.eventstreams.defaults.errorHandling", "Default error handling for newly created event streams", "'skip' or 'block'")
	ConfigEventStreamsDefaultsRetryTimeout              = ffc("config.eventstreams.defaults.retryTimeout", "Default retry timeout for newly created event streams", i18n.TimeDurationType)
	ConfigEventStreamsDefaultsBlockedRetryDelay         = ffc("config.eventstreams.defaults.blockedRetryDelay", "Default blocked retry delay for newly created event streams", i18n.TimeDurationType)
	ConfigEventStreamsDefaultsWebhookRequestTimeout     = ffc("config.eventstreams.defaults.webhookRequestTimeout", "Default WebHook request timeout for newly created event streams", i18n.TimeDurationType)
	ConfigEventStreamsDefaultsWebsocketDistributionMode = ffc("config.eventstreams.defaults.websocketDistributionMode", "Default WebSocket distribution mode for newly created event streams", "'load_balance' or 'broadcast'")
	ConfigEventStreamsRetryInitDelay                    = ffc("config.eventstreams.retry.initialDelay", "Initial retry delay", i18n.TimeDurationType)
	ConfigEventStreamsRetryMaxDelay                     = ffc("config.eventstreams.retry.maxDelay", "Maximum delay between retries", i18n.TimeDurationType)
	ConfigEventStreamsRetryFactor                       = ffc("config.eventstreams.retry.factor", "Factor to increase the delay by, between each retry", i18n.FloatType)

	ConfigPersistenceType              = ffc("config.persistence.type", "The type of persistence to use", "Only 'leveldb' currently supported")
	ConfigPersistenceLevelDBPath       = ffc("config.persistence.leveldb.path", "The path for the LevelDB persistence directory", i18n.StringType)
	ConfigPersistenceLevelDBMaxHandles = ffc("config.persistence.leveldb.maxHandles", "The maximum number of cached file handles LevelDB should keep open", i18n.IntType)
	ConfigPersistenceLevelDBSyncWrites = ffc("config.persistence.leveldb.syncWrites", "Whether to synchronously perform writes to the storage", i18n.BooleanType)

	ConfigWebhooksAllowPrivateIPs = ffc("config.webhooks.allowPrivateIPs", "Whether to allow WebHook URLs that resolve to Private IP address ranges (vs. internet addresses)", i18n.BooleanType)
	ConfigWebhooksURL             = ffc("config.webhooks.url", "Unused (overridden by the WebHook configuration of an individual event stream)", i18n.IgnoredType)
	ConfigWebhooksProxyURL        = ffc("config.webhooks.proxy.url", "Optional HTTP proxy to use when invoking WebHooks", i18n.StringType)
)
