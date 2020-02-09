package internal

type Conf struct {
	Action         string   `yaml:"action"`
	Name           string   `yaml:"eks_cluster_name"`
	Region         string   `yaml:"region"`
	Vpc_name       string   `yaml:"vpc_name"`
	Cidr           string   `yaml:"cidr"`
	NodeGroup      string   `yaml:"NodeGroup"`
	Public_Access  bool     `yaml:"public_access"`
	Private_Access bool     `yaml:"private_access"`
	Zones          []string `json:"zones"`
}
