const DATA = [
    {
        name: 'Petro',
        lastname: 'Kuper',
        nickname: 'petku'
    }, 

    {
        name: 'Ivan',
        lastname: 'Ivanov',
        nickname: 'ivanivanov'
    }, 


    {
        name: 'Ivan',
        lastname: 'Ivanov',
        nickname: 'ivanivanov'
    }, 
	
	
    {
        name: 'Ivan',
        lastname: 'Ivanov',
        nickname: 'ivanivanov'
    }, 
	
    {
        name: 'Vlad',
        lastname: 'Vladov',
        nickname: 'vlad'
    } 

];

let userProf = '<div class="user">';
for (let i = 0; i < DATA.length; i++) {
  userProf += "<a href='Чат.html'>" + "<h7>" + DATA[i].name  + ' ' + "</h7>" + 
    "<h7>"  + DATA[i].lastname  + "</h7>" + 
    "<br>" + "<h7>"  + DATA[i].nickname  + "</h7>"  + "</a>";
}
document.getElementById("result").innerHTML = userProf;

document.getElementById('search').onkeyup = function() {
  document.getElementById('result').innerHTML = '';
  let searchText = this.value.toLowerCase();
  let stringLength = searchText.length;
  if (stringLength > 1) {
    for (let i = 0; i < DATA.length; i++) {
      let userName = DATA[i].name.split('').slice(0, stringLength).join('').toLowerCase();
      let userLastname = DATA[i].lastname.split('').slice(0, stringLength).join('').toLowerCase();
      let userNickname = DATA[i].nickname.split('').slice(0, stringLength).join('').toLowerCase();
      if (userName == searchText || userLastname == searchText || userNickname == searchText) {
        document.getElementById('result').innerHTML += "<a href='Чат.html'>" + "<h7>" + DATA[i].name  + ' ' + "</h7>" + 
        "<h7>"  + DATA[i].lastname  + "</h7>" + 
        "<br>" + "<h7>"  + DATA[i].nickname  + "</h7>"  + "</a>";
      }
    }
  } else {
    document.getElementById("result").innerHTML = userProf;
  }
};