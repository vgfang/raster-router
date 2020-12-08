function update_img(filename, id) {
	// give the server some time to write the new image
	setTimeout(()=>{
		document.getElementById(id).src = filename+"_"+id
	},60)
}
function refresh_both(filename){
	var xhttp = new XMLHttpRequest()
	xhttp.onreadystatechange = () => {
		if (this.readyState==4 && this.status==200) {
			// pause to let file process playout
			setTimeout(()=>{
			}, 300)
			this.responseText;
		}
	};
	xhttp.open("GET", '', true);
	https.send();
}

console.log("fid: ", fid)