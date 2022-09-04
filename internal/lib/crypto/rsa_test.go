package crypto

import "testing"

func TestGeneratePair(t *testing.T) {
	privateKey, err := NewPrivateKey(1024)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PRIVATE KEY: %s\n", privateKey.GetBase64())
	t.Logf("PUBLIC KEY: %s\n", privateKey.Public().GetBase64())
}
