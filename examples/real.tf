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
resource "cybercontroller_alteon_real_server" "TestServer1" {
  clustername="terraform"  
  index="Real13"
  elements {
    	ipaddr="13.13.13.203"
      name="Sample-Real"
      timeout=18
      state=2        
    }
}

