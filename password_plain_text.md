
![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/c0075e09-b77e-46ba-8286-b858f2694c05)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/f2a6d98a-670b-480c-8a3b-a2e6ef835dd3)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a7ef25dd-0bb9-4796-ae88-5c571658073c)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/eafb4ec0-634f-4ada-91ea-51b14d4127f8)


Ref: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#password-hashing-algorithms


Go has input support using `crypto` for Argon2, bcrypt and PBKDF2

Using a slow, expensive hashing algorithm is recommended. These hashes introduce a calculation work factor which can then be scaled along with Moore’s Law. Some examples are Argon2, PBKDF2, bcrypt, and scrypt. While Argon2 is better, Bcrypt is still very good and scales with modern hardware due to the work factor. Go also provides us with a easy to implement library which makes this a good choice.
