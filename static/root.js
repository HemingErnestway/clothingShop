function Main() {
    this.__proto__ = new Util()
    const util = this.__proto__
    this.users = []
    this.modals = []

    this.edit = id => {
        const user = this.users.filter(usr => usr.id == id)[0]
        const bsOffcanvas = new bootstrap.Offcanvas(document.getElementById("auth"))
        bsOffcanvas.show()
    }

    this.rm = id => {

    }

    this.save = ()=> {
        const fd = new FormData(util.id("editForm"))
        const data = {}
        fd.forEach((val, key)=>{
            data[key] = val
        })
        console.log(JSON.stringify(data))
        util.post("http://localhost:8090/user", JSON.stringify(data), ()=>{
            this.reload()
        })
        // util.put()
    }

    this.init = () => {
        util.q(".action").forEach(el => {
            el.onclick = () => {
                this[el.dataset.action](el.dataset)
            }
        })
        if (util.emptyToken()) {
            this.auth()
        }
    }

    this.auth = () => {
        util.modal("auth", "show")
    }

    this.getUsers = () => {
        // util.get("http://localhost:8090/user", info => {
        //     if (info.error && info.error === "invalid_verifier") {
        //         console.log("invalid verifier")
        //         this.auth()
        //         return
        //     }
        //     this.users = info
        //     this.view()
        // })
        this.users = [
            {
                "uuid": 3,
                "name": "Ivan3",
                "surname": "Ivanov",
                "email": "ivan@example.com",
                "address": "Example, Address, 123, 32",
                "bonusPoints": 0,
                "birthDate": "31/01/2002",
                "login": "ivan2003",
                "password": "qwert123",
                "access": 0
            },
            {
                "uuid": 2,
                "name": "Petya",
                "surname": "Ivanov",
                "email": "ivan@example.com",
                "address": "Example, Address, 123, 32",
                "bonusPoints": 0,
                "birthDate": "31/01/2002",
                "login": "ivan2003",
                "password": "qwert123",
                "access": 0
            },
            {
                "uuid": 5,
                "name": "vasya",
                "surname": "vasyin",
                "email": "vv@mail.com",
                "address": "huh puh 123",
                "bonusPoints": 0,
                "birthDate": "12/12/1321",
                "login": "vavava123",
                "password": "qwerty",
                "access": 0
            }
        ]
        this.view()
    }

    this.view = () => {
        if (this.users.length < 1) {
            return
        }
        const users_str = this.users.map(st =>
            util.parse(util.tpl.tr, st)).join("")
        console.log(users_str)
        util.id("root").innerHTML = util.parse(util.tpl.table, {
            users: users_str,
        })
    }
    this.authIn = () => {
        util.modal("auth", "hide")
        util.post("/login", {
            login: util.id("authLogin").value,
            password: util.id("authPassword").value,
        }, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            util.setUser(resp)
            this.getUsers()
        })
    }
    this.init()
}
const main = new Main()
