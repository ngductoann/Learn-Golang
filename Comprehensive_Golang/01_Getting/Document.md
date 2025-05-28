# Document

<!--toc:start-->
- [Document](#document)
  - [Printing](#printing)
  - [Understanding text](#understanding-text)
<!--toc:end-->

## Printing

- General:
  - %v the value in a default format when printing structs, the plus flag (%+v) adds field names
  - %#v a Go-syntax representation of the value (floating-point infinities and NaNs print as ±Inf and NaN)
  - %T a Go-syntax representation of the type of the value
  - %% a literal percent sign; consumes no value

- Boolean:
  - %t the word true or false

- Integer:
  - %b base 2
  - %c the character represented by the corresponding Unicode code point
  - %d base 10
  - %o base 8
  - %O base 8 with 0o prefix
  - %q a single-quoted character literal safely escaped with Go syntax.
  - %x base 16, with lower-case letters for a-f
  - %X base 16, with upper-case letters for A-F
  - %U Unicode format: U+1234; same as "U+%04X"

- Floating-point and complex constituents:
  - %b decimalless scientific notation with exponent a power of two, in the manner of strconv.FormatFloat with the 'b' format, e.g. -123456p-78
  - %e scientific notation, e.g. -1.234456e+78
  - %E scientific notation, e.g. -1.234456E+78
  - %f decimal point but no exponent, e.g. 123.456
  - %F synonym for %f
  - %g %e for large exponents, %f otherwise. Precision is discussed below.
  - %G %E for large exponents, %F otherwise
  - %x hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
  - %X upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
  - The exponent is always a decimal integer. For formats other than %b the exponent is at least two digits.

- String and slice of bytes (treated equivalently with these verbs):
  - %s the uninterpreted bytes of the string or slice
  - %q a double-quoted string safely escaped with Go syntax
  - %x base 16, lower-case, two characters per byte
  - %X base 16, upper-case, two characters per byte

- Slice:
  - %p address of 0th element in base 16 notation, with leading 0x

- Pointer:
  - %p base 16 notation, with leading 0x
  - The %b, %d, %o, %x and %X verbs also work with pointers, formatting the value exactly as if it were an integer.

- The default format for %v is:
  - bool:                    %t
  - int, int8 etc.:          %d
  - uint, uint8 etc.:        %d, %#x if printed with %#v
  - float32, complex64, etc: %g
  - string:                  %s
  - chan:                    %p
  - pointer:                 %p

- After a backslash, certain single-character escapes represent special values:
  - \a   U+0007 alert or bell
  - \b   U+0008 backspace
  - \f   U+000C form feed
  - \n   U+000A line feed or newline
  - \r   U+000D carriage return
  - \t   U+0009 horizontal tab
  - \v   U+000B vertical tab
  - \\   U+005C backslash
  - \'   U+0027 single quote  (valid escape only within rune literals)
  - \"   U+0022 double quote  (valid escape only within string literals)

## Understanding text

- Ascii: American Standard Code for Information Interchange 2^8 (1 byte) = 256 unique values
- Unicode: 2^32 (4 bytes) = 4,294,967,296 unique values more than enough to account for every character in every language
- Utf-8: (up to 4 bytes) stores unicode as binary
  - If a character needs 1 byte that’s all it will use.
  - If a character needs 4 bytes it will use 4 bytes.
  - Variable length encoding = efficient memory use common characters like “C” take 8 bits, rare characters like “💕” take 32 bits. Other characters take 16 or 24 bits
