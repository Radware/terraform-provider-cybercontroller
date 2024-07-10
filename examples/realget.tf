terraform {
  required_providers {
    cybercontroller = {
      version = ">=0.0.4"
      source = "Radware/cybercontroller"
    }
  }
}

provider "cybercontroller" {
  username="radware"
  password="radware"
  ip="10.171.101.97" 
   
}

data "cybercontroller_alteon_real_server_data" "TestServer-data" {  
  clustername="terraform"
  index="Real13"   
}
output "Real_Server_GET" {
  description = "Real Server GET"  
  value = [data.cybercontroller_alteon_real_server_data.TestServer-data]
}
