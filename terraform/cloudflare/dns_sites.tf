resource "cloudflare_dns_record" "test_ajmuht" {
  zone_id = cloudflare_zone.ajmuht.id
  name = "test.ajmuht.eu"
  type = "A"
  content = var.servers.aquila
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "test_cmrlj" {
  zone_id = cloudflare_zone.cmrlj.id
  name = "test.cmrlj.eu"
  type = "A"
  content = var.servers.aquila
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "traefik_ajmuht" {
  zone_id = cloudflare_zone.ajmuht.id
  name = "traefik.ajmuht.eu"
  type = "A"
  content = var.servers.aquila
  proxied = true
  ttl = 1
}

resource "cloudflare_dns_record" "laravel_example" {
  zone_id = cloudflare_zone.ajmuht.id
  name = "laravel-example.ajmuht.eu"
  type = "A"
  content = var.servers.aquila
  proxied = true
  ttl = 1
}