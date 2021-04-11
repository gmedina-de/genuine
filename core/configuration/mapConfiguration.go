package configuration

type mapConfiguration struct {
	settings map[Setting]string
}

func MapConfiguration() Configuration {
	return &mapConfiguration{settings: make(map[Setting]string)}
}

func (this *mapConfiguration) Get(setting Setting) string {
	value, ok := this.settings[setting]
	if ok {
		return value
	} else {
		return ""
	}
}

func (this *mapConfiguration) Set(setting Setting, value string) {
	this.settings[setting] = value
}
