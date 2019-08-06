package config

import "go.zenithar.org/pkg/platform"

// Configuration contains spotigraph settings
type Configuration struct {
	Debug struct {
		Enable bool `toml:"enable" default:"false" comment:"allow debug mode"`
	} `toml:"Debug" comment:"###############################\n Debug \n##############################"`

	Instrumentation platform.InstrumentationConfig `toml:"Instrumentation" comment:"###############################\n Instrumentation \n##############################"`

	Core struct {
		BasePath    string `toml:"basePath" default:"" comment:"Define a NDS root directory"`
		ProductPath string `toml:"productPath" default:"" comment:"Define a NDS product folder"`
		FileName    string `toml:"fileName" default:"" comment:"Define a tar file name"`
	} `toml:"Core" comment:"###############################\n Core \n##############################"`

	Server struct {
		Network string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
		Listen  string `toml:"listen" default:":5555" comment:"Listen address for gRPC server"`
		UseTLS  bool   `toml:"useTLS" default:"false" comment:"Enable TLS listener"`
		TLS     struct {
			CertificatePath              string `toml:"certificatePath" default:"" comment:"Certificate path"`
			PrivateKeyPath               string `toml:"privateKeyPath" default:"" comment:"Private Key path"`
			CACertificatePath            string `toml:"caCertificatePath" default:"" comment:"CA Certificate Path"`
			ClientAuthenticationRequired bool   `toml:"clientAuthenticationRequired" default:"false" comment:"Force client authentication"`
		} `toml:"TLS" comment:"TLS Socket settings"`
	}
}
