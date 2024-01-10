function Main() {
    this.__proto__ = new Util()
    const util = this.__proto__
    this.users = []
    this.modals = []
    this.currentId = 0

    this.userEdit = id => {
        const user = this.users.filter(usr => usr.id === id)[0]
        this.currentId = id
        const bsOffcanvas = new bootstrap.Offcanvas(document.getElementById("userUpdate"))
        bsOffcanvas.show()
    }

    this.auth = () => {
        util.modal("auth", "show")
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

    this.userUpdate = () => {
        util.modal("userUpdate", "show")
    }

    this.userUpdateSave = () => {
        util.modal("userUpdate", "hide")
        let strData = JSON.stringify({
            name: util.id("userUpdateName").value,
            surname: util.id("userUpdateSurname").value,
            email: util.id("userUpdateEmail").value,
            address: util.id("userUpdateAddress").value,
            bonusPoints: +util.id("userUpdateBonusPoints").value,
            birthDate: util.id("userUpdateBirthDate").value,
            login: util.id("userUpdateLogin").value,
            password: util.id("userUpdatePassword").value,
            access: +util.id("userUpdateAccess").value,
        })
        console.log(strData)
        return util.put(`/user/${this.currentId}`, strData, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            this.getUsers()
        })
    }

    this.userDelete = id => {
        return util.delete(`/user/${id}`, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            this.getUsers()
        })
    }

    this.userCreate = () => {
        util.modal("userCreate", "show")
    }

    this.userCreateSave = () => {
        util.modal("userCreate", "hide")
        let strData = JSON.stringify({
            name: util.id("userCreateName").value,
            surname: util.id("userCreateSurname").value,
            email: util.id("userCreateEmail").value,
            address: util.id("userCreateAddress").value,
            birthDate: util.id("userCreateBirthDate").value,
            login: util.id("userCreateLogin").value,
            password: util.id("userCreatePassword").value,
        })
        console.log(strData)
        return util.post(`/user`, strData, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            this.getUsers()
        })
    }

    this.getUsers = () => {
        util.get("/user", info => {
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

    this.init = () => {
        console.log("init")
        console.log(localStorage.getItem("token"))
        util.q(".action").forEach(el => {
            el.onclick = () => {
                this[el.dataset.action](el.dataset)
            }
        })
        if (util.emptyToken()) {
            console.log("empty token")
            this.auth()
        }
        this.getUsers()
    }

    this.init()
}

const main = new Main()
