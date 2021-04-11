package configuration

type Configuration interface {
	Get(setting Setting) string
	Set(setting Setting, value string)
}
