# KaM Remake Server (Dockerized)

An a server (in a docker!) for fan-made mod for the game Knights and Merchants.

# Tree

```
kam-remake-server-dockerized
├── _kam-server-binaries    # (.gitignored directory, get server binaries at Releases page)
│   ├── r1xxxx (beta \ stable-ng)
│   │   ├── KaM_Remake_Server_linux_x86_64
│   │   └── KaM Remake Server Settings.ini
│   └── r6720 (stable)
│       ├── KaM_Remake_Server_x86_64
│       └── KaM_Remake_Settings.ini
├── docker-compose          # (docker-compose.yml example)
│   ├── .env
│   └── docker-compose.yml
├── entrypoint              # (init sctipt)
│   └── start.sh
├── metrics                 # (prometheus metrics exporter)
│   ├── exporter.go
│   └── metrics.sh
├── .gitignore
├── Dockerfile
├── Makefile
├── Makefile.env            # (build env's)
└── README.md
```

# Server Binaries

Get it from [Releases](https://github.com/ishad0w/kam_remake_server_dockerized/releases)

# Build

1. Customize __Makefile.env__
2. `make build`

# Push

1. Customize __Makefile.env__
2. `make push`
