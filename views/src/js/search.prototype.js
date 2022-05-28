async function getUsers(){
	let response = await fetch("/getUsers");
	let json = response.json();

	
}

const updateResult = query => {
	let resultList = document.querySelector(".result");
	resultList.innerHTML = "";

	arr.map(algo =>{
		query.split(" ").map(word =>{
			if(algo.toLowerCase().indexOf(word.toLowerCase()) != -1){
				resultList.innerHTML += `<li class="list-group-item">${algo}</li>`;
			}
		})
	})
}

updateResult("")