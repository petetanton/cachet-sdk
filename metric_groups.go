package cachet

import (
	"fmt"
)

const (
	// MetricGroupVisibilityPublic means "Viewable by public"
	MetricGroupVisibilityPublic = 1
	// MetricGroupVisibilityLoggedIn means "Only visible to logged in users"
	MetricGroupVisibilityLoggedIn = 0
)

// MetricGroupsService contains REST endpoints that belongs to cachet metric groups.
type MetricGroupsService struct {
	client *Client
}

// MetricGroup entity reflects one single metric group
type MetricGroup struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Order     int    `json:"order,omitempty"`
	Visible   int    `json:"visible,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// MetricGroupResponse reflects the response of /metrics/groups call
type MetricGroupResponse struct {
	Meta         Meta          `json:"meta,omitempty"`
	MetricGroups []MetricGroup `json:"data,omitempty"`
}

// MetricGroupsQueryParams contains fields to filter returned results
type MetricGroupsQueryParams struct {
	ID        int    `url:"id,omitempty"`
	Name      string `url:"name,omitempty"`
	Order     int    `url:"order,omitempty"`
	Collapsed bool   `url:"collapsed,omitempty"`
	Visible   int    `url:"visible,omitempty"`
	QueryOptions
}

// metricGroupAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the metric group in the "data" key.
type metricGroupAPIResponse struct {
	Data *MetricGroup `json:"data"`
}

// GetAll return all metric groups that have been created.
func (s *MetricGroupsService) GetAll(filter *MetricGroupsQueryParams) (*MetricGroupResponse, *Response, error) {
	u := "api/v1/metrics/groups"
	v := new(MetricGroupResponse)

	u, err := addOptions(u, filter)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get return a single metric group.
func (s *MetricGroupsService) Get(id int) (*MetricGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/metrics/groups/%d", id)
	v := new(metricGroupAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create creates a new metric group.
func (s *MetricGroupsService) Create(c *MetricGroup) (*MetricGroup, *Response, error) {
	u := "api/v1/metrics/groups"
	v := new(metricGroupAPIResponse)

	resp, err := s.client.Call("POST", u, c, v)
	return v.Data, resp, err
}

// Update updates a metric group.
func (s *MetricGroupsService) Update(id int, c *MetricGroup) (*MetricGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/metrics/groups/%d", id)
	v := new(metricGroupAPIResponse)

	resp, err := s.client.Call("PUT", u, c, v)
	return v.Data, resp, err
}

// Delete deletes a metric group.
func (s *MetricGroupsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/metrics/groups/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
