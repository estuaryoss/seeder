package constants

const (
	APP_NAME    = "seeder"
	APP_VERSION = "1.0.0"

	WORKSPACE                   = "workspace"
	DEPLOYMENTS_DIR_BEFORE_INIT = "deployments"
	CONFIG_YAML                 = "config.yaml"
	DEPLOYMENT_PLAN             = WORKSPACE + "/deployment_plan.json"
	DEPLOYMENT_STATE            = WORKSPACE + "/deployment_state.json"
	DEPLOYMENT_DIR_AFTER_INIT   = WORKSPACE + "/" + DEPLOYMENTS_DIR_BEFORE_INIT
)
