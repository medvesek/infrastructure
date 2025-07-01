// SSH KEYS

moved {
  from = hcloud_ssh_key.home_desktop
  to = module.hetzner.hcloud_ssh_key.home_desktop
}

moved {
  from = hcloud_ssh_key.github_actions
  to = module.hetzner.hcloud_ssh_key.github_actions
}

moved {
  from = hcloud_server.aquila
  to = module.hetzner.hcloud_server.aquila
}
