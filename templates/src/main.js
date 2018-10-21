new Vue({
    delimiters: ['$(', ')'],
    el: "#Main",
    data: {
        tasks: [],  
        displays: false
    },
    created: function () {
        this.load()
    },
    methods: {
        load: function () {
            axios({
                method: 'get',
                url: '/fetch'
            }).then((r) => {
                for (var i = 0; i < r.data.length; i++) {
                    this.tasks.push(r.data[i])
                }
            }).catch(function (e) {
                alert(e)
            })
        },
        rule: function (str, n) {
            var string = "";
            for (var i = 0; i < str.length; i++) {
                    if ((i % n == 0) && (i != 0)) {
                        string += str[i] + "<br>"
                    } else {
                        string += str[i]
                    }
            }
            return string
        },
        display: function (task) {
            this.displays = true
            var self = this;
            setTimeout(function(){
            document.getElementById("actualData").innerHTML =
                "<h3>" + self.rule(task.Name, 10) + "</h3>" +
                "<p>" + self.rule(task.Text, 90) + "</p>" +
                "<small> Last update: " + task.Updated + "</small>"
            },1);
        },
        newOne: function () {

            var title = document.getElementById("titleSome").value
            var text = document.getElementById("textSome").value

            if ((title == "") || (text == "")) {
                alert("Forget something")
                throw true
            }

            axios({
                method: 'post',
                url: '/new',
                data: $.param({
                    "name": title,
                    "text": text
                })
            }).then((r) => {
                this.load()
                alert(r.data)
            }).catch((error) => {
                alert("Post Error : " + error);
            });

            document.getElementById("titleSome").value = ""
            document.getElementById("textSome").value = ""
        }
    }
})