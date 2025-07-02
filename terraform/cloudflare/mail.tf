module "cmrlj_mail" {
  source = "./modules/mail"
  zone = cloudflare_zone.cmrlj
  ip = var.servers.aquila
}

module "ajmuht_mail" {
  source = "./modules/mail"
  zone = cloudflare_zone.ajmuht
  ip = var.servers.aquila
}
