// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package events

import "testing"

func TestMetaKey(t *testing.T) {
	m := Meta{EntityID: "asset-1", Version: 7}
	if got, want := m.Key(), "asset-1@7"; got != want {
		t.Errorf("Key() = %q, want %q", got, want)
	}
}
