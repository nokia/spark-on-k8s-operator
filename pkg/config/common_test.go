package config

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta1"
)

func TestCopyToFileForK8sConfigMap(t *testing.T) {
	app := &v1beta1.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-kerberos",
			Namespace: "default",
			UID:       "foo-123",
		},
		Spec: v1beta1.SparkApplicationSpec{
			HadoopConfigMap: stringptr("test-config"),
			Kerberos: v1beta1.Kerberos{
				KerberosEnabled:   boolptr(true),
				KerberosPrincipal: stringptr("TestUser@NOKIA.COM"),
				KeytabSecret:      stringptr("testuser-secret"),
				KeytabName:        stringptr("TestUser.keytab"),
				Krb5ConfigMap:     stringptr("krb5-test"),
			},
		},
	}

	krb5Content := `
        Some content for krb5.conf
        Some more content`
	newConfigMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *app.Spec.Kerberos.Krb5ConfigMap,
			Namespace: app.Namespace,
		},
		Data: map[string]string{
			"krb5.conf": krb5Content,
		},
	}

	krb5ConfigMap := getCm{configMap: newConfigMap}

	configMapFilesPath, err := copyToFile(krb5ConfigMap, app.Namespace, app.Name, *app.Spec.Kerberos.Krb5ConfigMap)
	if err != nil {
		t.Fatal(err)
	}
	expectedPath := configDir + app.Namespace + "/" + app.Name + "/" + *app.Spec.Kerberos.Krb5ConfigMap
	assert.Equal(t, expectedPath, configMapFilesPath, "ConfigMap directory created is not same as the expected configMap directory")

	krb5File := configMapFilesPath + "/krb5.conf"
	if _, err := os.Stat(krb5File); os.IsNotExist(err) {
		t.Errorf("%s does not exist", krb5File)
	}

	// Deleting the newly created config directory
	rmDir_err := removeDirAfterTest(configDir)
	if rmDir_err != nil {
		t.Errorf("Error in deleting %s. Not a part of feature testing", configDir)
	}
}

func TestCopyToFileForK8sSecret(t *testing.T) {
	app := &v1beta1.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-kerberos",
			Namespace: "default",
			UID:       "foo-123",
		},
		Spec: v1beta1.SparkApplicationSpec{
			HadoopConfigMap: stringptr("test-config"),
			Kerberos: v1beta1.Kerberos{
				KerberosEnabled:   boolptr(true),
				KerberosPrincipal: stringptr("TestUser@NOKIA.COM"),
				KeytabSecret:      stringptr("testuser-secret"),
				KeytabName:        stringptr("TestUser.keytab"),
				Krb5ConfigMap:     stringptr("krb5-test"),
			},
		},
	}
	keytabContent := []byte(`
        Some content for keytab
        Some more content`)
	newSecret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *app.Spec.Kerberos.KeytabSecret,
			Namespace: app.Namespace,
		},
		Data: map[string][]byte{
			*app.Spec.Kerberos.KeytabName: keytabContent,
		},
	}
	keytabSecret := getSecret{secret: newSecret}
	secretFilesPath, err := copyToFile(keytabSecret, app.Namespace, app.Name, *app.Spec.Kerberos.KeytabSecret)
	if err != nil {
		t.Fatal(err)
	}
	expectedPath := configDir + app.Namespace + "/" + app.Name + "/" + *app.Spec.Kerberos.KeytabSecret
	assert.Equal(t, expectedPath, secretFilesPath, "Secret directory is not same as the expected secret directory")
	keytabFile := secretFilesPath + "/" + *app.Spec.Kerberos.KeytabName
	if _, err := os.Stat(keytabFile); os.IsNotExist(err) {
		t.Errorf("%s does not exist", keytabFile)
	}

	// Deleting the newly created config directory
	rmDir_err := removeDirAfterTest(configDir)
	if rmDir_err != nil {
		t.Errorf("Error in deleting %s. Not a part of feature testing", configDir)
	}
}

func removeDirAfterTest(dir string) error {
	err := os.RemoveAll(dir)
	return err
}

func stringptr(s string) *string {
	return &s
}

func boolptr(b bool) *bool {
	return &b
}
