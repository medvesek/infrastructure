provider "hcloud" {
  token = var.HCLOUD_TOKEN
}

// SSH KEYS
resource "hcloud_ssh_key" "home_desktop" {
  name = "home_desktop"
  public_key = var.SSH_KEY_PUBLIC_HOME_DESKTOP
}

resource "hcloud_ssh_key" "github_actions" {
  name = "github_actions"
  public_key = var.SSH_KEY_PUBLIC_GITHUB_ACTIONS
}

// SERVERS
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