//go:generate mockgen -package=mock_grab -destination=../mocks/mock_grab/mock_cloud_interfaces.go . CloudProvider

package grab

import (
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"net/http"
)

type CloudProvider interface {

	// Does this provider split the full list of machines into
	// pages.  If this is the case that we need to loop through
	// until there are no more pages to process.
	RequiresPaging() bool

	// Generate the next HTTP URL to use to request data from the
	// providers instance endpoint.  For providers that don't need paging
	// this will be called once and return the single URL.  For the paging
	// case this will return an initial URL and will then update to the
	// next page once the response has been processed by calling
	// ProcessResponse() since the information about the size of the results
	// is embedded in the JSON.  If there are no more pages the done return
	// value will be "true".
	GenerateNextUrl() (url string, done bool)

	// Process the HTTP response from the cloud provider.  Convert machine
	// instance into the standard format and return these.  For providers
	// that support paging this will also update the page and URL to be used
	// next.  If there are any errors in processing the HTTP response they are
	// returned in err and no instances will be returned
	ProcessResponse(response *http.Response) (machines *[]reporter.MachineInstance, err error)

	// Return a function that can be used to reset this provider instance so
	// that it can be used again for subsequent requests.  For providers that
	// don't support paging this is essentially a no-op.  In the case of
	// a paging provider it will recent all the relevant counters and URL to
	// the initial state
	ResetFunction() func()
}