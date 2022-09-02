package api_version

type ApiVersionInfo struct {
	MajorVersion    int   `json:"major"`
	MinorVersion    int   `json:"minor"`
	SupportedMajors []int `json:"supportedMajors"`
}

var ApiVersion = ApiVersionInfo{
	MajorVersion:    2,
	MinorVersion:    1,
	SupportedMajors: []int{1, 2},
}
