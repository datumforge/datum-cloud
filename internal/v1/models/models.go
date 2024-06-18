package models

import (
	"github.com/datumforge/datum/pkg/rout"
	"github.com/mcuadros/go-defaults"
)

// =========
// WORKSPACE
// =========

// WorkspaceRequest is the request object for creating a workspace
type WorkspaceRequest struct {
	Name          string   `json:"name"`
	Description   string   `json:"description,omitempty"`
	Domains       []string `json:"domains,omitempty"`
	Environments  []string `json:"environments,omitempty" default:"[production,testing]"`
	Buckets       []string `json:"buckets,omitempty" default:"[assets,customers,orders,relationships,sales]"`
	Relationships []string `json:"relationships,omitempty" default:"[internal_users,marketing_subscribers,marketplaces,partners,vendors]"`
}

// WorkspaceReply is the response object for creating a workspace
type WorkspaceReply struct {
	rout.Reply
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description,omitempty"`
	Domains      []string      `json:"domains,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
}

type Environment struct {
	OrgDetails
	Buckets []Bucket `json:"buckets,omitempty"`
}

type Bucket struct {
	OrgDetails
	Relations []Relationship `json:"relations,omitempty"`
}

type Relationship struct {
	OrgDetails
}

type OrgDetails struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Validate ensures the required fields are set on the WorkspaceRequest request
func (r *WorkspaceRequest) Validate() error {
	// Required for all requests
	if r.Name == "" {
		return rout.MissingField("name")
	}

	// Set default values if not provided in the request
	defaultRequest := &WorkspaceRequest{}
	defaults.SetDefaults(defaultRequest)

	if r.Environments == nil {
		r.Environments = defaultRequest.Environments
	}

	if r.Buckets == nil {
		r.Buckets = defaultRequest.Buckets
	}

	if r.Relationships == nil {
		r.Relationships = defaultRequest.Relationships
	}

	return nil
}

// ExampleWorkspaceSuccessRequest is an example of a successful workspace request for OpenAPI documentation
var ExampleWorkspaceSuccessRequest = WorkspaceRequest{
	Name: "MITB Inc.",
}

// ExampleWorkspaceSuccessResponse is an example of a successful workspace response for OpenAPI documentation
var ExampleWorkspaceSuccessResponse = WorkspaceReply{
	Reply: rout.Reply{Success: true},
	ID:    "1234",
	Name:  "MITB Inc.",
}
