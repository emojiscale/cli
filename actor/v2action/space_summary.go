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

func (actor Actor) GetSpaceSummaryByOrganizationAndName(orgGUID string, name string) (SpaceSummary, Warnings, error) {
	var allWarnings Warnings

	org, warnings, err := actor.GetOrganization(orgGUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return SpaceSummary{}, nil, nil
	}

	spaceSummary := SpaceSummary{
		SpaceName: name,
		OrgName:   org.Name,
	}

	return spaceSummary, allWarnings, nil
}
