# focareg
[This tokenizer is still under development]

**focareg** is a Japanese tokenizer based on **Longest Match Principle**.  

This tokenizer is named after a famous example of mistokenization in a compound noun "外国人参政権".

The word originally means "voting rights for foreign residents" and it is tokenized into two words "外国人 参政権" in that meaning. However, some of Japanese tokenizers tokenizes it into three words "外国　人参　政権" which means "FOreign CArrot REGime".  

Since focareg is based on Longest Match Principle, it does not make this mistake if the compound noun "外国人参政権" is in the dictionary and any other longer words do not have it as a prefix.

