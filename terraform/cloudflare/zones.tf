// AJMUHT
resource "cloudflare_zone" "ajmuht" {
  name = "ajmuht.eu"
  account ={
    id = data.cloudflare_account.my_account.account_id
  }
}
resource "cloudflare_zone_setting" "ajmuht_ssl_automatic_mode" {
  zone_id = cloudflare_zone.ajmuht.id
  setting_id = "ssl_automatic_mode"
  value = "custom"
}
resource "cloudflare_zone_setting" "ajmuht_ssl_mode" {
  zone_id = cloudflare_zone.ajmuht.id
  setting_id = "ssl"
  value = "full"
}

// CMRLJ
resource "cloudflare_zone" "cmrlj" {
  name = "cmrlj.eu"
  type = "full"
  account ={
    id = data.cloudflare_account.my_account.account_id
  } 
}
resource "cloudflare_zone_setting" "cmrlj_ssl_automatic_mode" {
  zone_id = cloudflare_zone.cmrlj.id
  setting_id = "ssl_automatic_mode"
  value = "custom"
}
resource "cloudflare_zone_setting" "cmrlj_ssl_mode" {
  zone_id = cloudflare_zone.cmrlj.id
  setting_id = "ssl"
  value = "full"
}