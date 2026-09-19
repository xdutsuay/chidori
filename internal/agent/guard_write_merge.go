package agent

import "strings"

// mergeWriteContent returns the effective file body for write_file. Weak models
// often pass only the new line when the user asked to append; write_file
// otherwise replaces the whole file, which would wipe existing content.
func mergeWriteContent(before, content string) string { panic("fake") }

func looksLikeAppendFragment(before, content string) bool { panic("fake") }

func appendToFile(before, content string) string { panic("fake") }

func firstNonEmptyLine(s string) string { panic("fake") }
