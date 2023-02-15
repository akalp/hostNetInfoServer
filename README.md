# Host Network Information Server

This is a simple Go web service that returns information about the host's network interfaces, including their addresses and types.

## Installation

1. Clone the repository:
```bash
git clone https://github.com/akalp/hostNetInfoServer.git
```
2. Build the application:

* for linux x64:
```bash
GOOS=linux GOARCH=amd64 go build
```
* for windows x64:
```bash
GOOS=windows GOARCH=amd64 go build
```


3. Set the environment variables (optional):
```bash
export PORT=8080
export IF_PREFIX=en
```

4. Run the application:
```bash
./go-network-info
```

## Usage

By default, the web service listens on port 8080 and returns information for all network interfaces. You can customize the port and filter the interfaces by prefix using environment variables or command-line flags.

### Environment variables
* `HNIS_PORT` (default: 8080): The port on which the web service listens.
* `HNIS_IF_PREFIX` (default: ""): The prefix to match for network interface names.
### Command-line flags
* `port` (default: 8080): The port on which the web service listens.
* `if-prefix` (default: ""): The prefix to match for network interface names.

* ## API

The web service exposes a single endpoint (/) that returns a JSON-encoded list of objects representing the network interfaces and their addresses:

```json
[
  {
    "name": "en1",
    "mtu": 1500,
    "flags": "up|broadcast|multicast",
    "addresses": {
      "ipv4": null,
      "ipv6": null
    }
  },
  {
    "name": "en0",
    "mtu": 1500,
    "flags": "up|broadcast|multicast",
    "addresses": {
      "ipv4": [
        "127.0.0.1"
      ],
      "ipv6": [
        "::1"
      ]
    }
  }
]
```
Each object in the list represents a network interface, and includes the following fields:

* name: The name of the interface.
* mtu: The maximum transmission unit (MTU) of the interface.
* flags: A comma-separated list of flags that describe the state of the interface (e.g., "up", "broadcast", "multicast", etc.).
* addresses: An object of objects representing the IP addresses associated with the interface. Each ip address grouped by the type of it (IPv4 or IPv6).

### License

This code is released under the MIT License. See LICENSE.md for details.