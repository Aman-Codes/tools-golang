// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v1

import (
	"bytes"
	"testing"

	"github.com/spdx/tools-golang/v0/spdx"
)

// ===== Creation Info section Saver tests =====
func TestSaver2_1CISavesText(t *testing.T) {
	ci := &spdx.CreationInfo2_1{
		SPDXVersion:       "SPDX-2.1",
		DataLicense:       "CC0-1.0",
		SPDXIdentifier:    "SPDXRef-DOCUMENT",
		DocumentName:      "spdx-go-0.0.1.abcdef",
		DocumentNamespace: "https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1.abcdef.whatever",
		ExternalDocumentReferences: []string{
			"DocumentRef-spdx-go-0.0.1a https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1a.cdefab.whatever SHA1:0123456701234567012345670123456701234567",
			"DocumentRef-time-1.2.3 https://github.com/swinslow/spdx-docs/time/time-1.2.3.cdefab.whatever SHA1:0123456701234567012345670123456701234568",
		},
		LicenseListVersion: "2.0",
		CreatorPersons: []string{
			"John Doe",
			"Jane Doe (janedoe@example.com)",
		},
		CreatorOrganizations: []string{
			"John Doe, Inc.",
			"Jane Doe LLC",
		},
		CreatorTools: []string{
			"magictool1-1.0",
			"magictool2-1.0",
			"magictool3-1.0",
		},
		Created:         "2018-10-10T06:20:00Z",
		CreatorComment:  "this is a creator comment",
		DocumentComment: "this is a document comment",
	}

	// what we want to get, as a buffer of bytes
	want := bytes.NewBufferString(`SPDXVersion: SPDX-2.1
DataLicense: CC0-1.0
SPDXID: SPDXRef-DOCUMENT
DocumentName: spdx-go-0.0.1.abcdef
DocumentNamespace: https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1.abcdef.whatever
ExternalDocumentRef: DocumentRef-spdx-go-0.0.1a https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1a.cdefab.whatever SHA1:0123456701234567012345670123456701234567
ExternalDocumentRef: DocumentRef-time-1.2.3 https://github.com/swinslow/spdx-docs/time/time-1.2.3.cdefab.whatever SHA1:0123456701234567012345670123456701234568
LicenseListVersion: 2.0
Creator: Person: John Doe
Creator: Person: Jane Doe (janedoe@example.com)
Creator: Organization: John Doe, Inc.
Creator: Organization: Jane Doe LLC
Creator: Tool: magictool1-1.0
Creator: Tool: magictool2-1.0
Creator: Tool: magictool3-1.0
Created: 2018-10-10T06:20:00Z
CreatorComment: this is a creator comment
DocumentComment: this is a document comment

`)

	// render as buffer of bytes
	var got bytes.Buffer
	err := renderCreationInfo2_1(ci, &got)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c := bytes.Compare(want.Bytes(), got.Bytes())
	if c != 0 {
		t.Errorf("Expected %v, got %v", want.String(), got.String())
	}
}

func TestSaver2_1CIOmitsOptionalFieldsIfEmpty(t *testing.T) {
	// --- need at least one creator; do first for Persons ---
	ci1 := &spdx.CreationInfo2_1{
		SPDXVersion:       "SPDX-2.1",
		DataLicense:       "CC0-1.0",
		SPDXIdentifier:    "SPDXRef-DOCUMENT",
		DocumentName:      "spdx-go-0.0.1.abcdef",
		DocumentNamespace: "https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1.abcdef.whatever",
		CreatorPersons: []string{
			"John Doe",
		},
		Created: "2018-10-10T06:20:00Z",
	}

	// what we want to get, as a buffer of bytes
	want1 := bytes.NewBufferString(`SPDXVersion: SPDX-2.1
DataLicense: CC0-1.0
SPDXID: SPDXRef-DOCUMENT
DocumentName: spdx-go-0.0.1.abcdef
DocumentNamespace: https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1.abcdef.whatever
Creator: Person: John Doe
Created: 2018-10-10T06:20:00Z

`)

	// render as buffer of bytes
	var got1 bytes.Buffer
	err := renderCreationInfo2_1(ci1, &got1)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c1 := bytes.Compare(want1.Bytes(), got1.Bytes())
	if c1 != 0 {
		t.Errorf("Expected %v, got %v", want1.String(), got1.String())
	}

	// --- need at least one creator; now switch to organization ---
	ci2 := &spdx.CreationInfo2_1{
		SPDXVersion:       "SPDX-2.1",
		DataLicense:       "CC0-1.0",
		SPDXIdentifier:    "SPDXRef-DOCUMENT",
		DocumentName:      "spdx-go-0.0.1.abcdef",
		DocumentNamespace: "https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1.abcdef.whatever",
		CreatorOrganizations: []string{
			"John Doe, Inc.",
		},
		Created: "2018-10-10T06:20:00Z",
	}

	// what we want to get, as a buffer of bytes
	want2 := bytes.NewBufferString(`SPDXVersion: SPDX-2.1
DataLicense: CC0-1.0
SPDXID: SPDXRef-DOCUMENT
DocumentName: spdx-go-0.0.1.abcdef
DocumentNamespace: https://github.com/swinslow/spdx-docs/spdx-go/spdx-go-0.0.1.abcdef.whatever
Creator: Organization: John Doe, Inc.
Created: 2018-10-10T06:20:00Z

`)

	// render as buffer of bytes
	var got2 bytes.Buffer
	err = renderCreationInfo2_1(ci2, &got2)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c2 := bytes.Compare(want2.Bytes(), got2.Bytes())
	if c2 != 0 {
		t.Errorf("Expected %v, got %v", want2.String(), got2.String())
	}
}
