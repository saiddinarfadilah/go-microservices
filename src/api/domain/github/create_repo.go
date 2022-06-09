package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Url     string `json:"url"`
	Login   string `json:"login"`
	HtmlUrl string `json:"html_url"`
}

type RepoPermissions struct {
	IsAdmin bool `json:"is_admin"`
	HasPull bool `json:"has_pull"`
	HasPush bool `json:"has_push"`
}
