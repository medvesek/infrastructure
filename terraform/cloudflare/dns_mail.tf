// CMRLJ
resource "cloudflare_dns_record" "cmrlj_mail" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "mail.cmrlj.eu"
  type = "A"
  content = var.servers.aquila
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "cmrlj_mail_mx" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "cmrlj.eu"
  type = "MX"
  content = "mail.cmrlj.eu"
  priority = 1
  ttl = 1
}

resource "cloudflare_dns_record" "cmrlj_mail_spf" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "cmrlj.eu"
  content = "\"v=spf1 a mx ~all\""
  type = "TXT"
  ttl = 3600
}


// AJMUHT

resource "cloudflare_dns_record" "ajmuht_mail" {
  zone_id = cloudflare_zone.ajmuht.id
  name = "mail.ajmuht.eu"
  type = "A"
  content = var.servers.aquila
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
