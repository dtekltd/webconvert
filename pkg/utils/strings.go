package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func UrlToFilename(rawURL string) (string, error) {
	// Parse the URL
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}

	// Handle hostname (replace dots with underscores)
	hostname := strings.ReplaceAll(u.Hostname(), ".", "_")

	// Handle path (remove leading/trailing slashes, replace slashes with hyphens)
	path := strings.Trim(u.Path, "/")
	path = strings.ReplaceAll(path, "/", "-")

	// Handle query parameters
	var queryParts []string
	for key, values := range u.Query() {
		for _, value := range values {
			// Replace special characters in key/value with underscores
			cleanKey := SanitizeString(key)
			cleanValue := SanitizeString(value)
			queryParts = append(queryParts, fmt.Sprintf("%s_%s", cleanKey, cleanValue))
		}
	}

	// Combine all parts
	var filename string
	if path != "" {
		filename = fmt.Sprintf("%s-%s", hostname, path)
	} else {
		filename = hostname
	}

	if len(queryParts) > 0 {
		filename = fmt.Sprintf("%s-%s", filename, strings.Join(queryParts, "-"))
	}

	return filename, nil
}

// sanitizeString replaces special characters with underscores
func SanitizeString(s string) string {
	// Replace special characters with underscores
	reg := regexp.MustCompile(`[^\w-]`)
	s = reg.ReplaceAllString(s, "_")

	// Remove consecutive underscores
	reg = regexp.MustCompile(`_+`)
	s = reg.ReplaceAllString(s, "_")

	return strings.Trim(s, "_")
}
