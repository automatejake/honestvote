
var array = [86, 111, 116, 101, 67, 104, 101, 115, 116, 101, 114, 99, 111, 111, 108, 98, 101, 97, 110, 115]

function bin2string(array){
	var result = "";
	for(var i = 0; i < array.length; ++i){
		result+= (String.fromCharCode(array[i]));
	}
	return result;
}

console.log(bin2string(array))

