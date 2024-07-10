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

/*data "cybercontroller_alteon_real_server_data" "TestServer-data" {  
  clustername="terraform"
  index="Real13"   
}

/*data "cybercontroller_alteon_server_group_data" "TestGroup-data" {
  clustername="terraform"  
  index="grp-table"
  //realserverindex="single-Table1"
}


data "cybercontroller_alteon_apply_status_cluster_data" "ApplyStatus-data" {
  clustername="terraform"
  //depends_on = [null_resource.wait_time]   
}*/

/*data "cybercontroller_alteon_virtual_server_data" "VirtualServer-data" {
  clustername="terraform"
  //clustername="dzACMCluster2"   
  index="Virt1-1" 
}*/

/*data "cybercontroller_alteon_https_health_check_data" "HttpsHealthCheck-data" {
  clustername="terraform"  
  index="hc1" 
}*/

/*data "cybercontroller_alteon_ssl_policy_data" "SslPolicy-data" {
  clustername="terraform"   
  nameidindex="Policy1" 
}*/

/*data "cybercontroller_alteon_http2_policy_data" "Http2Policy-data" {
  clustername="terraform"   
  nameidindex="samplepolicy"
}*/

/*data "cybercontroller_alteon_virtual_service_data" "VirtualService-data" {
  clustername="terraform"   
  servindex="Virt1-1"
  index=2
}
 
output "Real_Server_GET" {
  description = "Real Server GET"  
  value = [data.cybercontroller_alteon_real_server_data.TestServer-data]
}


/*resource "cybercontroller_alteon_apply" "Test_apply1" {
  clustername="terraform"  
  /*lifecycle {    
    replace_triggered_by = [
      cybercontroller_alteon_real_server.TestServer1
    ]
  }
  depends_on = [cybercontroller_alteon_real_server.TestServer1,cybercontroller_alteon_real_server.TestServer2]
  //depends_on = [cybercontroller_alteon_server_group.Server-Grp1]
}

/*resource "cybercontroller_alteon_revert_apply" "Test_revert_apply" {
  clustername="terraform"  
}

/*resource "cybercontroller_alteon_save" "Test_save" {
  clustername="terraform"  
  depends_on = [cybercontroller_alteon_apply.Test_apply1]
}*/

resource "cybercontroller_alteon_revert" "Test_revert" {
  clustername="terraform"
}

/*resource "cybercontroller_alteon_cli_command" "Test_Cluster" {
  clustername="terraform"
  elements {
    agalteonclicommand="/c/slb/filt 3/ena/"    
    }
}*/

/*resource "cybercontroller_alteon_real_server" "TestServer1" {
  clustername="terraform"  
  index="Multi-Real1"
  elements {
    	ipaddr="111.1.1.20"
      name="maxsizeid"
      timeout=22
      state=2
      weight=18    
    }
    elements_2 {
    	proxy=2
      ldapwr=2    
      fasthealthcheck=1
      subdmac=2
    }
    elements_3 {
    	criticalconnthrsh=85
      highconnthrsh=75    
    }
}
resource "cybercontroller_alteon_real_server" "TestServer2" {
  clustername="terraform"  
  index="single-Table1"
  elements {
    	ipaddr="11.11.11.201"
      name="SingleTableReal"
      timeout=12
      state=2
      //weight=data.cybercontroller_alteon_real_server_data.TestServer-data.weight   
    }
}

resource "cybercontroller_alteon_real_server" "TestServer3" {
  clustername="terraform"  
  index="Real13"
  elements {
    	ipaddr="13.13.13.203"
      name="Sample-Real"
      timeout=18
      state=2        
    }
}

/*resource "cybercontroller_alteon_real_server" "TestServer4" {
  clustername="terraform"  
  index="Real4"
  elements {
    	ipaddr="14.14.14.204"
      name="Real4"
      timeout=18
      state=2
      //weight=data.cybercontroller_alteon_real_server_data.TestServer-data.weight   
    }
}

/*resource "cybercontroller_alteon_server_group" "Server-Grp1" {
  clustername="terraform"
  index="grp-table"
  elements {
    name="Group-Update"
    addserver="Real4"    
    healthcheckurl="content.html"  
    healthchecklayer=3
    metric=1
    //backupserver="Grp-bkup-name"    
    }
  //depends_on = [cybercontroller_alteon_real_server.TestServer4]
}

resource "cybercontroller_alteon_virtual_server" "TestVirtualServer1" {
  clustername="terraform"
  index="Virt1-1"
  elements {
    	virtserveripaddress="10.10.10.10"
      virtserverstate=2
      virtserverdname="virtual-Server-Domain-Update"
      //virtserverweight=data.cybercontroller_alteon_virtual_server_data.VirtualServer-data.virtserverweight
      virtservernat="14.24.4.5"
      //virtserverbwmcontract=1022
      virtserveravail=2
      virtservervname="VirtualServerVName"
    }
  //depends_on = [cybercontroller_alteon_server_group.LabServers]
}

/*resource "cybercontroller_alteon_ssl_policy" "TestSslPolicy" {
  clustername="terraform"
  nameidindex="Policy1"
  elements {
    	name="testsslPolicy-Update"      
      adminstatus=1
      //bessl=2
      fesslv3version=1
      passinfocomply=2
    }  
}

resource "cybercontroller_alteon_http2_policy" "TestHttp2Policy" {
  clustername="terraform"
  nameidindex="samplepolicy"
  elements {
    	name="testHttp2Policy-Update"      
      adminstatus=1
      backendstatus=1
      backendhpacksize="small"
      backendstreams=4
      streams=4
      hpacksize="small"
    }  
}
resource "cybercontroller_alteon_https_health_check" "TestAdvhcHttp" {
  clustername="terraform"
  index="hc1"
  elements {
    	name="HC_TEST-Update"      
      dport=3
      ipver=1
      hostname="advHChostName"
      invert=1
      authlevel=2
    }  
}

resource "cybercontroller_alteon_virtual_service" "TestVirtualService" {
  clustername="terraform"
  servindex="Virt1-1"
  index=1
  elements {
    	virtport=80
      realport=80
      dbind=2
      udpbalance=3
      pbind=3
      //cookiemode=3
    }
    elements_2 {
      connmgtstatus=2
      connmgttimeout=12
      cachepol="Virtservice secondtable-Update"
      servurlchangstatus=1
      servurlchanghosttype=1
    }
    elements_3 {
      servurlchanghostname="thirdtablhostname"
      servurlchangpathtype=1
      servurlchangpathmatch="UrlPathname"
    }
    elements_4 {
      servurlchangnewpgname="New-url-Page-Name"
      servpathhidehosttype=1
      servpathhidehostname="hst-nam-4thTable"
      servtextrepstatus=1
    }
    elements_5 {
      servtextrepmatchtext="Text-match-fifth-table-Update"
      udpage=1
      alwayson=1
    }
    elements_6 {
      hname="hstnamesixthtable"
      direct=1
    }
    elements_7 {
      realgroup="grp-nam-seventh-table-Update"
      sessionmirror=1
    }  
  depends_on = [cybercontroller_alteon_virtual_server.TestVirtualServer1]
}

resource "cybercontroller_alteon_virtual_service" "TestVirtualService2" {
  clustername="terraform"
  servindex="Virt1-1"
  index=2
  elements {
    	virtport=53
      realport=53
      //dbind=3
      sideband="name-sideband-test-Update"
    }
  depends_on = [cybercontroller_alteon_virtual_server.TestVirtualServer1]
}*/
