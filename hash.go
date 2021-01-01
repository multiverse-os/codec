package codec

// Hashing
////////////////////////////////////////////////////////////////////////////////
// In cryptography, hashing refers to the mathematical functions that are
// one-way-functions which takes a message, and produces a digest, d such that
// the inverse function can't be found to restore m from d.
//
// Here are the other properties of a well-designed hashing function:
//
//
////////////////////////////////////////////////////////////////////////////////
//
//  **deterministic** the same message always results in the same hash
//
//  **computationally efficient** to compute the hash value for any given
//                                message
//
//  **infeasible** to generate a message from its hash value except by
//                 brute-force
//
//  **small changes** to a message result in new hash value uncorrelated with
//                    the old hash value
//
//  **difficult** to find two different messages with the same hash value.
//
////////////////////////////////////////////////////////////////////////////////
