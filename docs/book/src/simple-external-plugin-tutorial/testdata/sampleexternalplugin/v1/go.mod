module v1

go 1.20

require (
	github.com/spf13/pflag v1.0.5
	sigs.k8s.io/kubebuilder/v3 v3.11.1
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/gobuffalo/flect v1.0.2 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	golang.org/x/tools v0.12.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace sigs.k8s.io/kubebuilder/v3 => ../../../../../../../
