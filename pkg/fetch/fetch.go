package fetch

import (
	"net/url"
	"os"
	"path"

	getter "github.com/hashicorp/go-getter"
	log "github.com/sirupsen/logrus"
)

// GetPipeline parses the URL and pulls the pipeline repo to put under the AppDir
func GetPipeline(pipelineURI string, appName string) error {
	log.Printf("Fetching pipeline: %s...", pipelineURI)

	// URL to fetch from
	pipelineURL, err := url.Parse(pipelineURI)
	if err != nil {
		log.Errorf("Error parsing module URL: %v", err)
		return err
	}
	log.Printf("Pipeline URL: %v", pipelineURL)

	// Destination for pipeline module
	pwd, err := os.Getwd()
	if err != nil {
		log.Errorf("Error getting current working directory: %v", err)
		return err
	}

	destination := path.Join(pwd, appName+"/src")
	err = getter.GetAny(destination, pipelineURI)
	if err != nil {
		log.Errorf("Error getting Pipeline from URL: %v", err)
		return err
	}

	return nil
}
