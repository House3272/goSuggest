document.addEventListener('DOMContentLoaded', function(){


var searchBox = document.getElementById('searchBox');
var inputBox = document.getElementById('searchInput');
var suggestBox = document.getElementById('suggestBox');

var suggestion0 = document.getElementById('suggestion0');
var suggestionTemplate = suggestion0.cloneNode(true);
suggestionTemplate.removeAttribute("id");
suggestionTemplate.children[0].innerHTML = "";
suggestionTemplate.children[0].setAttribute("onclick", "window.open(this.title);");
suggestBox.removeChild(suggestion0);


var maxResults = 13;




inputBox.addEventListener("input", function() {

	var value = inputBox.value.trim();
	value = value.replace(/\s\s+/g, ' ');

	if( value.length < 1 ){
		suggestBox.style.display = "none";
		searchBox.style.borderRadius = "4px";
		return;
	}
	suggestBox.style.display = "block";
	searchBox.style.borderRadius = "4px 4px 0 0";
	suggestBox.innerHTML = "";
	
	value = value[0].toUpperCase() + value.slice(1);	

	var xhr = new XMLHttpRequest();
	xhr.open("get", "go/suggest/"+value);
	xhr.setRequestHeader("Content-Type", "application/json");
	xhr.onload = function(e) {

		console.log(xhr.status);
		console.log(xhr.response);

		if( xhr.status>300 || xhr.status<199 ){
			suggestion0.children[0].innerHTML = xhr.response;
			suggestBox.appendChild(suggestion0);
			return;
		}

		var response = JSON.parse(xhr.response);
		for (var i=0; (i<response.length) && (i<maxResults) ; i++){
			var tempElem = suggestionTemplate.cloneNode(true);
			tempElem.children[0].innerHTML = response[i];
			tempElem.children[0].setAttribute("title", "https://wikipedia.org/wiki/"+response[i]);
			suggestBox.appendChild(tempElem);
		}

	};
	xhr.send();

});




inputBox.addEventListener("keypress", function(e) {
	var key = e.which || e.keyCode;
	if (key === 13)
		console.log("Enter");
});












// end on ready wrapper
});




	/*test

var temp = ["abbreviators","abc","bananas","abdicate","abdicated","kitten",
"bengal cat","abdication","abdications","abdicator","abdomen","abdomens","abdominal"];

	for (var i=0; i<value.length && i<maxResults ;i++) {
		var tempElem = suggestionTemplate.cloneNode(true);
		tempElem.children[0].innerHTML = temp[i];
		tempElem.children[0].setAttribute("title", "https://wikipedia.org/wiki/"+temp[i]);
		suggestBox.appendChild(tempElem);
	}
	return;
	//*/