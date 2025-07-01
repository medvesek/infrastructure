terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5"
    }
  }
}

resource "cloudflare_dns_record" "a" {
  zone_id = var.zone_id
  name = var.hostname
  type = "A"
  content = var.ip
  proxied = true
  ttl = 1
}
