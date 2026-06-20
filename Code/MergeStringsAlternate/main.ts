function mergeAlternately(word1: string, word2: string): string {
    let result = "";
    let index = 0;
    const len1 = word1.length;
    const len2 = word2.length;

    while (index < len1 || index < len2) {
        if (index < len1) {
            result += word1[index];
        }

        if (index < len2) {
            result += word2[index];
        }

        index++;
    }

    return result;
}

// Example usage
console.log(mergeAlternately("abc", "pqr")); // Output: "apbqcr"
