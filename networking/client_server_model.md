# Introduction to TCP/IP (Part 3) - Client Server Model

> **Source:** [Microchip Developer Help](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/client-server/)

## Client Server Model Defined

The **client-server programming model** is a distributed computing architecture that segregates information users (clients) from information providers (servers).

### Client
- An application that **needs something** (like a web page or IP address) from a server
- May contact a server at any time
- Information **users**

### Server
- An application that **provides information or resources** to clients
- Must be always up and running, waiting for requests
- Information **providers**

### Key Principle
- Client applications communicate **only with server applications** and vice versa
- Clients do **not** communicate directly with other clients

### Alternative: Peer to Peer (P2P)
In P2P architectures, two or more hosts communicate directly with each other (no client-server distinction).

---

## Client Server Examples

### Example: DHCP Client Server

<img width="576" height="255" alt="image" src="https://github.com/user-attachments/assets/35b58583-d88f-45e5-b496-8ee518b99880" />


The **Dynamic Host Configuration Protocol (DHCP)** is responsible for requesting and offering IP addresses.

| Component | Role | Description |
|-----------|------|-------------|
| DHCP Client | Requests IP | Automatically requests an IP address when a network is detected |
| DHCP Server | Provides IP | Always active, ready to respond to client requests |

- DHCP server typically exists in a **router**
- May also run on a network server for larger networks

### Example: HTTP Client and Server

<img width="575" height="345" alt="image" src="https://github.com/user-attachments/assets/abf15033-b7c6-4c14-a0ae-c2fd6656ea08" />


**Scenario:** Using an HTTP client on a PC to control home lights

1. HTTP client runs on a home lighting control board
2. Control board monitors a lighting control website on an Internet web server
3. User browses to the same webpage, enters credentials
4. User changes the webpage settings
5. Control board checks the webpage and controls lights accordingly

### Example: HTTP Client and Server in the Same Host

<img width="577" height="342" alt="image" src="https://github.com/user-attachments/assets/70b40d8e-446e-4737-b63c-878be2aeedee" />


A network host can be **both a client and a server** simultaneously.

**Example:**
- Control board runs an **HTTP Client** (to check lighting control website)
- Control board also runs an **HTTP Server** (to serve a setup/configuration page)
- Configuration page allows changing the website and login info used by the HTTP client

### Example: Local Network HTTP Server

<img width="577" height="345" alt="image" src="https://github.com/user-attachments/assets/64445013-69fa-4d68-b3d6-7228d941b7e6" />


An HTTP server on an embedded device can directly control the device:
- Eliminates need for HTTP client application and Internet web server
- **Easiest solution** if HTTP client is on the same local network
- **Challenge:** Accessing from the Internet is not trivial

---

## Internet Server vs. Local Network Server

### Internet Server

<img width="577" height="302" alt="image" src="https://github.com/user-attachments/assets/920cef5c-51b9-4560-87cc-2f2256035045" />


Accessing an HTTP server on the Internet is **effortless**.

**Architecture:**
- Webpage runs on an HTTP server in the Internet
- Embedded HTTP client posts status and polls for commands
- Web browser (PC/smartphone) monitors and controls via the webpage

**Hosting Options:**
- Shared web hosting (GoDaddy™, Network Solutions®)
- Cloud services (Amazon Web Services, Microsoft Azure®)
- Service providers allow you to choose a website name

### Local Network Server

<img width="577" height="288" alt="image" src="https://github.com/user-attachments/assets/bc19c6b8-c5f1-498f-90f5-56a211ef99be" />


HTTP server located on a local network (running on the embedded device itself).

**Architecture:**
- Webpage runs directly on the embedded HTTP server
- Web browser accesses the embedded device directly

**Challenges** that Internet servers don't have:
- Firewall restrictions
- No website name by default
- Dynamic IP address issues

---

## Local Network Server Obstacles and Solutions

### Obstacles

1. **Firewall** - Router restricts public access to local network
2. **No website name** - Remote clients must use IP address
3. **Dynamic IP address** - ISP may change your IP address


Firewall
<img width="1600" height="900" alt="image" src="https://github.com/user-attachments/assets/5498ae25-c6cc-4308-ba29-b17627d55b4e" />


<img width="577" height="313" alt="image" src="https://github.com/user-attachments/assets/d715ff05-579b-4095-8966-7ad4765bc300" />


### Solution 1: Port Forwarding

When you receive incoming traffic on port 80 (HTTP), don't block it. Instead, forward it to the internal device at 192.168.1.101

<img width="577" height="291" alt="image" src="https://github.com/user-attachments/assets/b1935f06-5a1d-43f4-b830-2a93622d4926" />


Port forwarding allows Internet packets destined for a particular port to be forwarded to a specific local network IP address.

**Configuration:**
- Configure router to forward incoming HTTP server packets to the local IP address
- Require username/password for unauthorized access prevention
- Use **SSL encryption** to prevent packet sniffing

> **Note:** Properly implemented and secure port forwarding is also called a "**firewall pinhole**"

### Solution 2: Website Name (Domain Name)

<img width="577" height="341" alt="image" src="https://github.com/user-attachments/assets/2872d535-cb12-40d0-a94a-8c627fd61d5b" />


**Advantages of a website name:**
- Easier to remember than IP addresses
- If IP address changes, DNS servers can be automatically updated

**Domain name providers:**
- GoDaddy
- DynDNS

### Solution 3: Dynamic IP Address Handling

<img width="577" height="339" alt="image" src="https://github.com/user-attachments/assets/24df549b-3add-41b1-8d9c-8d80b20b45f6" />


**Understanding Dynamic vs Static IP:**

| Type | Description | Typical Use |
|------|-------------|-------------|
| Dynamic IP | Can change at any time | Assigned by ISPs to customers |
| Static IP | Fixed, never changes | Internet routers and servers |

> Static IP addresses cost more from ISPs

**Solution for Dynamic IP:**
1. Use website name instead of IP address to access your server
2. Update DNS server with new IP address when it changes
3. Some routers have this capability built-in
4. Embedded device can periodically check its IP and update DNS if changed

**DNS Service Providers:**
- DynDNS
- ZoneEdit

<img width="577" height="223" alt="image" src="https://github.com/user-attachments/assets/bb6f25b4-f3f5-45b3-a479-9ffb2ed2af59" />


These provide domain names and ability to update DNS servers when IP changes.

---

## Learn More

- [All TCP/IP Protocol Suite Topics](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/)
- [Part 4 - Sockets and Ports](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/sockets-ports/)

