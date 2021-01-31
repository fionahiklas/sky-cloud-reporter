//go:generate mockgen -package=mock_grab -destination=../mocks/mock_grab/mock_cloud_interfaces.go . CloudProvider

package grab

import (
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"net/http"
)

type CloudProvider interface {
	GetInstanceUrl() string
	ConvertResponseToMachineInstances(response *http.Response) (machines *[]reporter.MachineInstance, err error)
}