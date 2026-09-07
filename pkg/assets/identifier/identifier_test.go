// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package identifier

import "testing"

// TestAllHasNoDuplicates guards against a copy-paste slip in All (two constants
// sharing a value, or the same one listed twice): every claim type is distinct.
func TestAllHasNoDuplicates(t *testing.T) {
	seen := map[Type]struct{}{}
	for _, id := range All {
		if id == "" {
			t.Error("All contains an empty identifier type")
		}
		if _, dup := seen[id]; dup {
			t.Errorf("All lists %q more than once", id)
		}
		seen[id] = struct{}{}
	}
}

// TestAllContainsEveryConstant fails if a new constant is declared but not added
// to All, keeping the iterable set complete for downstream validation.
func TestAllContainsEveryConstant(t *testing.T) {
	declared := []Type{
		Hostname, FQDN, IPv4, IPv6, MACAddress, BiosUUID, SerialNumber, ProviderResourceID,
	}
	if len(All) != len(declared) {
		t.Fatalf("All has %d entries, declared constants %d: add the new type to All", len(All), len(declared))
	}
	inAll := map[Type]struct{}{}
	for _, id := range All {
		inAll[id] = struct{}{}
	}
	for _, id := range declared {
		if _, ok := inAll[id]; !ok {
			t.Errorf("constant %q is missing from All", id)
		}
	}
}

// TestNoUnknownSentinel pins the removal of the "unknown" sentinel: an identifier
// is a claim about identity, so "identity unknown" is not a claim, and keeping it
// in All made Valid("unknown") report true while every consumer rejected it.
func TestNoUnknownSentinel(t *testing.T) {
	if Valid(Type("unknown")) {
		t.Error(`Valid("unknown") = true, want false: the sentinel is not an identity claim`)
	}
	if Precedence(Type("unknown")) != 0 {
		t.Error(`Precedence("unknown") should be 0 like any unrecognized type`)
	}
}

// Both services bound provider_resource_id by this constant, so the test asserts
// the DERIVATION rather than the literal: above the longest id discovery emits,
// and below the btree entry size Postgres refuses to index. Raising or lowering
// it means restating why, and a new collector with a longer id family fails here
// instead of in production.
func TestMaxProviderResourceIDLen(t *testing.T) {
	// The longest id discovery emits today: GCP documents BigQuery dataset ids at
	// 1,024 characters and the collector stores the CAI asset name verbatim.
	const bigQueryWorstCase = len("//bigquery.googleapis.com/projects/") + 30 + len("/datasets/") + 1024
	// The hard wall is Postgres refusing to index a btree entry over 2,704 bytes,
	// measured at 2,700 characters with a 13-character resource_type. This leaves
	// room for a longer resource_type or another index column.
	const indexCeiling = 2560

	if MaxProviderResourceIDLen < bigQueryWorstCase {
		t.Fatalf("MaxProviderResourceIDLen = %d, below the longest id discovery emits (%d, a max-length BigQuery dataset CAI name): a legitimate resource would cost its collector every scan",
			MaxProviderResourceIDLen, bigQueryWorstCase)
	}
	if MaxProviderResourceIDLen > indexCeiling {
		t.Fatalf("MaxProviderResourceIDLen = %d exceeds the btree entry ceiling %d: the composite unique indexes on this column would reject rows at insert time",
			MaxProviderResourceIDLen, indexCeiling)
	}
}

// TestValidAcceptsEveryDeclaredType pairs with TestNoUnknownSentinel: every type
// in All is accepted, and anything outside it is not, so a constant added to the
// catalog without being listed in All is caught as a rejected claim rather than
// silently dropped by a consumer.
func TestValidAcceptsEveryDeclaredType(t *testing.T) {
	for _, id := range All {
		if !Valid(id) {
			t.Errorf("Valid(%q) = false, want true: the type is declared in All", id)
		}
	}
	for _, notAType := range []Type{"", "hostname ", "HOSTNAME", "uuid", "provider-resource-id"} {
		if Valid(notAType) {
			t.Errorf("Valid(%q) = true, want false: the values are matched as exact strings", notAType)
		}
	}
}

// TestPrecedenceOrder asserts the RANKING rather than the numbers: what callers
// depend on is that a hardware-bound claim outranks a name-bound one and that a
// name outranks an address, so the weights can be re-spaced without touching
// this test, while a reordering that changes which claim wins a conflict fails.
func TestPrecedenceOrder(t *testing.T) {
	// Strongest to weakest. IPv4 and IPv6 are deliberately absent: they tie.
	ranked := []Type{BiosUUID, SerialNumber, ProviderResourceID, MACAddress, FQDN, Hostname, IPv4}
	for i := 1; i < len(ranked); i++ {
		stronger, weaker := ranked[i-1], ranked[i]
		if Precedence(stronger) <= Precedence(weaker) {
			t.Errorf("Precedence(%q) = %d, not above Precedence(%q) = %d: the conflict winner would change",
				stronger, Precedence(stronger), weaker, Precedence(weaker))
		}
	}

	// An address is an address: neither IP family is a stronger identity signal.
	if Precedence(IPv4) != Precedence(IPv6) {
		t.Errorf("Precedence(IPv4) = %d and Precedence(IPv6) = %d differ: an asset would match differently by IP family",
			Precedence(IPv4), Precedence(IPv6))
	}

	// A declared claim always carries weight; 0 is reserved for unrecognized types.
	for _, id := range All {
		if Precedence(id) <= 0 {
			t.Errorf("Precedence(%q) = %d: a declared claim type must outweigh an unrecognized one", id, Precedence(id))
		}
	}
}
