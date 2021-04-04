## NMOS Node Generator

A tool to generate fake NMOS nodes for general or scale testing of a registry. Written in go.

### Example usage

```bash
./nmos-node-generator --config config.json
```

#### Command line arguments

```
--config string
        Conifg file containing resource generation info (default "config.json")
```

The config file should have a syntax similar to below

```json
{
  "resource": {
    "nodes": 1,
    "devices": 1,
    "sources": {
      "generic": 1,
      "audio": 1,
      "data": 1
    },
    "senders": {
      "video": 1,
      "audio": 1,
      "data": 1
    },
    "receivers": {
      "video": 1,
      "audio": 1,
      "data": 1
    },
    "flows": 1
  },
  "registry": {
    "ip": "localhost",
    "port": 8010,
    "scheme": "http",
    "version": "1.2"
  },
  "resource_post_delay": 0
}
```