// Copyright (c) 2026 ArthIQ Labs LLC. All rights reserved.
// SPDX-License-Identifier: BUSL-1.1

// Package enrichment provides shared helpers for LLM tag processing.
package enrichment

import "strings"

// BuildTagSentence concatenates enrichment tags into a single embeddable string.
// Moods, genres, lyric themes, and keywords are included (in that order).
// Instruments and characters are excluded: instruments describe sound (captured
// by audio_embedding) and character names are proper nouns that embed poorly.
func BuildTagSentence(moods, genres, lyricTags, keywords []string) string {
	var parts []string
	parts = append(parts, moods...)
	parts = append(parts, genres...)
	parts = append(parts, lyricTags...)
	parts = append(parts, keywords...)
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, ", ")
}

// BuildTagsText builds a space-separated string of all taxonomy values for FTS5 indexing.
// All six tag dimensions are included so that full-text search covers every tag surface.
func BuildTagsText(moods, genres, instruments, lyricTags, keywords, characters []string) string {
	var parts []string
	parts = append(parts, moods...)
	parts = append(parts, genres...)
	parts = append(parts, instruments...)
	parts = append(parts, lyricTags...)
	parts = append(parts, keywords...)
	parts = append(parts, characters...)
	return strings.Join(parts, " ")
}
