resource "scaleway_lb_ip" "ip" {}

resource "scaleway_lb" "lb" {
  ip_id = scaleway_lb_ip.ip.id
  type  = "LB-S"

  tags = [
    "app=test",
    "env=dev"
  ]
}

resource "scaleway_lb_backend" "bkd" {
  lb_id            = scaleway_lb.lb.id
  forward_protocol = "tcp"
  forward_port     = 80
  proxy_protocol   = "none"
}

resource "scaleway_lb_frontend" "frt" {
  lb_id        = scaleway_lb.lb.id
  backend_id   = scaleway_lb_backend.bkd.id
  inbound_port = 80
}
