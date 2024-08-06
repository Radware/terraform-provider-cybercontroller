<img src="https://www.radware.com/RadwareSite/MediaLibraries/Images/logo.svg" width="300px">

# Terraform Provider for Cyber Controller
The Terraform Provider for Cyber Controller helps automate provisioning and updating applications on Alteon clusters managed by Cyber Controller Alteon Cluster Manager module. 

## Requirements

- Terraform > 1.7.x
- Go v1.22.0 (To build the provider)
- CyberController >= 10.6.0.


# Building the Provider

Create a directory to clone the provider repository .

Link to clone the repository :git clone https://github.com/Radware/terraform-provider-cybercontroller.git

In command Prompt navigate to the cloned directory and use below command to build exe
		go build .

# Using the Provider

If you're building the provider, follow the instructions to install it as a plugin. After placing it into your plugins directory, run terraform init to initialize it.

## Copyright

Copyright 2024 Radware LTD

## License
GNU General Public License v3.0