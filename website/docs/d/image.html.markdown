---
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_image"
sidebar_current: "docs-azurerm-datasource-image"
description: |-
  Use this data source to access information about an Image.

---

# Data Source: azurerm_image

Use this data source to access information about an Image.

## Argument Reference



* `name` - (Optional) The name of the Image.
* `name_regex` - (Optional) Regex pattern of the image to match.
* `resource_group_name` - (Required) The Name of the Resource Group for this resource.
* `sort_descending` - (Optional) By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.


## Attributes Reference

* `data_disk` - a collection of `data_disk` blocks as defined below.
* `location` - The Azure Location for this resource.
* `os_disk` - a collection of `os_disk` blocks as defined below.
* `tags` - A mapping of tags assigned to the resource.



`data_disk` supports the following:

* `blob_uri` - The URI in Azure storage of the blob used to create the image.
* `caching` - The caching mode for the Data Disk, such as `ReadWrite`, `ReadOnly`, or `None`.
* `lun` - The logical unit number of the data disk.
* `managed_disk_id` - The ID of the Managed Disk used as the Data Disk Image.
* `size_gb` - The size of this Data Disk in GB.

`os_disk` supports the following:

* `blob_uri` - The URI in Azure storage of the blob used to create the image.
* `caching` - The caching mode for the OS Disk, such as `ReadWrite`, `ReadOnly`, or `None`.
* `managed_disk_id` - The ID of the Managed Disk used as the OS Disk Image.
* `os_state` - The State of the OS used in the Image, such as `Generalized`.
* `os_type` - The type of Operating System used on the OS Disk. such as `Linux` or `Windows`.
* `size_gb` - The size of the OS Disk in GB.

