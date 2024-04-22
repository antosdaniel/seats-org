// Adapted from https://raw.githubusercontent.com/google/safehtml/3c4cd5b5d8c9a6c5882fba099979e9f50b65c876/style.go

// Copyright (c) 2017 The Go Authors. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

package safehtml

import (
	"net/url"
	"regexp"
	"strings"
)

// SanitizeCSS attempts to sanitize CSS properties.
func SanitizeCSS(property, value string) (string, string) {
	property = SanitizeCSSProperty(property)
	if property == InnocuousPropertyName {
		return InnocuousPropertyName, InnocuousPropertyValue
	}
	return property, SanitizeCSSValue(property, value)
}

func SanitizeCSSValue(property, value string) string {
	if sanitizer, ok := cssPropertyNameToValueSanitizer[property]; ok {
		return sanitizer(value)
	}
	return sanitizeRegular(value)
}

func SanitizeCSSProperty(property string) string {
	if !identifierPattern.MatchString(property) {
		return InnocuousPropertyName
	}
	return strings.ToLower(property)
}

// identifierPattern matches a subset of valid <ident-token> values defined in
// https://www.w3.org/TR/css-syntax-3/#ident-token-diagram. This pattern matches all generic family name
// keywords defined in https://drafts.csswg.org/css-fonts-3/#family-name-value.
var identifierPattern = regexp.MustCompile(`^[-a-zA-Z]+$`)

var cssPropertyNameToValueSanitizer = map[string]func(string) string{
	"background-image":    sanitizeBackgroundImage,
	"font-family":         sanitizeFontFamily,
	"display":             sanitizeEnum,
	"background-color":    sanitizeRegular,
	"background-position": sanitizeRegular,
	"background-repeat":   sanitizeRegular,
	"background-size":     sanitizeRegular,
	"color":               sanitizeRegular,
	"height":              sanitizeRegular,
	"width":               sanitizeRegular,
	"left":                sanitizeRegular,
	"right":               sanitizeRegular,
	"top":                 sanitizeRegular,
	"bottom":              sanitizeRegular,
	"font-weight":         sanitizeRegular,
	"padding":             sanitizeRegular,
	"z-index":             sanitizeRegular,
}

var validURLPrefixes = []string{
	`url("`,
	`url('`,
	`url(`,
}

var validURLSuffixes = []string{
	`")`,
	`')`,
	`)`,
}

func sanitizeBackgroundImage(v string) string {
	// Check for <> as per https://github.com/google/safehtml/blob/be23134998433fcf0135dda53593fc8f8bf4df7c/style.go#L87C2-L89C3
	if strings.ContainsAny(v, "<>") {
		return InnocuousPropertyValue
	}
	for _, u := range strings.Split(v, ",") {
		u = strings.TrimSpace(u)
		var found bool
		for i, prefix := range validURLPrefixes {
			if strings.HasPrefix(u, prefix) && strings.HasSuffix(u, validURLSuffixes[i]) {
				found = true
				u = strings.TrimPrefix(u, validURLPrefixes[i])
				u = strings.TrimSuffix(u, validURLSuffixes[i])
				break
			}
		}
		if !found || !urlIsSafe(u) {
			return InnocuousPropertyValue
		}
	}
	return v
}

func urlIsSafe(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		return false
	}
	if u.IsAbs() {
		if strings.EqualFold(u.Scheme, "http") || strings.EqualFold(u.Scheme, "https") || strings.EqualFold(u.Scheme, "mailto") {
			return true
		}
		return false
	}
	return true
}

var genericFontFamilyName = regexp.MustCompile(`^[a-zA-Z][- a-zA-Z]+$`)

func sanitizeFontFamily(s string) string {
	for _, f := range strings.Split(s, ",") {
		f = strings.TrimSpace(f)
		if strings.HasPrefix(f, `"`) {
			if !strings.HasSuffix(f, `"`) {
				return InnocuousPropertyValue
			}
			continue
		}
		if !genericFontFamilyName.MatchString(f) {
			return InnocuousPropertyValue
		}
	}
	return s
}

func sanitizeEnum(s string) string {
	if !safeEnumPropertyValuePattern.MatchString(s) {
		return InnocuousPropertyValue
	}
	return s
}

func sanitizeRegular(s string) string {
	if !safeRegularPropertyValuePattern.MatchString(s) {
		return InnocuousPropertyValue
	}
	return s
}

// InnocuousPropertyName is an innocuous property generated by a sanitizer when its input is unsafe.
const InnocuousPropertyName = "zTemplUnsafeCSSPropertyName"

// InnocuousPropertyValue is an innocuous property generated by a sanitizer when its input is unsafe.
const InnocuousPropertyValue = "zTemplUnsafeCSSPropertyValue"

// safeRegularPropertyValuePattern matches strings that are safe to use as property values.
// Specifically, it matches string where every '*' or '/' is followed by end-of-text or a safe rune
// (i.e. alphanumerics or runes in the set [+-.!#%_ \t]). This regex ensures that the following
// are disallowed:
//   - "/*" and "*/", which are CSS comment markers.
//   - "//", even though this is not a comment marker in the CSS specification. Disallowing
//     this string minimizes the chance that browser peculiarities or parsing bugs will allow
//     sanitization to be bypassed.
//   - '(' and ')', which can be used to call functions.
//   - ',', since it can be used to inject extra values into a property.
//   - Runes which could be matched on CSS error recovery of a previously malformed token, such as '@'
//     and ':'. See http://www.w3.org/TR/css3-syntax/#error-handling.
var safeRegularPropertyValuePattern = regexp.MustCompile(`^(?:[*/]?(?:[0-9a-zA-Z+-.!#%_ \t]|$))*$`)

// safeEnumPropertyValuePattern matches strings that are safe to use as enumerated property values.
// Specifically, it matches strings that contain only alphabetic and '-' runes.
var safeEnumPropertyValuePattern = regexp.MustCompile(`^[a-zA-Z-]*$`)