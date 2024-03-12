# cidrlist

Simple cli to get CIDR list from well known providers

Currently list following providers:

- akamai
- aws
- cloudflare
- gcp
- google

# Usage

```bash
usage: ip-range-provider <command> [<args> ...]

A command-line tool to fetch IP address ranges from well known providers.


Flags:
  --[no-]help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
help [<command>...]
    Show help.

get <provider>
    Get IP from provider (cloudflare, akamai, aws, gcp, and google).
```