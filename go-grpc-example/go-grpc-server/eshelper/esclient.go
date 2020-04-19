package eshelper

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESClientInterface interface {
	CreateIndex() ()

}

type ESClientAdapter struct {
	esclient *elasticsearch.Client
}

func NewESClientWrapper() (*ESClientWrapper, error) {
	esclientwrapper := new(ESClientWrapper)
	esclientwrapper.esclient = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Printf("Error creating the client: %s", err)
		return nil, err
	  }
	return esclientwrapper, nil
}

func (eswrapper *ESWrapper) 