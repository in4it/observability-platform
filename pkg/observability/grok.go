package observability

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-grok"
	"github.com/in4it/go-devops-platform/storage"
)

func compilePatternDefinition(patternDefinition PatternDefinition) (*grok.Grok, error) {
	g := grok.New()
	g.AddPatterns(patternDefinition.Patterns)
	err := g.Compile("%{"+patternDefinition.Name+"}", true)
	if err != nil {
		return nil, fmt.Errorf("grok compile error: %s", err)
	}
	return g, nil
}

func getPatternDefinitions(storage storage.Iface) ([]PatternDefinition, error) {
	patternDefinition := []PatternDefinition{}
	logConfig, err := getLogConfig(storage)
	if err != nil {
		return patternDefinition, fmt.Errorf("getLogConfig error: %s", err)
	}
	return logConfig.PatternDefinitions, nil
}

func getLogConfig(storage storage.Iface) (LogConfig, error) {
	var logConfig LogConfig
	contents, err := storage.ReadFile("log-config.json")
	if err != nil {
		return logConfig, fmt.Errorf("can't read log-config.json: %s", err)
	}
	err = json.Unmarshal(contents, &logConfig)
	if err != nil {
		return logConfig, fmt.Errorf("json unmarshal error: %s", err)
	}
	return logConfig, err
}
