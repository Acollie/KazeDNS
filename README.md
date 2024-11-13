# 🌬️ KazeDNS

KazeDNS is a lightweight DNS server written in Go, designed for simplicity and efficiency. It supports both UDP and TCP protocols, allowing for customizable DNS root servers and blocklists.

✨ Features

    DNS over UDP and TCP: Handles DNS queries over both protocols seamlessly.
    Custom DNS Root Server: Configure your own DNS root server for tailored domain resolution.
    Blocklist Support: Implement custom blocklists to filter unwanted domains.
    Prometheus Metrics: Integrates with Prometheus for monitoring and metrics collection.

## 🚀 Getting Started

Follow these steps to set up KazeDNS:

    Clone the Repository:
```bash
git clone https://github.com/Acollie/KazeDNS.git
cd KazeDNS
```

Build the Application:
```
go build
```

Run the Server:
```bash
    ./KazeDNS
```
