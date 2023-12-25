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
        this.getUsers()
    }

    this.auth = () => {
        util.modal("auth", "show")
    }

    this.getUsers = () => {
        util.get("http://localhost:8090/user", info => {
            console.log(info)
            if (info.error && info.error === "invalid_verifier") {
                console.log("invalid verifier")
                this.auth()
                return
            }
            this.users = info
            this.view()
        })
    }

    this.view = () => {
        if (this.users.length < 1) {
            return
        }
        const users_str = this.users.map(st =>
            util.parse(util.tpl.tr, st)).join("")
        console.log(users_str)
        console.log(util.parse(util.tpl.table, {users: users_str}))
        util.id("root").innerHTML
            = util.parse(util.tpl.table, {users: users_str})
    }

    this.authIn = () => {
        util.modal("auth", "hide")
        let strData = JSON.stringify({
            login: util.id("authLogin").value,
            password: util.id("authPassword").value,
        })
        return util.post("/login", strData, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            util.setUser(resp)
            this.getUsers()
        })
    }

    this.init()
}

const main = new Main()
