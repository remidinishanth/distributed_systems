Guice (pronounced 'juice') is a lightweight dependency injection framework for Java 8 and above, brought to you by Google.

Without using Guice, this is how you will need to do things in java:

<img width="680" alt="image" src="https://github.com/remidinishanth/distributed_systems/assets/19663316/7784dbec-ca4d-4ae9-bd70-5225cc854a91">

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

Now we a module which extends AbstractModule as follows

```java
import com.google.inject.AbstractModule;

public class GuiceModule extends AbstractModule {
    @Override
    protected void configure() {
        bind(GreetingService.class).to(GreetingServiceImpl.class);
    }
}
```



In the previous example, the dependency GreetingService was passed through the constructor of MyClass. If you want to inject the dependency without passing it through the constructor, you can use the @Inject annotation on a field instead. Here's an updated version:

```
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

In this updated code, the GreetingService dependency is injected using the @Inject annotation on the field greetingService directly. The field is automatically populated by Guice when an instance of MyClass is created.

By using the @Inject annotation on the field, you don't need to pass the dependency through the constructor explicitly. Guice will take care of injecting the appropriate instance of GreetingService when creating an instance of MyClass.
