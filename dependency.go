package main

import (
	`strings`
)

type dependency struct {
	module  string
	version string
}

func parseDependencies(originals ...string) (dependencies []*dependency) {
	dependencies = make([]*dependency, 0, len(originals))
	for _, original := range originals {
		dependencies = append(dependencies, parseDependency(original))
	}

	return
}

func parseDependency(original string) (dependency *dependency) {
	var _configs []string
	defer newDependency(_configs)

	if _configs = strings.Split(original, "@"); 2 == len(_configs) {
		return
	}
	if _configs = strings.Split(original, " "); 2 == len(_configs) {
		return
	}

	return
}

func newDependency(configs []string) (_dependency *dependency) {
	if nil != configs && 2 == len(configs) {
		_dependency = &dependency{
			module:  configs[0],
			version: configs[1],
		}
	}

	return
}
