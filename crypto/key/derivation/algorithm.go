package keyderivation

////////////////////////////////////////////////////////////////////////////////
//  Key Derivation Function (KDF)
////////////////////////////////////////////////////////////////////////////////
//   is a cryptographic hash function that derives one or more secret keys
//   from a secret value such as a main key, a password, or a passphrase
//   using a pseudorandom function.
//
////////////////////////////////////////////////////////////////////////////////

type Algorithm int

const (
	PBKDF2 Algorithm = iota
	Argon2
	BCrypt
	SCru
	Catena
	Lyra2
	Makwa
	YesCrypt
	Balloon
	HKDF
	// Insecure
	MD5
	//
	Blowfish
	Eksblowfish
	//
	SHA256
	SHA512
	//
	DES
	BSDi
	NTHASH
)

// Aliasing
const Blowfish2, Blowfish2a, Blowfish2x, Blowfish2y, Blowfish2b = Blowfish

////////////////////////////////////////////////////////////////////////////////
//
// Key derivation functions supported by crypt
//
//   Over time various algorithms have been introduced. To enable backward
//   compatibility, each scheme started using some convention of serializing
//   the password hashes that was later called the Modular Crypt Format (MCF).
//   Old crypt(3) hashes generated before the de facto MCF standard may vary
//   from scheme to scheme. A well-defined subset of the Modular Crypt Format
//   was created during the Password Hashing Competition.
//
// The format is defined as:
//
//         $<id>[$<param>=<value>(,<param>=<value>)*][$<salt>[$<hash>]]
//
// where
//
//     id: an identifier representing the hashing algorithm (such as 1 for MD5, 5 for SHA-256 etc.)
//
//     param name and its value: hash complexity parameters, like rounds/iterations count
//
//     salt: Base64-like encoded salt
//
//     hash: Base64-like encoded result of hashing the password and salt
//
////////////////////////////////////////////////////////////////////////////////
//
//    Scheme    id 	Schema 	Example
// 	  DES   	  Kyq4bCxAXJkbg
// _ 	BSDi 	    _EQ0.jzhSVeUyoSqLupI
// 1 	MD5 	    $1$etNnh7FA$OlM7eljE/B7F1J4XYNnk81
// 2,           2a, 2x, 2y 	bcrypt 	$2a$10$VIhIOofSMqgdGlL4wzE//e.77dAQGqntF/1dT7bqCrVtquInWy2qi
// 3 	NTHASH 	  $3$$8846f7eaee8fb117ad06bdd830b7586c
// 5 	SHA-256 	$5$9ks3nNEqv31FX.F$gdEoLFsCRsn/WRN3wxUnzfeZLoooVlzeF4WjLomTRFD
// 6 	SHA-512 	$6$qoE2letU$wWPRl.PVczjzeMVgjiA8LLy2nOyZbf7Amj3qLIL978o18gbMySdKZ7uepq9tmMQXxyTIrS12Pln.2Q/6Xscao0
//    md5 	Solaris MD5 	$md5,rounds=5000$GUBv0xjJ$$mSwgIswdjlTY0YxV7HBVm0
//    sha1 	PBKDF1 with SHA-1 	$sha1$40000$jtNX3nZ2$hBNaIXkt4wBI2o5rsi8KejSjNqIq
//
// The PHC subset covers a majority of MCF hashes. A number of extra application-defined methods exist.[3]
//
////////////////////////////////////////////////////////////////////////////////
//
//   $id$salt$hashed
//
//       the printable form of a password hash as produced by crypt (C), where
//       $id is the algorithm used. Other Unix-like systems may have different
//       values, like NetBSD. Key stretching is used to increase password
//       cracking difficulty, using by default 1000 rounds of modified MD5,
//       64 rounds of Blowfish, 5000 rounds of SHA-256 or SHA-512.
//
//       The number of rounds may be varied for Blowfish, or for SHA-256 and
//       SHA-512 by using $A$rounds=X$, where "A" and "X" are the algorithm IDs
//       and the number of rounds.
//
////////////////////////////////////////////////////////////////////////////////
//
// $1$ – MD5
// $2a$ – Blowfish
// $2y$ – Eksblowfish
// $5$ – SHA-256
// $6$ – SHA-512
func (self Algorithm) SchemeId() string {
	switch self {
	case MD5:
		return "1"
	case Blowfish, Blowfish2a, Blowfish2x, Blowfish2y, Eksblowfish, Blowfish2b:
		return BlowfishSchemeId.String()
	case SHA256:
		return "5"
	case SHA512:
		return "6"
	default:
		return ""
	}
}

////////////////////////////////////////////////////////////////////////////////
//
//  Blowfish-based scheme
//
//    Niels Provos and David Mazières designed a crypt() scheme called bcrypt
//    based on Blowfish, and presented it at USENIX in 1999. The printable form
//    of these hashes starts with `$2$`, `$2a$`, `$2b$`, `$2x$` or `$2y$`
//    depending on which variant of the algorithm is used:
//
//      $2$ – Obsolete.
//
//      $2a$ – The current key used to identify this scheme. Since a major
//             security flaw was discovered in 2011 in a non-OpenBSD
//             crypt_blowfish implementation of the algorithm, hashes indicated
//             by this string are now ambiguous and might have been generated
//             by the flawed implementation, or a subsequent fixed,
//             implementation. The flaw may be triggered by some password
//             strings containing non-ASCII (8th-bit-set) characters.
//
//      $2b$ – Used by recent OpenBSD implementations to include a mitigation
//             to a wraparound problem. Previous versions of the algorithm have
//             a problem with long passwords. By design, long passwords are
//             truncated at 72 characters, but there is a byte integer
//             wraparound problem with certain password lengths resulting in
//             weak hashes.
//
//      $2x$ – A flag added after the crypt_blowfish bug discovery. Old hashes
//             can be renamed to be $2x$ to indicate that they were generated
//             with the broken algorithm. These hashes are still weak, but at
//             least it's clear which algorithm was used to generate them.
//
//      $2y$ – A flag in crypt_blowfish to unambiguously use the new, corrected
//             algorithm. On an older implementation suffering from the bug,
//             $2y$ simply won't work. On a newer, fixed implementation, it will
//             produce the same result as using $2a$.
//
////////////////////////////////////////////////////////////////////////////////

func BlowfishSchemeId() Algorithm { return Blowfish2b }

func (self Algorithm) Insecure() bool {
	switch self {
	case Blowfish2a, Blowfish2, Blowfish2x, Blowfish2y, MD5:
		return true
	default:
		return false
	}
}
