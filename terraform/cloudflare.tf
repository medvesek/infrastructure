provider "cloudflare" {
  api_token = var.CLOUDFLARE_API_TOKEN
}

// DATA
data "cloudflare_account" "my_account" {
  filter = {}
}

// ZONE
resource "cloudflare_zone" "cmrlj" {
  name = "cmrlj.eu"
  type = "full"
  account ={
    id = data.cloudflare_account.my_account.account_id
  } 
}

// DNS RECORDS
resource "cloudflare_dns_record" "testme" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "testme"
  type = "A"
  content = hcloud_server.aquila.ipv4_address
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "mail" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "mail"
  type = "A"
  content = hcloud_server.aquila.ipv4_address
  proxied = true
  ttl = 1
}


resource "cloudflare_dns_record" "mail_mx" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "@"
  type = "mx"
  content = "mail.cmrlj.eu"
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "mail_spf" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "v=spf1 a mx ~all"
  type = "txt"
  proxied = true
  ttl = 1
}