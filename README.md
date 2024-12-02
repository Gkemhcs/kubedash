# KUBEDASH

[![Release Pipeline for kubedash](https://github.com/Gkemhcs/kubedash/actions/workflows/release.yaml/badge.svg)](https://github.com/Gkemhcs/kubedash/actions/workflows/release.yaml)
[![CI Pipeline for kubedash](https://github.com/Gkemhcs/kubedash/actions/workflows/ci.yaml/badge.svg?event=push)](https://github.com/Gkemhcs/kubedash/actions/workflows/ci.yaml)

**kubedash** is a terminal-based Kubernetes dashboard that provides a simple and interactive way to manage Kubernetes resources directly from the command line. It allows you to list, search, describe, and delete resources with ease. Built with Go using the `tview`, `cobra`, and `logrus` libraries, kubedash offers a clean, user-friendly terminal interface for Kubernetes management.

## Features

- **List Resources**: Quickly list Kubernetes resources such as Pods, Deployments, ReplicaSets, and more.
- **Search Resources**: Easily search for resources across your clusters using simple queries.
- **Describe Resources**: View detailed information about a resource to better understand its configuration and status.
- **Delete Resources**: Delete Kubernetes resources from the command line with simple actions.
- **Customizable**: Configure the tool to suit your preferences with simple flags and settings.

## Prerequisites

- **Go**: kubedash is built using Go 1.18+.
- **Kubernetes Cluster**: You need access to a running Kubernetes cluster (local or remote) with valid `kubectl` credentials.
- **Libraries**: This tool relies on the following Go libraries:
  - `tview`: For the terminal UI.
  - `cobra`: For command-line argument parsing and creating commands.
  - `logrus`: For structured logging.

## Demo Video
[🎥videolink🎥](https://drive.google.com/file/d/1Mi8OSCNpTuJ_DfJIvd4O5X7vRrDssTcZ/view?usp=sharing) 


## Images 
**Home layout**
!["root"](./assets/root.png)

**Search Option**
!["search"](./assets/search.png)

**Describe Option**
!["describe"](./assets/describe.png)

**Delete Option**
!["delete"](./assets/delete.png)


## Installation

To install **kubedash**, follow these steps:

### Method 1: Build from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/Gkemhcs/kubedash.git
   cd kubedash
### Method 2:  using `go install`
```bash
go install github.com/Gkemhcs/kubedash@latest
```
### Method 3:  Download from Github Releases


1. Visit the [Releases Page](https://github.com/Gkemhcs/kubedash/releases) and download the appropriate binary for your operating system:
   - **Linux**: `kubedash-linux-amd64`
   - **macOS**: `kubedash-darwin-amd64`
   - **Windows**: `kubedash-windows-amd64.exe`

2. Make the binary executable (for Linux/macOS):

   ```bash
   chmod +x kubedash-linux-amd64
   sudo mv kubedash-linux-amd64 /usr/local/bin/kubedash
   ```
3. On Windows, add the binary location to your system PATH or run it directly

Here’s the full README.md content in code form:

## Running Application 
```bash
# prints the info 
kubedash info 
# prints the version of cli tool
kubedash version 
# start application
kubedash
```
## Keyboard Shortcuts

    Ctrl+s: Start a search query.
    Ctrl+d: Delete the selected resource.
    d: Describe the selected resource.
    Ctrl+c: Quit the kubedash application.
