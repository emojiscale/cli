package v2action

type SpaceSummary struct {
	SpaceName            string
	OrgName              string
	AppNames             []string
	ServiceInstanceNames []string
	SpaceQuotaName       string
	SecurityGroupNames   []string
	SecurityGroupRules   []SecurityGroupRule
}

type SecurityGroupRule struct {
	Description string
	Destination string
	Lifecycle   string
	Name        string
	Port        int
	Protocol    string
}

// func (actor Actor) GetSpaceSummaryByOrganizationAndName(orgGUID string, name string) (SpaceSummary, Warnings, error) {

// }
