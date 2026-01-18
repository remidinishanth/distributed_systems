## File and Inode
<img width="1174" height="803" alt="image" src="https://github.com/user-attachments/assets/9f8f3088-4230-45b9-9371-1a6321350ac0" />

## File metadata
<img width="1174" height="726" alt="image" src="https://github.com/user-attachments/assets/e1464d40-a0bd-40a0-bc7d-7f58b5d97ba6" />

<img width="986" height="842" alt="image" src="https://github.com/user-attachments/assets/30b1697b-3c63-4345-9718-e809d8cbdc7b" />

Redirection in case of large files
<img width="1241" height="435" alt="image" src="https://github.com/user-attachments/assets/916bc27f-30c4-4914-b881-68991929c523" />

Maximum file size possible

<img width="1292" height="502" alt="image" src="https://github.com/user-attachments/assets/0f4b73be-d125-4afa-a5fd-77edf7c630b3" />

## Hard links and soft links

<img width="1003" height="412" alt="image" src="https://github.com/user-attachments/assets/0c640908-98ab-4513-a78e-75f61cfa20fd" />

<img width="1025" height="574" alt="image" src="https://github.com/user-attachments/assets/cf02d2e9-af11-453f-a069-1459e507f364" />

Soft links
<img width="1226" height="615" alt="image" src="https://github.com/user-attachments/assets/7a4ec608-9737-46a5-b740-57b44b5b2e15" />

Difference with example

<img width="2186" height="1086" alt="image" src="https://github.com/user-attachments/assets/99c78452-69ac-45b5-9374-e9752c3668d4" />

<img width="1098" height="374" alt="image" src="https://github.com/user-attachments/assets/659423a7-2172-4613-bc06-b406bdb36aa8" />

<img width="1098" height="374" alt="image" src="https://github.com/user-attachments/assets/567a4845-e416-4299-af56-bbe33553e44c" />


<img width="1024" height="559" alt="image" src="https://github.com/user-attachments/assets/5e0c423a-7587-4544-baf3-3bf58e9c3c1a" />


### 🔗 Hard Links vs. Soft Links

A **Hard Link** is a direct reference to the physical data on the disk (the inode). A **Soft Link** (or Symbolic Link) is a special file that points to another file's path (like a shortcut).

| Feature | Hard Link | Soft Link (Symlink) | Explanation |
| :--- | :---: | :---: | :--- |
| **Inode Number** | 🆔 Same | 🆕 Different | Hard links share the exact same inode as the original; soft links get a new unique inode. |
| **Cross Filesystem** | ❌ No | ✅ Yes | Hard links must stay on the same partition; soft links can point anywhere (even network drives). |
| **Link to Directory** | ❌ No | ✅ Yes | Hard linking directories is restricted (to prevent loops); soft links can point to directories easily. |
| **Delete Original** | 🛡️ File Stays | 💔 Link Breaks | If the original is deleted, hard links still work (data remains until last link is gone). Soft links become "dangling." |
| **File Size** | 💾 Original Size | 🤏 Path Length | Hard links show the size of the actual data; soft links are tiny (size = length of the path string). |
| **Speed** | ⚡ Fastest | 🐢 Slightly Slower | Hard links access data directly; soft links require an extra step to resolve the path. |
| **Command** | `ln target link` | `ln -s target link` | The `-s` flag is the key switch to create a soft link. |

## Directory and entry

Everything is a file, even directory
<img width="1104" height="375" alt="image" src="https://github.com/user-attachments/assets/8566bcb0-bd25-47c3-b655-7c00a1e3d971" />


<img width="1174" height="726" alt="image" src="https://github.com/user-attachments/assets/aa978c47-3ba3-475c-a47f-dd80aed591a1" />

* Directory: A type of file that acts as a container for other files.
* Dentry (Directory Entry): A kernel data structure representing a specific component of a path.
<img width="1024" height="559" alt="image" src="https://github.com/user-attachments/assets/bd412db3-4c26-414a-bf29-89119221d33f" />


<img width="1174" height="726" alt="image" src="https://github.com/user-attachments/assets/dcf3f7f5-6419-4a3a-accc-04612683745f" />

## 📂 What Data is Stored in an Inode?

An **inode** (index node) is a data structure in a Unix-style file system that describes a file-system object like a file or a directory. Below is a breakdown of what attributes are actually stored within the inode itself versus what is stored elsewhere.

| Data Attribute | Stored? | Explanation / Notes |
| :--- | :---: | :--- |
| **Filename** | ❌ No | Filenames are stored in **directories**, mapping names to inode numbers. |
| **Containing Directory** | ❌ No | A file can be in multiple directories (via hard links), so the inode does not track a specific parent. |
| **File Size** | ✅ Yes | Stores the size of the file in bytes. |
| **File Type** | ✅ Yes | Identifies if it is a regular file, directory, character device, etc. |
| **# of Soft Links** | ❌ No | Soft links are distinct files; the target inode does not track how many soft links point to it. |
| **Location of Soft Links** | ❌ No | The inode is unaware of where soft links pointing to it are located. |
| **# of Hard Links** | ✅ Yes | Used to track reference counts. The file is only deleted when this count reaches 0. |
| **Location of Hard Links** | ❌ No | The inode knows *how many* exist, but not *where* they are in the directory tree. |
| **Access Rights** | ✅ Yes | Stores permissions (e.g., Read, Write, Execute for User/Group/Others). |
| **Timestamps** | ✅ Yes | Tracks creation (ctime), modification (mtime), and access (atime) times. |
| **File Contents** | ⚠️ Sometimes | Generally **No** (data is in blocks), but some file systems store very small files directly in the inode (inline data). |
| **Ordered List of Data Blocks** | ✅ Yes | Contains pointers to the disk blocks where the actual file content resides. |
