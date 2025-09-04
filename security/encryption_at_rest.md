## Encryption at rest

or DARE - Data at Rest encryption. 

* The data is encrypted with a DEK(Data Encryption Key) by a symmetric encryption algorithm such as AES256.
* The DEK is then encrypted using an asymmetric encryption algorithm such as RSA and this key is referred to as the KEK (Key Encryption Key).
* These encryption operations, and the keys used to support them, are typically performed on some HSM (hardware security module) the cloud provider manages behind the scenes.
* This model results in a good balance of performance (key used for bulk operations is small, KEK can be rotated without having to decrypt and re-encrypt the data) and security (larger key used to secure smaller key).
* This process is called **envelope encryption**.
<img width="1022" height="633" alt="image" src="https://github.com/user-attachments/assets/5101ac88-d060-4417-8579-e77646ee537e" />
