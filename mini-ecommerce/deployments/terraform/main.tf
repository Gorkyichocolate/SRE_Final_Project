terraform {

  required_version = ">= 1.5.0"

  required_providers {

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.30"
    }

    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

provider "kubernetes" {

  config_path = "~/.kube/config"
}

provider "docker" {}

resource "kubernetes_namespace" "mini_ecommerce" {

  metadata {
    name = var.namespace
  }
}

resource "docker_image" "mini_ecommerce" {

  name = "${var.dockerhub_username}/mini-ecommerce:latest"

  build {
    context    = "../.."
    dockerfile = "../../deployments/docker/Dockerfile"
  }
}

resource "kubernetes_deployment" "mini_ecommerce" {

  metadata {

    name      = "mini-ecommerce"
    namespace = kubernetes_namespace.mini_ecommerce.metadata[0].name

    labels = {
      app = "mini-ecommerce"
    }
  }

  spec {

    replicas = 1

    selector {

      match_labels = {
        app = "mini-ecommerce"
      }
    }

    template {

      metadata {

        labels = {
          app = "mini-ecommerce"
        }
      }

      spec {

        container {

          image = docker_image.mini_ecommerce.name
          name  = "mini-ecommerce"

          port {
            container_port = 8080
          }

          resources {

            limits = {
              cpu    = "500m"
              memory = "512Mi"
            }

            requests = {
              cpu    = "100m"
              memory = "128Mi"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "mini_ecommerce" {

  metadata {

    name      = "mini-ecommerce-service"
    namespace = kubernetes_namespace.mini_ecommerce.metadata[0].name
  }

  spec {

    selector = {
      app = "mini-ecommerce"
    }

    port {

      port        = 80
      target_port = 8080
    }

    type = "NodePort"
  }
}