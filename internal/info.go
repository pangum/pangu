package internal

import (
	"runtime"

	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/os"
)

var (
	Name     = os.GetEnvironment(constant.EnvironmentName)
	Version  = os.GetEnvironment(constant.EnvironmentVersion)
	Build    = os.GetEnvironment(constant.EnvironmentBuild)
	Compiled = os.GetEnvironment(constant.EnvironmentCompiled)
	Revision = os.GetEnvironment(constant.EnvironmentRevision)
	Branch   = os.GetEnvironment(constant.EnvironmentBranch)
	Runtime  = runtime.Version()
)
