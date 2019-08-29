package fetch

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"path"
	"sync"

	getter "github.com/hashicorp/go-getter"
	log "github.com/sirupsen/logrus"
)

// GetPipeline parses the URL and pulls the pipeline repo to put under the AppDir
func GetPipeline(pipelineURI string, appName string) {
	log.Printf("Fetching pipeline: %s...", pipelineURI)

	// Create a go-getter client
	fetchMode := getter.ClientModeDir

	// URL to fetch from
	pipelineURL, err := url.Parse(pipelineURI)
	if err != nil {
		log.Errorf("Error parsing module URL: %v", err)
	}
	log.Printf("Pipeline URL: %v", pipelineURL)
	// Destination for pipeline module
	pwd, err := os.Getwd()
	if err != nil {
		log.Errorf("Error getting current working directory: %v", err)
	}
	destination := path.Join(pwd, appName+"/pipeline")

	opts := []getter.ClientOption{}
	ctx, cancel := context.WithCancel(context.Background())
	client := &getter.Client{
		Ctx:     ctx,
		Src:     pipelineURI,
		Dst:     destination,
		Pwd:     pwd,
		Mode:    fetchMode,
		Options: opts,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	errChan := make(chan error, 2)
	go func() {
		defer wg.Done()
		defer cancel()
		if err := client.Get(); err != nil {
			errChan <- err
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	select {
	case sig := <-c:
		signal.Reset(os.Interrupt)
		cancel()
		wg.Wait()
		log.Printf("signal %v", sig)
	case <-ctx.Done():
		wg.Wait()
		log.Printf("git fetch succeeded! cd into your app directory")
	case err := <-errChan:
		wg.Wait()
		log.Fatalf("Error downloading: %s", err)
	}
}
