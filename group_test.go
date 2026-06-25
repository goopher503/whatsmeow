package whatsmeow

import (
	"testing"

	"go.mau.fi/whatsmeow/types"
)

func TestBuildSubGroupSuggestionsActionNodeIncludesMultipleSuggestions(t *testing.T) {
	creatorA := types.NewJID("111", types.DefaultUserServer)
	creatorB := types.NewJID("222", types.DefaultUserServer)
	groupA := types.NewJID("120363000000001", types.GroupServer)
	groupB := types.NewJID("120363000000002", types.GroupServer)

	node := buildSubGroupSuggestionsActionNode("approve", []SubGroupSuggestion{
		{JID: groupA, Creator: creatorA},
		{JID: groupB, Creator: creatorB},
	})

	if node.Tag != "sub_group_suggestions_action" {
		t.Fatalf("unexpected root tag: %s", node.Tag)
	}
	actionChildren := node.GetChildren()
	if len(actionChildren) != 1 {
		t.Fatalf("expected one action child, got %d", len(actionChildren))
	}
	action := actionChildren[0]
	if action.Tag != "approve" {
		t.Fatalf("unexpected action tag: %s", action.Tag)
	}
	suggestionChildren := action.GetChildren()
	if len(suggestionChildren) != 2 {
		t.Fatalf("expected two suggestion children, got %d", len(suggestionChildren))
	}

	tests := []struct {
		name    string
		child   int
		group   types.JID
		creator types.JID
	}{
		{name: "first", child: 0, group: groupA, creator: creatorA},
		{name: "second", child: 1, group: groupB, creator: creatorB},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suggestion := suggestionChildren[tt.child]
			if suggestion.Tag != "sub_group_suggestion" {
				t.Fatalf("unexpected suggestion tag: %s", suggestion.Tag)
			}
			if got := suggestion.Attrs["jid"]; got != tt.group {
				t.Fatalf("unexpected group jid: got %v want %v", got, tt.group)
			}
			if got := suggestion.Attrs["creator"]; got != tt.creator {
				t.Fatalf("unexpected creator jid: got %v want %v", got, tt.creator)
			}
		})
	}
}
