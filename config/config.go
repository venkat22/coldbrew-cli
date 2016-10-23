package config

type Config struct {
	Name         *string            `json:"name,omitempty" yaml:"name,omitempty"`
	ClusterName  *string            `json:"cluster,omitempty" yaml:"cluster,omitempty"`
	Port         *uint16            `json:"port,omitempty" yaml:"port,omitempty"`
	CPU          *float64           `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory       *string            `json:"memory,omitempty" yaml:"memory,omitempty"`
	Units        *uint16            `json:"units,omitempty" yaml:"units,omitempty"`
	Env          map[string]string  `json:"env" yaml:"env"`
	LoadBalancer ConfigLoadBalancer `json:"load_balancer" yaml:"load_balancer"`
	AWS          ConfigAWS          `json:"aws" yaml:"aws"`
	Docker       ConfigDocker       `json:"docker" yaml:"docker"`
}

type ConfigLoadBalancer struct {
	Enabled     *bool                         `json:"enabled" yaml:"enabled"`
	IsHTTPS     *bool                         `json:"https,omitempty" yaml:"https,omitempty"`
	Port        *uint16                       `json:"port,omitempty" yaml:"port,omitempty"`
	HealthCheck ConfigLoadBalancerHealthCheck `json:"health_check,omitempty" yaml:"health_check,omitempty"`
}

type ConfigLoadBalancerHealthCheck struct {
	Interval       *string `json:"interval,omitempty" yaml:"interval,omitempty"`
	Path           *string `json:"path,omitempty" yaml:"path,omitempty"`
	Status         *string `json:"status,omitempty" yaml:"status,omitempty"`
	Timeout        *string `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	HealthyLimit   *uint16 `json:"healthy_limit,omitempty" yaml:"healthy_limit,omitempty"`
	UnhealthyLimit *uint16 `json:"unhealthy_limit,omitempty" yaml:"unhealthy_limit,omitempty"`
}

type ConfigAWS struct {
	ELBLoadBalancerName  *string `json:"elb_name,omitempty" yaml:"elb_name,omitempty"`
	ELBTargetGroupName   *string `json:"elb_target_group_name,omitempty" yaml:"elb_target_group_name,omitempty"`
	ELBSecurityGroupName *string `json:"elb_security_group_name,omitempty" yaml:"elb_security_group_name,omitempty"`
	ECRRepositoryName    *string `json:"ecr_repo_name,omitempty" yaml:"ecr_repo_name,omitempty"`
}

type ConfigDocker struct {
	Bin *string `json:"bin,omitempty" yaml:"bin,omitempty"`
}
