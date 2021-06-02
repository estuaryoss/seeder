package constants

import "os"

const (
	APP_NAME    = "seeder"
	APP_VERSION = "1.0.0"

	WORKSPACE                   = "workspace"
	DEPLOYMENTS_DIR_BEFORE_INIT = "deployments"
	CONFIG_YAML                 = "config.yaml"
	CONFIG_YAML_AFTER_INIT      = WORKSPACE + string(os.PathSeparator) + "config.yaml"
	DEPLOYMENT_PLAN             = WORKSPACE + string(os.PathSeparator) + "deployment_plan.json"
	DEPLOYMENT_STATE            = WORKSPACE + string(os.PathSeparator) + "deployment_state.json"
	DEPLOYMENT_DIR_AFTER_INIT   = WORKSPACE + string(os.PathSeparator) + DEPLOYMENTS_DIR_BEFORE_INIT

	NA = "NA"
)
