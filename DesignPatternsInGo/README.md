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
