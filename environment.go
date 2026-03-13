package main

import "fmt"

//environment struct

type Environment struct {
	values map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{
		values: make(map[string]interface{}),
	}
}

func (e *Environment) define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) get(name string) interface{} {
	value, exists := e.values[name]
	if !exists {
		fmt.Println("Undefined variable:", name)
		return nil
	}

	return value
}
