# Introduction to TCP/IP (Part 1) - Routers, Switches and Addressing

> **Source:** [Microchip Developer Help](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/routers-switches-addressing/)

## Routers in Local Networks

Routers connect one network to another. They:
- Create local networks and control access to them
- Route TCP/IP traffic on local networks
- Enable local network traffic to reach the Internet using **Network Address Translation (NAT)**
- Use firewalls to restrict public Internet access to the local network

## What is an IP Address?

IP addresses uniquely identify every host (network node) on a TCP/IP network.

- **Virtual addresses** assigned by routers
- Each of the four 8-bit fields is represented by a decimal number (0-255)
- Controlled by a **DHCP server** running in the router
- Assigned IP addresses can change at any time

### IPv4 vs IPv6

| Version | Address Size | Total Addresses |
|---------|-------------|-----------------|
| IPv4 | 32-bit (4 octets) | ~4.3 billion (about 1 per person) |
| IPv6 | 128-bit (8 x 16-bit fields) | 3.4 × 10³⁸ (5 × 10²⁸ per person) |

## Obtaining IP Addresses (DHCP Process)

When a device connects to a network, it automatically requests an IP address from the router.

### Step 1: PC Generates Request for IP Address
- PC's IP Address is `0.0.0.0` before assignment
- PC and router are physically connected via Ethernet cable
- Each end of the cable connects to a network **PHY** (transceiver for generating/receiving signals)
- Broadcast IP address: `255.255.255.255`
- **DHCP** (Dynamic Host Configuration Protocol) is used to request and grant IP addresses

### Step 2: IP Address Request Received in Router
- Router receives the broadcast packet
- Router sees sender needs an IP address and creates a new one
- Other hosts discard the packet once they realize they cannot provide an IP address

### Step 3: Router Offers an IP Address to the PC
- Router allocates a new IP address and broadcasts a message containing it
- Uses broadcast IP because PC doesn't know its IP yet and can't filter packets
- Frame includes the destination MAC address of the PC (used by switch)

### Step 4: PC Receives and Configures IP Address
- PC receives the broadcast packet
- Opens it and finds the IP address assignment message
- Configures its network interface with the assigned IP

## Switches in Local Networks

A switch enables connection of multiple devices to the same network.

- Each network interface has its own dedicated **PHY**
- **Uplink port**: Tx and Rx signals are reversed (no crossover cable needed)
- Most new switches have **Auto-MDIX** interfaces (automatically switch Tx/Rx if needed)

### Switches Inside Routers

Most home/small business routers have a built-in switch.

## Switches use MAC Addresses

Switches use **Media Access Controller (MAC)** addresses to forward and filter data.

- MAC controls **Layer 2** network functions
- Forwards and filters frames based on MAC addresses

### Two Addresses per Network Host

| Address Type | Layer | Description |
|--------------|-------|-------------|
| IP Address | Layer 3 | Virtual address |
| MAC Address | Layer 2 | Physical address |

### Key Differences: Switches vs Routers

| Device | Layer | Address Type | MAC Address? |
|--------|-------|--------------|--------------|
| Switch | Layer 2 | MAC only | No (transparent to network) |
| Router | Layer 3 | IP | Yes - two: LAN and WAN |

### MAC Functions

- Generates frames to send to the network
- Receives frames from the network
- Frames without matching MAC or broadcast address (`FF:FF:FF:FF:FF:FF`) are not forwarded to Layer 3
- Enables multiple devices to access the same physical network using:
  - **CSMA/CD** - Carrier Sense Multiple Access with Collision Detection (Ethernet)
  - **CSMA/CA** - Carrier Sense Multiple Access with Collision Avoidance (WLAN)

## What is a MAC Address?

- Fixed hardware-based address that **never changes**
- Programmed during manufacturing
- **Globally unique** (managed by IEEE registration authority)
- Six 8-bit fields expressed as hex numbers (e.g., `00:1A:2B:3C:4D:5E`)

## Switch Operation on a Local Network

A switch uses a **routing table** to associate port numbers with MAC addresses.

### Step 1: PC Sends a Frame to the Switch
- Host attempts to communicate with router to obtain an IP address
- Creates a packet with broadcast IP address
- Encapsulates packet into a frame with broadcast MAC address (`FF:FF:FF:FF:FF:FF`)

### Step 2: Switch Receives Frame
- Uses routing table to associate host's MAC address with the receiving interface
- Frame contains source MAC address

### Step 3: Switch Broadcasts Frame to All Nodes
- MAC sees this is a broadcast frame
- Switch forwards frame to all connected hosts

### Step 4: Router Sends Reply to PC
- Switch associates router's MAC address with the receiving interface

### Step 5: Switch Forwards Frame to PC
- Switch looks up destination MAC address in routing table
- Forwards frame only to the appropriate port
- No other port sees this frame

## Example: Simplified Local Network TCP/IP Communication

### Scenario
- Embedded network device with a web page for monitoring/control
- PC accesses the web page via browser

### Step 1: Open Web Browser and Enter Development Board's IP Address
- Both PC and development board have IP addresses
- Switch's routing table has been updated with MAC addresses for each port

### Step 2: PC Generates and Transmits a Frame
- Web browser creates a message requesting the web page at IP `192.168.1.102`
- Source and destination IP addresses added to create a **packet**
- Source and destination MAC addresses added to create a **frame**
- Frame sent to PHY for transmission

### Step 3: Frame is Forwarded Through the Switch
- Switch opens frame to find destination MAC address
- Finds MAC in routing table, sends to appropriate port
- Switch only cares about Layer 2 (MAC) addresses, not IP addresses

### Step 4: Frame Arrives at Development Board
1. Destination MAC address checked → matches device's MAC
2. Frame opened, destination IP address checked → matches device's IP
3. Packet opened to see message → destined for web server
4. Message sent to web server application

**Filtering at each layer:**
- MAC mismatch → frame discarded at Layer 2 (Data Link)
- IP mismatch → packet discarded at Layer 3 (Network)
- Application not running → discarded at Layer 4 (Transport)

### Step 5: Web Server Sends the Webpage to the PC
1. Web server generates message containing the web page
2. Source/destination IP addresses added → packet
3. Source/destination MAC addresses added → frame
4. Frame sent to PHY for transmission
5. Switch forwards to PC using routing table
6. PC opens frame, checks MAC, opens packet, checks IP
7. PC opens message and receives the HTML file

---

## Learn More

- [All TCP/IP Protocol Suite Topics](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/)
- [Part 2 - Five Layer Model and Applications](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/five-layer-model-and-apps/)

