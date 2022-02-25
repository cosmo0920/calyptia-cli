package main

import (
	"context"

	"github.com/calyptia/api/types"
)

type Client interface {
	Project(ctx context.Context, projectID string) (types.Project, error)
	UpdateProject(ctx context.Context, projectID string, opts types.UpdateProject) error

	CreateInvitation(ctx context.Context, projectID string, payload types.CreateInvitation) error

	Members(ctx context.Context, projectID string, params types.MembersParams) ([]types.Membership, error)

	Token(ctx context.Context, tokenID string) (types.Token, error)
	UpdateToken(ctx context.Context, tokenID string, opts types.UpdateToken) error
	DeleteToken(ctx context.Context, tokenID string) error

	RegisterAgent(ctx context.Context, payload types.RegisterAgent) (types.RegisteredAgent, error)
	Agents(ctx context.Context, projectID string, params types.AgentsParams) ([]types.Agent, error)
	Agent(ctx context.Context, agentID string) (types.Agent, error)
	UpdateAgent(ctx context.Context, agentID string, payload types.UpdateAgent) error
	DeleteAgent(ctx context.Context, agentID string) error

	AgentConfigHistory(ctx context.Context, agentID string, params types.AgentConfigHistoryParams) ([]types.AgentConfig, error)

	CreateAggregator(ctx context.Context, payload types.CreateAggregator) (types.CreatedAggregator, error)
	Aggregators(ctx context.Context, projectID string, params types.AggregatorsParams) ([]types.Aggregator, error)
	Aggregator(ctx context.Context, aggregatorID string) (types.Aggregator, error)
	UpdateAggregator(ctx context.Context, aggregatorID string, payload types.UpdateAggregator) error
	DeleteAggregator(ctx context.Context, aggregatorID string) error

	CreateResourceProfile(ctx context.Context, aggregatorID string, payload types.CreateResourceProfile) (types.CreatedResourceProfile, error)
	ResourceProfiles(ctx context.Context, aggregatorID string, params types.ResourceProfilesParams) ([]types.ResourceProfile, error)
	ResourceProfile(ctx context.Context, resourceProfileID string) (types.ResourceProfile, error)
	UpdateResourceProfile(ctx context.Context, resourceProfileID string, opts types.UpdateResourceProfile) error
	DeleteResourceProfile(ctx context.Context, resourceProfileID string) error

	CreatePipeline(ctx context.Context, aggregatorID string, payload types.CreatePipeline) (types.CreatedPipeline, error)
	Pipelines(ctx context.Context, aggregatorID string, params types.PipelinesParams) ([]types.Pipeline, error)
	ProjectPipelines(ctx context.Context, projectID string, params types.PipelinesParams) ([]types.Pipeline, error)
	Pipeline(ctx context.Context, pipelineID string) (types.Pipeline, error)
	UpdatePipeline(ctx context.Context, pipelineID string, opts types.UpdatePipeline) (types.UpdatedPipeline, error)
	DeletePipeline(ctx context.Context, pipelineID string) error

	PipelineConfigHistory(ctx context.Context, pipelineID string, params types.PipelineConfigHistoryParams) ([]types.PipelineConfig, error)

	PipelineStatusHistory(ctx context.Context, pipelineID string, params types.PipelineStatusHistoryParams) ([]types.PipelineStatus, error)

	CreatePipelineFile(ctx context.Context, pipelineID string, payload types.CreatePipelineFile) (types.CreatedPipelineFile, error)
	PipelineFiles(ctx context.Context, pipelineID string, params types.PipelineFilesParams) ([]types.PipelineFile, error)
	PipelineFile(ctx context.Context, fileID string) (types.PipelineFile, error)
	UpdatePipelineFile(ctx context.Context, fileID string, opts types.UpdatePipelineFile) error
	DeletePipelineFile(ctx context.Context, fileID string) error

	CreatePipelineSecret(ctx context.Context, pipelineID string, payload types.CreatePipelineSecret) (types.CreatedPipelineSecret, error)
	PipelineSecrets(ctx context.Context, pipelineID string, params types.PipelineSecretsParams) ([]types.PipelineSecret, error)
	PipelineSecret(ctx context.Context, secretID string) (types.PipelineSecret, error)
	UpdatePipelineSecret(ctx context.Context, secretID string, opts types.UpdatePipelineSecret) error
	DeletePipelineSecret(ctx context.Context, secretID string) error

	CreatePipelinePort(ctx context.Context, pipelineID string, payload types.CreatePipelinePort) (types.CreatedPipelinePort, error)
	PipelinePorts(ctx context.Context, pipelineID string, params types.PipelinePortsParams) ([]types.PipelinePort, error)
	PipelinePort(ctx context.Context, portID string) (types.PipelinePort, error)
	UpdatePipelinePort(ctx context.Context, portID string, opts types.UpdatePipelinePort) error
	DeletePipelinePort(ctx context.Context, portID string) error

	ValidateConfig(ctx context.Context, agentType types.AgentType, payload types.ValidatingConfig) (types.ValidatedConfig, error)

	ProjectMetrics(ctx context.Context, projectID string, params types.MetricsParams) (types.ProjectMetrics, error)
	AgentMetrics(ctx context.Context, agentID string, params types.MetricsParams) (types.AgentMetrics, error)
	PipelineMetrics(ctx context.Context, pipelineID string, params types.MetricsParams) (types.AgentMetrics, error)
}
