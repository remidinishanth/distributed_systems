<img width="1477" height="809" alt="image" src="https://github.com/user-attachments/assets/bd801437-aa11-46bd-9043-cc8b08befb37" />


## File Storage

<img width="1279" height="470" alt="image" src="https://github.com/user-attachments/assets/3d6dbbd5-3b92-4e94-a04c-46c74b06766e" />


<img width="1123" height="809" alt="image" src="https://github.com/user-attachments/assets/4a64eb8b-dd82-430c-9c8f-95433f228006" />


<img width="1123" height="718" alt="image" src="https://github.com/user-attachments/assets/0e0bc003-9fe9-41ac-9ff7-05297ffbc3e9" />


<img width="1528" height="857" alt="image" src="https://github.com/user-attachments/assets/728a3480-4502-466b-92f5-4870040b48c6" />

## Heap File

<img width="1141" height="663" alt="image" src="https://github.com/user-attachments/assets/a96f4733-8626-4107-bc5f-e7f4cdb21f0e" />


<img width="1553" height="823" alt="image" src="https://github.com/user-attachments/assets/09afa63d-fb21-4cf6-8266-2e1a5a7a1d19" />
<img width="1473" height="336" alt="image" src="https://github.com/user-attachments/assets/32ef94dc-40bf-4fd1-a782-bab783a60b78" />

### Page directory

<img width="1521" height="788" alt="image" src="https://github.com/user-attachments/assets/15a98e36-3982-4844-aed8-9321ee352549" />


## Page Layout

<img width="1435" height="799" alt="image" src="https://github.com/user-attachments/assets/5b8e1042-35cf-4598-9b01-82044cabf80d" />

<img width="1362" height="670" alt="image" src="https://github.com/user-attachments/assets/98ea027f-3ecc-4e72-971f-45ac99349082" />

### Approach 1

<img width="1024" height="559" alt="image" src="https://github.com/user-attachments/assets/17b4f871-4d84-4887-83f8-0a83bddc4443" />

### Slotted Pages

<img width="1452" height="823" alt="image" src="https://github.com/user-attachments/assets/a8257f74-af0c-42af-82c3-4d4a28b04ad9" />

<img width="2786" height="2551" alt="image" src="https://github.com/user-attachments/assets/48824a24-8de4-42bd-bace-0536803b436d" />


<img width="1238" height="380" alt="image" src="https://github.com/user-attachments/assets/42c7c080-8109-412f-81c2-efa41b2c4c86" />

Ref: https://www.cs.swarthmore.edu/~soni/cs44/f18/Labs/lab2.html

### More concrete example

<img width="2626" height="960" alt="image" src="https://github.com/user-attachments/assets/4a0cb782-1e01-4056-983b-b58ba9438c4b" />

Ref: Internal Layout of a Heap Table File of Postgres SQL https://www.interdb.jp/pg/pgsql01/03.html

#### Writing of a Tuple

Suppose a table composed of one page that contains just one heap tuple. The pd_lower of this page points to the first line pointer, and both the line pointer and the pd_upper point to the first heap tuple. 

<img width="2578" height="542" alt="image" src="https://github.com/user-attachments/assets/863d177b-7c29-401f-87c1-1eb7b6d5254b" />

When the second tuple is inserted, it is placed after the first one. The second line pointer is appended to the first one, and it points to the second tuple. The pd_lower changes to point to the second line pointer, and the pd_upper to the second heap tuple. 

#### Reading Heap tuples

Two typical access methods, sequential scan and B-tree index scan, are outlined here:
* (a) **Sequential scan** – It reads all tuples in all pages sequentially by scanning all line pointers in each page.
* (b) **B-tree index scan** – It reads an index file that contains index tuples, each of which is composed of an index key and a TID that points to the target heap tuple.
If the index tuple with the key that you are looking for has been found1, PostgreSQL reads the desired heap tuple using the obtained TID value.

<img width="2412" height="1608" alt="image" src="https://github.com/user-attachments/assets/5e3d5b86-e9d2-43c5-8589-04304504680c" />


### Record IDs

<img width="1492" height="838" alt="image" src="https://github.com/user-attachments/assets/a7353355-8cc1-4b47-80cc-c506ae7fa5a3" />

## Tuple Layout

<img width="2382" height="990" alt="image" src="https://github.com/user-attachments/assets/51400ece-f87c-4b69-9639-c74f32372950" />

<img width="1500" height="570" alt="image" src="https://github.com/user-attachments/assets/2e04e24d-ac5f-4d0f-93d6-acd99cb0a8a1" />

<img width="1500" height="754" alt="image" src="https://github.com/user-attachments/assets/c2fa8aec-b84b-4691-8c55-52e41c7a52e1" />

### DENORMALIZED TUPLE DATA

<img width="1024" height="518" alt="image" src="https://github.com/user-attachments/assets/7a21bfb3-061d-4b27-9e5e-e66b4dc35c1e" />

<img width="1500" height="801" alt="image" src="https://github.com/user-attachments/assets/99ec8946-e550-40f5-8f87-ce9215967e49" />

## CONCLUSION
* Database is organized in pages.
* Different ways to track pages.
* Different ways to store pages.
* Different ways to store tuples.
