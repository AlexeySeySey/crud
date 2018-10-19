Vue.prototype.$http = axios;
    new Vue({
        delimiters: ['$(', ')'],
        el:"#Main",
        data: {
            tasks : [{
                    name:'Name1',
                    text:'Text1',
                    updatedAt:'Time1'  
                }],
        },
        methods: {
            display: function(task) {
                 document.getElementById("actualData").innerHTML = 
                 "<h3>"+task.name+"</h3>"+
                 "<p>"+task.text+"</p>"+
                 "<small>"+task.updatedAt+"</small>"
            },
            newOne: function() {
                      
                var title = document.getElementById("titleSome").value
                var text = document.getElementById("textSome").value

                      if ((title == "") || (text == "")){
                          alert("Forget something")
                          throw true
                      }
alert("Starting...")
                // Sending data ???
                //var vm = this;
                /*this.$http
                        .get('/new')
                        .then(function (response) {
                            alert(response.data)
                        }).catch( error => {
                    alert(error.response)
                });*/
                this.$http.post('http://localhost/new', {
                    name: "lol",
                    text: "kek",
                    updatedAt: new Date().toLocaleString()
                }).then(response => {
                   var arr = response.data;
                    for (var i=0; i<response.data.length; i++) {
                        alert(response.data[i])
                    }
                }).catch(error => {
                    alert(error.response)
                });

alert("fine")

                this.tasks.push({
                    name: title,
                    text: text,
                    updateAt: new Date().toLocaleString()
                })
                //document.getElementById("titleSome").value = ""
                //document.getElementById("textSome").value = ""

                },
            }
        })