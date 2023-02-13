# netscan

Scan network ports

## Usage

Provide a CSV file as argument with the networks, hosts and ports that need to be scanned.
Example:
```
network, host, port
tcp4, www.google.com, 443
tcp4, localhost, 21
udp4, 8.8.8.8, 53
udp4, 8.8.8.8, 56
tcp4, 8.8.8.8, 80
tcp4, www.amayas.ca, 80
tcp4, www.amayas.ca, 443
```
