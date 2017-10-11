package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"portchecker/models"
)

const YAML = `
format_version: 0.1

default:
    mode: check
routes:
    - from: vm-from
      to: vm-target
      port: 9200
      mode: check

...`

func TestParseConfig(t *testing.T) {
	config, err := Unmarshall([]byte(YAML))
	assert.Nil(t, err)
	assert.Equal(t, "0.1", config.FormatVersion)
	assert.Equal(t, 1, len(config.Routes))
	//assert.Equal(t, "1.4.2", len(manifest.Applications[0].Spark.Version))
}

func TestParseConfigRoutes(t *testing.T) {
	config, err := Unmarshall([]byte(YAML))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(config.Routes))
	assert.Equal(t, "vm-from", config.Routes[0].From)
	assert.Equal(t, "vm-target", config.Routes[0].To)
	assert.Equal(t, 9200, config.Routes[0].Port)
	assert.Equal(t, "check", config.Routes[0].Mode)

}

func TestUnmarshallFromFile(t *testing.T) {
	config, err := UnmarshallFromFile("test/basic_yaml_file.yml")
	assert.Nil(t, err)

	assert.Equal(t, "0.1", config.FormatVersion)
	assert.Equal(t, 1, len(config.Routes))
	assert.Equal(t, "vm-from", config.Routes[0].From)
	assert.Equal(t, "vm-target", config.Routes[0].To)
	assert.Equal(t, 9200, config.Routes[0].Port)
	assert.Equal(t, "check", config.Routes[0].Mode)
	assert.Equal(t, "check", config.MainParams.Mode)

}

func TestMarshall(t *testing.T) {

	expectYaml := `format_version: "2.3"
routes:
- from: from
  to: to
  port: 9999
  mode: test
main_params:
  mode: ""
`
	array := []models.Route {
		models.Route{From:"from", To:"to", Mode:"test", Port:9999},
	}
	api := models.Config{FormatVersion: "2.3", Routes:array}

	result := Marshall(api)
	assert.NotNil(t, result)
	assert.Equal(t, expectYaml, result)
}
