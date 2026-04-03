---
subcategory: "Container Apps"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_container_app_environment_http_route_config"
description: |-
  Manages a Container App Environment HTTP Route Config.
---

# azurerm_container_app_environment_http_route_config

Manages a Container App Environment HTTP Route Config.

## Example Usage

```hcl
resource "azurerm_container_app_environment_http_route_config" "example" {
  name = "example"

  rules {

    targets {
      container_app = "TODO"      
    }    
  }
  container_app_environment_id = "TODO"
}
```

## Arguments Reference

The following arguments are supported:

* `container_app_environment_id` - (Required) The ID of the TODO. Changing this forces a new resource to be created.

* `name` - (Required) The name which should be used for this Container App Environment HTTP Route Config. Changing this forces a new resource to be created.

* `rules` - (Required) One or more `rules` blocks as defined below.

---

* `custom_domains` - (Optional) One or more `custom_domains` blocks as defined below.

---

A `action` block supports the following:

* `prefix_rewrite` - (Optional) TODO.

---

A `custom_domains` block supports the following:

* `name` - (Required) The name which should be used for this TODO.

* `binding_type` - (Optional) TODO.

* `certificate_id` - (Optional) The ID of the TODO.

---

A `match` block supports the following:

* `case_sensitive` - (Optional) TODO. Defaults to `true`.

* `path` - (Optional) TODO.

* `path_separated_prefix` - (Optional) TODO.

* `prefix` - (Optional) TODO.

---

A `routes` block supports the following:

* `match` - (Required) A `match` block as defined above.

* `action` - (Optional) A `action` block as defined above.

---

A `rules` block supports the following:

* `targets` - (Required) One or more `targets` blocks as defined below.

* `description` - (Optional) TODO.

* `routes` - (Optional) One or more `routes` blocks as defined above.

---

A `targets` block supports the following:

* `container_app` - (Required) TODO.

* `label` - (Optional) TODO.

* `revision` - (Optional) TODO.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported: 

* `id` - The ID of the Container App Environment HTTP Route Config.

* `fqdn` - TODO.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Container App Environment HTTP Route Config.
* `read` - (Defaults to 5 minutes) Used when retrieving the Container App Environment HTTP Route Config.
* `update` - (Defaults to 30 minutes) Used when updating the Container App Environment HTTP Route Config.
* `delete` - (Defaults to 30 minutes) Used when deleting the Container App Environment HTTP Route Config.

## Import

Container App Environment HTTP Route Configs can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_container_app_environment_http_route_config.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myEnvironment/httpRouteConfigs/myhttproute
```