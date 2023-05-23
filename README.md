# natsfx
NATS fx module

## Getting Started

### Example Code
```go
package main

import (
	"github.com/tachunwu/natsfx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		natsfx.Module,
	)

	app.Run()
}

```

### Result
```
[Fx] PROVIDE    *server.Server <= github.com/tachunwu/natsfx.New() from module "nats-server"
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             github.com/tachunwu/natsfx.StartServer() from module "nats-server"
[Fx] HOOK OnStart               github.com/tachunwu/natsfx.New.func1() executing (caller: github.com/tachunwu/natsfx.New)
[65595] [INF] Starting nats-server
[65595] [INF]   Version:  2.9.17
[65595] [INF]   Git:      [not set]
[65595] [INF]   Name:     NDTHRSNW7KRWJDUIQKZ6WKBJPE3IRB5QEU24MXGM3MRDICNTACIXZTA7
[65595] [INF]   Node:     i4lyN3pJ
[65595] [INF]   ID:       NDTHRSNW7KRWJDUIQKZ6WKBJPE3IRB5QEU24MXGM3MRDICNTACIXZTA7
[65595] [INF] Starting JetStream
[65595] [INF]     _ ___ _____ ___ _____ ___ ___   _   __  __
[65595] [INF]  _ | | __|_   _/ __|_   _| _ \ __| /_\ |  \/  |
[65595] [INF] | || | _|  | | \__ \ | | |   / _| / _ \| |\/| |
[65595] [INF]  \__/|___| |_| |___/ |_| |_|_\___/_/ \_\_|  |_|
[65595] [INF] 
[65595] [INF]          https://docs.nats.io/jetstream
[65595] [INF] 
[65595] [INF] ---------------- JETSTREAM ----------------
[65595] [INF]   Max Memory:      12.00 GB
[65595] [INF]   Max Storage:     256.94 GB
[65595] [INF]   Store Directory: "data/jetstream"
[65595] [INF] -------------------------------------------
[65595] [INF] Listening for client connections on 0.0.0.0:4222
[65595] [INF] Server is ready
[Fx] HOOK OnStart               github.com/tachunwu/natsfx.New.func1() called by github.com/tachunwu/natsfx.New ran successfully in 4.35875ms
[Fx] RUNNING
```
