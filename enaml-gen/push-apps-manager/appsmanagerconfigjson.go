package push_apps_manager 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type AppsManagerConfigJson struct {

	/*NavLinks - Descr: White labeling for navigation links Default: [map[name:Marketplace href:/marketplace] map[name:Docs href:http://docs.run.pivotal.io] map[name:Tools href:/tools]]
*/
	NavLinks interface{} `yaml:"nav_links,omitempty"`

	/*ProductName - Descr: White labeling for product name Default: Apps Manager
*/
	ProductName interface{} `yaml:"product_name,omitempty"`

	/*DisplayPlanPrices - Descr: Display Marketplace Service Plan Prices Default: false
*/
	DisplayPlanPrices interface{} `yaml:"display_plan_prices,omitempty"`

	/*MarketplaceName - Descr: White labeling for marketplace name Default: Marketplace
*/
	MarketplaceName interface{} `yaml:"marketplace_name,omitempty"`

}