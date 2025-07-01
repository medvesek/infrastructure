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



moved {
  from = cloudflare_dns_record.ajmuht_mail
  to = module.ajmuht_mail.cloudflare_dns_record.a
}

moved {
  from = cloudflare_dns_record.ajmuht_mail_mx
  to = module.ajmuht_mail.cloudflare_dns_record.mx
}

moved {
  from = cloudflare_dns_record.ajmuht_mail_spf
  to = module.ajmuht_mail.cloudflare_dns_record.spf
}


moved {
  from = cloudflare_dns_record.cmrlj_mail
  to = module.cmrlj_mail.cloudflare_dns_record.a
}

moved {
  from = cloudflare_dns_record.cmrlj_mail_mx
  to = module.cmrlj_mail.cloudflare_dns_record.mx
}

moved {
  from = cloudflare_dns_record.cmrlj_mail_spf
  to = module.cmrlj_mail.cloudflare_dns_record.spf
}
