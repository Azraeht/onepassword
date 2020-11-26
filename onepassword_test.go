package onepassword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockItemResponse = `
{
  "uuid": "test-item",
  "templateUuid": "102",
  "trashed": "N",
  "createdAt": "2019-05-18T14:58:54Z",
  "updatedAt": "2019-05-18T15:04:56Z",
  "itemVersion": 2,
  "vaultUuid": "test-vault",
  "details": {
    "fields": [
      {
        "designation": "username",
        "name": "username",
        "type": "T",
        "value": "root"
      },
      {
        "designation": "password",
        "name": "password",
        "type": "P",
        "value": "rootpassword"
      }
    ],
    "notesPlain": "",
    "sections": []
  },
  "overview": {
    "URLs": [],
    "ainfo": "redshift.company.io",
    "pbe": 0,
    "pgrng": false,
    "ps": 0,
    "tags": [],
    "title": "Redshift",
    "url": ""
  }
}
`

var fields = []Field{
	Field{
		Designation: "username",
		Name:        "username",
		Value:       "root",
		Type:        "T",
	},
	Field{
		Designation: "password",
		Name:        "password",
		Value:       "rootpassword",
		Type:        "P",
	},
}

var sections = []Sections{}

var details = Details{
	Fields:   fields,
	Sections: sections,
}

var expectedItem = parsedItem{
	UUID:      "test-item",
	CreatedAt: "2019-05-18T14:58:54Z",
	Details:   details,
}

func TestParseItemResponse(t *testing.T) {
	actualItem, err := parseItemResponse([]byte(mockItemResponse))
	if assert.Nil(t, err) {
		assert.Equal(t, actualItem, expectedItem, "item should equal")
	}
}

func TestGetItem(t *testing.T) {
	opPath, err := buildMockOnePassword()
	if err != nil {
		t.Errorf("failed to build mock 1Password CLI: %s", err)
	}
	client, err := NewClient(opPath, "test-subdomain", "test@subdomain.com", "test-password", "test-secret-key")
	if err != nil {
		t.Errorf("failed to create Client: %s", err)
	}
	assert.Equal(t, "test-session", client.Session)
	actualItem, err := client.GetItem(VaultName("test-vault"), ItemName("test-item"))
	if err != nil {
		t.Errorf("error getting item: %s", err)
	}
	assert.Equal(t, actualItem.UUID, expectedItem.UUID)
}
