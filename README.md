# block

## Name

*block* - blocks domains by using pi-hole's block lists.

## Description

The block plugin will block any domain that is on the block lists. The block lists are downloaded on
startup or otherwise once a week.

For a domain that is blocked we will return a NXDOMAIN response.

THIS IS A PROOF OF CONCEPT. IT IS NOT PRODUCTION QUALITY.

## Syntax

~~~ txt
block
~~~

## Metrics

If monitoring is enabled (via the *prometheus* directive) the following metric is exported:

* `coredns_block_count_total{server}` - counter of total number of blocked domains.

The `server` label indicates which server handled the request, see the *metrics* plugin for details.

## Examples

Block all domain on the block list.

``` corefile
. {
  forward . 9.9.9.9
  block
}
```

On startup the block lists are downloaded, and assuming `005.example.org` is on the list, it will
be blocked, including any subdomains.

~~~
2018/09/30 08:40:09 [INFO] plugin/block: Block lists updated: 226126 domains added
2018/09/30 08:40:12 [INFO] plugin/block: Blocked 005.example.org.
2018/09/30 08:41:41 [INFO] plugin/block: Blocked www.005.example.org.
~~~

## Bugs

*Block* currently requires a **working** resolver to fetch the downloads. This should be re-worked
to use the proxy/forwarder (if defined).
