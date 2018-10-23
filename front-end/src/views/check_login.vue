<template>
    <div id="login">
        <h1>Login</h1>
        <input type="text" name="name" v-model="input.name" placeholder="Username" />
        <input type="password" name="password" v-model="input.password" placeholder="Password" />
        <button type="button" v-on:click="postData">Login</button>
    </div>
</template>

<script>
import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'
export default {
  name: 'Login',
  data () {
    return {
      input: {
        name: '',
        password: ''
      }
    }
  },
  methods: {
    postData: function () {
      let mock = new MockAdapter(axios)
      let _this = this
      mock.onPost('/user').reply(200, {
        id: 2,
        name: 'Hello World'
      })
      axios({
        method: 'post',
        url: '/user',
        headers: {
          'Content-type': 'application/x-www-form-urlencoded'
        },
        data: {
          name: this.input.name,
          password: this.input.password
        }
      })
        .then(function (response) {
          console.log(response)
          alert(`成功送出資料，使用者姓名是 ${_this.input.name}，密碼是 ${_this.input.password}`)
        })
        .catch(error => {
          console.log(error.response)
        })
    }
  }
}
</script>

<style scoped>
    #login {
        width: 500px;
        border: 1px solid #CCCCCC;
        background-color: #FFFFFF;
        margin: auto;
        margin-top: 200px;
        padding: 20px;
    }
</style>
