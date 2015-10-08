# Installing

Pick the binary that matches your operating system and architecture. For
example, on my Macbook Pro I would use `ftree-darwin-amd64`. Copy that somewhere
in your path and rename to `ftree`. Presto bango.

## Verifying

The best way to verify the release is to examine the PGP signature. You can find
my public key at http://calebthompson.io/pubkey.asc,
<https://pgp.mit.edu/pks/lookup?op=get&search=0x1621ADC2A0ACE70A>, and via GnuPG
with `gpg --recv-keys A0ACE70A`.

You can find a signed and verifiable statement of my ownership of that key at
<https://github.com/calebthompson/i-am>.

Once you have that key, verification with:

```shell
$ gpg --verify ftree-<os>-<arch>.sig ftree-<os>-<arch>
```

If you don't have GnuPG, you can verify the SHA sums.

```shell
$ shasum --algorithm 256 --check ftree-<os>-<arch>.sha256
```
