<img width="2254" height="1192" alt="image" src="https://github.com/user-attachments/assets/f4688580-9958-4003-926b-85c553eefc35" />

<img width="3342" height="2310" alt="image" src="https://github.com/user-attachments/assets/f01b3652-58ee-4796-9a64-e51723fdd05c" />

The advantages of data structures make Redis a good choice for:
- Recording the number of clicks and comments for each post (hash)
- Sorting the commented user list and deduping the users (zset)
- Caching user behavior history and filtering malicious behaviors (zset, hash)
- Storing boolean information of extremely large data into small space. For example, login status, membership status. (bitmap)
