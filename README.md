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

Use port 443 for https requests

```json
{
  "resource": {
    "nodes": 1,
    "devices": 1,
    "sources": {
      "generic": {
        "count": 1,
        "flows": {
          "media_type": "raw",
          "sender": {
            "iface": [1, 2]
          }
        }
      },
      "audio": {
        "count": 1,
        "flows": {
          "media_type": "audio/L16",
          "sender": {
            "iface": [1, 2]
          }
        }
      },
      "data": {
        "count": 1,
        "flows": {
          "media_type": "smpte291",
          "sender": {
            "iface": [1, 2]
          }
        }
      }
    },
    "receivers": {
      "video": {
        "count": 1,
        "iface": [1, 2],
        "media_type": "raw"
      },
      "audio": {
        "count": 1,
        "iface": [1, 2],
        "media_type": "audio/L16"
      },
      "data": {
        "count": 1,
        "iface": [1, 2],
        "media_type": "smpte291"
      }
    }
  },
  "registry": {
    "ip": "localhost",
    "port": 8010,
    "scheme": "http",
    "version": "1.2"
  }
}
```
