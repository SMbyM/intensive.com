const app = new Vue({
    el: '#app',
    data() {
      return {
        form: {}
      }
    },
    methods: {
      sendForm(e) {
  
        e.preventDefault()
        console.log('Отправка JSON данных', this.form)

        fetch(
          e.target.action,
          {
            method: 'POST', 
            mode: 'cors',
            cache: 'no-cache', 
            credentials: 'same-origin', 
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.form) 
          }
        ).then(
          function(response) {
            console.log('Ответ сервера', response);
          },
          function(error) {
            console.error(error);
          }
        );
  
      }
    },
  })