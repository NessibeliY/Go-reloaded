# Go-reloaded
<p>The Go-reloaded project receives as arguments the  name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output).</p>
<p>The modifications that the program executes:</p>
<ul>
<li>Every instance of (hex) should replace the word before with the decimal version of the word.</li>
<li>Every instance of (bin) should replace the word before with the decimal version of the word.</li>
<li>Every instance of (up) converts the word before with the Uppercase version of it.</li>
<li>Every instance of (low) converts the word before with the Lowercase version of it.</li>
<li>Every instance of (cap) converts the word before with the capitalized version of it.</li>
<ul>
	<li>For (low), (up), (cap) if a number appears next to it, like so: (low, &lt;number&gt;) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly.</li>
	</ul>
<li>Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one.</li>
<ul>
	<li>Except if there are groups of punctuation like: ... or !?.</li>
	</ul>
<li>The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces.</li>
<li>Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h.</li>
</ul>

<h2 style="font-size:200%;text-align:left;">To run the program</h2>
<ul class="code">
<li><code>go run . sample.txt result.txt</code></li>
</ul>


<h2 style="font-size:150%;text-align:left;">To run the tests</h2>
<ol class="code">
<li><code>cd modification/</code></li>
<li><code>go test -v</code></li>
</ol>