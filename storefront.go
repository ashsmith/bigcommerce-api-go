package bigcommerce

type StorefrontService struct {
	Status    StorefrontStatusService
	Seo       StorefrontSeoSettingsService
	Security  StorefrontSecuritySettingsService
	Search    StorefrontSearchSettingsService
	Category  StorefrontCategorySettingsService
	RobotsTxt StorefrontRobotsTxtSettingsService
}
