# Given a string s and an integer k, break up the string into multiple texts such that each text has a length of k or less. You must break it up so that words don't break across lines. If there's no way to break the text up, then return null.
# You can assume that there are no spaces at the ends of the string and that there is exactly one space between each word.
# For example, given the string "the quick brown fox jumps over the lazy dog" and k = 10, you should return: ["the quick", "brown fox", "jumps over", "the lazy", "dog"]. No string in the list has a length of more than 10.


def split_text(input, limit):
    words = input.split(" ")
    results = []
    current_text = ""
    for w in words:
        if len(w) > limit:
            return None
        if current_text:
            if len(current_text) + len(w) + 1 <= limit:
                current_text += " " + w
            else:
                results.append(current_text)
                current_text = w
        else:
            current_text = w
    results.append(current_text)
    return results


print(split_text("the quick brown fox jumps over the lazy dog", 20))
