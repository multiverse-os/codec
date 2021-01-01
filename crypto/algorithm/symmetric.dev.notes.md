
## TL/DR

Some key points:

    * Prefer NaCl where you can. Use AES-GCM if AES is required and you have a choice. Use AES-CBC and AES-CTR for compatibility.
    * Always encrypt-then-MAC. Don’t ever just encrypt.
    * Always check assumptions about the message, including its authenticity, before decrypting the message.
    * Think about how you’re getting nonces and IVs, and whether it’s the appropriate method.


## Authenticity and integrity

Why are the integrity and authenticity components required? One common thing that crops up when people try to build a secure system is that only the confidentiality aspect is used: the programmer will use just AES with some non-authenticated mode (like the CBC, OFB, or CTR modes). AES ciphertexts are malleable: they can be modified and the receiver is often none the wiser. In the context of an encrypted instant message or email or web page, it might seem this modification would become obvious. However, an attacker can exploit the different responses between invalid ciphertext (when decryption fails) and invalid message (the plaintext is wrong) to extract the key. Perhaps the most well-known such attack is the padding oracle attack. In other cases, the invalid plaintext might be used to exploit bugs in the message handler. This is especially problematic in systems that use encryption to secure messages intended for automated systems.

Appending an HMAC or using an authenticated mode (like GCM) requires that the attacker prove they have the key used to authenticate the message. Rejecting a message that fails the MAC reduces the possibility of an invalid message. It also means that an attacker that is just sending invalid data has their message dropped before even wasting processor time on decryption.

To effectively authenticate with an HMAC, the HMAC key should be a different key than the AES key. In this book, we use HMACs with AES, and we’ll append the HMAC key to the AES key for the full encryption or decryption key. For AES-256 in CBC or CTR mode with an HMAC-SHA-256, that means 32 bytes of AES key and 32 bytes of HMAC key for a total key size of 64 bytes; the choice of HMAC-SHA-256 is clarified in a later section.

There are several choices for how to apply a MAC. The right answer is to encrypt-then-MAC:

    * **Encrypt-and-MAC** in this case, we would apply a MAC to the plaintext, then send the encrypted plaintext and MAC. In order to verify the MAC, the receiver has to decrypt the message; this still permits an attacker to submit modified ciphertexts with the same problems described earlier. This presents a surface for IND-CCA attacks.
    * **MAC-then-encrypt** a MAC is applied and appended to the plaintext, and both are encrypted. Note that the receiver still has to decrypt the message, and the MAC can be modified by modifying the resulting ciphertext, which is a surface for IND-CCA attacks as well.
    * **Encrypt-then-MAC** encrypt the message and append a MAC of the ciphertext. The receiver verifies the MAC, and does not proceed to decrypt if the MAC is invalid. This removes the IND-CCA surface.

Moxie Marlinspike’s Cryptographic Doom Principle [Moxie11] contains a more thorough discussion of this.

Always use either an authenticated mode (like GCM), or encrypt-then-MAC.



================================================================================

______
#### NaCl

NaCl is the Networking and Cryptography library that has a symmetric library (secretbox) and an asymmetric library (box), and was designed by Daniel J. Bernstein. The additional Go cryptography packages contain an implementation of NaCl. It uses a 32-byte key and 24-byte nonces. A nonce is a number used once: a nonce should never be reused in a set of messages encrypted with the same key, or else a compromise may occur. In some cases, a randomly generated nonce is suitable. In other cases, it will be part of a stateful system; perhaps it is a message counter or sequence number.

The secretbox system uses a stream cipher called “XSalsa20” to provide confidentiality, and a MAC called “Poly1305”. The package uses the data types *[32]byte for the key and *[24]byte for the nonce. Working with these data types may be a bit unfamiliar; the code below demonstrates generating a random key and a random nonce and how to interoperate with functions that expect a []byte.

```
 1 const (
 2         KeySize   = 32
 3         NonceSize = 24
 4 )
 5 
 6 // GenerateKey creates a new random secret key.
 7 func GenerateKey() (*[KeySize]byte, error) {
 8 	    key := new([KeySize]byte)
 9 	    _, err := io.ReadFull(rand.Reader, key[:])
10 	    if err != nil {
11 	    	return nil, err
12 	    }
13 
14 	    return key, nil
15 }
16 
17 // GenerateNonce creates a new random nonce.
18 func GenerateNonce() (*[NonceSize]byte, error) {
19 	    nonce := new([NonceSize]byte)
20 	    _, err := io.ReadFull(rand.Reader, nonce[:])
21 	    if err != nil {
22 	    	return nil, err
23 	    }
24 
25 	    return nonce, nil
26 }
```

NaCl uses the term seal to mean securing a message (such that it now is confidential, and that its integrity and authenticity may be verified), and open to mean recovering a message (verifying its integrity and authenticity, and decrypting the message).

In this code, randomly generated nonces will be used; in the key exchange chapter, this choice will be clarified. Notably, the selected key exchange methods will permit randomly chosen nonces as secure means of obtaining a nonce. In other use cases, this may not be the case! The recipient will need some way of recovering the nonce, so it will be prepended to the message. If another means of getting a nonce is used, there might be a different way to ensure the recipient has the same nonce used to seal the message.

```
 1 var (
 2         ErrEncrypt = errors.New("secret: encryption failed")
 3         ErrDecrypt = errors.New("secret: decryption failed")
 4 )
 5 
 6 // Encrypt generates a random nonce and encrypts the input using
 7 // NaCl's secretbox package. The nonce is prepended to the ciphertext.
 8 // A sealed message will the same size as the original message plus
 9 // secretbox.Overhead bytes long.
10 func Encrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
11 	    nonce, err := GenerateNonce()
12 	    if err != nil {
13 	    	return nil, ErrEncrypt
14 	    }
15 
16 	    out := make([]byte, len(nonce))
17 	    copy(out, nonce[:])
18 	    out = secretbox.Seal(out, message, nonce, key)
19 	    return out, nil
20 }
```

Decryption expects that the message contains a prepended nonce, and we verify this assumption by checking the length of the message. A message that is too short to be a valid encryption message is dropped right away.

```
 1 // Decrypt extracts the nonce from the ciphertext, and attempts to
 2 // decrypt with NaCl's secretbox.
 3 func Decrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
 4 	    if len(message) < (NonceSize + secretbox.Overhead) {
 5 	    	return nil, ErrDecrypt
 6 	    }
 7 
 8 	    var nonce [NonceSize]byte
 9 	    copy(nonce[:], message[:NonceSize])
10 	    out, ok := secretbox.Open(nil, message[NonceSize:], &nonce, key)
11 	    if !ok {
12 	    	return nil, ErrDecrypt
13 	    }
14 
15 	    return out, nil
16 }
```


Keep in mind that random nonces are not always the right choice. We’ll talk more about this in a chapter on key exchanges, where we’ll talk about how we actually get and share the keys that we’re using.

Behind the scenes, NaCl will encrypt a message, then apply a MAC algorithm to this ciphertext to get the final message. This procedure of “encrypt-then-MAC” is how to properly combine an encryption cipher and a MAC.

______
#### AES-GCM

If AES is required or chosen, AES-GCM is often the best choice; it pairs the AES block cipher with the GCM block cipher mode. It is an AEAD cipher: authenticated encryption with additional data. It encrypts some data, which will be authenticated along with some optional additional data that is not encrypted. The key length is 16 bytes for AES-128, 24 bytes for AES-192, or 32 bytes for AES-256. It also takes a nonce as input, and the same caveats apply to the nonce selection here. Another caveat is that GCM is difficult to implement properly, so it is important to vet the quality of the packages that may be used in a system using AES-GCM.

Which key size should you choose? That depends on the application. Generally, if there’s a specification, use the key size indicated. Cryptography Engineering ([Ferg10]) recommends using 256-bit keys; that’s what we’ll use here. Again, the security model for your system should dictate these parameters. In the AES examples in this chapter, changing the key size to 16 will suffice to switch to AES-128 (and 24 for AES-192). The nonce size does not change across the three versions.

Unlike most block cipher modes, GCM provides authentication. It also allows the for the authentication of some additional, unencrypted data along with the ciphertext. Given that it is an AEAD mode (which provides integrity and authenticity), an HMAC does not need to be appended for this mode.

The AEAD type in the crypto/cipher package uses the same “open” and “seal” terms as NaCl. The AES-GCM analogue of the NaCl encryption above would be:

```
 1 // Encrypt secures a message using AES-GCM.
 2 func Encrypt(key, message []byte) ([]byte, error) {
 3         c, err := aes.NewCipher(key)
 4         if err != nil {
 5                 return nil, ErrEncrypt
 6         }
 7 
 8         gcm, err := cipher.NewGCM(c)
 9         if err != nil {
10                 return nil, ErrEncrypt
11         }
12 
13         nonce, err := GenerateNonce()
14         if err != nil {
15                 return nil, ErrEncrypt
16         }
17 
18         // Seal will append the output to the first argument; the usage
19         // here appends the ciphertext to the nonce. The final parameter
20         // is any additional data to be authenticated.
21         out := gcm.Seal(nonce, nonce, message, nil)
22         return out, nil
23 }
```

This version does not provide any additional (unencrypted but authenticated) data in the ciphertext.

Perhaps there is a system in which the message is prefixed with a 32-bit sender ID, which allows the receiver to select the appropriate decryption key. The following example will authenticate this sender ID:

```
 1 // EncryptWithID secures a message and prepends a 4-byte sender ID
 2 // to the message.
 3 func EncryptWithID(key, message []byte, sender uint32) ([]byte, error) {
 4         buf := make([]byte, 4)
 5         binary.BigEndian.PutUint32(buf, sender)
 6         
 7         c, err := aes.NewCipher(key)
 8         if err != nil {
 9                 return nil, ErrEncrypt
10         }
11         
12         gcm, err := cipher.NewGCM(c)
13         if err != nil {
14                 return nil, ErrEncrypt
15         }
16         
17         nonce, err := GenerateNonce()
18         if err != nil {
19                 return nil, ErrEncrypt
20         }
21 
22         buf = append(buf, nonce)
23         buf := gcm.Seal(buf, nonce, message, message[:4])
24         return buf, nil
25 }
```

In order to decrypt the message, the receiver will need to provide the appropriate sender ID as well. As before, we check some basic assumptions about the length of the message first, accounting for the prepended message ID and nonce.

```
 1 func DecryptWithID(message []byte) ([]byte, error) {
 2         if len(message) <= NonceSize+4 {
 3                 return nil, ErrDecrypt
 4         }
 5 
 6     // SelectKeyForID is a mock call to a database or key cache.
 7         id := binary.BigEndian.Uint32(message[:4])
 8         key, ok := SelectKeyForID(id)
 9         if !ok {
10                 return nil, ErrDecrypt
11         }
12 
13         c, err := aes.NewCipher(key)
14         if err != nil {
15                 return nil, ErrDecrypt
16         }
17 
18         gcm, err := cipher.NewGCM(c)
19         if err != nil {
20                 return nil, ErrDecrypt
21         }
22 
23         nonce := make([]byte, NonceSize)
24         copy(nonce, message[4:])
25 
26         // Decrypt the message, using the sender ID as the additional
27         // data requiring authentication.
28         out, err := gcm.Open(nil, nonce, message[4+NonceSize:], message[:4])
29         if err != nil {
30                 return nil, ErrDecrypt
31         }
32         return out, nil
33 }
```

If the message header is altered at all, even if the new sender ID returns the same key, the message will fail to decrypt: any alteration to the additional data results in a decryption failure.
AES-CTR with HMAC

The last options you should consider, if you have a choice, are AES-CTR and AES-CBC with an HMAC. In these ciphersuites, data is first encrypted with AES in the appropriate mode, then an HMAC is appended. In this book, we assume the use of these ciphersuites only when required as part of a specification or for compatibility.

CTR also uses a nonce; again, the nonce must be only ever used once with the same key. Reusing a nonce can be catastrophic, and will leak information about the message; the system will now fail the indistinguishability requirements and therefore becomes insecure. If there is any question as to whether a nonce is unique, a random nonce should be generated. If this is being used for compatibility with an existing system, you’ll need to consider how that system handles nonces.

If you’re using AES-CTR, you’re probably following along with some sort of specification that should specify which HMAC construction to use. The general rule of thumb from the FIPS guidelines is HMAC-SHA-256 for AES-128 and HMAC-SHA-384 for AES-256; Cryptography Engineering ([Ferg10]) and [Perc09] recommend HMAC-SHA-256. We’ll use HMAC-SHA-256 with AES-256.

Here, we’ll encrypt by selecting a random nonce, encrypting the data, and computing the MAC for the ciphertext. The nonce will be prepended to the message and the MAC appended. The message will be encrypted in-place. The key is expected to be the HMAC key appended to the AES key.

```
 1 const (
 2         NonceSize = aes.BlockSize
 3         MACSize = 32 // Output size of HMAC-SHA-256
 4         CKeySize = 32 // Cipher key size - AES-256
 5         MKeySize = 32 // HMAC key size - HMAC-SHA-256
 6 )
 7 
 8 var KeySize = CKeySize + MKeySize
 9 
10 func Encrypt(key, message []byte) ([]byte, error) {
11         if len(key) != KeySize {
12                 return nil, ErrEncrypt
13         }
14 
15         nonce, err := util.RandBytes(NonceSize)
16         if err != nil {
17                 return nil, ErrEncrypt
18         }
19 
20         ct := make([]byte, len(message))
21 
22         // NewCipher only returns an error with an invalid key size,
23         // but the key size was checked at the beginning of the function.
24         c, _ := aes.NewCipher(key[:CKeySize])
25         ctr := cipher.NewCTR(c, nonce)
26         ctr.XORKeyStream(ct, message)
27 
28         h := hmac.New(sha256.New, key[CKeySize:])
29         ct = append(nonce, ct...)
30         h.Write(ct)
31         ct = h.Sum(ct)
32         return ct, nil
33 }
```

In order to decrypt, the message length is checked to make sure it has a nonce, MAC, and a non-zero message size. Then, the MAC is checked. If it’s valid, the message is decrypted.

```
 1 func Decrypt(key, message []byte) ([]byte, error) {
 2         if len(key) != KeySize {
 3                 return nil, ErrDecrypt
 4         }
 5 
 6         if len(message) <= (NonceSize + MACSize) {
 7                 return nil, ErrDecrypt
 8         }
 9 
10         macStart := len(message) - MACSize
11         tag := message[macStart:]
12         out := make([]byte, macStart-NonceSize)
13         message = message[:macStart]
14 
15         h := hmac.New(sha256.New, key[CKeySize:])
16         h.Write(message)
17         mac := h.Sum(nil)
18         if !hmac.Equal(mac, tag) {
19                 return nil, ErrDecrypt
20         }
21 
22         c, _ := aes.NewCipher(key[:CKeySize])
23         ctr := cipher.NewCTR(c, message[:NonceSize])
24         ctr.XORKeyStream(out, message[NonceSize:])
25         return out, nil
26 }
```

______
#### AES-CBC
The previous modes mask the underlying nature of the block cipher: AES operates on blocks of data, and a full block is needed to encrypt or decrypt. The previous modes act as stream ciphers, where messages lengths do not need to be a multiple of the block size. CBC, however, does not act in this way, and requires messages be padded to the appropriate length. CBC also does not use nonces in the same way.

In CBC mode, each block of ciphertext is XOR’d with the previous block. This leads to the question of what the first block is XOR’d with. In CBC, we use a sort of dummy block called an initialisation vector. It may be randomly generated, which is often the right choice. We also noted that with the other encryption schemes, it was possible to use a message or sequence number as the IV: such numbers should not be directly used with CBC. They should be encrypted (using AES-ECB) with a separate IV encryption key. An IV should never be reused with the same message and key.

The standard padding scheme used is the PKCS #7 padding scheme. We pad the remaining bytes with a byte containing the number of bytes of padding: if we have to add three bytes of padding, we’ll append `0x03 0x03 0x03` to the end of our plaintext.

```
1 func pad(in []byte) []byte {
2         padding := 16 - (len(in) % 16)
3         for i := 0; i < padding; i++ {
4                 in = append(in, byte(padding))
5         }
6         return in
7 }

When we unpad, we’ll take the last byte, check to see if it makes sense (does it indicate padding longer than the message? Is the padding more than a block length?), and then make sure that all the padding bytes are present. Always check your assumptions about the message before accepting the message. Once that’s done, we strip the padding characters and return the plain text.

```
 1 func unpad(in []byte) []byte {
 2         if len(in) == 0 {
 3                 return nil
 4         }
 5 
 6         padding := in[len(in)-1]
 7         if int(padding) > len(in) || padding > aes.BlockSize {
 8                 return nil
 9         } else if padding == 0 {
10                 return nil
11         }
12 
13         for i := len(in) - 1; i > len(in)-int(padding)-1; i-- {
14                 if in[i] != padding {
15                         return nil
16                 }
17         }
18         return in[:len(in)-int(padding)]
19 }
```

The padding takes place outside of encryption: we pad before encrypting data and unpad after decrypting.

Encryption is done by padding the message and generating a random IV.

```
 1 func Encrypt(key, message []byte) ([]byte, error) {
 2         if len(key) != KeySize {
 3                 return nil, ErrEncrypt
 4         }
 5 
 6         iv, err := util.RandBytes(NonceSize)
 7         if err != nil {
 8                 return nil, ErrEncrypt
 9         }
10 
11         pmessage := pad(message)
12         ct := make([]byte, len(pmessage))
13 
14         // NewCipher only returns an error with an invalid key size,
15         // but the key size was checked at the beginning of the function.
16         c, _ := aes.NewCipher(key[:CKeySize])
17         ctr := cipher.NewCBCEncrypter(c, iv)
18         ctr.CryptBlocks(ct, pmessage)
19 
20         h := hmac.New(sha256.New, key[CKeySize:])
21         ct = append(iv, ct...)
22         h.Write(ct)
23         ct = h.Sum(ct)
24         return ct, nil
25 }
```

Encryption proceeds much as with CTR mode, with the addition of message padding.

In decryption, we validate two of our assumptions:

    * The message length should be a multiple of the AES block size (which is 16). HMAC-SHA-256 produces a 32-byte MAC, which is also a multiple of the block size; we can check the length of the entire message and not try to check only the ciphertext. A message that isn’t a multiple of the block size wasn’t padded prior to encryption, and therefore is an invalid message.
    * The message must be at least four blocks long: one block for the IV, one block for the message, and two blocks for the HMAC. If an HMAC function is used with a larger output size, this assumption will need to be revisited.

The decryption also checks the HMAC before actually decrypting the message, and verifies that the plaintext was properly padded.

```
 1 func Decrypt(key, message []byte) ([]byte, error) {
 2         if len(key) != KeySize {
 3                 return nil, ErrEncrypt
 4         }
 5 
 6         // HMAC-SHA-256 returns a MAC that is also a multiple of the
 7         // block size.
 8         if (len(message) % aes.BlockSize) != 0 {
 9                 return nil, ErrDecrypt
10         }
11 
12         // A message must have at least an IV block, a message block,
13         // and two blocks of HMAC.
14         if len(message) < (4 * aes.BlockSize) {
15                 return nil, ErrDecrypt
16         }
17 
18         macStart := len(message) - MACSize
19         tag := message[macStart:]
20         out := make([]byte, macStart-NonceSize)
21         message = message[:macStart]
22 
23         h := hmac.New(sha256.New, key[CKeySize:])
24         h.Write(message)
25         mac := h.Sum(nil)
26         if !hmac.Equal(mac, tag) {
27                 return nil, ErrDecrypt
28         }
29 
30         // NewCipher only returns an error with an invalid key size,
31         // but the key size was checked at the beginning of the function.
32         c, _ := aes.NewCipher(key[:CKeySize])
33         ctr := cipher.NewCBCDecrypter(c, message[:NonceSize])
34         ctr.CryptBlocks(out, message[NonceSize:])
35 
36         pt := unpad(out)
37         if pt == nil {
38                 return nil, ErrDecrypt
39         }
40 
41         return pt, nil
42 }
```

_________
#### Messages v. streams

In this book, we operate on messages: discrete-sized chunks of data. Processing streams of data is more difficult due to the authenticity requirement. How do you supply authentication information? Let’s think about encrypting a stream, trying to provide the same security properties we’ve employed in this chapter.

We can’t encrypt-then-MAC: by it’s nature, we usually don’t know the size of a stream. We can’t send the MAC after the stream is complete, as that usually is indicated by the stream being closed. We can’t decrypt a stream on the fly, because we have to see the entire ciphertext in order to check the MAC. Attempting to secure a stream adds enormous complexity to the problem, with no good answers. The solution is to break the stream into discrete chunks, and treat them as messages. Unfortunately, this means we can’t encrypt or decrypt io.Readers and io.Writers easily, and must operate on []byte messages. Dropping the MAC is simply not an option.
Conclusions

In this chapter, we’ve elided discussion about how we actually get the keys (usually, generating a random key isn’t useful). This is a large enough topic to warrant discussion in a chapter of its own.

