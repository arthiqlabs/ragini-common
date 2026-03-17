# ragini-common

Shared Go modules for the Ragini ecosystem.

| Package | Contents |
|---|---|
| `pkg/schema` | SurrealDB schema types (Track, Profile, PlayEvent, SwipeEvent, TasteVector) |
| `pkg/sargam` | Sargam Synthesis Package format structures |
| `pkg/qdrant` | Qdrant collection definitions and query helpers |
| `pkg/score` | Sargam Score algorithm (Go reference implementation) |

## Usage

```go
import (
    "github.com/ragini-audio/ragini-common/pkg/schema"
    "github.com/ragini-audio/ragini-common/pkg/score"
)
```

## Licence

Business Source Licence 1.1. Copyright © 2026 ArthIQ Labs LLC.
