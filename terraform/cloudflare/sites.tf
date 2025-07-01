module "site_test_ajmuht" {
  source = "./modules/site"
  zone_id = cloudflare_zone.ajmuht.id
  ip = var.servers.aquila
  hostname = "test.ajmuht.eu"
}

module "site_test_cmrlj" {
  source = "./modules/site"
  zone_id = cloudflare_zone.cmrlj.id
  ip = var.servers.aquila
  hostname = "test.cmrlj.eu"
}

module "site_traefik_ajmuht" {
  source = "./modules/site"
  zone_id = cloudflare_zone.ajmuht.id
  ip = var.servers.aquila
  hostname = "traefik.ajmuht.eu"
}

module "site_laravel_example_ajmuht" {
  source = "./modules/site"
  zone_id = cloudflare_zone.ajmuht.id
  ip = var.servers.aquila
  hostname = "laravel-example.ajmuht.eu"
}


moved {
  from = cloudflare_dns_record.test_ajmuht
  to = module.site_test_ajmuht.cloudflare_dns_record.a
}

moved {
  from = cloudflare_dns_record.test_cmrlj
  to = module.site_test_cmrlj.cloudflare_dns_record.a
}

moved {
  from = cloudflare_dns_record.traefik_ajmuht
  to = module.site_traefik_ajmuht.cloudflare_dns_record.a
}

moved {
  from = cloudflare_dns_record.laravel_example
  to = module.site_laravel_example_ajmuht.cloudflare_dns_record.a
}
