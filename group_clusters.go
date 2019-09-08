package gitlab

import (
	"fmt"
	"time"
)

// GroupClustersService handles communication with the
// group clusters related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html
type GroupClustersService struct {
	client *Client
}

// GroupCluster represents a GitLab Group Cluster.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/group_clusters.html
type GroupCluster struct {
	ID                 int                      `json:"id"`
	Name               string                   `json:"name"`
	Domain             string                   `json:"domain"`
	CreatedAt          *time.Time               `json:"created_at"`
	ProviderType       string                   `json:"provider_type"`
	PlatformType       string                   `json:"platform_type"`
	EnvironmentScope   string                   `json:"environment_scope"`
	ClusterType        string                   `json:"cluster_type"`
	User               *User                    `json:"user"`
	PlatformKubernetes *GroupPlatformKubernetes `json:"platform_kubernetes"`
	Group              *Group                   `json:"group"`
}

// PlatformKubernetes represents a GitLab Group Cluster PlatformKubernetes.
type GroupPlatformKubernetes struct {
	APIURL            string `json:"api_url"`
	Token             string `json:"token"`
	CaCert            string `json:"ca_cert"`
	AuthorizationType string `json:"authorization_type"`
}

func (v GroupCluster) String() string {
	return Stringify(v)
}

// ListClusters gets a list of all clusters in a group.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html#list-group-clusters
func (s *GroupClustersService) ListClusters(gid interface{}, options ...OptionFunc) ([]*GroupCluster, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters", pathEscape(group))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var pcs []*GroupCluster
	resp, err := s.client.Do(req, &pcs)
	if err != nil {
		return nil, resp, err
	}

	return pcs, resp, err
}

// GetCluster gets a cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html#get-a-single-group-cluster
func (s *GroupClustersService) GetCluster(gid interface{}, cluster int, options ...OptionFunc) (*GroupCluster, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/%d", pathEscape(group), cluster)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pc := new(GroupCluster)
	resp, err := s.client.Do(req, &pc)
	if err != nil {
		return nil, resp, err
	}

	return pc, resp, err
}

// AddGroupClusterOptions represents the available AddCluster() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html#add-existing-cluster-to-group
type AddGroupClusterOptions struct {
	Name               *string                            `url:"name,omitempty" json:"name,omitempty"`
	Domain             *string                            `url:"domain,omitempty" json:"domain,omitempty"`
	Enabled            *bool                              `url:"enabled,omitempty" json:"enabled,omitempty"`
	Managed            *bool                              `url:"managed,omitempty" json:"managed,omitempty"`
	EnvironmentScope   *string                            `url:"environment_scope,omitempty" json:"environment_scope,omitempty"`
	PlatformKubernetes *AddGroupPlatformKubernetesOptions `url:"platform_kubernetes_attributes,omitempty" json:"platform_kubernetes_attributes,omitempty"`
}

// AddGroupPlatformKubernetesOptions represents the available PlatformKubernetes options for adding a Group Cluster.
type AddGroupPlatformKubernetesOptions struct {
	APIURL            *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	Token             *string `url:"token,omitempty" json:"token,omitempty"`
	CaCert            *string `url:"ca_cert,omitempty" json:"ca_cert,omitempty"`
	AuthorizationType *string `url:"authorization_type,omitempty" json:"authorization_type,omitempty"`
}

// AddCluster adds an existing cluster to the group.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html#add-existing-cluster-to-group
func (s *GroupClustersService) AddCluster(gid interface{}, opt *AddGroupClusterOptions, options ...OptionFunc) (*GroupCluster, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/user", pathEscape(group))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pc := new(GroupCluster)
	resp, err := s.client.Do(req, pc)
	if err != nil {
		return nil, resp, err
	}

	return pc, resp, err
}

// EditGroupClusterOptions represents the available EditCluster() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html#edit-group-cluster
type EditGroupClusterOptions struct {
	Name               *string                        `url:"name,omitempty" json:"name,omitempty"`
	Domain             *string                        `url:"domain,omitempty" json:"domain,omitempty"`
	EnvironmentScope   *string                        `url:"environment_scope,omitempty" json:"environment_scope,omitempty"`
	PlatformKubernetes *EditGroupPlatformKubernetesOptions `url:"platform_kubernetes_attributes,omitempty" json:"platform_kubernetes_attributes,omitempty"`
}

// EditGroupPlatformKubernetesOptions represents the available PlatformKubernetes options for editing a Group Cluster.
type EditGroupPlatformKubernetesOptions struct {
	APIURL    *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	Token     *string `url:"token,omitempty" json:"token,omitempty"`
	CaCert    *string `url:"ca_cert,omitempty" json:"ca_cert,omitempty"`
}

// EditCluster updates an existing group cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html#edit-group-cluster
func (s *GroupClustersService) EditCluster(gid interface{}, cluster int, opt *EditGroupClusterOptions, options ...OptionFunc) (*GroupCluster, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/%d", pathEscape(group), cluster)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pc := new(GroupCluster)
	resp, err := s.client.Do(req, pc)
	if err != nil {
		return nil, resp, err
	}

	return pc, resp, err
}

// DeleteCluster deletes an existing group cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_clusters.html#delete-group-cluster
func (s *GroupClustersService) DeleteCluster(gid interface{}, cluster int, options ...OptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/clusters/%d", pathEscape(group), cluster)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
