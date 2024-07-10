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


resource "cybercontroller_alteon_cli_command" "Test_Cluster" {
  clustername="terraform"
  elements {
    agalteonclicommand="/c/slb/filt 3/ena/"    
    }
}

