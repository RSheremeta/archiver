# archiver

#### A tiny archiver for text data compression using [Shannon-Fano][sf-wiki] algorithm
Most valuable to large data.
Supported formats:
- ``txt`` to ``txt``
- ``txt`` to ``rtf``
- ``rtf`` to ``txt``
- ``rtf`` to ``rtf``

### Usage

There are two ways:
- Clone the repo and build the project using ``go build``
- Download a released binary [here][bin]

#### Archiving
1. ``./archiver pack --method sf <filename>`` or shorted ``./archiver pack -m sf <filename>``
2. This will generate a new compressed file with ``.sf`` extension

#### Unarchiving
1. ``./archiver unpack --method sf <filename>`` or shorted ``./archiver unpack -m sf <filename>``
2. Optional flag: ``--extension`` or ``-e`` used for decompressing target file to a specific extension
3. Available options: ``txt``(default), ``rtf``
4. This will generate a new decompressed file

#### Examples
Feel free to play around with files in the ``/examples`` dir

#### Efficiency
Depending on a data content compression efficiency could be up to 60%.

Compression and decompression results:

![eff-img](https://i.ibb.co/qmy80KS/eff.png)



[sf-wiki]: https://en.wikipedia.org/wiki/Shannon%E2%80%93Fano_coding
[bin]: https://github.com/RSheremeta/archiver/releases/tag/1.0