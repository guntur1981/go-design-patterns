# Design Patterns in Go

A collection of design patterns concept and their implementation in Go from various sources.

## Table of Contents

1. [Single Responsibility Principle (SOLID)](/01-single-responsibility-principle)
   > A type should be made for one specific purpose.
2. [Open-Closed Principle (SOLID)](/02-open-closed-principle)
   > A type should be open for extension, but closed for modification.
3. [Liskov Substitution Principle (SOLID)](/03-liskov-substitution-principle)
   > All behaviours of an interface type should be implemented without any problems.
4. [Interface Segregation Principle (SOLID)](/04-interface-segregation-principle)
   > An interface type should not impose any behavior that is not required.
5. [Dependency Inversion Principle (SOLID)](/05-dependency-inversion-principle)
   > High-level modules/packages should not depend on low-level modules/packages except through interfaces.
6. [Builder Design Pattern](/06-builder-design-pattern)
   > Creates complex objects **step by step** (**method chaining**).
7. [Factory Design Pattern](/07-factory-design-pattern)
   > Creates objects **in one go** without exposing the creation logic to the user.
8. [Prototype Design Pattern](/08-prototype-design-pattern)
   > Creates objects by copying from an existing object (**deep copying**).
9. [Singleton Design Pattern](/09-singleton-design-pattern)
   > Makes sure that a struct only has one instance.
10. [Adapter Design Pattern](/10-adapter-design-pattern)
    > A struct which adapts an existing interface X to conform to the required interface Y
11. [Bridge Design Pattern](/11-bridge-design-pattern)
    > Decouples an interface from an implementation, preventing complexity explosion.
12. [Composite Design Pattern](/12-composite-design-pattern)
    > Treats individual (scalar) objects and compositions of objects in a uniform manner.
13. [Decorator Design Pattern](/13-decorator-design-pattern)
    > Facilitates the addition of behaviors to individual objects through embedding.
14. [Facade Design Pattern](/14-facade-design-pattern)
    > Provides a simplified interface to a complex and large body of code.
15. [Flyweight Design Pattern](/15-flyweight-design-pattern)
    > A space optimization technique, use less memory by avoiding redundancy when storing data.
16. [Proxy Design Pattern](/16-proxy-design-pattern)
    > Controls access to the original object, perform additional actions before or after accessing the object.
17. [Chain of Responsibility Design Pattern](/17-chain-of-responsibility-design-pattern)
    > Allows an object to pass a request along a chain of potential handlers until the request is handled.
18. [Command Design Pattern](/18-command-design-pattern)
    > A stand-alone object represents a command to perform a particular action.
19. [Interpreter Design Pattern](/19-interpreter-design-pattern)
    > Provides a way to interpret and evaluate sentences or expressions in a language.
20. [Iterator Design Pattern](/20-iterator-design-pattern)
    > Facilitates the traversal of various data structures in a controlled manner.
21. [Mediator Design Pattern](/21-mediator-design-pattern)
    > Facilitates communication between other components without them having direct access to each other.
22. [Memento Design Pattern](/22-memento-design-pattern)
    > Rollbacks the system states arbitrarily.
23. [Observer Design Pattern](/23-observer-design-pattern)
    > Propagates a change in an object to other objects without tightly coupling them.
24. [State Design Pattern](/24-state-design-pattern)
    > Changes object's behaviour when its internal state change.
25. [Strategy Design Pattern](/25-strategy-design-pattern)
    > Separates a mechanism into its abstract and concrete implementation steps, which can be varied at run time.
26. [Visitor Design Pattern](/26-visitor-design-pattern)
    > Allows to add further operations to objects without having to modify them.
