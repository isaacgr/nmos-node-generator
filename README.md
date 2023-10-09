## NMOS Node Generator

A tool to generate fake NMOS nodes for general or scale testing of a registry. Written in go.

### Download

<a>http://dev.irowell.io/nmos-node-generator/</a>

### Example usage

```bash
./nmos-node-generator --config config.json
```

#### Command line arguments

```
-config string
        Conifg file containing resource generation info (default "config.json")
-random-device-uuid
        Whether to use a random UUID for the device or not (default false)
-random-device-uuid
        Whether to use a random UUID for the devices resource or not (default false)
```

The config file should have a syntax similar to below

Use port 443 for https requests

```json
{
  "registry": {
    "ip": "localhost",
    "port": 8010,
    "scheme": "http",
    "version": "v1.3"
  },
  "resource": {
    "nodes": {
      "count": 1,
      "num_interfaces": 2,
      "name_prefix": "evNode"
    },
    "devices": 1,
    "name_prefix": "evDevice",
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
  }
}
```
