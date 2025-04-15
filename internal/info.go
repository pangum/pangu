package internal

import (
	"runtime"

	"github.com/pangum/core/internal/internal/constant"
	"github.com/pangum/core/internal/internal/os"
)

var (
	Name     = constant.EnvironmentNotSet
	Version  = constant.EnvironmentNotSet
	Build    = constant.EnvironmentNotSet
	Compiled = constant.EnvironmentNotSet
	Revision = constant.EnvironmentNotSet
	Branch   = constant.EnvironmentNotSet
	Runtime  = runtime.Version()
)

func GetName() string {
	return os.GetEnvironment(constant.EnvironmentName, Name, constant.EnvironmentNotSet)
}

func GetVersion() string {
	return os.GetEnvironment(constant.EnvironmentVersion, Version, constant.EnvironmentNotSet)
}

func GetBuild() string {
	return os.GetEnvironment(constant.EnvironmentBuild, Build, constant.EnvironmentNotSet)
}

func GetCompiled() string {
	return os.GetEnvironment(constant.EnvironmentCompiled, Compiled, constant.EnvironmentNotSet)
}

func GetRevision() string {
	return os.GetEnvironment(constant.EnvironmentRevision, Revision, constant.EnvironmentNotSet)
}

func GetBranch() string {
	return os.GetEnvironment(constant.EnvironmentBranch, Branch, constant.EnvironmentNotSet)
}
