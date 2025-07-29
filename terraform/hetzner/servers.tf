locals {
  user_data = templatefile(
    "${path.module}/init.sh", 
    { 
      SSH_KEY_PUBLIC_HOME_DESKTOP: var.SSH_KEY_PUBLIC_HOME_DESKTOP,
      SSH_KEY_PUBLIC_GITHUB_ACTIONS: var.SSH_KEY_PUBLIC_GITHUB_ACTIONS,
    }
  )
}

resource "hcloud_server" "aquila" {
  name        = "aquila"
  image       = "ubuntu-24.04"
  server_type = "cx22"
  location    = "nbg1"
  user_data   = local.user_data
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }
}
