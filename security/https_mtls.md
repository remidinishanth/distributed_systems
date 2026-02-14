Recall:	two	kinds	of	encryption	schemes.

* E is encrypt, D is decrypt
* Symmetric key cryptography means same key is used to encrypt & decrypt
  * ciphertext = E_k(plaintext)
  * plaintext = D_k(ciphertext)
* Asymmetric key(public key) cryptography: encrypt & decrypt keys differ
  * ciphertext = E_PK(plaintext)
  * plaintext = D_SK(ciphertext)
  * PK and SK are public and secret(private) keys repectively.
* Public key cryptography	is	orders	of	magnitude	slower	than	symmetric	

## How key exchange works today

The most famous cryptographic protocol for key exchange is Diffie–Hellman, published in 1976 by Whitfield Diffie and Martin Hellman. Diffie–Hellman allows the creation of a shared secret between a sender and receiver. This shared secret is unable to be deduced by an eavesdropper who is observing the messages between the sender and receiver, except via a brute force attack. If the keyspace for the shared secret is large enough and the secret generated is sufficiently random, brute force attacks become nearly impossible. source: <https://developer.okta.com/books/api-security/tls/key-exchange/>

## Terminology

* Certificate (cert)

The public half of a public/private key pair with some additional metadata about who issued it etc. It may be freely given to anyone.

* Private Key

A private key can verify that its corresponding certificate/public key was used to encrypt data. It is never given out publicly.

* Certificate Authority (CA)

A company that issues digital certificates. For SSL/TLS certificates, there are a small number of providers (e.g. Symantec/Versign/Thawte, Comodo, GoDaddy, LetsEncrypt) whose certificates are included by most browsers and Operating Systems. They serve the purpose of a “trusted third party”.

* Certificate Signing Request (CSR)

A file generated with a private key. A CSR can be sent to a CA to request to be signed. The CA uses its private key to digitally sign the CSR and create a signed cert. Browsers can then use the CA’s cert to validate the new cert has been approved by the CA.

* HTTPS

Also called HTTP over SSL/TLS, is an extension of HTTP which encrypts communication. HTTPS URLs begin with "https://" and use port 443 by default. This is an improvement over HTTP, which is vulnerable to eavesdropping and man-in-the-middle attacks.

* SSL

Secure Sockets Layer was released by Netscape in 1995. SSL adoption increased after the redesigned SSL 3.0 was released in 1996. The IETF prohibited SSL 2.0 in 2011. SSL 3.0 was prohibited in 2015 after the IETF identified various security vulnerabilities which affected all SSL 3.0 ciphers.

* TLS

Transport Layer Security is the successor to SSL. In fact, the documentation for TLS 1.0 describes it as an "upgrade" of SSL 3.0. The current TLS version is 1.3. Although virtually all HTTPS-secured traffic uses TLS due to problems with SSL, the SSL nomenclature persists in internet culture. These days, when somebody says SSL, it is likely they mean TLS.

## HTTPS

The **Hypertext Transfer Protocol (HTTP)** is the basic communication protocol that both clients and servers must implement in order to be able to communicate. It covers things such as requests and responses, sessions, caching, authentication and more. 

The protocol transfers information between the browser and the server in clear text, allowing the network, through which the information passes, to see the information transmitted. This is a security concern, so **HTTP Secure (HTTPS)** was introduced, allowing the client and the server to first establish an encrypted communication channel, and then pass the clear text HTTP messages through it, effectively protecting them from eavesdropping.

The encrypted channel is created using the Transport Layer Security (TLS) protocol, previously called Secure Socket Layer (SSL).

Browser to Youtube.com

![image](https://user-images.githubusercontent.com/19663316/116873789-e038e000-ac35-11eb-84f0-7ffea094363c.png)

How Youtube.com creates his cert and how it is signed by Google CA. Every computer maintains a list of certificates which it trusts and are known to be legitimate.

![image](https://user-images.githubusercontent.com/19663316/116874052-49205800-ac36-11eb-8938-3147947d097a.png)

Certification authorities (CAs) are organizations trusted to sign certificates. Operating systems, such as Windows, macOS, iOS and Android, as well as the Firefox browser, have a list of trusted certificates. Chain of Trust

![image](https://user-images.githubusercontent.com/19663316/116968787-cd281d80-acd2-11eb-87e7-e9e948d1e201.png)

![image](https://user-images.githubusercontent.com/19663316/116874842-994bea00-ac37-11eb-8ce9-6cd99688b5e0.png)

![image](https://user-images.githubusercontent.com/19663316/116964662-b6c99400-acc9-11eb-845f-4e51e40e0e5b.png)


Managing our own CA

![image](https://user-images.githubusercontent.com/19663316/116874430-f1362100-ac36-11eb-8d98-46fab9eb6fae.png)

There are lots of ways for a client to authenticate itself against a server, including basic authentication, form-based authentication, and OAuth.

To prevent exposing user credentials over the wire, the client communicates with the server over HTTPS, and the server’s identify is confirmed by validating its SSL certificate. The server doesn’t necessarily care who the client is, just as long as they have the correct credentials.

An even higher level of security can be gained with using SSL certificates for both the client and the server. source: <http://www.robinhowlett.com/blog/2016/01/05/everything-you-ever-wanted-to-know-about-ssl-but-were-afraid-to-ask/>

In a traditional TLS handshake, the client authenticates the server, and the server doesn’t know too much about the client. In Client side TLS(TLS with client authentication), the server additionally authenticates that the client connecting to it is authorized to connect. TLS Client Authentication is useful in cases where a server is keeping track of hundreds of thousands or millions of clients. In the case of a mobile banking app, where the bank wants to ensure customers’ secure financial data doesn’t get stolen by bots spoofing their mobile app, they can issue a unique certificate to every app install and in the TLS handshake validate requests are coming from their mobile app.

SYN, SYN ACK, ACK happens as part of TCP, Cert handshake happens as part of TLS

![image](https://user-images.githubusercontent.com/19663316/116966486-eaa6b880-accd-11eb-9723-ad20d393bdc0.png)


![image](https://user-images.githubusercontent.com/19663316/116966365-9ac7f180-accd-11eb-9f34-3ee3b5f76194.png)


![image](https://user-images.githubusercontent.com/19663316/116964475-415dc380-acc9-11eb-8585-cc4877f502d6.png)
