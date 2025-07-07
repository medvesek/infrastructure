resource "cloudflare_dns_record" "freeunlimitedparking_ajmuht_cname" {
  zone_id = module.ajmuht_eu_domain.id
  name = "freeunlimitedparking.ajmuht.eu"
  type = "CNAME"
  content = "medvesekg.github.io"
  proxied = true
  ttl = 1
}