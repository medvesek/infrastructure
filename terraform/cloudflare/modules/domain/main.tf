terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5"
    }
  }
}

resource "cloudflare_zone" "zone" {
  name = var.domain
  account ={
    id = var.account_id
  }
}
resource "cloudflare_zone_setting" "ssl_automatic_mode" {
  zone_id = cloudflare_zone.zone.id
  setting_id = "ssl_automatic_mode"
  value = "custom"
}
resource "cloudflare_zone_setting" "ssl_mode" {
  zone_id = cloudflare_zone.zone.id
  setting_id = "ssl"
  value = "full"
}