package models

type ApiResponse struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Description interface{} `json:"description"`
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Timestamp   string      `json:"timestamp"`
	Path        string      `json:"path"`
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{}
}

func (apiResponse *ApiResponse) GetCode() int {
	return apiResponse.Code
}

func (apiResponse *ApiResponse) SetCode(code int) {
	apiResponse.Code = code
}

func (apiResponse *ApiResponse) GetMessage() string {
	return apiResponse.Message
}

func (apiResponse *ApiResponse) SetMessage(message string) {
	apiResponse.Message = message
}

func (apiResponse *ApiResponse) GetName() string {
	return apiResponse.Name
}

func (apiResponse *ApiResponse) SetName(name string) {
	apiResponse.Name = name
}

func (apiResponse *ApiResponse) GetDescription() interface{} {
	return apiResponse.Description
}

func (apiResponse *ApiResponse) SetDescription(description interface{}) {
	apiResponse.Description = description
}

func (apiResponse *ApiResponse) GetVersion() string {
	return apiResponse.Version
}

func (apiResponse *ApiResponse) SetVersion(version string) {
	apiResponse.Version = version
}

func (apiResponse *ApiResponse) GetTimestamp() string {
	return apiResponse.Timestamp
}

func (apiResponse *ApiResponse) SetTimestamp(ts string) {
	apiResponse.Timestamp = ts
}

func (apiResponse *ApiResponse) GetPath() string {
	return apiResponse.Version
}

func (apiResponse *ApiResponse) SetPath(path string) {
	apiResponse.Path = path
}
