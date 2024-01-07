package models

type Project struct {
	ID                int        `json:"id"`
	Description       string     `json:"description"`
	DescriptionHTML   string     `json:"description_html"`
	Name              string     `json:"name"`
	NameWithNamespace string     `json:"name_with_namespace"`
	Path              string     `json:"path"`
	PathWithNamespace string     `json:"path_with_namespace"`
	CreatedAt         string     `json:"created_at"`
	DefaultBranch     string     `json:"default_branch"`
	TagList           []struct{} `json:"tag_list"`
	SSHURLToRepo      string     `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string     `json:"http_url_to_repo"`
	WebURL            string     `json:"web_url"`
	ReadmeURL         string     `json:"readme_url"`
	ForksCount        int        `json:"forks_count"`
	AvatarURL         string     `json:"avatar_url"`
	StarCount         int        `json:"star_count"`
	LasActivityAt     string     `json:"last_activity_at"`
	Namespace         struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Path      string `json:"path"`
		Kind      string `json:"kind"`
		FullPath  string `json:"full_path"`
		ParentID  int    `json:"parent_id"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"namespace"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix"`
	Links                        struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
		ClusterAgents string `json:"cluster_agents"`
	} `json:"_links"`
	CodeSuggestions                bool   `json:"code_suggestions"`
	EmptyRepo                      bool   `json:"empty_repo"`
	Archived                       bool   `json:"archived"`
	Visibility                     string `json:"visibility"`
	ResolveOutdatedDiffDiscussions bool   `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy      struct {
		Cadence       string `json:"cadence"`
		Enabled       bool   `json:"enabled"`
		KeepN         int    `json:"keep_n"`
		OlderThan     string `json:"older_than"`
		NameRegex     string `json:"name_regex"`
		NameRegexKeep string `json:"name_regex_keep"`
		NextRunAt     string `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	IssuesEnabled                               bool     `json:"issues_enabled"`
	MergeRequestsEnabled                        bool     `json:"merge_requests_enabled"`
	WikiEnabled                                 bool     `json:"wiki_enabled"`
	JobsEnabled                                 bool     `json:"jobs_enabled"`
	SnippetsEnabled                             bool     `json:"snippets_enabled"`
	ContainerRegistryEnabled                    bool     `json:"container_registry_enabled"`
	ServiceDeskEnabled                          bool     `json:"service_desk_enabled"`
	ServiceDeskAddress                          string   `json:"service_desk_address"`
	CanCreateMergeRequestIn                     bool     `json:"can_create_merge_request_in"`
	IssuesAccessLevel                           string   `json:"issues_access_level"`
	RepositoryAccessLevel                       string   `json:"repository_access_level"`
	MergeRequestsAccessLevel                    string   `json:"merge_requests_access_level"`
	ForkingAccessLevel                          string   `json:"forking_access_level"`
	WikiAccessLevel                             string   `json:"wiki_access_level"`
	BuildsAccessLevel                           string   `json:"builds_access_level"`
	SnippetsAccessLevel                         string   `json:"snippets_access_level"`
	PagesAccessLevel                            string   `json:"pages_access_level"`
	AnalyticsAccessLevel                        string   `json:"analytics_access_level"`
	ContainerRegistryAccessLevel                string   `json:"container_registry_access_level"`
	SecurityAndComplianceAccessLevel            string   `json:"security_and_compliance_access_level"`
	ReleaseAccessLevel                          string   `json:"release_access_level"`
	EnvironmentAccessLevel                      string   `json:"environment_access_level"`
	FeatureFlagsAccessLevel                     string   `json:"feature_flags_access_level"`
	InfrastructureAccessLevel                   string   `json:"infrastructure_access_level"`
	MonitorAccessLevel                          string   `json:"monitor_access_level"`
	ModelExpirementAccessLevel                  string   `json:"model_expirement_access_level"`
	ModelRegistryAccessLevel                    string   `json:"model_registry_access_level"`
	EmailsDisabled                              bool     `json:"emails_disabled"`
	EmailsEnabled                               bool     `json:"emails_enabled"`
	SharedRunnersEnabled                        bool     `json:"shared_runners_enabled"`
	LfsEnabled                                  bool     `json:"lfs_enabled"`
	CreatedID                                   int      `json:"creator_id"`
	ImportURL                                   string   `json:"import_url"`
	ImportType                                  string   `json:"import_type"`
	ImportStatus                                string   `json:"import_status"`
	OpenIssuesCount                             int      `json:"open_issues_count"`
	UpdatedAt                                   string   `json:"updated_at"`
	CIDefaultGitDepth                           int      `json:"ci_default_git_depth"`
	CIForwardDeploymentEnabled                  bool     `json:"ci_forward_deployment_enabled"`
	CIForwardDeploymentRollbackAllowed          bool     `json:"ci_forward_deployment_rollback_allowed"`
	CIJobTokenScopeEnabled                      bool     `json:"ci_job_token_scope_enabled"`
	CISeparatedCaches                           bool     `json:"ci_separated_caches"`
	CIAllowForkPipelineJobsToRunInParentProject bool     `json:"ci_allow_fork_pipeline_jobs_to_run_in_parent_project"`
	BuildGitStrategy                            string   `json:"build_git_strategy"`
	KeepLastArtifact                            bool     `json:"keep_last_artifact"`
	RestrictUserDefinedVariables                bool     `json:"restrict_user_defined_variables"`
	RunnersToken                                string   `json:"runners_token"`
	RunnerTokenExpirationInterval               int      `json:"runner_token_expiration_interval"`
	GroupRunnersEnabled                         bool     `json:"group_runners_enabled"`
	AutoCancelPendingPipelines                  string   `json:"auto_cancel_pending_pipelines"`
	BuildTimeout                                int      `json:"build_timeout"`
	AutoDevopsEnabled                           bool     `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                    string   `json:"auto_devops_deploy_strategy"`
	CIConfigPath                                string   `json:"ci_config_path"`
	PublicJobs                                  bool     `json:"public_jobs"`
	SharedWithGroups                            []string `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds            bool     `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline                 bool     `json:"allow_merge_on_skipped_pipeline"`
	RequestAccessEnabled                        bool     `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved   bool     `json:"only_allow_merge_if_all_discussions_are_resolved"`
	PrintingMergeRequestLinkEnabled             bool     `json:"printing_merge_request_link_enabled"`
	MergeMethod                                 string   `json:"merge_method"`
	SquashOption                                string   `json:"squash_option"`
	EnforceAuthChecksOnUploads                  bool     `json:"enforce_auth_checks_on_uploads"`
	SuggestionCommitMessage                     string   `json:"suggestion_commit_message"`
	MergeCommitTemplate                         string   `json:"merge_commit_template"`
	SquashCommitMessage                         string   `json:"squash_commit_message"`
	IssueBranchTemplate                         string   `json:"issue_branch_template"`
	AutocloseReferencedIssues                   bool     `json:"autoclose_referenced_issues"`
	ExternalAuthorizationClassificationLabel    string   `json:"external_authorization_classification_label"`
	RequirementsEnabled                         bool     `json:"requirements_enabled"`
	RequirementsAccessLevel                     string   `json:"requirements_access_level"`
	SecurityAndComplianceEnabled                bool     `json:"security_and_compliance_enabled"`
	ComplianceFrameworks                        []string `json:"compliance_frameworks"`
	Permissions                                 struct {
		ProjectAccess bool `json:"project_access"`
		GroupAccess   struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		}
	}
}