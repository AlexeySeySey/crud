var V = new Vue({
    delimiters: ['$(', ')'],
    el: "#Main",
    data: {
        tasks: [],
        displays: false
    },
    beforeCreate: function () {
        this.tasks = []
        document.getElementsByTagName("body")[0].style.cssText = "visibility:hidden"
    },
    created: function () {
        axios({
            method: 'get',
            url: '/fetch'
        }).then((r) => {
            for (var i = 0; i < r.data.length; i++) {
                this.tasks.push(r.data[i])
            }
            document.getElementsByTagName("body")[0].style.cssText = "visibility:visible"
        }).catch(function (e) {
            alert(e)
        })
    },
    methods: {
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
        load: function () {
            this.tasks = []
            axios({
                method: 'get',
                url: '/fetch'
            }).then((r) => {
                for (var i = 0; i < r.data.length; i++) {
                    this.tasks.push(r.data[i])
                }
            }).catch(function (e) {
                alert(e.data)
            })
        },
        display: function (task) {
            this.displays = true
            var self = this
            setTimeout(function () {
                document.getElementById("actualData").innerHTML =
                    "<div class='btn-group'>" +
                    "<h3 id='tasksName'>" + self.rule(task.Name, 10) + "</h3>" +
                    "<span id='editNoteButton'>" +
                    "<button class='btn btn-success upInfo' onclick='V.changeOne(" + JSON.stringify(task) + ")'>" +
                    "<i class='fa fa-edit'></i></button></span></div>" +
                    "<p id='taskText'>" + self.rule(task.Text, 90) + "</p>" +
                    "<small>Created: " + task.Created + "</small><br>" +
                    "<small>Updated: " + task.Updated + "</small>"
            }, 1)
        },
        dropOne: function (num) {
            axios({
                method: 'delete',
                url: '/delete/' + num
            }).then((r) => {
                location.reload()
            }).catch((e) => {
                alert(e.data)
            })
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
            }).catch((error) => {
                alert("Post Error : " + error);
            });

            document.getElementById("titleSome").value = ""
            document.getElementById("textSome").value = ""
        },
        changeOne: function (task) {
            var self = this
            document.getElementById("tasksName").innerHTML = "<input id='newName' style='width:810px' type='text' class='form-control' name='new_name' placeholder='Title...' value='" + self.rule(task.Name, 10) + "'>"
            document.getElementById("taskText").innerHTML = "<textarea id='newText' type='text' class='form-control' name='text' placeholder='Text...'>" + self.rule(task.Text, 90) + "</textarea>"
            document.getElementById("editNoteButton").innerHTML = "<button class='btn btn-primary' onclick='V.changeAccept(" + JSON.stringify(task) + ")'><i class='fa fa-check'></i></button>"
        },
        changeAccept: function (task) {
            alert($("#newName").val()+"   "+$("#newText").val()+"   id:"+task.Id)
            axios({
                method: 'put',
                url: '/update/' + task.Id,
                data: $.param({
                    "name": $("#newName").val(),
                    "text": $("#newText").val()
                })
            }).then((r) => {
                alert(r.data)
                //this.displays = false
                this.load()
                //return non-editing state back by changing styles back
                //click automaticly on this specific note by its id
            }).catch((e) => {
                alert(e.data)
            })
        }
    }
})