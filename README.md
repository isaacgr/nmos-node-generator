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
        Whether to use a random UUID for the device or not (default true)
-random-resource-uuid
        Whether to use a random UUID for the device resources or not (default true)
```

The config file should have a syntax similar to below

Use port 443 for https requests

```json
{
  "registry": {
    "ip": "10.2.0.10",
    "port": 8090,
    "scheme": "http",
    "version": "v1.3"
  },
  "resource": {
    "nodes": {
      "count": 1,
      "num_interfaces": 2,
      "name_prefix": "evNode",
      "attached_network_devices": [
        {
          "chassis_id": "2c-dd-e9-49-2d-8e",
          "port_id": "Ethernet1/1"
        },
        {
          "chassis_id": "2c-dd-e9-49-2d-8e",
          "port_id": "Ethernet2/1"
        }
      ]
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
