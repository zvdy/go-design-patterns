# Go Design Patterns

This project demonstrates the implementation of various design patterns in Go. The design patterns are referenced from the book
[Design Patterns: Elements of Reusable Object-Oriented Software](https://en.wikipedia.org/wiki/Design_Patterns)

![Design Patterns: Elements of Reusable Object-Oriented Software](https://upload.wikimedia.org/wikipedia/en/7/78/Design_Patterns_cover.jpg)

## Design Patterns

### Behavioral Patterns

- **Iterator Pattern**: This pattern provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

- **Observer Pattern**: This pattern defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

- **Strategy Pattern**: This pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

### Creational Patterns

- **Builder Pattern**: This pattern separates the construction of a complex object from its representation so that the same construction process can create different representations.

- **Factory Pattern**: This pattern provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

- **Singleton Pattern**: This pattern ensures a class only has one instance, and provides a global point of access to it.

### Structural Patterns

- **Adapter Pattern**: This pattern converts the interface of a class into another interface clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.

- **Facade Pattern**: This pattern provides a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.
