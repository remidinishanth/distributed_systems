# Introduction to TCP/IP (Part 1) - Routers, Switches and Addressing

> **Source:** [Microchip Developer Help](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/routers-switches-addressing/)

## Routers in Local Networks

<img width="506" height="100" alt="image" src="https://github.com/user-attachments/assets/427b8f3e-2fba-489a-8db8-713e2f57e417" />


Routers connect one network to another. They:
- Create local networks and control access to them
- Route TCP/IP traffic on local networks
- Enable local network traffic to reach the Internet using **Network Address Translation (NAT)**
- Use firewalls to restrict public Internet access to the local network

## What is an IP Address?

<img width="540" height="202" alt="image" src="https://github.com/user-attachments/assets/c37ef387-9610-4afb-83be-7b2cfa8825ee" />


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

<img width="559" height="158" alt="image" src="https://github.com/user-attachments/assets/79d05c45-f090-43e4-bc01-e7d424f6b6d3" />


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

<img width="550" height="133" alt="image" src="https://github.com/user-attachments/assets/038910a1-1c55-46dd-aa23-5e5e6b0c4acb" />


### Step 3: Router Offers an IP Address to the PC
- Router allocates a new IP address and broadcasts a message containing it
- Uses broadcast IP because PC doesn't know its IP yet and can't filter packets
- Frame includes the destination MAC address of the PC (used by switch)

### Step 4: PC Receives and Configures IP Address
- PC receives the broadcast packet
- Opens it and finds the IP address assignment message
- Configures its network interface with the assigned IP

<img width="568" height="137" alt="image" src="https://github.com/user-attachments/assets/8957f061-c8f5-4a2f-83b7-7c3c71507da7" />


## Switches in Local Networks

<img width="525" height="196" alt="image" src="https://github.com/user-attachments/assets/b3284b69-42df-4d39-bca6-ec5af8b7ff02" />


A switch enables connection of multiple devices to the same network.

- Each network interface has its own dedicated **PHY**
- **Uplink port**: Tx and Rx signals are reversed (no crossover cable needed)
- Most new switches have **Auto-MDIX** interfaces (automatically switch Tx/Rx if needed)

### Switches Inside Routers

<img width="535" height="201" alt="image" src="https://github.com/user-attachments/assets/cdca5403-a395-4400-a94b-df976d9f56c7" />


Most home/small business routers have a built-in switch.

## Switches vs Routers

<img width="758" height="877" alt="image" src="https://github.com/user-attachments/assets/66de59dd-5e96-4533-9d99-c4688754e460" />

<img width="1024" height="768" alt="image" src="https://github.com/user-attachments/assets/5a6eb627-2692-4d36-bbe6-0e95e32420e9" />

<img width="800" height="800" alt="image" src="https://github.com/user-attachments/assets/37cb15eb-4a07-4e42-9157-14ba1301c60e" />

<img width="2258" height="1254" alt="image" src="https://github.com/user-attachments/assets/1722d0f9-d076-4c1f-adca-93ff116666ea" />


Ref: https://labs.iximiuz.com/courses/computer-networking-fundamentals/bridge-vs-switch

## Switches use MAC Addresses

<img width="576" height="219" alt="image" src="https://github.com/user-attachments/assets/5e4c1450-c3a7-4364-a56f-ba172eeaa4ae" />


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

<img width="630" height="373" alt="image" src="https://github.com/user-attachments/assets/88a51488-678a-4a1b-b421-faab6f6427aa" />


### Step 1: PC Sends a Frame to the Switch
- Host attempts to communicate with router to obtain an IP address
- Creates a packet with broadcast IP address
- Encapsulates packet into a frame with broadcast MAC address (`FF:FF:FF:FF:FF:FF`)

<img width="571" height="338" alt="image" src="https://github.com/user-attachments/assets/f5bed39e-08a6-4991-b14f-9bd5e41e2ad3" />


### Step 2: Switch Receives Frame
- Uses routing table to associate host's MAC address with the receiving interface
- Frame contains source MAC address

<img width="571" height="338" alt="image" src="https://github.com/user-attachments/assets/7f3f28e6-40f2-458d-8701-82ab1980dfa7" />


### Step 3: Switch Broadcasts Frame to All Nodes
- MAC sees this is a broadcast frame
- Switch forwards frame to all connected hosts

<img width="571" height="338" alt="image" src="https://github.com/user-attachments/assets/ffbcd5fc-8eb7-46ae-b56e-8ce146e12b9d" />


### Step 4: Router Sends Reply to PC
- Switch associates router's MAC address with the receiving interface

<img width="571" height="338" alt="image" src="https://github.com/user-attachments/assets/f4f13cf8-808b-4922-8204-7974bd84f01f" />


### Step 5: Switch Forwards Frame to PC
- Switch looks up destination MAC address in routing table
- Forwards frame only to the appropriate port
- No other port sees this frame

## Example: Simplified Local Network TCP/IP Communication

<img width="498" height="294" alt="image" src="https://github.com/user-attachments/assets/480d5b00-a77e-455c-bf18-51691506a0a7" />


### Scenario
- Embedded network device with a web page for monitoring/control
- PC accesses the web page via browser

<img width="533" height="363" alt="image" src="https://github.com/user-attachments/assets/508a48aa-013a-4b37-85ce-0385b876bb17" />


### Step 1: Open Web Browser and Enter Development Board's IP Address
- Both PC and development board have IP addresses
- Switch's routing table has been updated with MAC addresses for each port

<img width="551" height="328" alt="image" src="https://github.com/user-attachments/assets/3546e74d-29de-4e07-91de-7ee89cf0b52a" />


### Step 2: PC Generates and Transmits a Frame
- Web browser creates a message requesting the web page at IP `192.168.1.102`
- Source and destination IP addresses added to create a **packet**
- Source and destination MAC addresses added to create a **frame**
- Frame sent to PHY for transmission

<img width="558" height="336" alt="image" src="https://github.com/user-attachments/assets/8d9e3dc0-7448-4e19-aeaf-300550cd5b50" />


### Step 3: Frame is Forwarded Through the Switch
- Switch opens frame to find destination MAC address
- Finds MAC in routing table, sends to appropriate port
- Switch only cares about Layer 2 (MAC) addresses, not IP addresses

<img width="540" height="320" alt="image" src="https://github.com/user-attachments/assets/897428b7-ecaf-480c-b9f0-ef810824c56d" />


### Step 4: Frame Arrives at Development Board
1. Destination MAC address checked → matches device's MAC
2. Frame opened, destination IP address checked → matches device's IP
3. Packet opened to see message → destined for web server
4. Message sent to web server application

**Filtering at each layer:**
- MAC mismatch → frame discarded at Layer 2 (Data Link)
- IP mismatch → packet discarded at Layer 3 (Network)
- Application not running → discarded at Layer 4 (Transport)


<img width="554" height="328" alt="image" src="https://github.com/user-attachments/assets/567ba4b6-67c8-41f2-ba73-f50d2d5819ad" />


### Step 5: Web Server Sends the Webpage to the PC
1. Web server generates message containing the web page
2. Source/destination IP addresses added → packet
3. Source/destination MAC addresses added → frame
4. Frame sent to PHY for transmission
5. Switch forwards to PC using routing table
6. PC opens frame, checks MAC, opens packet, checks IP
7. PC opens message and receives the HTML file

---

Question: What happens next?
Answer:

    The frame is received at the switch.
    The switch finds the frame’s destination MAC address and uses its routing table to determine what port to forward the frame to.
    The frame is sent to the PC
    The PC opens the frame and checks the destination MAC address to determine if it needs to pay attention to it.
    The PC opens the packet and checks the destination IP address.
    The PC opens the message and finds the web page (which is just an HTML file) it requested.

## Learn More

- [All TCP/IP Protocol Suite Topics](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/)
- [Part 2 - Five Layer Model and Applications](https://developerhelp.microchip.com/xwiki/bin/view/applications/tcp-ip/five-layer-model-and-apps/)

