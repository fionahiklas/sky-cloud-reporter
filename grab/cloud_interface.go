//go:generate mockgen -package=mock_grab -destination=../mocks/mock_grab/mock_cloud_interfaces.go . CloudProvider

package grab

type CloudProvider interface {
	GetInstanceUrl() string

}