package cloudtwo

type provider struct {
	BaseUrl string
}

func NewProvider(baseUrl string) (*provider) {
	return &provider{
		BaseUrl: baseUrl,
	}
}