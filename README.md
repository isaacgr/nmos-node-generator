## NMOS Node Generator

A tool to generate fake NMOS nodes for general or scale testing of a registry. Written in go.

### Example usage

```bash
./nmos-node-generator --config config.json
```

#### Command line arguments

```
--config string
        Conifg file containing resource generation info (default "configs/config.json")
```

The config file should have a syntax similar to below

```json
{
  "resource": {
    "nodes": 1,
    "devices": 1,
    "sources": 1,
    "senders": 1,
    "receivers": 1,
    "flows": 1
  },
  "registry": {
    "ip": "172.16.10.10",
    "port": 7893,
    "scheme": "http",
    "version": "1.2"
  },
  "post_delay": 0 // delay between resource posts
}
```