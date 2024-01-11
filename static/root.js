function Main() {
    this.__proto__ = new Util()
    const util = this.__proto__
    this.users = []
    this.products = []
    this.modals = []
    this.currentUserId = 0
    this.currentProductId = 0

    // AUTH

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

    // USER CREATE

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

    // USER READ

    this.getUsers = () => {
        util.get("/user", info => {
            console.log(info)
            if (info.error && info.error === "invalid_verifier") {
                console.log("invalid verifier")
                this.auth()
                return
            }
            this.users = info
            this.viewUsers()
        })
    }

    this.viewUsers = () => {
        if (this.users.length < 1) {
            return
        }
        const users_str = this.users.map(st =>
            util.parse(util.tplUsers.tr, st)).join("")

        util.id("root_users").innerHTML
            = util.parse(util.tplUsers.table, {users: users_str})
    }

    // USER UPDATE

    this.userEdit = id => {
        // const user = this.users.filter(usr => usr.id === id)[0]
        this.currentUserId = id
        const bsOffcanvas = new bootstrap.Offcanvas(util.id("userUpdate"))
        bsOffcanvas.show()
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
        return util.put(`/user/${this.currentUserId}`, strData, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            this.getUsers()
        })
    }

    // USER DELETE

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


    // PRODUCT CREATE

    this.productCreate = () => {
        util.modal("productCreate", "show")
    }

    this.productCreateSave = () => {
        util.modal("productCreate", "hide")
        let strData = JSON.stringify({
            name: util.id("productCreateName").value,
            description: util.id("productCreateDescription").value,
            price: +util.id("productCreatePrice").value,
            quantity: +util.id("productCreateQuantity").value,
            categoryId: +util.id("productCreateCategoryId").value,
            seasonId: +util.id("productCreateSeasonId").value,
            colorId: +util.id("productCreateColorId").value,
            countryId: +util.id("productCreateCountryId").value,
            genderId: +util.id("productCreateGenderId").value,
            ageGroupId: +util.id("productCreateAgeGroupId").value,
            brandId: +util.id("productCreateBrandId").value,
            sizeId: +util.id("productCreateSizeId").value,
        })
        console.log(strData)
        return util.post(`/product`, strData, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            this.getProducts()
        })
    }

    // PRODUCT READ

    this.getProducts = () => {
        util.get("/product", info => {
            console.log(info)
            if (info.error && info.error === "invalid_verifier") {
                console.log("invalid verifier")
                this.auth()
                return
            }
            this.products = info
            this.viewProducts()
        })
    }

    this.viewProducts = () => {
        if (this.products.length < 1) {
            return
        }
        const products_str = this.products.map(st =>
            util.parse(util.tplProducts.tr, st)).join("")

        util.id("root_products").innerHTML
            = util.parse(util.tplProducts.table, {products: products_str})
    }

    // PRODUCT UPDATE

    this.productEdit = id => {
        this.currentProductId = id
        const bsOffcanvas = new bootstrap.Offcanvas(util.id("productUpdate"))
        bsOffcanvas.show()
    }

    this.productUpdateSave = () => {
        util.modal("productUpdate", "hide")
        console.log(this.currentProductId)
        let strData = JSON.stringify({
            name: util.id("productUpdateName").value,
            description: util.id("productUpdateDescription").value,
            price: +util.id("productUpdatePrice").value,
            quantity: +util.id("productUpdateQuantity").value,
            categoryId: +util.id("productUpdateCategoryId").value,
            seasonId: +util.id("productUpdateSeasonId").value,
            colorId: +util.id("productUpdateColorId").value,
            countryId: +util.id("productUpdateCountryId").value,
            genderId: +util.id("productUpdateGenderId").value,
            ageGroupId: +util.id("productUpdateAgeGroupId").value,
            brandId: +util.id("productUpdateBrandId").value,
            sizeId: +util.id("productUpdateSizeId").value,
        })
        console.log(strData)
        return util.put(`/product/${this.currentProductId}`, strData, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            this.getProducts()
        })
    }

    // USER DELETE

    this.productDelete = id => {
        return util.delete(`/product/${id}`, resp => {
            if (resp.error) {
                this.auth()
                return
            }
            console.log(resp)
            this.getProducts()
        })
    }


    // INIT

    this.init = () => {
        console.log("init")
        console.log(localStorage.getItem("token"))

        if (util.emptyToken()) {
            console.log("empty token")
            this.auth()
        }

        const page = window.location.pathname.split("/").pop()
        if (page === "") {
            console.log("index")
            this.getUsers()
        } else if (page === "admin_products.html") {
            console.log("ap")
            this.getProducts()
        }
    }

    this.init()
}

const main = new Main()
