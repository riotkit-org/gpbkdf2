gpbkdf2
=======

Single-binary, tiny and fast shell application for encoding passwords using [Password Based Key Deriviation Function 2 (PBKDF2)](https://en.wikipedia.org/wiki/PBKDF2).

```
PBKDF2 applies a pseudorandom function, such as hash-based message authentication code (HMAC), to the input password or
passphrase along with a salt value and repeats the process many times to produce a derived key, which can then be used as
a cryptographic key in subsequent operations. The added computational work makes password cracking much more difficult, 
and is known as key stretching.
```

Source: Wikipedia

*NOTE: gpbkdf2 has stable API, it is safe to use it in the scripts.*

Usage
-----

```bash
./gpbkdf2 --passphrase=my-secret-passphrase --salt=my-secret-salt --digest-algorithm=sha512 --digest-rounds=10000 --length=128
b9f3da942744045a247ed9e8d2b5cac7987beee58441c56fae44871652534e0e2b78aaa2a1d199fbf460bcade837a0fd43a8d41774c0883726884a4955da196684a81342683f4984b25d3f0284d40b91809c9a2f79bba0cd406a8a48bba6a4c2c12a69036fe3d8dbcada2f9b3f8b19c7b1d4538d36daf75ff1b525ee85ce4606
```

Thanks to
---------

This project strongly uses library [crypto/pbkdf2](golang.org/x/crypto/pbkdf2) library.

Copyleft
--------

Created by **RiotKit Collective**.
Project was created by RiotKit as part of working on grassroot initiatives for social change, the anarchist movement.
