## README

### Overview

This program processes an input file (`input.txt`) line-by-line and applies specific text transformations based on keywords found in each line. The transformations are as follows:

1. **CAPS Transformation**
   Any line containing the marker `(cap)` will be converted to Titlecase ignoring small words like a, an, the, at, if etc  if found within the sentence and not outside.

2. **LOW Transformation**
   Any line containing the marker `(low)` will be converted to lowercase.

3. **TODO -> ACTION Transformation**
   Lines starting with `TODO:` will have the word `TODO` replaced with `ACTION`.

4. **REVERSE -> Reverse Transformation**
   Any line containing the word `(Reverse)` will produce a reversed output.

5. **Truncate Lines Longer Than 80 Characters**
   If a line exceeds 80 characters in length, it will be truncated to 80 characters.

### Transformation Logic

Each line in the input file is processed according to the markers or conditions provided:

* **CAPS and LOW transformations**: These are case-conversion operations. Lines containing `(cap)` are fully converted to Title, and lines containing `(low)` are fully converted to lowercase.

* **TODO -> ACTION transformation**: The `TODO:` tag will be replaced with `ACTION`. For example, `TODO: Complete the task` becomes `ACTION: Complete the task`.

* **CLASSIFIED -> REDACTED transformation**: This operation removes occurrences of the word `(Reverse)` and reverses the whole string.

* **Truncation**: If a line exceeds 80 characters, it is truncated to ensure the text does not exceed this length.

### Conflict and Error Handling

**Overlapping Transformations**:
If a line contains **multiple applicable transformations**, such as both `(cap)` and `(low)`, or both `TODO:` and `(Reverse)`, the program will raise an error. Only **one transformation** can be applied to a line at a time.

For example:

* A line containing both `(cap)` and `(low)` will result in an error.
* A line containing both `TODO:` and `(Reverse)` will also cause an error.

### Error Handling

In the case of conflicting transformations, the program will stop and display the following error message:

```
Error: Only one transformation is allowed per line. Please check the input file.
```

### Input and Output

* **Input File**:
  The program reads the contents of `input.txt`.

* **Output File**:
  After processing, the program writes the modified lines to `output.txt`.

### Summary of Rules Applied

When the program runs, the following rules are applied to each line of the input:

1. **Conversion to CAPS**: Converts the line to Title if it contains `(cap)`.
2. **Conversion to lowercase**: Converts the line to lowercase if it contains `(low)`.
3. **Replace `TODO` with `ACTION`**: Replaces the word `TODO:` with `ACTION:`.
4. **String reversal**: Reverses the string if it contains `(Reverse)`.
5. **Lines longer than 80 characters are truncated**: Any line exceeding 80 characters is cut off at 80 characters.

### Example

**Input:**

```
this is a normal line (low)
THIS SHOULD BE CAPS (cap)
TODO: Complete the report (low)
Information must be secured (Reverse)
```

**Output:**

```
this is a normal line
This Should be Caps
ACTION: Complete the report
secured be must information
```

**Error Case:**

If the input has conflicting transformations, such as:

```
this should be CAPS (cap) and lowercase (low)
```

The program will output an error and not process the file.

---

This README summarizes the behavior of the program, transformation rules, and error handling. If you encounter any issues with the transformations, ensure that only one transformation marker is used per line in the input file.
