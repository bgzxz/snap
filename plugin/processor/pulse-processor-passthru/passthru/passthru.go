package passthru

import (
	"log"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
	"github.com/intelsdi-x/pulse/core/ctypes"
)

const (
	name       = "passthru"
	version    = 1
	pluginType = plugin.ProcessorPluginType
)

// Meta returns a plugin meta data
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(name, version, pluginType, []string{plugin.PulseGOBContentType}, []string{plugin.PulseGOBContentType})
}

func NewPassthruPublisher() *passthruProcessor {
	return &passthruProcessor{}
}

type passthruProcessor struct{}

func (p *passthruProcessor) GetConfigPolicyNode() cpolicy.ConfigPolicyNode {
	config := cpolicy.NewPolicyNode()
	return *config
}

func (p *passthruProcessor) Process(contentType string, content []byte, config map[string]ctypes.ConfigValue, logger *log.Logger) (string, []byte, error) {
	logger.Println("Processor started")
	//just passing through
	return contentType, content, nil
}
