package grab

import "github.com/fionahiklas/sky-cloud-reporter/common/reporter"

type Grabber interface {
	GrabInstances() (instances []reporter.MachineInstance, err error)
}
