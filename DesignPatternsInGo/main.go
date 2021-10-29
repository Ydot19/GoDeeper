package main

import (
	"github.com/Ydot19/GoDeeper/DesignPatternsInGo/factories"
	"github.com/Ydot19/GoDeeper/DesignPatternsInGo/solid"
)

/**
solidPrinciples Output from solid principles directory
 */
func solidPrinciples(){
	solid.SingleResponsibilityPrinciple()
	solid.OpenClosedSolid()
	solid.LiskovSubstitutionPrinciple()
	solid.InterfaceSegregationPrinciple()
	solid.DependencyInversionPrinciple()
}

/**
factoryDesignPatter Out of builder design pattern
 */
func factoryDesignPatter(){
	factories.FactoryFunction()
	factories.InterfaceFactories()
	factories.FactoryGeneration()
}

func main(){
	factoryDesignPatter()
}
