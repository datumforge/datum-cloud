package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/rout"
	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/datumforge/datum-cloud/internal/v1/models"
)

const (
	relationBucketName = "relationships"
)

// WorkspaceHandler is the handler for the workspace endpoint
func (h *Handler) WorkspaceHandler(ctx echo.Context) error {
	var in models.WorkspaceRequest
	if err := ctx.Bind(&in); err != nil {
		return h.InvalidInput(ctx, err)
	}

	if err := in.Validate(); err != nil {
		return h.InvalidInput(ctx, err)
	}

	h.Logger.Debugw("creating workspace",
		"name", in.Name,
		"environments", in.Environments,
		"buckets", in.Buckets,
		"relationships", in.Relationships,
	)

	// create root organization
	rootOrgName := in.Name
	input := datumclient.CreateOrganizationInput{
		Name:        strings.ToLower(rootOrgName),
		DisplayName: &rootOrgName,
	}

	if in.Description != "" {
		input.Description = &in.Description
	}

	if len(in.Domains) > 0 {
		input.CreateOrgSettings = &datumclient.CreateOrganizationSettingInput{
			Domains: in.Domains,
		}
	}

	ws, err := h.DatumClient.CreateOrganization(ctx.Request().Context(), input)
	if err != nil {
		return h.BadRequest(ctx, err)
	}

	workspace := ws.CreateOrganization.Organization

	out := models.WorkspaceReply{
		Reply:       rout.Reply{Success: true},
		ID:          workspace.ID,
		Name:        workspace.DisplayName,
		Description: *workspace.Description,
		Domains:     workspace.Setting.Domains,
	}

	// create environments
	envOrgs, err := h.createEnvironments(ctx.Request().Context(), workspace.ID, in.Environments, input)
	if err != nil {
		return h.BadRequest(ctx, err)
	}

	// for each environment, create buckets
	for i, envOrg := range envOrgs {
		// add environments to the response
		out.Environments = append(out.Environments, models.Environment{
			OrgDetails: models.OrgDetails{
				ID:   envOrg.ID,
				Name: envOrg.DisplayName,
			},
			Buckets: []models.Bucket{},
		})

		// create buckets
		bucketOrgs, err := h.createBuckets(ctx.Request().Context(), envOrg.ID, envOrg.DisplayName, in.Buckets, input)
		if err != nil {
			return h.BadRequest(ctx, err)
		}

		for j, bucketOrg := range bucketOrgs {
			// add buckets to the response
			out.Environments[i].Buckets = append(out.Environments[i].Buckets, models.Bucket{
				OrgDetails: models.OrgDetails{
					ID:   bucketOrg.ID,
					Name: bucketOrg.DisplayName,
				},
			})

			// create relationships under the relationships bucket
			if bucketOrg.DisplayName == relationBucketName {
				relationshipOrgs, err := h.createRelationships(ctx.Request().Context(), bucketOrg.ID, envOrg.DisplayName, in.Relationships, input)
				if err != nil {
					return h.BadRequest(ctx, err)
				}

				out.Environments[i].Buckets[j].Relations = []models.Relationship{}

				for _, relationshipOrg := range relationshipOrgs {
					// add relationships to the response
					out.Environments[i].Buckets[j].Relations = append(out.Environments[i].Buckets[j].Relations, models.Relationship{
						OrgDetails: models.OrgDetails{
							ID:   relationshipOrg.ID,
							Name: relationshipOrg.DisplayName,
						},
					})
				}
			}
		}
	}

	return h.Success(ctx, out)
}

// BindWorkspaceHandler is used to bind the workspace endpoint to the OpenAPI schema
func (h *Handler) BindWorkspaceHandler() *openapi3.Operation {
	register := openapi3.NewOperation()
	register.Description = "Workspace creates an opinionated organization hierarchy for the new organization"
	register.OperationID = "WorkspaceHandler"
	register.Security = &openapi3.SecurityRequirements{}

	h.AddRequestBody("WorkspaceRequest", models.ExampleWorkspaceSuccessRequest, register)
	h.AddResponse("WorkspaceReply", "success", models.ExampleWorkspaceSuccessResponse, register, http.StatusOK)
	register.AddResponse(http.StatusInternalServerError, internalServerError())
	register.AddResponse(http.StatusBadRequest, badRequest())

	return register
}

// createChildOrganizations creates the child organizations for the workspace
func (h *Handler) createChildOrganizations(ctx context.Context, namePrefix, parentOrgID string, childNames, additionalTags []string) ([]datumclient.CreateOrganization_CreateOrganization_Organization, error) {
	var orgs []datumclient.CreateOrganization_CreateOrganization_Organization

	for _, childName := range childNames {
		h.Logger.Debugw("creating child organization", "childName", childName)

		orgName := childName

		input := datumclient.CreateOrganizationInput{
			Name:        strings.ToLower(fmt.Sprintf("%s.%s", namePrefix, orgName)),
			DisplayName: &orgName,
			ParentID:    &parentOrgID,
			Tags:        append(additionalTags, orgName),
		}

		// create child organization
		o, err := h.DatumClient.CreateOrganization(ctx, input)
		if err != nil {
			return nil, err
		}

		orgs = append(orgs, o.CreateOrganization.Organization)
	}

	return orgs, nil
}

// createEnvironments creates the environments for the workspace
func (h *Handler) createEnvironments(ctx context.Context, rootOrgID string, environments []string, input datumclient.CreateOrganizationInput) ([]datumclient.CreateOrganization_CreateOrganization_Organization, error) {
	return h.createChildOrganizations(ctx, input.Name, rootOrgID, environments, []string{})
}

// createBuckets creates the buckets for the workspace for each environment
func (h *Handler) createBuckets(ctx context.Context, envOrgID, environment string, buckets []string, input datumclient.CreateOrganizationInput) ([]datumclient.CreateOrganization_CreateOrganization_Organization, error) {
	return h.createChildOrganizations(ctx, fmt.Sprintf("%s.%s", input.Name, environment), envOrgID, buckets, []string{environment})
}

// createRelationships creates the relationships for the workspace for each environment
func (h *Handler) createRelationships(ctx context.Context, relationshipOrgID, environment string, relationships []string, input datumclient.CreateOrganizationInput) ([]datumclient.CreateOrganization_CreateOrganization_Organization, error) {
	return h.createChildOrganizations(ctx, fmt.Sprintf("%s.%s.%s", input.Name, environment, "relationships"), relationshipOrgID, relationships, []string{environment, "relationships"})
}
