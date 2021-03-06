package ccv3

import (
	"bytes"
	"encoding/json"
	"net/url"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
)

// IsolationSegment represents a Cloud Controller Isolation Segment.
type IsolationSegment struct {
	Name string `json:"name"`
	GUID string `json:"guid,omitempty"`
}

// CreateIsolationSegment will create an Isolation Segment on the Cloud
// Controller. Note: This will not validate that the placement tag exists in
// the diego cluster.
func (client *Client) CreateIsolationSegment(name string) (IsolationSegment, Warnings, error) {
	body, err := json.Marshal(IsolationSegment{Name: name})
	if err != nil {
		return IsolationSegment{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.NewIsolationSegmentRequest,
		Body:        bytes.NewBuffer(body),
	})
	if err != nil {
		return IsolationSegment{}, nil, err
	}

	var isolationSegment IsolationSegment
	response := cloudcontroller.Response{
		Result: &isolationSegment,
	}

	err = client.connection.Make(request, &response)
	return isolationSegment, response.Warnings, err
}

// GetIsolationSegments lists applications with optional filters.
func (client *Client) GetIsolationSegments(query url.Values) ([]IsolationSegment, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetIsolationSegmentsRequest,
		Query:       query,
	})
	if err != nil {
		return nil, nil, err
	}

	var fullIsolationSegmentsList []IsolationSegment
	warnings, err := client.paginate(request, IsolationSegment{}, func(item interface{}) error {
		if isolationSegment, ok := item.(IsolationSegment); ok {
			fullIsolationSegmentsList = append(fullIsolationSegmentsList, isolationSegment)
		} else {
			return cloudcontroller.UnknownObjectInListError{
				Expected:   IsolationSegment{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullIsolationSegmentsList, warnings, err
}

// DeleteIsolationSegment removes an isolation segment from the cloud
// controller. Note: This will only remove it from the cloud controller
// database. It will not remove it from diego.
func (client *Client) DeleteIsolationSegment(guid string) (Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.DeleteIsolationSegmentRequest,
		URIParams:   map[string]string{"guid": guid},
	})
	if err != nil {
		return nil, err
	}

	var response cloudcontroller.Response
	err = client.connection.Make(request, &response)
	return response.Warnings, err
}
