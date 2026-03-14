package main

import "fmt"

type Environment struct {
	values    map[string]interface{}
	enclosing *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		values: make(map[string]interface{}),
	}
}

func (e *Environment) define(name string, value interface{}) {
	e.values[name] = value
}

// FIX: Added proper assignment logic that checks if the variable exists
func (e *Environment) assign(name string, value interface{}) {
	if _, exists := e.values[name]; exists {
		e.values[name] = value
		return
	}
	if e.enclosing != nil {
		e.enclosing.assign(name, value)
		return
	}
	fmt.Println("Undefined variable for assignment:", name)
}
func (e *Environment) get(name string) interface{} {
	value, exists := e.values[name]
	if exists {
		return value
	}
	if e.enclosing != nil {
		return e.enclosing.get(name)
	}

	fmt.Println("Undefined variable: ", name)
	return nil
}
