package bus

import (
	"encoding/json"
	"testing"

	"github.com/keepdevops/cofiswarm-observer-sdk/pkg/servicecomponent"
)

func TestInfoRouteReturnsEngine(t *testing.T) {
	out, err := Routes("ollama")[SubjInfo](nil)
	if err != nil {
		t.Fatal(err)
	}
	r := out.(infoReply)
	if !r.OK || r.Engine != "ollama" || !r.Stub {
		t.Fatalf("got %+v", r)
	}
}

func TestHealthRouteOK(t *testing.T) {
	out, _ := Routes("ollama")[SubjHealth](nil)
	if r := out.(healthReply); !r.OK || r.Status != "ok" {
		t.Fatalf("got %+v", r)
	}
}

func TestReplyCarriesSchemaVersion(t *testing.T) {
	out, _ := Routes("ollama")[SubjInfo](nil)
	b, _ := json.Marshal(out)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	if m["schema_version"] != servicecomponent.SchemaVersion {
		t.Fatalf("schema_version = %v, want %s", m["schema_version"], servicecomponent.SchemaVersion)
	}
}

func TestSubjectsNamespaced(t *testing.T) {
	if SubjInfo != "swarm.observer.infer.ollama.info" || SubjHealth != "swarm.observer.infer.ollama.health" {
		t.Fatalf("subjects drifted: %s %s", SubjInfo, SubjHealth)
	}
}
