# Go Design Patterns

A comprehensive collection of design patterns implemented in Go, organized by category. This repository demonstrates practical implementations of SOLID principles, creational patterns, structural patterns, and behavioral patterns.

## Table of Contents

- [SOLID Principles](#solid-principles)
- [Creational Patterns](#creational-patterns)
- [Structural Patterns](#structural-patterns)
- [Behavioral Patterns](#behavioral-patterns)
- [Getting Started](#getting-started)

---

## SOLID Principles

The foundation of object-oriented design. These five principles help create more maintainable and scalable software.

| Pattern | Description | Location |
|---------|-------------|----------|
| **S**ingle Responsibility Principle | A class should have only one reason to change | [01.SOLID/01.SingleResponsabilitiyPrinciple](01.SOLID/01.SingleResponsabilitiyPrinciple) |
| **O**pen Closed Principle | Software entities should be open for extension but closed for modification | [01.SOLID/02.OpenClosedPrinciple](01.SOLID/02.OpenClosedPrinciple) |
| **L**iskov Substitution Principle | Objects of superclasses should be replaceable with objects of subclasses | [01.SOLID/03.LiskovSubstitutionPrinciple](01.SOLID/03.LiskovSubstitutionPrinciple) |
| **I**nterface Segregation Principle | Clients should not depend on interfaces they don't use | [01.SOLID/04.InterfaceSegregationPrinciple](01.SOLID/04.InterfaceSegregationPrinciple) |
| **D**ependency Injection Principle | High-level modules should not depend on low-level modules; both should depend on abstractions | [01.SOLID/05.DependencyInjectionPrinciple](01.SOLID/05.DependencyInjectionPrinciple) |

---

## Creational Patterns

Patterns that deal with object creation mechanisms, trying to create objects in a manner suitable to the situation.

### Builder Pattern
Separate the construction of a complex object from its representation.

| Variant | Description | Location |
|---------|-------------|----------|
| Classic Builder | Traditional builder implementation | [02.CreationalBuilder/01.Builder](02.CreationalBuilder/01.Builder) |
| Builder with Facets | Using multiple builders for different aspects | [02.CreationalBuilder/02.BuilderWithFacets](02.CreationalBuilder/02.BuilderWithFacets) |
| Builder Parameter | Using builder as function parameters | [02.CreationalBuilder/03.BuilderParameter](02.CreationalBuilder/03.BuilderParameter) |
| Functional Builder | Builder using functional approaches | [02.CreationalBuilder/04.BuilderFunctional](02.CreationalBuilder/04.BuilderFunctional) |

### Factory Patterns
Create objects without specifying their concrete classes.

| Variant | Description | Location |
|---------|-------------|----------|
| Factory Function | Simple factory using functions | [03.CreationalFactories/01.FactoryFunction](03.CreationalFactories/01.FactoryFunction) |
| Interface Factory | Factory based on interfaces | [03.CreationalFactories/02.InterfaceFactory](03.CreationalFactories/02.InterfaceFactory) |
| Factory Generator | Function that generates factories | [03.CreationalFactories/03.FactoryGenerator](03.CreationalFactories/03.FactoryGenerator) |
| Prototype Factory | Using prototypes to create objects | [03.CreationalFactories/04.PrototypeFactory](03.CreationalFactories/04.PrototypeFactory) |

### Prototype Pattern
Create new objects by copying an existing object (prototype).

| Variant | Description | Location |
|---------|-------------|----------|
| Deep Copying | Creating deep copies of objects | [04.CreationalPrototype/01.DeepCopying](04.CreationalPrototype/01.DeepCopying) |
| Copy Method | Using dedicated copy methods | [04.CreationalPrototype/02.CopyMethod](04.CreationalPrototype/02.CopyMethod) |
| Copy Through Serialization | Cloning objects via serialization | [04.CreationalPrototype/03.CopyThroughSerialization](04.CreationalPrototype/03.CopyThroughSerialization) |
| Prototype Factory | Factory pattern combined with prototypes | [04.CreationalPrototype/04.PrototypeFactory](04.CreationalPrototype/04.PrototypeFactory) |

---

## Structural Patterns

Patterns that deal with object composition, creating relationships between entities.

| Pattern | Description | Location |
|---------|-------------|----------|
| **Singleton** | Ensure a class has only one instance | [05.StructuralSingleton](05.StructuralSingleton) |
| **Adapter** | Convert the interface of a class into another interface clients expect | [06.StructuralAdapter](06.StructuralAdapter) |
| **Bridge** | Decouple an abstraction from its implementation | [07.StructuralBridge](07.StructuralBridge) |
| **Composite** | Compose objects into tree structures to represent part-whole hierarchies | [08.StructuralComposite](08.StructuralComposite) |
| **Decorator** | Attach additional responsibilities to an object dynamically | [09.StructuralDecorator](09.StructuralDecorator) |
| **Facade** | Provide a unified interface to a set of interfaces | [10.StructuralFacade](10.StructuralFacade) |
| **Flyweight** | Use sharing to support large numbers of fine-grained objects efficiently | [11.StructuralFlyweight](11.StructuralFlyweight) |
| **Proxy** | Provide a surrogate or placeholder for another object | [12.StructuralProxy](12.StructuralProxy) |

---

## Behavioral Patterns

Patterns that deal with object collaboration and responsibility distribution.

| Pattern | Description | Location |
|---------|-------------|----------|
| **Chain of Responsibility** | Pass requests along a chain of handlers | [13.BehavioralChainOfResponsability](13.BehavioralChainOfResponsability) |
| **Command** | Encapsulate a request as an object | [14.BehavioralCommand](14.BehavioralCommand) |
| **Interpreter** | Define a grammar for a language and interpret sentences | [15.BehavioralInterpreter](15.BehavioralInterpreter) |
| **Iterator** | Provide a way to access elements sequentially | [16.BehavioralIterator](16.BehavioralIterator) |
| **Mediator** | Define an object that encapsulates how a set of objects interact | [17.BehavioralMediator](17.BehavioralMediator) |
| **Memento** | Capture and externalize an object's internal state | [18.BehavioralMemento](18.BehavioralMemento) |
| **Observer** | Define a one-to-many dependency between objects | [19.BehaviloralObserver](19.BehaviloralObserver) |
| **State** | Allow an object to alter its behavior when its state changes | [20.BehavioralState](20.BehavioralState) |
| **Strategy** | Define a family of algorithms and encapsulate each one | [21.BehavioralStrategy](21.BehavioralStrategy) |
| **Template Method** | Define the skeleton of an algorithm in a method | [22.BehavioralTemplateMethod](22.BehavioralTemplateMethod) |
| **Visitor** | Represent an operation to be performed on elements | [23BehavioralVisitor](23BehavioralVisitor) |

---

## Getting Started

### Prerequisites

- Go 1.16 or later
- Basic understanding of design patterns

### Running Tests

Each pattern implementation includes tests. To run tests for a specific pattern:

```bash
# Navigate to a pattern directory
cd 01.SOLID/01.SingleResponsabilitiyPrinciple

# Run the tests
go test -v
```

To run all tests in the repository:

```bash
go test ./...
```

### Exploring the Code

Each pattern directory contains:
- `main.go` - Implementation of the pattern
- `main_test.go` - Test cases demonstrating the pattern

Start by reading the code and tests to understand how each pattern works in Go.

---

## Structure

```
DesignPatterns/
├── 01.SOLID/                          # SOLID Principles
├── 02.CreationalBuilder/              # Builder Pattern variations
├── 03.CreationalFactories/            # Factory Pattern variations
├── 04.CreationalPrototype/            # Prototype Pattern variations
├── 05.StructuralSingleton/            # Singleton Pattern
├── 06.StructuralAdapter/              # Adapter Pattern
├── 07.StructuralBridge/               # Bridge Pattern
├── 08.StructuralComposite/            # Composite Pattern
├── 09.StructuralDecorator/            # Decorator Pattern
├── 10.StructuralFacade/               # Facade Pattern
├── 11.StructuralFlyweight/            # Flyweight Pattern
├── 12.StructuralProxy/                # Proxy Pattern
├── 13.BehavioralChainOfResponsability/# Chain of Responsibility Pattern
├── 14.BehavioralCommand/              # Command Pattern
├── 15.BehavioralInterpreter/          # Interpreter Pattern
├── 16.BehavioralIterator/             # Iterator Pattern
├── 17.BehavioralMediator/             # Mediator Pattern
├── 18.BehavioralMemento/              # Memento Pattern
├── 19.BehavioralObserver/             # Observer Pattern
├── 20.BehavioralState/                # State Pattern
├── 21.BehavioralStrategy/             # Strategy Pattern
├── 22.BehavioralTemplateMethod/       # Template Method Pattern
├── 23.BehavioralVisitor/              # Visitor Pattern
└── go.mod                             # Go module file
```

---

## References

- [Design Patterns: Elements of Reusable Object-Oriented Software](https://en.wikipedia.org/wiki/Design_Patterns) - The Gang of Four
- [SOLID Principles](https://en.wikipedia.org/wiki/SOLID) - Object-oriented design principles
- [Go Documentation](https://golang.org/doc/)
