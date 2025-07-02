module "ajmuht_eu_domain" {
  source = "./modules/domain"
  account_id = data.cloudflare_account.my_account.account_id
  domain = "ajmuht.eu"
}
module "cmrlj_eu_domain" {
  source = "./modules/domain"
  account_id = data.cloudflare_account.my_account.account_id
  domain = "cmrlj.eu"
}

moved {
  from = cloudflare_zone.ajmuht
  to = module.ajmuht_eu_domain.cloudflare_zone.zone
}

moved {
  from = cloudflare_zone.cmrlj
  to = module.cmrlj_eu_domain.cloudflare_zone.zone
}

moved {
  from = cloudflare_zone_setting.ajmuht_ssl_automatic_mode
  to = module.ajmuht_eu_domain.cloudflare_zone_setting.ssl_automatic_mode
}

moved {
  from = cloudflare_zone_setting.ajmuht_ssl_mode
  to = module.ajmuht_eu_domain.cloudflare_zone_setting.ssl_mode
}

moved {
  from = cloudflare_zone_setting.cmrlj_ssl_automatic_mode
  to = module.cmrlj_eu_domain.cloudflare_zone_setting.ssl_automatic_mode
}

moved {
  from = cloudflare_zone_setting.cmrlj_ssl_mode
  to = module.cmrlj_eu_domain.cloudflare_zone_setting.ssl_mode
}

