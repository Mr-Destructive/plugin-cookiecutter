package {{cookiecutter.plugin_name}}

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type {{cookiecutter.plugin_name}}Config struct {
}

var ConfigSchema = map[string]*schema.Attribute{
}

func ConfigInstance() interface{} {
	return &{{cookiecutter.plugin_name}}Config{}
}

func GetConfig(connection *plugin.Connection) {{cookiecutter.plugin_name}}Config {
	if connection == nil || connection.Config == nil {
		return {{cookiecutter.plugin_name}}Config{}
	}
	config, _ := connection.Config.({{cookiecutter.plugin_name}}Config)
	return config
}
