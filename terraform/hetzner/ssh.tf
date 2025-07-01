resource "hcloud_ssh_key" "home_desktop" {
  name = "home_desktop"
  public_key = var.SSH_KEY_PUBLIC_HOME_DESKTOP
}

resource "hcloud_ssh_key" "github_actions" {
  name = "github_actions"
  public_key = var.SSH_KEY_PUBLIC_GITHUB_ACTIONS
}