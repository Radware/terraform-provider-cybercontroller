<img src="https://www.radware.com/RadwareSite/MediaLibraries/Images/logo.svg" width="300px">

# Terraform Provider for Cyber Controller
The Terraform Provider for Cyber Controller project provides a Terraform provider code for managing Alteon Cluster Manager Configurations in Terraform. 

## Requirements

- Terraform > 1.7.x
- Go v1.22.0 (To build the provider)

# Building the Provider

Clone repository to: $GOPATH/src/github.com/Radware/terraform-provider-cybercontroller

```
$ mkdir -p $GOPATH/src/github.com/Radware; cd $GOPATH/src/github.com/Radware
$ git clone https://github.com/Radware/terraform-provider-cybercontroller.git

```
Enter the provider directory and build the provider

```
$ cd $GOPATH/src/github.com/Radware/terraform-provider-cybercontroller
$ make build

```
# Using the Provider

If you're building the provider, follow the instructions to install it as a plugin. After placing it into your plugins directory, run terraform init to initialize it.

## Copyright

Copyright 2024 Radware LTD

## License
GNU General Public License v3.0