/* Go Help Syntax */
/*To get more information on a package/library open a command prompt and use:*/
go doc <pkgname> /*replacing <pkgname> with the name of what you're searching for ex: go doc fmt*/

/*Data-types and Definitions
Integer Types - Based on bit size and sign (+/-)
  Signed types are: int8, int16, int32, & int64
    Unsigned types are: uint8, uint16, uint32, & uint64
      Pointer types are written as: uintptr*/
/*The three types are machine dependent and are written without the bit width at the end. The bit-width size(s) are dependent on the system.*/

/*Two aliases are used for integers in Go:
  Byte - Aliased to uint8; Used to [essentially] get to / convert between ASCII characters.
    Rune - Aliased to uint32; Used to [essentially] get to / convert between unicode characters.*/

/* There is NO automatic conversion between int. types. Almost always, int will be used over the others. */
