<template>
    <div id="login">
        <h1>Login</h1>
        <input type="text" name="name" v-model="input.name" placeholder="Username" />
        <input type="password" name="password" v-model="input.password" placeholder="Password" />
        <button type="button" v-on:click="postData">Login</button>
        <p>{{passwd}}</p>
        <p>{{token}}</p>
    </div>
</template>

<script>
import axios from 'axios'
import QS from 'qs'

export default {
  name: 'Login',
  data () {
    return {
      input: {
        name: '',
        password: ''
      },
      passwd: '',
      token: ''
    }
  },
  methods: {
    postData: function () {
      axios({
        method: 'post',
        url: 'http://127.0.0.1:8000/ans',
        headers: {
          'Content-type': 'application/x-www-form-urlencoded'
        },
        data: QS.stringify({
          name: this.input.name,
          password: this.input.password
        })
      })
        .then((resp) => {
          this.passwd = resp.data.password
          this.token = resp.data.token
          this.$cookies.set('Auth', this.token, '1h')
          if (resp != null) {
            this.$router.push('/secure')
          }
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
