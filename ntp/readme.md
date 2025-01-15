
Chrony is an implementation of the Network Time Protocol (NTP). You can use Chrony:

* to synchronize the system clock with NTP servers,
* to synchronize the system clock with a reference clock, for example a GPS receiver,
* to synchronize the system clock with a manual time input,
* as an NTPv4(RFC 5905) server or peer to provide a time service to other computers in the network.

Ref: https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/7/html/system_administrators_guide/ch-configuring_ntp_using_the_chrony_suite


Chrony is an alternative implementation of the Network Time Protocol (NTP), primarily designed for systems that are not continuously connected to the network or have intermittent connectivity. It is known for being robust and suitable for mobile and virtual systems.

Chrony consists of two main programs:

* chronyd - A daemon that runs in userspace in the background and adjusts the system clock.
* chronyc - A command-line interface for monitoring and controlling chronyd.

Advantages of Chrony
* Better Performance: Faster and more accurate synchronization.
* Handle Intermittent Connectivity: Suitable for systems that are not always connected to a network.
* Faster Recovery: More responsive to sudden changes in the clock frequency.
* Versatile: Can work well in virtual machines and mobile devices.
