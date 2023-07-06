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
