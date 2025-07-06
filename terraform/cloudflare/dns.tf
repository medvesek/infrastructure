resource "cloudflare_dns_record" "freeunlimitedparking_ajmuht_cname" {
  zone_id = module.ajmuht_eu_domain.id
  name = "freeunlimitedparking.ajmiht.eu"
  type = "CNAME"
  content = "medvesekg.github.io/freeunlimitedparking"
  proxied = true
  ttl = 1
}