package ccv2

type SpaceQuota struct {
	GUID string
	Name string
}

func (client *Client) GetSpaceQuota(guid string) (SpaceQuota, Warnings, error) {
	return SpaceQuota{}, nil, nil
}