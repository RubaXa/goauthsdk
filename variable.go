package oauth

import (
	"os"
)

type EnvVar struct {
	Name    string
	Default string
}

func (ev EnvVar) String() string {
	v := os.Getenv(ev.Name)
	if v == "" {
		v = ev.Default
	}

	return v
}

func (ev EnvVar) Get() string {
	return ev.String()
}

func GetEnvVar(name string, def ...string) EnvVar {
	v := EnvVar{Name: name}
	if len(def) == 1 {
		v.Default = def[0]
	}

	return v
}
