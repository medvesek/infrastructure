module "site_test_ajmuht" {
  source = "./modules/site"
  zone_id = module.ajmuht_eu_domain.id
  ip = var.servers.aquila
  hostname = "test.ajmuht.eu"
}
module "site_traefik_ajmuht" {
  source = "./modules/site"
  zone_id = module.ajmuht_eu_domain.id
  ip = var.servers.aquila
  hostname = "traefik.ajmuht.eu"
}

module "site_laravel_example_ajmuht" {
  source = "./modules/site"
  zone_id = module.ajmuht_eu_domain.id
  ip = var.servers.aquila
  hostname = "laravel-example.ajmuht.eu"
}

module "site_fastify_example_ajmuht" {
  source = "./modules/site"
  zone_id = module.ajmuht_eu_domain.id
  ip = var.servers.aquila
  hostname = "fastify-example.ajmuht.eu"
}

module "site_rubicon_combat_ajmuht" {
  source = "./modules/site"
  zone_id = module.ajmuht_eu_domain.id
  ip = var.servers.aquila
  hostname = "rubicon-websockets.ajmuht.eu"
}
