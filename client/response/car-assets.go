package response

type CarAssets struct {
	CarID                  int        `json:"car_id"`
	CarRules               []CarRules `json:"car_rules"`
	DetailCopy             string     `json:"detail_copy"`
	DetailScreenShotImages string     `json:"detail_screen_shot_images"`
	DetailTechspecsCopy    string     `json:"detail_techspecs_copy"`
	Folder                 string     `json:"folder"`
	GalleryImages          string     `json:"gallery_images"`
	GalleryPrefix          string     `json:"gallery_prefix"`
	GroupImage             string     `json:"group_image"`
	GroupName              string     `json:"group_name"`
	LargeImage             string     `json:"large_image"`
	Logo                   string     `json:"logo"`
	SmallImage             string     `json:"small_image"`
	SponsorLogo            string     `json:"sponsor_logo"`
	TemplatePath           string     `json:"template_path"`
}
