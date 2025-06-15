terraform {

  cloud {
    organization = "medvesekg"
    workspaces {
      name = "infrastructure"
    }
  }

  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "~> 1.45"
    }
  }
}

variable "HCLOUD_TOKEN" {
  sensitive = true
}

provider "hcloud" {
  token = var.HCLOUD_TOKEN
}

resource "hcloud_server" "aquila" {
  name        = "aquila"
  image       = "ubuntu-24.04"
  server_type = "cx22"
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }
}