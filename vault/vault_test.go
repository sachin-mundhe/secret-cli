package vault

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func init() {
	os.Remove(secretpath())
}

func secretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".vault.test")
}

func TestSet(t *testing.T) {
	file := secretpath()
	v := NewVault("demo", file)
	err := v.Set("u1", "P123456")
	if err != nil {
		t.Error("Expected nil but got err:", err)
	}

	//Lets add one more key value
	err = v.Set("u2", "P1234567890")
	if err != nil {
		t.Error("Expected nil but got err:", err)
	}

}

func TestSetNegative(t *testing.T) {
	file := secretpath()
	vault := NewVault("", file)
	err := vault.Set("demo", "testing")
	if err == nil {
		t.Error("Expected  Error but got nil")
	}
}

func TestGet(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	_, err := vault.Get("u1")
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestGetNegative(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	_, err := vault.Get("abc")
	if err == nil {
		t.Error("Expected Error but got nil")
	}
	vault = NewVault("", file)
	_, err = vault.Get("abc")
	if err == nil {
		t.Error("Expected Error but got nil ")
	}
}

func TestLoad(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	err := vault.load()
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestLoadNegative(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, "")
	vault := NewVault("abc", file)
	err := vault.load()
	if err == nil {
		t.Error("Expected error but got nil", err)
	}
	os.Remove(file)
}

func TestSave(t *testing.T) {
	var v Vault
	err := v.save()
	if err == nil {
		t.Error("Expected error but got nil ")
	}
	deleteFile()
}

func deleteFile() {
	file := secretpath()
	os.Remove(file)
}
