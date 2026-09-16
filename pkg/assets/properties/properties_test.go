// SPDX-FileCopyrightText: 2026 Greenbone AG
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package properties

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestIsCanonicalMatchesAll(t *testing.T) {
	for _, k := range All {
		if !IsCanonical(k) {
			t.Errorf("IsCanonical(%q) = false, want true (it is in All)", k)
		}
	}
	for _, k := range []string{"", "displayName", "parent", "projectNumber", "vpcId", "ParentID"} {
		if IsCanonical(k) {
			t.Errorf("IsCanonical(%q) = true, want false (not a canonical key)", k)
		}
	}
}

func TestAllHasNoDuplicates(t *testing.T) {
	seen := map[string]struct{}{}
	for _, k := range All {
		if k == "" {
			t.Error("All contains an empty key")
		}
		if _, dup := seen[k]; dup {
			t.Errorf("All lists %q more than once", k)
		}
		seen[k] = struct{}{}
	}
}

func TestDataPlaneGateKeysRegistered(t *testing.T) {
	required := []string{
		PublicNetworkAccess,
		AuthorizedCidrs,
		APIServerPublic,
		APIServerAuthorizedCidrs,
		PrivateNodes,
	}
	for _, k := range required {
		if !IsCanonical(k) {
			t.Errorf("data-plane gate key %q is not canonical; register it in All before producers emit it", k)
		}
	}
}

func TestDerivedGraphFieldsAreNotCanonical(t *testing.T) {
	for _, k := range []string{"ingressAllowSet", "effectiveInbound"} {
		if IsCanonical(k) {
			t.Errorf("derived graph-only field %q must not be canonical (it is written by exposure, never emitted by a producer)", k)
		}
	}
}

// TestAllListsEveryConstant reads the package source so that a constant added
// without an All entry fails here instead of being dropped by every producer
// that filters on IsCanonical.
func TestAllListsEveryConstant(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "properties.go", nil, parser.SkipObjectResolution)
	if err != nil {
		t.Fatal(err)
	}
	declared := map[string]string{}
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.CONST {
			continue
		}
		for _, spec := range gd.Specs {
			vs := spec.(*ast.ValueSpec)
			for i, name := range vs.Names {
				lit, ok := vs.Values[i].(*ast.BasicLit)
				if !ok || lit.Kind != token.STRING {
					t.Fatalf("%s is not a string literal", name.Name)
				}
				declared[lit.Value[1:len(lit.Value)-1]] = name.Name
			}
		}
	}
	for key, name := range declared {
		if !IsCanonical(key) {
			t.Errorf("%s (%q) is declared but missing from All", name, key)
		}
	}
	for _, key := range All {
		if _, ok := declared[key]; !ok {
			t.Errorf("All lists %q, which is not a declared constant", key)
		}
	}
}
