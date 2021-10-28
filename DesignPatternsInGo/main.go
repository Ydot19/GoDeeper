package main

import (
	"github.com/Ydot19/GoDeeper/DesignPatternsInGo/solid"
)

/**
Output from solid principles directory
 */
func solidPrinciples(){
	solid.SingleResponsibilityPrinciple()
	solid.OpenClosedSolid()
	solid.LiskovSubstitutionPrinciple()
	solid.InterfaceSegregationPrinciple()
	solid.DependencyInversionPrinciple()
}

func main(){
	solidPrinciples()
}
