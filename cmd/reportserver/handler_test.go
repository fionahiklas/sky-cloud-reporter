package main

import (
	"github.com/fionahiklas/sky-cloud-reporter/grab"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGrabHandler(t *testing.T) {
	assert := assert.New(t)

	testProviderFactory := func() grab.CloudProvider { return nil }
	testProviderFactories := []func() grab.CloudProvider {testProviderFactory}

	testHttpClientFactory := func() grab.HttpClient { return nil }
	testGrabberFactory := func(client grab.HttpClient, provider grab.CloudProvider) grab.Grabber { return nil }

	result := NewGrabHandler(testProviderFactories, testGrabberFactory, testHttpClientFactory)

	assert.NotNil(result)

	assert.NotNil(result.CloudProviderFactories)
	assert.Equal(1, len(result.CloudProviderFactories))
	assert.NotNil(result.GrabberFactory)
	assert.NotNil(result.HttpClientFactory)
}

func TestGrabHandlerReturnsHttpHandler(t *testing.T) {
	assert := assert.New(t)

	testProviderFactory := func() grab.CloudProvider { return nil }
	testProviderFactories := []func() grab.CloudProvider {testProviderFactory}
	testHttpClientFactory := func() grab.HttpClient { return nil }
	testGrabberFactory := func(client grab.HttpClient, provider grab.CloudProvider) grab.Grabber { return nil }

	grabHandler := NewGrabHandler(testProviderFactories, testGrabberFactory, testHttpClientFactory)
	result := grabHandler.HttpHandler()

	assert.NotNil(result)
}

// TODO: Add way more tests!