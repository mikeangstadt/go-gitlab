//
// Copyright 2017, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"time"
)

// ResourceLabelEventsService handles communication with the event related methods of
// the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/resource_label_events.html
type ResourceLabelEventsService struct {
	client *Client
}

// ContributionEvent represents a user's contribution
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/resource_label_events.html#get-single-issue-label-event
type ResourceLabelEvent struct {
	ID              int        `json:"id"`
	Action          string     `json:"action"`
	CreatedAt       *time.Time `json:"created_at"`
	ResourceType    string     `json:"resource_type"`
	ResourceID      int        `json:"resource_id"`
	User struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"user"`
	Label struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Color       string `json:"color"`
		TextColor   string `json:"text_color"`
		Description string `json:"description"`
	} `json:"label"`
}

// ListResourceLabelEventsOptions represents the options for 
type ListResourceLabelEventsOptions struct {
	ListOptions
}

// ListIssueResourceLabelEvents retrieves resource label events
// for the specified project and issue.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/resource_label_events.html#list-project-issue-label-events
func (s *ResourceLabelEventsService) ListIssueResourceLabelEvents(pid interface{}, issue int, opt *ListResourceLabelEventsOptions, options ...OptionFunc) ([]*ResourceLabelEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/issues/%d/resource_label_events", pathEscape(project), issue)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var rs []*ResourceLabelEvent
	resp, err := s.client.Do(req, &rs)
	if err != nil {
		return nil, resp, err
	}

	return rs, resp, err
}

// GetIssueResourceLabelEvent gets a single issue-label-event.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/resource_label_events.html#get-single-issue-label-event
func (s *ResourceLabelEventsService) GetIssueResourceLabelEvent(pid interface{}, issue int, event int, options ...OptionFunc) (*ResourceLabelEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/resource_label_events/%d", pathEscape(project), issue, event)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	e := new(ResourceLabelEvent)
	resp, err := s.client.Do(req, e)
	if err != nil {
		return nil, resp, err
	}

	return e, resp, err
}

// ListGroupEpicResourceLabelEvents retrieves resource label events
// for the specified group and epic.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/resource_label_events.html#list-group-epic-label-events
func (s *ResourceLabelEventsService) ListGroupEpicResourceLabelEvents(gid interface{}, epic int, opt *ListResourceLabelEventsOptions, options ...OptionFunc) ([]*ResourceLabelEvent, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("groups/%s/epics/%d/resource_label_events", pathEscape(group), epic)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var rs []*ResourceLabelEvent
	resp, err := s.client.Do(req, &rs)
	if err != nil {
		return nil, resp, err
	}

	return rs, resp, err
}

// GetGroupEpicResourceLabelEvent gets a single group epic label event.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/resource_label_events.html#get-single-epic-label-event
func (s *ResourceLabelEventsService) GetGroupEpicResourceLabelEvent(gid interface{}, epic int, event int, options ...OptionFunc) (*ResourceLabelEvent, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/epics/%d/resource_label_events/%d", pathEscape(group), epic, event)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	e := new(ResourceLabelEvent)
	resp, err := s.client.Do(req, e)
	if err != nil {
		return nil, resp, err
	}

	return e, resp, err
}

// ListMergeRequestResourceLabelEvents retrieves resource label events
// for the specified project and merge request.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/resource_label_events.html#list-project-merge-request-label-events
func (s *ResourceLabelEventsService) ListMergeRequestResourceLabelEvents(pid interface{}, request int, opt *ListResourceLabelEventsOptions, options ...OptionFunc) ([]*ResourceLabelEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/merge_requests/%d/resource_label_events", pathEscape(project), request)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var rs []*ResourceLabelEvent
	resp, err := s.client.Do(req, &rs)
	if err != nil {
		return nil, resp, err
	}

	return rs, resp, err
}

// GetMergeRequestResourceLabelEvent gets a single merge request label event.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/resource_label_events.html#get-single-merge-request-label-event
func (s *ResourceLabelEventsService) GetMergeRequestResourceLabelEvent(pid interface{}, request int, event int, options ...OptionFunc) (*ResourceLabelEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/resource_label_events/%d", pathEscape(project), request, event)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	e := new(ResourceLabelEvent)
	resp, err := s.client.Do(req, e)
	if err != nil {
		return nil, resp, err
	}

	return e, resp, err
}
