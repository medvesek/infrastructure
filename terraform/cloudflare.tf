provider "cloudflare" {
  api_token = var.CLOUDFLARE_API_TOKEN
}

// DATA
data "cloudflare_account" "my_account" {
  filter = {}
}

// ZONE
resource "cloudflare_zone" "ajmuht" {
  name = "ajmuht.eu"
  account ={
    id = data.cloudflare_account.my_account.account_id
  }
}
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
  name = "testme.cmrlj.eu"
  type = "A"
  content = hcloud_server.aquila.ipv4_address
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "mail" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "mail.cmrlj.eu"
  type = "A"
  content = hcloud_server.aquila.ipv4_address
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "mail_mx" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "cmrlj.eu"
  type = "MX"
  content = "mail.cmrlj.eu"
  priority = 1
  ttl = 1
}

resource "cloudflare_dns_record" "mail_spf" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "cmrlj.eu"
  content = "\"v=spf1 a mx ~all\""
  type = "TXT"
  ttl = 3600
}

resource "cloudflare_dns_record" "ajmuht_mail" {
  zone_id = cloudflare_zone.ajmuht.id
  name = "mail.ajmuht.eu"
  type = "A"
  content = hcloud_server.aquila.ipv4_address
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "ajmuht_mail_mx" {
  zone_id = cloudflare_zone.ajmuht.id
  name = "ajmuht.eu"
  type = "MX"
  content = "mail.ajmuht.eu"
  priority = 1
  ttl = 1
}

resource "cloudflare_dns_record" "ajmuht_mail_spf" {
  zone_id = cloudflare_zone.ajmuht.id
  name = "ajmuht.eu"
  content = "\"v=spf1 a mx ~all\""
  type = "TXT"
  ttl = 3600
}
