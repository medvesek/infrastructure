module "cmrlj_mail" {
  source = "./modules/mail"
  zone = module.cmrlj_eu_domain
  ip = var.servers.aquila
}

module "ajmuht_mail" {
  source = "./modules/mail"
  zone = module.ajmuht_eu_domain
  ip = var.servers.aquila
}
