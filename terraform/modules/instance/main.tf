resource "scaleway_instance_ip" "ip" {}

resource "scaleway_instance_server" "srv" {
  type  = "DEV1-S"
  image = "ubuntu_jammy"

  tags = [
    "app=test",
    "env=dev"
  ]

  ip_id = scaleway_instance_ip.ip.id

  user_data = {
    cloud-init = <<-EOT
    #cloud-config
    runcmd:
      - apt-get update
      - apt-get install nginx -y
      - systemctl enable --now nginx
      - echo "Hello i'm $(hostname)!" > /var/www/html/index.nginx-debian.html
    EOT
  }
}
