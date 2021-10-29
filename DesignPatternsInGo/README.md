# Design Patterns In Golang

## Solid
Basic overview of solid in Golang

### S (Single Responsibility Principle)
- A type should only have one reason to change
- _Separation of Concerns_ - Different types/packages handling different, independent tasks/problems

[See Go File](./solid/single_responsibility_principle.go)

### O (Open-Closed Principle)
- types should be open for extension but closed for modification

[See Go File](./solid/open_closed.go)

### L (Liskov Substitution Principle)
- Write code such that there is an abilityu to substitute an embedding type in place of its embedded part

[See Go File](./solid/liskov_subsition_principle.go)

### Interface Segregation Principle 
- Don't do too much in one interface
- Split interfaces

[See Go File](./solid/interface_segregation_principle.go)

### Dependency Inversion Principle
- High-level modules should not depend upon low-level ones
- Use abstractions

[See Go File](./solid/dependency_inversion_principle.go)


## Factory Design Pattern

Motivation:
- Some objects are simple and can be created in a single constructor call
- Other objects require a lot of ceremony to create
  - Example: Factory function with 10 arguments is counter-productive

Solution:
- Opt to build objects piece-by-piece
- Builder provides an API for constructing an object step-by-step

Note: This for any entity that deals with object creation
  - Structs
  - Interfaces etc

### Types
  - [Factory Function](./factories/factory_function.go)
  - [Factory Generation](./factories/factory_generation.go)
  - [Factory Prototypes](./factories/factory_prototypes.go)
  - [Interface Factories](./factories/factory_prototypes.go)
    - Good for encapsulation of lower level interfaces/structs
  
