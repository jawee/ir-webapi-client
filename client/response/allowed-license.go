package response

type AllowedLicense struct {
	ParentID        int    `json:"parent_id"`
	LicenseGroup    int    `json:"license_group"`
	MinLicenseLevel int    `json:"min_license_level"`
	MaxLicenseLevel int    `json:"max_license_level"`
	GroupName       string `json:"group_name"`
}
