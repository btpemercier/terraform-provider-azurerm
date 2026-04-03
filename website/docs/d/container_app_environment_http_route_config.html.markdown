---
subcategory: "Container Apps"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_container_app_environment_http_route_config"
description: |-
  Gets information about a Container App Environment HTTP Route Config.
---

# Data Source: azurerm_container_app_environment_http_route_config

Use this data source to access information about an existing Container App Environment HTTP Route Config.

## Example Usage

```hcl
data "azurerm_container_app_environment" "example" {
  name                = "existing-environment"
  resource_group_name = "existing-resources"
}

data "azurerm_container_app_environment_http_route_config" "example" {
  name                         = "existing-route"
  container_app_environment_id = data.azurerm_container_app_environment.example.id
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of the Container App Environment HTTP Route Config.

* `container_app_environment_id` - (Required) The ID of the Container App Environment.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the Container App Environment HTTP Route Config.

* `fqdn` - The FQDN of the HTTP Route Config.

* `custom_domains` - A list of `custom_domains` blocks as defined below.

* `rules` - A list of `rules` blocks as defined below.

---

A `custom_domains` block exports the following:

* `binding_type` - The binding type.

* `certificate_id` - The ID of the Certificate bound to this hostname.

* `name` - The hostname.

---

A `rules` block exports the following:

* `description` - The description of the rule.

* `routes` - A list of `routes` blocks as defined below.

* `targets` - A list of `targets` blocks as defined below.

---

A `routes` block exports the following:

* `action` - An `action` block as defined below.

* `match` - A `match` block as defined below.

---

An `action` block exports the following:

* `prefix_rewrite` - The prefix rewrite.

---

A `match` block exports the following:

* `case_sensitive` - Whether path matching is case sensitive.

* `path` - The exact path to match on.

* `path_separated_prefix` - The path separated prefix to match on.

* `prefix` - The prefix to match on.

---

A `targets` block exports the following:

* `container_app` - The Container App name.

* `label` - The label.

* `revision` - The revision.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the Container App Environment HTTP Route Config.

## API Providers
<!-- This section is generated, changes will be overwritten -->
This data source uses the following Azure API Providers:

* `Microsoft.App` - 2025-07-01
