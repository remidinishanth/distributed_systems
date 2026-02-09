---
layout: page
title: "Password Plain Text"
category: "general"
---


![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/c0075e09-b77e-46ba-8286-b858f2694c05)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/f2a6d98a-670b-480c-8a3b-a2e6ef835dd3)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/a7ef25dd-0bb9-4796-ae88-5c571658073c)

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/eafb4ec0-634f-4ada-91ea-51b14d4127f8)


Ref: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#password-hashing-algorithms

𝐖𝐡𝐚𝐭 𝐢𝐬 𝐬𝐚𝐥𝐭?
According to OWASP guidelines, “a salt is a unique, randomly generated string that is added to each password as part of the hashing process”.
 
𝐇𝐨𝐰 𝐭𝐨 𝐬𝐭𝐨𝐫𝐞 𝐚 𝐩𝐚𝐬𝐬𝐰𝐨𝐫𝐝 𝐚𝐧𝐝 𝐬𝐚𝐥𝐭?
1️ A salt is not meant to be secret and it can be stored in plain text in the database. It is used to ensure the hash result is unique to each password.
 
2️ The password can be stored in the database using the following format: 𝘩𝘢𝘴𝘩( 𝘱𝘢𝘴𝘴𝘸𝘰𝘳𝘥 + 𝘴𝘢𝘭𝘵).

![image](https://github.com/remidinishanth/distributed_systems/assets/19663316/8e304c2a-82b1-4b8f-824d-f62d331eb7bd)

If you see in the above, password is never sent over wire to the server.

Go has input support using `crypto` for Argon2, bcrypt and PBKDF2

Using a slow, expensive hashing algorithm is recommended. These hashes introduce a calculation work factor which can then be scaled along with Moore’s Law. Some examples are Argon2, PBKDF2, bcrypt, and scrypt. While Argon2 is better, Bcrypt is still very good and scales with modern hardware due to the work factor. Go also provides us with a easy to implement library which makes this a good choice.
