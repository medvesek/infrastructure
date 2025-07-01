resource "hcloud_server" "aquila" {
  name        = "aquila"
  image       = "ubuntu-24.04"
  server_type = "cx22"
  location    = "nbg1"
  ssh_keys    = ["home_desktop", "github_actions"]
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }
}