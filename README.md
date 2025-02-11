# osquery-nix-extension

[osquery](https://github.com/facebook/osquery) exposes an operating system as a high-performance relational database. This allows you to write SQL-based queries to explore operating system data. With osquery, SQL tables represent abstract concepts such as running processes, loaded kernel modules, open network connections, browser plugins, hardware events or file hashes.

If you're interested in learning more about osquery, visit the [GitHub project](https://github.com/facebook/osquery), the [website](https://osquery.io), and the [users guide](https://osquery.readthedocs.io).

## What is osquery-nix-extension?

osquery supports nix natively already. Many tables are filled with the proper values. However, it does not have a list of nix packages that are installed. There is a `deb_packages` table and a `rpm_packages` table, but no `nix_packages` table. Hence, this extension adds a `nix_packages` table with 3 columns, `name`, `version`, and `store_path`. I'm sure there are more that could be added, but these fit my needs.

I'm also quite sure there are more tables that might be valuable to have, but I currently didn't need any others.

## Install

To install:

```bash
go install github.com/craiggwilson/osquery-nix-extension/cmd/osquery-nix-extension
```

After the binary exists, it can be loaded into `osqueryi` as follows:
```bash
osqueryi --extension $GOPATH/bin/osquery-nix-extension
```

To use the extension with `osqueryd`, have a look at the [extensions documentation](https://osquery.readthedocs.io/en/latest/deployment/extensions/) directly from `osquery`.