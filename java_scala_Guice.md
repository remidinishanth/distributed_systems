---
layout: page
title: "Java Scala Guice"
category: "general"
---

Guice (pronounced 'juice') is a lightweight dependency injection framework for Java 8 and above, brought to you by Google.

Dependency Injection Using Guice in Java https://github.com/google/guice/wiki/GettingStarted

> Dependency injection is a design pattern wherein classes declare their dependencies as arguments instead of creating those dependencies

## Without Dependency Injection, Bad code

```java
// GreetingService.java
public interface GreetingService {
    String greet();
}

// GreetingServiceImpl.java
public class GreetingServiceImpl implements GreetingService {
    public String greet() {
        return "Hello, Normal without Guice!";
    }
}
```

```java
// MyClass.java
public class MyClass {
    private final GreetingService greetingService = new GreetingServiceImpl();

    public void performGreeting() {
        String message = greetingService.greet();
        System.out.println(message);
    }

    public static void main(String[] args) {
        MyClass myClass = new MyClass();
        myClass.performGreeting();
    }
}
```

However, this approach has some disadvantages:

* Tight coupling: By directly instantiating the dependency within the class, you create tight coupling between MyClass and GreetingServiceImpl. This can make it harder to change or swap out the implementation later on.

* Lack of flexibility: With direct instantiation, you cannot easily switch to a different implementation of GreetingService without modifying the MyClass code. This reduces the flexibility and maintainability of the codebase.

* Difficulty in testing: It can be challenging to write unit tests for MyClass when the dependency is directly instantiated within the class. Testing may require more complex setup and mocking to isolate the behavior of MyClass.

## With Dependency Injection, Without Guice
Without using Guice, this is how you will need to do things in java:

<img width="680" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/7784dbec-ca4d-4ae9-bd70-5225cc854a91">

```java
// MyClass.java
public class MyClass {
    private final GreetingService greetingService;

    public MyClass(GreetingService greetingService) {
        this.greetingService = greetingService;
    }

    public void performGreeting() {
        String message = greetingService.greet();
        System.out.println(message);
    }

    public static void main(String[] args) {
        GreetingService greetingService = new GreetingServiceImpl();
        MyClass myClass = new MyClass(greetingService);
        myClass.performGreeting();
    }
}
```

By using dependency injection, the benefits include:

Loose coupling: MyClass doesn't need to know about the specific implementation of GreetingService. It only depends on the abstraction defined by the interface GreetingService.

Flexibility: The implementation of GreetingService can be easily changed or swapped out by providing a different implementation while creating an instance of MyClass.

Testability: With dependency injection, it becomes easier to write unit tests for MyClass. You can provide a mock or stub implementation of GreetingService during testing, enabling isolated testing of MyClass behavior.

## Using Factory

Refer: https://github.com/google/guice/wiki/Motivation#factories

A factory class decouples the client and implementing class. A simple factory uses static methods to get and set mock implementations for interfaces. A factory is implemented with some boilerplate code:
```java
public interface GreetingService {
    String greet();
}

public class DefaultGreetingService implements GreetingService {
    private final String name;

    public DefaultGreetingService(String name) {
        this.name = name;
    }

    public String greet() {
        return "Hello, " + name + "!";
    }
}

public class GreetingServiceFactory {
    private static GreetingService instance;

    public static void setInstance(GreetingService service) {
        instance = service;
    }

    public static GreetingService getInstance() {
        if (instance == null) {
            return new DefaultGreetingService("John Doe");
        }

        return instance;
    }
}

public class MyClass {
    private final GreetingService greetingService;

    public MyClass() {
        this.greetingService = GreetingServiceFactory.getInstance();
    }

    public void performGreeting() {
        String message = greetingService.greet();
        System.out.println(message);
    }

    public static void main(String[] args) {
        MyClass myClass = new MyClass();
        myClass.performGreeting();
    }
}
```

So we can now test it like the following

```java
public class MyClassTest {
    @Test
    public void testPerformGreetingWithFakeService() {
        // Create a fake implementation of GreetingService
        GreetingService fakeService = new FakeGreetingService();

        // Set the fake implementation in the factory
        GreetingServiceFactory.setInstance(fakeService);

        // Create an instance of MyClass
        MyClass myClass = new MyClass();

        // Perform the greeting
        myClass.performGreeting();

        // Verify the output
        assertEquals("Hello, Fake User!", getConsoleOutput());
    }

    // Helper method to capture the console output
    private String getConsoleOutput() {
        // Replace this with your own implementation to capture the console output during the test
        return ""; // Placeholder implementation
    }

    // Fake implementation of GreetingService
    private static class FakeGreetingService implements GreetingService {
        public String greet() {
            return "Hello, Fake User!";
        }
    }
}
```

## With Guice

By using a dependency injection framework like Guice, you can further automate and streamline the injection process, manage the lifecycle of dependencies, and handle complex dependency graphs. It simplifies the management of dependencies and promotes modular design.

When we use Guice, We will need to modify the main to use Injector to get the instance of `MyClass` but we don't need to set up `GreetingService` explicitly in the main

```java
import com.google.inject.Guice;
import com.google.inject.Inject;
import com.google.inject.Injector;

public class MyClass {
    private final GreetingService greetingService;

    @Inject
    public MyClass(GreetingService greetingService) {
        this.greetingService = greetingService;
    }

    public void performGreeting() {
        String message = greetingService.greet();
        System.out.println(message);
    }

    public static void main(String[] args) {
        Injector injector = Guice.createInjector(new GuiceModule());
        MyClass myClass = injector.getInstance(MyClass.class);
        myClass.performGreeting();
    }
}
```

Now we need module which extends inject.AbstractModule as follows, this tell that whenever we need GreetingService, pass GreetingServiceImpl, that way, we need not explicitly instantiate `GreetingServiceImpl` inside main class, so this can solve our problem when we need to recursively inject many classes

```java
import com.google.inject.AbstractModule;

public class GuiceModule extends AbstractModule {
    @Override
    protected void configure() {
        bind(GreetingService.class).to(GreetingServiceImpl.class);
    }
}
```


### Without using Constructor

In the previous example, the dependency `GreetingService` was passed through the constructor of `MyClass`. If you want to inject the dependency without passing it through the constructor, you can use the `@Inject` annotation on a field instead. Here's an updated version:

```java
import com.google.inject.Guice;
import com.google.inject.Inject;
import com.google.inject.Injector;

public class MyClass {
    @Inject
    private GreetingService greetingService;

    public void performGreeting() {
        String message = greetingService.greet();
        System.out.println(message);
    }

    public static void main(String[] args) {
        Injector injector = Guice.createInjector(new GuiceModule());
        MyClass myClass = injector.getInstance(MyClass.class);
        myClass.performGreeting();
    }
}
```

In this updated code, the `GreetingService` dependency is injected using the `@Inject` annotation on the field `greetingService` directly. The field is automatically populated by Guice when an instance of `MyClass` is created.

By using the `@Inject` annotation on the field, you don't need to pass the dependency through the constructor explicitly. Guice will take care of injecting the appropriate instance of `GreetingService` when creating an instance of MyClass.
