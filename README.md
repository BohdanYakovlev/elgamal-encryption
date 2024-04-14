# ElGamal Encryption

This console application implements the ElGamal encryption algorithm for secure communication of messages. ElGamal encryption is an asymmetric key encryption algorithm used for public-key cryptography.

## How to Run the Code

To run this console app, follow these steps:

1. Clone this repository to your local machine:

```
git clone <repository_url>
```

2. Navigate to the directory containing the Go code:

```
cd elgamal-encryption
```

3. Ensure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).

4. Execute the Go code by running the following command:

```
go run main.go
```
5. The program will generate a pair of public and private keys, encrypt two sample messages, and then decrypt them.
## Overview
The ElGamal encryption system consists of three algorithms:

  - Key Generation: This algorithm generates the public and private keys used for encryption and decryption.
  - Encryption: It takes a message and the recipient's public key to produce a ciphertext.
  - Decryption: It takes a ciphertext and the recipient's private key to recover the original message.

## Key Generation
The key generation process in ElGamal encryption involves the generation of a public and private key pair. Here's a summary of the steps involved:

  - Generate Prime Number: A large prime number, denoted as p, is generated.
  - Find Primitive Root: A primitive root modulo p, denoted as g, is found. A primitive root is an integer g such that every integer from 1 to p-1 can be expressed as a power of g modulo p.
  - Generate Private Key: A random integer, denoted as cKey, is generated as the private key.
  - Compute Public Key: The public key, denoted as (y, g, p), is computed as (y = (g^cKey) mod p).
## Encryption
In ElGamal encryption, a message is encrypted using the recipient's public key. Here's a summary of the encryption process:

  - Choose Random Key: A random integer, denoted as tempKey, is chosen.
  - Compute Ciphertext Components: The ciphertext components, denoted as (a, b), are computed as follows:
  - a = (g^tempKey) mod p
  - b = (message * (y^tempKey)) mod p
## Decryption
Decryption in ElGamal encryption involves using the recipient's private key to recover the original message. Here's a summary of the decryption process:

  - Recover Message: Given a ciphertext (a, b), the original message is recovered using the recipient's private key cKey and the ciphertext components as follows: message = (b * (a^(p-1-cKey))) mod p

## Conclusion

This console application demonstrates the implementation of the ElGamal encryption algorithm for secure communication. It generates keys, encrypts messages, and decrypts them, ensuring confidentiality in communication.

Feel free to reach out if you have any questions or encounter any issues while running the code.
