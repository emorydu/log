module github.com/emorydu/log

go 1.19

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	k8s.io/klog v1.0.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace go.uber.org/zap v1.25.0 => ./third_party/forked/zap
