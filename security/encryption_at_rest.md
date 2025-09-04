## Encryption at rest

or DARE - Data at Rest encryption. 

* The data is encrypted with a DEK(Data Encryption Key) by a symmetric encryption algorithm such as AES256.
* The DEK is then encrypted using an asymmetric encryption algorithm such as RSA and this key is referred to as the KEK (Key Encryption Key).
* These encryption operations, and the keys used to support them, are typically performed on some HSM (hardware security module) the cloud provider manages behind the scenes.
* This model results in a good balance of performance (key used for bulk operations is small, KEK can be rotated without having to decrypt and re-encrypt the data) and security (larger key used to secure smaller key).
* This process is called **envelope encryption**.
<img width="1022" height="633" alt="image" src="https://github.com/user-attachments/assets/5101ac88-d060-4417-8579-e77646ee537e" />


### Symmetric encryption and Key rotation

Also refer to https://www.lambrospetrou.com/articles/encryption/ 

* Symmetric encryption with a data encryption key.


<img width="1151" height="669" alt="image" src="https://github.com/user-attachments/assets/4a3565c8-dfed-498e-9059-8d9eafa0d8da" />


<img width="806" height="429" alt="image" src="https://github.com/user-attachments/assets/bf29e8e1-792d-444a-ac38-3b8d40bbcc53" />

<img width="806" height="409" alt="image" src="https://github.com/user-attachments/assets/2278f68d-596e-47b6-ad5b-a99d7a692aee" />
