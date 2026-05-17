output "namespace" {

  value = kubernetes_namespace.mini_ecommerce.metadata[0].name
}

output "docker_image" {

  value = docker_image.mini_ecommerce.name
}

output "service_name" {

  value = kubernetes_service.mini_ecommerce.metadata[0].name
}