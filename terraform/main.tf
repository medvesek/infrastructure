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


module "cloudflare" {
  source = "./cloudflare"
  CLOUDFLARE_API_TOKEN = var.CLOUDFLARE_API_TOKEN
  servers = module.hetzner.servers
}

module "hetzner" {
  source = "./hetzner"
  HCLOUD_TOKEN = var.HCLOUD_TOKEN
  SSH_KEY_PUBLIC_HOME_DESKTOP = var.SSH_KEY_PUBLIC_HOME_DESKTOP
  SSH_KEY_PUBLIC_GITHUB_ACTIONS = var.SSH_KEY_PUBLIC_GITHUB_ACTIONS
}
