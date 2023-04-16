# Key-Value database benchmarks

This repository contains benchmarks for the next key-value databases:

- [bbolt](https://github.com/etcd-io/bbolt)
- [badger](https://github.com/dgraph-io/badger)
- [leveldb](https://github.com/syndtr/goleveldb)

I compared them by the following metrics:

- **Read speed** (duration per operation)
- **Write speed** (duration per operation)
- **Size after filling** for 100, 1000, 10_000, 100_000 elems with value size 10, 100, 1000 bytes
- **Size after filling and deleting each 10 elems** for the same data as above

## Results

You can read them at
the [Google Sheet](https://docs.google.com/spreadsheets/d/11gZrCdfpd4cZcnuycVIDFFNy2jvtoDCJP6TG-O77xzc/edit?usp=sharing).

## Contact

Feel free to contact me if you have any questions or suggestions.

Just can't say that this repo is for my [YouTube Channel](https://youtube.com/@VyacheArt).

## License

MIT