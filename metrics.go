package cachet

import (
	"fmt"
)

const (
	// MetricsViewLastHour means "Default view: Last Hour"
	MetricsViewLastHour = 0
	// MetricsViewLast12Hours means "Default view: Last 12 Hours"
	MetricsViewLast12Hours = 1
	// MetricsViewLastWeek means "Default view: Week"
	MetricsViewLastWeek = 2
	// MetricsViewLastMonth means "Default view: Month"
	MetricsViewLastMonth = 3

	// MetricsCalculationSum means "Calculation of Metrics: Sum"
	MetricsCalculationSum = 0
	// MetricsCalculationAverage means "Calculation of Metrics: Average"
	MetricsCalculationAverage = 1

	// MetricsVisibilityLoggedIn means "Visibility: Visible to authenticated users"
	MetricsVisibilityLoggedIn = 0
	// MetricsVisibilityPublic means "Visibility: Visible to everybody"
	MetricsVisibilityPublic = 1
	// MetricsVisibilityHidden means "Visibility: Always hidden"
	MetricsVisibilityHidden = 2

	// MetricsNoDisplayChart means to not display chart in Status Page
	MetricsNoDisplayChart = 0
	// MetricsDisplayChart means to display chart in Status Page
	MetricsDisplayChart = 1
)

// MetricsService contains REST endpoints that belongs to cachet metrics.
type MetricsService struct {
	client *Client
}

// Metric entity reflects one single metric
type Metric struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Suffix          string `json:"suffix,omitempty"`
	Description     string `json:"description,omitempty"`
	DefaultValue    int    `json:"default_value"`
	CalcType        int    `json:"calc_type"`
	DisplayChart    bool   `json:"display_chart"`
	Places          int    `json:"places,omitempty"`
	DefaultView     int    `json:"default_view"`
	Threshold       int    `json:"threshold,omitempty"`
	Order           int    `json:"order,omitempty"`
	Visible         int    `json:"visible"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	DefaultViewName string `json:"default_view_name,omitempty"`
	GroupID         int    `json:"group_id,omitempty"`
}

// Point is a single point in a Metric
type Point struct {
	ID              int    `json:"id,omitempty"`
	MetricID        int    `json:"metric_id,omitempty"`
	Value           int    `json:"value,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	Counter         int    `json:"counter,omitempty"`
	CalculatedValue int    `json:"calculated_value,omitempty"`
}

// MetricResponse reflects the response of /metric call
type MetricResponse struct {
	Meta    Meta     `json:"meta,omitempty"`
	Metrics []Metric `json:"data,omitempty"`
}

// MetricQueryParams contains fields to filter returned results
type MetricQueryParams struct {
	QueryOptions
}

// metricAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the metric in the "data" key.
type metricAPIResponse struct {
	Data *Metric `json:"data"`
}

// metricPointsAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the metric points in the "data" key.
type metricPointsAPIResponse struct {
	Data *[]Point `json:"data"`
}

// metricPointAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the metric point in the "data" key.
type metricPointAPIResponse struct {
	Data *Point `json:"data"`
}

// GetAll returns all metrics that have been setup.
//
// Docs: https://docs.cachethq.io/reference#get-metrics
func (s *MetricsService) GetAll(filter *MetricQueryParams) (*MetricResponse, *Response, error) {
	u := "api/v1/metrics"
	v := new(MetricResponse)

	u, err := addOptions(u, filter)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single metric, without points.
//
// Docs: https://docs.cachethq.io/reference#get-a-metric
func (s *MetricsService) Get(id int) (*Metric, *Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d", id)
	v := new(metricAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create a new metric.
//
// Docs: https://docs.cachethq.io/reference#metrics
func (s *MetricsService) Create(m *Metric) (*Metric, *Response, error) {
	u := "api/v1/metrics"
	v := new(metricAPIResponse)

	resp, err := s.client.Call("POST", u, m, v)
	return v.Data, resp, err
}

// Update a metric.
//
// Docs: this is not documented in the API docs but is supported
func (s *MetricsService) Update(id int, m *Metric) (*Metric, *Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d", id)
	v := new(metricAPIResponse)

	resp, err := s.client.Call("PUT", u, m, v)
	return v.Data, resp, err
}

// Delete a metric.
//
// Docs: https://docs.cachethq.io/reference#delete-a-metric
func (s *MetricsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}

// GetPoints return a list of metric points.
//
// Docs: https://docs.cachethq.io/reference#get-metric-points
func (s *MetricsService) GetPoints(id int) (*[]Point, *Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d/points", id)
	v := new(metricPointsAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// AddPoint adds a metric point to a given metric.
//
// Docs: https://docs.cachethq.io/reference#post-metric-points
func (s *MetricsService) AddPoint(id int, value int, timestamp string) (*Point, *Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d/points", id)
	v := new(metricPointAPIResponse)

	p := struct {
		Value     int    `json:"value"`
		Timestamp string `json:"timestamp"`
	}{
		Value:     value,
		Timestamp: timestamp,
	}

	resp, err := s.client.Call("POST", u, p, v)
	return v.Data, resp, err
}

// DeletePoint deletes a metric point.
//
// Docs: https://docs.cachethq.io/reference#delete-a-metric-point
func (s *MetricsService) DeletePoint(id, pointID int) (*Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d/points/%d", id, pointID)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
