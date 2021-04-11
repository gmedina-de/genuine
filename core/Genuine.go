package core

import (
	"fmt"
	. "reflect"
	"strings"
)

var (
	implementations = make(map[Type]interface{})
	instances       = make(map[Type]Value)
	currentLevel    = -1
)

func Constructor(constructor interface{}) {
	constructorReturn := ValueOf(constructor).Type().Out(0)
	implementations[constructorReturn] = constructor
}

func Genuine(applicationConstructor interface{}) Application {
	return inject(applicationConstructor).Interface().(Application)
}

func inject(constructor interface{}) Value {
	currentLevel++

	constructorValue := ValueOf(constructor)
	constructorType := constructorValue.Type()
	parameters := make([]Value, constructorType.NumIn())

	fmt.Println(strings.Repeat("\t", currentLevel),"Constructor function is", constructorType)
	for i := 0; i < len(parameters); i++ {
		parameterType := constructorType.In(i)
		parameterName := parameterType.Name()
		instance, found := instances[parameterType]
		if !found {
			fmt.Println(strings.Repeat("\t", currentLevel),parameterName, "was not yet instantiated")
			constructor, found := implementations[parameterType]
			if !found {
				panic("No constructor found for " + parameterName + ", required for dependency injection, please provide one")
			}

			instance = inject(constructor)
			instances[parameterType] = instance
		}
		parameters[i] = instance
	}

	fmt.Println(strings.Repeat("\t", currentLevel),"Instantiating with parameters", parameters)
	currentLevel--
	return constructorValue.Call(parameters)[0]
}
