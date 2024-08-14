# digitalocean-go

A comprehensive Go CLI application for managing various DigitalOcean resources including droplets, VPCs, Kubernetes clusters, databases, and domains.

## Features

- **Droplets**: List, create, and delete droplets.
- **VPCs**: List, create, and delete Virtual Private Clouds.
- **Kubernetes**: List, create, and delete Kubernetes clusters.
- **Databases**: List, create, and delete managed databases.
- **Domains**: List, create, and delete domains and DNS records.
- **Billing**: Retrieve billing information.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/felipepimentel/digitalocean-go.git
   ```

2. Change to the project directory:

   ```bash
   cd digitalocean-go
   ```

3. Build the application:

   ```bash
   go build -o digitalocean-cli ./cmd/digitalocean-cli
   ```

4. Copy the `.env.example` file to `.env` and add your DigitalOcean API token:

   ```bash
   cp .env.example .env
   ```

## Usage

### Droplets

- List all droplets:
  
  ```bash
  ./digitalocean-cli droplet list
  ```

- Create a new droplet:
  
  ```bash
  ./digitalocean-cli droplet create --name my-droplet --region nyc3 --size s-1vcpu-1gb --image ubuntu-20-04-x64
  ```

- Delete a droplet:
  
  ```bash
  ./digitalocean-cli droplet delete [droplet_id]
  ```

### VPCs

- List all VPCs:
  
  ```bash
  ./digitalocean-cli vpc list
  ```

- Create a new VPC:
  
  ```bash
  ./digitalocean-cli vpc create --name my-vpc --region nyc3 --ip-range 10.10.10.0/24
  ```

- Delete a VPC:
  
  ```bash
  ./digitalocean-cli vpc delete [vpc_id]
  ```

### Kubernetes

- List all Kubernetes clusters:
  
  ```bash
  ./digitalocean-cli kubernetes list
  ```

- Create a new Kubernetes cluster:
  
  ```bash
  ./digitalocean-cli kubernetes create --name my-cluster --region nyc3 --version 1.21.5-do.0 --nodes 3
  ```

- Delete a Kubernetes cluster:
  
  ```bash
  ./digitalocean-cli kubernetes delete [cluster_id]
  ```

### Databases

- List all managed databases:
  
  ```bash
  ./digitalocean-cli database list
  ```

- Create a new managed database:
  
  ```bash
  ./digitalocean-cli database create --name my-database --engine pg --version 13 --size db-s-1vcpu-1gb --region nyc3
  ```

- Delete a managed database:
  
  ```bash
  ./digitalocean-cli database delete [database_id]
  ```

### Domains

- List all domains:
  
  ```bash
  ./digitalocean-cli domain list
  ```

- Create a new domain:
  
  ```bash
  ./digitalocean-cli domain create --name example.com
  ```

- Delete a domain:
  
  ```bash
  ./digitalocean-cli domain delete [domain_name]
  ```

### Billing

- Get billing information:
This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.