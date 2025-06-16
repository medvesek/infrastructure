output "instance_ip" {
  description = "IP address of the server"
  value       = hcloud_server.aquila.ipv4_address               
}