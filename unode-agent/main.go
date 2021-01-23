// WIP
package main

import (
    "fmt"
    "os/exec"
    "os"
    //"strings"
    "runtime"

    //hash
    "crypto/sha256"
    "io"
    "log"

    // rsa ssh
    "time"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "golang.org/x/crypto/ssh"
    "io/ioutil"

    "crypto"

    //nacl
    "github.com/kevinburke/nacl"
    "github.com/kevinburke/nacl/secretbox"
    "encoding/base64"
    // golang.org/x/crypto/nacl/secretbox
    // golang.org/x/crypto/nacl/box

)



func msg_nacl() {
    key, err := nacl.Load("6368616e676520746869732070617373776f726420746f206120736563726574")
    if err != nil {
        panic(err)
    }
    encrypted := secretbox.EasySeal([]byte("hello world"), key)
    fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
}


func msg_sign(msg []byte) {

  privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
        panic(err)
  }

  publicKey := privateKey.PublicKey

  // msg := []byte("verifiable message")

  // Before signing, we need to hash our message
  // The hash is what we actually sign
  msgHash := sha256.New()
  _, err = msgHash.Write(msg)
  if err != nil {
        panic(err)
  }
  msgHashSum := msgHash.Sum(nil)

  // In order to generate the signature, we provide a random number generator,
  // our private key, the hashing algorithm that we used, and the hash sum
  // of our message
  signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
  if err != nil {
        panic(err)
  }

  // To verify the signature, we provide the public key, the hashing algorithm
  // the hash sum of our message and the signature we generated previously
  // there is an optional "options" parameter which can omit for now
  err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
  if err != nil {
        fmt.Println("could not verify signature: ", err)
        return
  }
  // If we don't get any error from the `VerifyPSS` method, that means our
  // signature is valid
  fmt.Println("signature verified")

}


func msg_encrypt(msg []byte) {

  // The GenerateKey method takes in a reader that returns random bits, and
  // the number of bits
  privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
        panic(err)
  }



  // The public key is a part of the *rsa.PrivateKey struct
  publicKey := privateKey.PublicKey

  // use the public and private keys
  // ...

  encryptedBytes, err := rsa.EncryptOAEP(
        sha256.New(),
        rand.Reader,
        &publicKey,
        // []byte("super secret message"),
        msg,
        nil)
  if err != nil {
        panic(err)
  }

  fmt.Println("encrypted bytes: ", encryptedBytes)

  // The first argument is an optional random data generator (the rand.Reader we used before)
  // we can set this value as nil
  // The OAEPOptions in the end signify that we encrypted the data using OAEP, and that we used
  // SHA256 to hash the input.
  decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
  if err != nil {
        panic(err)
  }

  // We get back the original information in the form of bytes, which we
  // the cast to a string and print
  fmt.Println("decrypted message: ", string(decryptedBytes))
}



// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
        // Private Key generation
        privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
        if err != nil {
                return nil, err
        }

        // Validate Private Key
        err = privateKey.Validate()
        if err != nil {
                return nil, err
        }

        log.Println("Private Key generated")
        return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
        // Get ASN.1 DER format
        privDER := x509.MarshalPKCS1PrivateKey(privateKey)

        // pem.Block
        privBlock := pem.Block{
                Type:    "RSA PRIVATE KEY",
                Headers: nil,
                Bytes:   privDER,
        }

        // Private key in PEM format
        privatePEM := pem.EncodeToMemory(&privBlock)

        return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
        publicRsaKey, err := ssh.NewPublicKey(privatekey)
        if err != nil {
                return nil, err
        }

        pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

        log.Println("Public key generated")
        return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
        err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
        if err != nil {
                return err
        }

        log.Printf("Key saved to: %s", saveFileTo)
        return nil
}

func save_sshkey() {
        savePrivateFileTo := "./id_rsa_test"
        savePublicFileTo := "./id_rsa_test.pub"
        bitSize := 4096

        privateKey, err := generatePrivateKey(bitSize)
        if err != nil {
                log.Fatal(err.Error())
        }

        publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
        if err != nil {
                log.Fatal(err.Error())
        }

        privateKeyBytes := encodePrivateKeyToPEM(privateKey)

        err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
        if err != nil {
                log.Fatal(err.Error())
        }

        err = writeKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
        if err != nil {
                log.Fatal(err.Error())
        }
}










func shell_cmd(name string, args []string) {
    cmd := exec.Command(name, args...)
    out, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(out))
}


func shellcmd() {
    system_os := runtime.GOOS
    if system_os == "linux" {
        name := os.Args[2]
        args := os.Args[3:]
        shell_cmd(name, args)
    }else if system_os == "windows" {
        if len(os.Args) < 3 {
            fmt.Printf("Usage: %v <cmd> <args>\n", os.Args[0])
            fmt.Printf("Example: %v ls -lhat\n", os.Args[0])
            os.Exit(3)
        }
        var name string = "powershell.exe"
        args := os.Args[2:]
        shell_cmd(name, args)
    }else{
        fmt.Printf("OS %v not supported.\n", system_os)
    }

}


func hash_file() []byte {
	f, err := os.Open("/tmp/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%x", h.Sum(nil))
	return h.Sum(nil)
}


func service() {

    // while true loop
    for {
        fmt.Println("Infinite Loop 1")
        time.Sleep(time.Second)
    }

    // Alternative Version
    for true {
        fmt.Println("Infinite Loop 2")
        time.Sleep(time.Second)
    }
}


func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage: %v <fuction>\n", os.Args[0])
        fmt.Printf("Example: %v shellcmd ls -lhat\n", os.Args[0])
        os.Exit(3)
    }
    function := os.Args[1]
    if function == "shellcmd" {
        shellcmd()
    }else if function == "hash_file" {
        // file_hash, err := hash_file()
        file_hash := hash_file()
        fmt.Printf("%x\n", file_hash)
    }else if function == "service" {
        service()
    }else if function == "save_sshkey" {
        save_sshkey()
    }else if function == "msg_sign" {
        msg := []byte("Jeremy verifiable message")
        msg_sign(msg)
    }else if function == "msg_encrypt" {
        msg := []byte("Jeremy encrypted message")
        msg_encrypt(msg)
    }else if function == "msg_nacl" {
        // msg := []byte("Jeremy encrypted message")
        // msg_encrypt(msg)
        msg_nacl()
    }else{
        fmt.Printf("E: Unsupported function %v\n", function)
    }
}
