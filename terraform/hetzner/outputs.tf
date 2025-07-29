output "servers" {
  description = "IP addresses of servers"
  value       = {
    aquila = hcloud_server.aquila.ipv4_address   
  }          
}