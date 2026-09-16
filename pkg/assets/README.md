![Greenbone Logo](https://www.greenbone.net/wp-content/uploads/gb_new-logo_horizontal_rgb_small.png)

# Overview

The canonical asset taxonomy: one value set per concept, shared so every service
stores and matches the same strings instead of translating between its own enums.
There is no code at this level, only the subpackages.

Subpackages:
* [assetcategory](assetcategory/README.md) - the coarse grouping an asset type belongs to
* [assettype](assettype/README.md) - the provider-neutral asset type catalog, each type carrying its category
* [properties](properties/README.md) - the canonical keys of the normalized fields a discovered resource carries
* [identifier](identifier/README.md) - the asset-identity claim types
