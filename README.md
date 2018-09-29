# block

## Name

*block* - blocks domains by using pi-hole's block lists.

## Description

This block plugin will block any domain looks for names that are on the block lists. The block list
are downloaded on startup or otherwise once a week.

## Syntax

~~~ txt
block
~~~

## Metrics

If monitoring is enabled (via the *prometheus* directive) the following metric is exported:

* `coredns_block_count_total{server}` - counter of total number of blocked domains.

The `server` label indicated which server handled the request, see the *metrics* plugin for details.

## Examples

Block all domain on the block list.

``` corefile
. {
  forward . 9.9.9.9
  block
}
```
