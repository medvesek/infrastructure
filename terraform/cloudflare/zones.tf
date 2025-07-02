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
