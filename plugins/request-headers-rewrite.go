package plugins

import (
	"encoding/json"
	"net/http"

	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
)

const requestBodyRewriteName = "request-body-rewrite"

func init() {
	if err := plugin.RegisterPlugin(&RequestBodyRewrite{}); err != nil {
		log.Fatalf("failed to register plugin %s: %s", requestBodyRewriteName, err.Error())
	}
}

type RequestBodyRewrite struct {
	plugin.DefaultPlugin
}

type RequestBodyRewriteConfig struct {
	NewBody string `json:"new_body"`
}

func (*RequestBodyRewrite) Name() string {
	return requestBodyRewriteName
}

func (p *RequestBodyRewrite) ParseConf(in []byte) (interface{}, error) {
	conf := RequestBodyRewriteConfig{}
	err := json.Unmarshal(in, &conf)
	if err != nil {
		log.Errorf("failed to parse config for plugin %s: %s", p.Name(), err.Error())
	}
	return conf, err
}

func (*RequestBodyRewrite) RequestFilter(conf interface{}, _ http.ResponseWriter, r pkgHTTP.Request) {
	newBody := conf.(RequestBodyRewriteConfig).NewBody

	if newBody == "" {
		return
	}

	// Log the "Hello World" message
	log.Infof("Hello World from the RequestBodyRewrite plugin")

	// Read the current body
	body, err := r.Body()

	if err != nil {
		log.Errorf("failed to read request body: %s", err)
		return
	}

	// Log the original body
	log.Infof("Original request body: %s", string(body))

	// Replace the body with the new body
	// Unfortunately, the go-runner does not support direct modification of request body
	// This might require handling at the application level or through a different approach

	// Log the new body to be set
	log.Infof("New request body: %s", newBody)
}
