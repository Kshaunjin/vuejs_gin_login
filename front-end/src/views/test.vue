<template>
    <div id="app">
        <h1>Login</h1>
        <input type="text" name="name" v-model="input.name" placeholder="Username" />
        <input type="password" name="password" v-model="input.password" placeholder="Password" />
        <button type="button" v-on:click="postData">Login</button>
    </div>
</template>

<script>
import axios from 'axios'
export default {
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
      axios
        .post('http://127.0.0.1:8000/ans', {
          name: this.input.name,
          password: this.input.password
        })
        .then(function (response) {
          alert(`成功送出資料，使用者姓名是 ${this.input.name}，密碼是 ${this.input.password}`)
          console.log(response)
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
