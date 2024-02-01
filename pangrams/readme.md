#  Pangram
- A pangram is a string that contains every letter of the alphabet. Given a sentence determine whether it is a pangram in the English alphabet. Ignore case. Return either pangram or not pangram as appropriate.

![Screenshot of the Challange from Hackerank](screenshot.png)

## My Thinking
1. Use a set to track the alphabet 
2. Convert the input string to lowercase to ensure case-insensitivity.
3. Iterate over the input string
4. Update the set: check if the current character is a letter.
- If it is, remove this character from the set. If the character is not in the set, it either means it has already been encountered or it's not a letter, so no action is needed.
5. Check if the set is empty: After the loop, if the set is empty, it means all letters of the alphabet were found in the input string, and you can return "pangram".
6. Return the result: If the set is not empty, it means at least one letter of the alphabet was not found in the input string, so you return "not pangram".


