terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5"
    }
  }
}

resource "cloudflare_dns_record" "a" {
  zone_id = var.zone.id
  name = "mail.${var.zone.name}"
  type = "A"
  content = var.ip
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "mx" {
  zone_id = var.zone.id
  name = var.zone.name
  type = "MX"
  content = "mail.${var.zone.name}"
  priority = 1
  ttl = 1
}

resource "cloudflare_dns_record" "spf" {
  zone_id = var.zone.id
  name = var.zone.name
  content = "\"v=spf1 a mx ~all\""
  type = "TXT"
  ttl = 3600
}