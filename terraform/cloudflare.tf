moved {
  from = cloudflare_dns_record.ajmuht_mail
  to = module.cloudflare.cloudflare_dns_record.ajmuht_mail
}
moved {
  from = cloudflare_zone.ajmuht
  to = module.cloudflare.cloudflare_zone.ajmuht
}

moved {
  from = cloudflare_zone_setting.ajmuht_ssl_automatic_mode
  to = module.cloudflare.cloudflare_zone_setting.ajmuht_ssl_automatic_mode
}

moved {
  from = cloudflare_zone_setting.ajmuht_ssl_mode
  to = module.cloudflare.cloudflare_zone_setting.ajmuht_ssl_mode
}

moved {
  from = cloudflare_zone.cmrlj
  to = module.cloudflare.cloudflare_zone.cmrlj
}

moved {
  from = cloudflare_zone_setting.cmrlj_ssl_automatic_mode
  to = module.cloudflare.cloudflare_zone_setting.cmrlj_ssl_automatic_mode
}

moved {
  from = cloudflare_zone_setting.cmrlj_ssl_mode
  to = module.cloudflare.cloudflare_zone_setting.cmrlj_ssl_mode
}

moved {
  from = cloudflare_dns_record.test_ajmuht
  to = module.cloudflare.cloudflare_dns_record.test_ajmuht
}

moved {
  from = cloudflare_dns_record.test_cmrlj
  to = module.cloudflare.cloudflare_dns_record.test_cmrlj
}

moved {
  from = cloudflare_dns_record.traefix_ajmuht
  to = module.cloudflare.cloudflare_dns_record.traefik_ajmuht
}

moved {
  from = cloudflare_dns_record.laravel_example
  to = module.cloudflare.cloudflare_dns_record.laravel_example
}

moved {
  from = cloudflare_dns_record.mail
  to = module.cloudflare.cloudflare_dns_record.cmrlj_mail
}

moved {
  from = cloudflare_dns_record.mail_mx
  to = module.cloudflare.cloudflare_dns_record.cmrlj_mail_mx
}

moved {
  from = cloudflare_dns_record.mail_spf
  to = module.cloudflare.cloudflare_dns_record.cmrlj_mail_spf
}

moved {
  from = cloudflare_dns_record.ajmuht_mail
  to = module.cloudflare.cloudflare_dns_record.ajmuht_mail
}

moved {
  from = cloudflare_dns_record.ajmuht_mail_mx
  to = module.cloudflare.cloudflare_dns_record.ajmuht_mail_mx
}

moved {
  from =  cloudflare_dns_record.ajmuht_mail_spf
  to = module.cloudflare.cloudflare_dns_record.ajmuht_mail_spf
}