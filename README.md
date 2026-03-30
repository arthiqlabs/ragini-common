# ragini-common

Shared Go modules for the Ragini ecosystem.

| Package | Contents |
|---|---|
| `pkg/schema` | SurrealDB schema types (Track, Profile, PlayEvent, SwipeEvent, TasteVector, Genre) |
| `pkg/sargam` | Sargam Music Passport format structures |
| `pkg/qdrant` | Qdrant collection definitions and query helpers |
| `pkg/score` | Sargam Score algorithm (Go reference implementation) |

## Usage

```go
import (
    "github.com/ragini-audio/ragini-common/pkg/schema"
    "github.com/ragini-audio/ragini-common/pkg/score"
)
```

## Track struct changes (schema_version 2, 2026-03-29)

The `schema.Track` struct was updated to support the expanded taxonomy schema:

| Old field | New field | Notes |
|-----------|-----------|-------|
| `Embedding []float32` | `AudioEmbedding []float32` | Renamed; 384-dim |
| *(new)* | `LyricEmbedding []float32` | 384-dim lyric embedding |
| `Tags []string` | `Moods []TagWeight` | Join table; multi-label |
| `Genre *string` | `Genres []TagWeight` | Join table; multi-label |
| `HasLyrics bool` + `LyricText string` | `Lyrics *string` | nil=not fetched; ""=fetched/not found |
| `Wishlist bool` | *(removed)* | |
| `MBID *string` | *(removed)* | Not in current schema |
| *(new)* | `Instruments []TagWeight` | Join table |
| *(new)* | `LyricTags []TagWeight` | Join table; lyric themes, situational tags |
| *(new)* | `Keywords []string` | Free-form extracted keywords |
| *(new)* | `Characters []string` | Named characters from lyrics |
| *(new)* | `Description string` | LLM-generated prose description |
| *(new)* | `AnalysisTier int` | 0=none 1=fingerprint 2=acoustic 3=lyrics 4=LLM |
| `AnalysisTier *int` | `AnalysisTier int` | Was pointer, now value |

New types added:
- `TagWeight{Tag string, Weight float64, Source string}` — pairs a label with confidence and provenance (`"acoustic"|"llm"|"id3"|"manual"`)
- `Session{ID, ProfileID, TokenHash, CreatedAt, ExpiresAt}` — authenticated session record

`SwipeEvent.Timestamp` renamed to `SwipeEvent.RatedAt`. `SwipeEvent.Source string` added.

## Licence

Business Source Licence 1.1. Copyright © 2026 ArthIQ Labs LLC.
