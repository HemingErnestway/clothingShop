function Util() {
    let token = ""
    let user = {}
    const ls = localStorage.getItem("token")

    if (ls != null) token = ls

    this.emptyToken = () => {
        return token === ""
    }

    this.get = (url, callback) => {
        console.log(token)
        fetch(url, {
            method: "GET",
            headers: {
                "Content-type": "application/json",
                "Authorization": token,
            }
        }).then(data => data.json()).then(callback)
    }

    this.post = (url, data, callback) => {
        fetch(url, {
            method: "POST",
            body: data,
            headers: {
                "Content-Type": "application/json",
                "Authorization": token,
            }
        }).then(data => data.json()).then(callback)
    }

    this.put = (url, data, callback) => {
        fetch(url, {
            method: "PUT",
            body: data,
            headers: {
                "Content-Type": "application/json",
                "Authorization": token,
            }
        }).then(data => data.json()).then(callback)
    }

    this.delete = (url, callback) => {
        fetch(url, {
            method: "DELETE",
            headers: {
                "Authorization": token,
            }
        }).then(data => data.json()).then(callback)
    }

    this.id = el => {
        return document.getElementById(el)
    }

    this.$ = el => {
        return document.querySelector(el)
    }

    this.q = el => {
        return document.querySelectorAll(el)
    }

    this.modals = {}

    this.modal = (id, action) => {
        if (!this.modals[id]) {
            this.modals[id] = new bootstrap.Offcanvas('#' + id)
        }
        this.modals[id][action]()
    }

    this.setUser = usr => {
        user = usr
        token = user.token
        localStorage.setItem("token", token)
    }

    this.parse = (content, params) => {
        let param = Object.assign({}, params)
        return content.replace(/{{(\w+)}}/g, (str) => {
            str = str.substring(2, str.length - 2)
            if (param[str] === undefined) {
                return ''
            }
            return param[str]
        })
    }

    this.tpl = {
        table: `
<div class="container">
  <div class="row">
  <button type="button" 
          class="btn btn-success action col-1"  
          data-action="add"
          onClick="main.userCreate()">Create
  </button>
  
  <div class="col-12">
    <table class="table table-bordered">
      <thead>
        <tr>
          <th scope="col">Id</th>
          <th scope="col">Name</th>
          <th scope="col">Surname</th>
          <th scope="col">Email</th>
          <th scope="col">Address</th>
          <th scope="col">BP</th>
          <th scope="col">B-Date</th>
          <th scope="col">Login</th>
          <th scope="col">Password</th>
          <th scope="col">Access</th>
          <th scope="col">Actions</th>
        </tr>
      </thead>
      <tbody id="tableContent">
        {{users}}
      </tbody>
    </table>
  </div>
  </div>
</div>
        `,
        tr: `
<tr>
    <th scope="row">{{uuid}}</th>
    <td>{{name}}</td>
    <td>{{surname}}</td>
    <td>{{email}}</td>
    <td>{{address}}</td>
    <td>{{bonusPoints}}</td>
    <td>{{birthDate}}</td>
    <td>{{login}}</td>
    <td>{{password}}</td>
    <td>{{access}}</td>
    <td>
        <button type="button" 
                class="btn btn-primary" 
                data-bs-toggle="offcanvas"
                data-bs-target="#offcanvasright" 
                aria-controls="offcanvasright" 
                onClick="main.userEdit({{uuid}})">Edit
        </button>
        <button type="button" 
                class="btn btn-danger" 
                data-bs-toggle="offcanvas"
                data-bs-target="#offcanvasright" 
                aria-controls="offcanvasright" 
                onClick="main.userDelete({{uuid}})">Delete
        </button>
    </td>
</tr>
        `
    }
}
