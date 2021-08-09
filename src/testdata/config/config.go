package config

// TTMakeConf represents table test structure of MakeConfig test
type TTMakeConf struct {
	Name    string
	EnvVar  string
	IsProd  bool
	IsDev   bool
}

// CreateTTMakeConf creates table test for MakeConfig test
func CreateTTMakeConf() []TTMakeConf {
	tt := []TTMakeConf{
		{
			Name:   "dev config",
			EnvVar: "dev",
			IsDev:  true,
		},
		{
			Name:   "prod config",
			EnvVar: "prod",
			IsProd: true,
		},
	}
	return tt
}
