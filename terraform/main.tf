module "srv" {
  source = "./modules/instance"
  count  = 3
}

module "lb" {
  source = "./modules/loadbalancer"
}
