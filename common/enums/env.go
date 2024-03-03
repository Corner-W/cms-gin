package enums

import (
	"fmt"
	"strings"
)

type EnvEnum struct {
	Code        string
	NamespaceId string
	Desc        string
}

func (e *EnvEnum) String() string {
	return fmt.Sprintf("-env=%s for %s", e.Code, e.Desc)
}

var (
	LOCAL = EnvEnum{"local", "", "本地"}
	DEV   = EnvEnum{"dev", "", "开发"}
	TEST2 = EnvEnum{"test2", "", "测试"}
	PROD  = EnvEnum{"prod", "", "生产"}

	DEF_ENV = LOCAL

	envs = []EnvEnum{LOCAL, DEV, TEST2, PROD}
)

func ShowEnvs() string {
	var builder strings.Builder
	for _, env := range envs {
		builder.WriteString(env.String())
		builder.WriteString(";")
	}
	return builder.String()
}

func ToEnv(env string) EnvEnum {
	for _, e := range envs {
		if e.Code == env {
			return e
		}
	}
	return DEF_ENV
}
