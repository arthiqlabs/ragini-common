// Copyright (c) 2026 ArthIQ Labs LLC. All rights reserved.
// SPDX-License-Identifier: BUSL-1.1

// Package schema defines the core data types for the Ragini ecosystem.
package schema

import "time"

// TrackSource describes the origin and confidence of a track record.
type TrackSource string

const (
	SourceLocalFLAC            TrackSource = "local_flac"
	SourceAirPlayKnown         TrackSource = "airplay_known"
	SourceAirPlayStream        TrackSource = "airplay_stream"
	SourceAirPlayFingerprinted TrackSource = "airplay_fingerprinted"
	SourceSargam               TrackSource = "sargam" // imported from a .sargam package
)

// SwipeSignal is one of the four swipe vocabulary signals.
type SwipeSignal string

const (
	SignalLove            SwipeSignal = "love"
	SignalDislike         SwipeSignal = "dislike"
	SignalSituationalSkip SwipeSignal = "situational_skip"
	SignalDeepen          SwipeSignal = "deepen"
)

// Track is a music track record.
// ID is SHA-256(chromaprint_fingerprint) — stable across renames and re-encodes.
type Track struct {
	ID           string      `json:"id"`
	Title        string      `json:"title"`
	Artist       string      `json:"artist"`
	Album        string      `json:"album"`
	Year         *int        `json:"year,omitempty"`
	DurationMS   int64       `json:"duration_ms"`
	FilePath     *string     `json:"file_path,omitempty"`
	Source       TrackSource `json:"source"`
	AnalysisTier int         `json:"analysis_tier"`  // 0=none 1=fingerprint 2=acoustic 3=lyrics 4=llm
	AnalysedAt   *time.Time  `json:"analysed_at,omitempty"`
	AddedAt      time.Time   `json:"added_at"`

	// Acoustic features (Tier 2)
	BPM     *float64 `json:"bpm,omitempty"`
	KeyName *string  `json:"key_name,omitempty"`
	Valence *float32 `json:"valence,omitempty"`
	Arousal *float32 `json:"arousal,omitempty"`
	Energy  *float32 `json:"energy,omitempty"`

	// Embeddings (384-dim float32, little-endian)
	AudioEmbedding []float32 `json:"audio_embedding,omitempty"`
	LyricEmbedding []float32 `json:"lyric_embedding,omitempty"`

	// Description (auto-generated prose)
	Description string `json:"description,omitempty"`

	// Lyrics: nil = not yet fetched; pointer to empty string = fetched, not found
	Lyrics *string `json:"lyrics,omitempty"`

	// External canonical IDs
	ISRC     *string `json:"isrc,omitempty"`
	AcoustID *string `json:"acoustid,omitempty"`

	// Taxonomy join fields — populated by store when requested
	Moods       []TagWeight `json:"moods,omitempty"`
	Genres      []TagWeight `json:"genres,omitempty"`
	Instruments []TagWeight `json:"instruments,omitempty"`
	LyricTags   []TagWeight `json:"lyric_tags,omitempty"`
	Keywords    []string    `json:"keywords,omitempty"`
	Characters  []string    `json:"characters,omitempty"`
}

// TagWeight pairs a label with its confidence weight and source provenance.
type TagWeight struct {
	Tag    string  `json:"tag"`
	Weight float64 `json:"weight"`
	Source string  `json:"source,omitempty"` // "acoustic"|"llm"|"id3"|"manual"
}

// Profile is a named user profile within a Ragini instance.
type Profile struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	AvatarColour string  `json:"avatar_colour"`
	PINHash      *string `json:"pin_hash,omitempty"` // bcrypt; never logged
	IsDefault    bool    `json:"is_default"`
}

// SwipeEvent records a single swipe gesture from a user.
type SwipeEvent struct {
	ID        string      `json:"id"`
	TrackID   string      `json:"track_id"`
	ProfileID string      `json:"profile_id"`
	Signal    SwipeSignal `json:"signal"`
	RatedAt   time.Time   `json:"rated_at"`
	Source    string      `json:"source"` // default "local"
}

// PlayEvent records a single playback session for a track.
// CompletedPct is 0.0 at start; updated toward 1.0 as playback progresses.
type PlayEvent struct {
	ID           string    `json:"id"`
	TrackID      string    `json:"track_id"`
	ProfileID    string    `json:"profile_id"`
	StartedAt    time.Time `json:"started_at"`
	CompletedPct float64   `json:"completed_pct"`
	SourcePath   string    `json:"source_path"`
}

// Session is an authenticated user session.
type Session struct {
	ID        string    `json:"id"`
	ProfileID string    `json:"profile_id"`
	TokenHash string    `json:"token_hash"` // bcrypt or SHA-256 of bearer token; never logged
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
